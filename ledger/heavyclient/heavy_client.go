//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package heavyclient

import (
	"context"
	"sync"
	"time"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/insmetrics"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/ledger/storage/blob"
	"github.com/insolar/insolar/ledger/storage/drop"
	"github.com/insolar/insolar/ledger/storage/object"
	"github.com/insolar/insolar/ledger/storage/pulse"
	"github.com/insolar/insolar/utils/backoff"
	"github.com/pkg/errors"
	"go.opencensus.io/stats"
)

// Options contains heavy client configuration params.
type Options struct {
	SyncMessageLimit int
	PulsesDeltaLimit int
	BackoffConf      configuration.Backoff
}

// JetClient heavy replication client. Replicates records for one jet.
type JetClient struct {
	bus                    insolar.MessageBus
	replicaStorage         storage.ReplicaStorage
	db                     storage.DBContext
	dropAccessor           drop.Accessor
	blobCollectionAccessor blob.CollectionAccessor
	pulseAccessor          pulse.Accessor
	pulseCalculator        pulse.Calculator
	recSyncAccessor        object.RecordCollectionAccessor
	idxCollectionAccessor  object.IndexCollectionAccessor
	idxCleaner             object.IndexCleaner

	opts Options

	// life cycle control
	//
	startOnce sync.Once
	cancel    context.CancelFunc
	signal    chan struct{}
	// syncdone closes when syncloop is gracefully finished
	syncdone chan struct{}

	// state:
	jetID       insolar.JetID
	muPulses    sync.Mutex
	leftPulses  []insolar.PulseNumber
	syncbackoff *backoff.Backoff
}

// NewJetClient heavy replication client constructor.
//
// First argument defines what jet it serve.
func NewJetClient(
	replicaStorage storage.ReplicaStorage,
	mb insolar.MessageBus,
	pulseAccessor pulse.Accessor,
	pulseCalculator pulse.Calculator,
	dropAccessor drop.Accessor,
	blobSyncAccessor blob.CollectionAccessor,
	recSyncAccessor object.RecordCollectionAccessor,
	idxCollectionAccessor object.IndexCollectionAccessor,
	idxCleaner object.IndexCleaner,
	db storage.DBContext,
	jetID insolar.ID,
	opts Options,
) *JetClient {
	jsc := &JetClient{
		bus:                    mb,
		replicaStorage:         replicaStorage,
		dropAccessor:           dropAccessor,
		blobCollectionAccessor: blobSyncAccessor,
		pulseCalculator:        pulseCalculator,
		pulseAccessor:          pulseAccessor,
		recSyncAccessor:        recSyncAccessor,
		idxCollectionAccessor:  idxCollectionAccessor,
		idxCleaner:             idxCleaner,
		db:                     db,
		jetID:                  insolar.JetID(jetID),
		syncbackoff:            backoffFromConfig(opts.BackoffConf),
		signal:                 make(chan struct{}, 1),
		syncdone:               make(chan struct{}),
		opts:                   opts,
	}
	return jsc
}

// should be called from protected by mutex code
func (c *JetClient) updateLeftPulsesMetrics(ctx context.Context) {
	// instrumentation
	var pn insolar.PulseNumber
	if len(c.leftPulses) > 0 {
		pn = c.leftPulses[0]
	}
	ctx = insmetrics.InsertTag(ctx, tagJet, c.jetID.DebugString())
	stats.Record(ctx,
		statUnsyncedPulsesCount.M(int64(len(c.leftPulses))),
		statFirstUnsyncedPulse.M(int64(pn)),
	)
}

// addPulses add pulse numbers for syncing.
func (c *JetClient) addPulses(ctx context.Context, pns []insolar.PulseNumber) {
	c.muPulses.Lock()
	c.leftPulses = append(c.leftPulses, pns...)

	if err := c.replicaStorage.SetSyncClientJetPulses(ctx, insolar.ID(c.jetID), c.leftPulses); err != nil {
		inslogger.FromContext(ctx).Errorf(
			"attempt to persist jet sync state failed: jetID=%v: %v", c.jetID, err.Error())
	}

	c.updateLeftPulsesMetrics(ctx)
	c.muPulses.Unlock()
}

func (c *JetClient) pulsesLeft() int {
	c.muPulses.Lock()
	defer c.muPulses.Unlock()
	return len(c.leftPulses)
}

// unshiftPulse removes and returns pulse number from head of processing queue.
func (c *JetClient) unshiftPulse(ctx context.Context) *insolar.PulseNumber {
	c.muPulses.Lock()
	defer c.muPulses.Unlock()

	if len(c.leftPulses) == 0 {
		return nil
	}
	result := c.leftPulses[0]

	// shift array elements on one position to left
	shifted := c.leftPulses[:len(c.leftPulses)-1]
	copy(shifted, c.leftPulses[1:])
	c.leftPulses = shifted

	if err := c.replicaStorage.SetSyncClientJetPulses(ctx, insolar.ID(c.jetID), c.leftPulses); err != nil {
		inslogger.FromContext(ctx).Errorf(
			"attempt to persist jet sync state failed: jetID=%v: %v", c.jetID, err.Error())
	}

	c.updateLeftPulsesMetrics(ctx)
	return &result
}

func (c *JetClient) nextPulseNumber() (insolar.PulseNumber, bool) {
	c.muPulses.Lock()
	defer c.muPulses.Unlock()

	if len(c.leftPulses) == 0 {
		return 0, false
	}
	return c.leftPulses[0], true
}

func (c *JetClient) runOnce(ctx context.Context) {
	// retrydelay = m.syncbackoff.ForAttempt(attempt)
	c.startOnce.Do(func() {
		// resets TraceID from and other fields in context
		// (TraceID is mostly meaningless in async sync loop)
		ctx, cancel := context.WithCancel(context.Background())
		c.cancel = cancel
		go c.syncloop(ctx)
	})
}

func (c *JetClient) syncloop(ctx context.Context) {
	inslog := inslogger.FromContext(ctx)
	defer close(c.syncdone)

	var (
		syncPN     insolar.PulseNumber
		hasNext    bool
		retrydelay time.Duration
	)

	finishpulse := func() {
		_ = c.unshiftPulse(ctx)
		c.syncbackoff.Reset()
		retrydelay = 0
	}

	for {
		select {
		case <-time.After(retrydelay):
			// for first try delay should be zero
		case <-ctx.Done():
			if c.pulsesLeft() == 0 {
				// got cancel signal and have nothing to do
				return
			}
			// client in canceled state signal but has smth to do
		}

		for {
			// if we have pulses to sync, process it
			syncPN, hasNext = c.nextPulseNumber()
			if hasNext {
				inslog.Debugf("synchronization next sync pulse num: %v (left=%v)", syncPN, c.leftPulses)
				break
			}

			inslog.Debug("synchronization waiting signal what new pulse happens")
			_, ok := <-c.signal
			if !ok {
				inslog.Info("stop is called, so we are should just stop syncronization loop")
				return
			}
		}

		if isPulseNumberOutdated(ctx, c.pulseAccessor, c.pulseCalculator, syncPN, c.opts.PulsesDeltaLimit) {
			inslog.Infof("pulse %v on jet %v is outdated, skip it", syncPN, c.jetID)
			finishpulse()
			continue
		}

		inslog.Infof("start synchronization to heavy for pulse %v", syncPN)

		shouldretry := false
		syncerr := c.HeavySync(ctx, syncPN)
		inslog := inslog.WithFields(map[string]interface{}{
			"jet_id":  c.jetID.DebugString(),
			"pulse":   syncPN,
			"attempt": c.syncbackoff.Attempt(),
		})
		if syncerr != nil {
			if heavyerr, ok := syncerr.(*reply.HeavyError); ok {
				shouldretry = heavyerr.IsRetryable()
			}

			syncerr = errors.Wrap(syncerr, "HeavySync failed")
			inslog.WithFields(map[string]interface{}{
				"err":       syncerr.Error(),
				"retryable": shouldretry,
			}).Error("sync failed")

			if shouldretry {
				retrydelay = c.syncbackoff.Duration()
				stats.Record(ctx, statSyncedRetries.M(1))
				continue
			}
		} else {
			ctx = insmetrics.InsertTag(ctx, tagJet, c.jetID.DebugString())
			stats.Record(ctx,
				statSyncedPulsesCount.M(1),
			)
			inslog.Info("sync completed")
		}

		finishpulse()
	}

}

// Stop stops heavy client replication
func (c *JetClient) Stop(ctx context.Context) {
	// cancel should be set if client has started
	if c.cancel != nil {
		// two signals for sync loop to stop
		c.cancel()
		close(c.signal)
		// waits sync loop to stop
		<-c.syncdone
	}
}

func backoffFromConfig(bconf configuration.Backoff) *backoff.Backoff {
	return &backoff.Backoff{
		Jitter: bconf.Jitter,
		Min:    bconf.Min,
		Max:    bconf.Max,
		Factor: bconf.Factor,
	}
}

func isPulseNumberOutdated(
	ctx context.Context,
	pulseAccessor pulse.Accessor,
	pulseCalculator pulse.Calculator,
	pn insolar.PulseNumber,
	delta int,
) bool {
	current, err := pulseAccessor.Latest(ctx)
	if err != nil {
		panic(err)
	}

	currentPulse, err := pulseAccessor.ForPulseNumber(ctx, current.PulseNumber)
	if err != nil {
		panic(err)
	}

	pnPulse, err := pulseAccessor.ForPulseNumber(ctx, pn)
	if err != nil {
		inslogger.FromContext(ctx).Errorf("Can't get pulse by pulse number: %v", pn)
		return true
	}

	backPN, err := pulseCalculator.Backwards(ctx, currentPulse.PulseNumber, delta)
	if err != nil {
		return false
	}

	return backPN.PulseNumber >= pnPulse.PulseNumber
}
