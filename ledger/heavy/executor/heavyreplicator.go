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

package executor

import (
	"context"
	"fmt"
	"sync"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/jet"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/insolar/pulse"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/drop"

	"github.com/insolar/insolar/ledger/object"
	"github.com/pkg/errors"
)

// HeavyReplicator is a base interface for a heavy sync component
type HeavyReplicator interface {
	// NotifyAboutMessage is method for notifying a sync component about new data
	NotifyAboutMessage(context.Context, *payload.Replication)

	// Stop stops the component
	Stop()
}

// HeavyReplicatorDefault is a base impl for HeavyReplicator
type HeavyReplicatorDefault struct {
	once sync.Once
	done chan struct{}

	records         object.RecordModifier
	indexes         object.IndexModifier
	pcs             insolar.PlatformCryptographyScheme
	pulseCalculator pulse.Calculator
	drops           drop.Modifier
	keeper          JetKeeper
	backuper        BackupMaker
	jets            jet.Modifier

	syncWaitingData chan *payload.Replication
}

func NewHeavyReplicatorDefault(
	records object.RecordModifier,
	indexes object.IndexModifier,
	pcs insolar.PlatformCryptographyScheme,
	pulseCalculator pulse.Calculator,
	drops drop.Modifier,
	keeper JetKeeper,
	backuper BackupMaker,
	jets jet.Modifier,
) *HeavyReplicatorDefault {
	return &HeavyReplicatorDefault{
		records:         records,
		indexes:         indexes,
		pcs:             pcs,
		pulseCalculator: pulseCalculator,
		drops:           drops,
		keeper:          keeper,
		backuper:        backuper,
		jets:            jets,
	}
}

func (h *HeavyReplicatorDefault) NotifyAboutMessage(ctx context.Context, msg *payload.Replication) {
	h.once.Do(func() {
		go h.sync(context.Background())
	})

	logger := inslogger.FromContext(ctx).WithFields(map[string]interface{}{
		"jet_id": msg.JetID.DebugString(),
		"pulse":  msg.Pulse,
	})
	logger.Info("heavy replicator got a new message")
}

func (h *HeavyReplicatorDefault) Stop() {
	close(h.done)
}

func (h *HeavyReplicatorDefault) sync(ctx context.Context) {
	work := func(msg *payload.Replication) {
		logger := inslogger.FromContext(ctx).WithFields(map[string]interface{}{
			"jet_id": msg.JetID.DebugString(),
			"pulse":  msg.Pulse,
		})
		logger.Info("heavy replicator starts replication")

		logger.Debug("storing records")
		if err := storeRecords(ctx, h.records, h.pcs, msg.Pulse, msg.Records); err != nil {
			logger.Error(errors.Wrap(err, "failed to store records"))
			return
		}

		logger.Debug("storing indexes")
		if err := storeIndexes(ctx, h.indexes, msg.Indexes, msg.Pulse); err != nil {
			logger.Error(errors.Wrap(err, "failed to store indexes"))
			return
		}

		logger.Debug("storing drop")
		dr, err := storeDrop(ctx, h.drops, msg.Drop)
		if err != nil {
			logger.Error(errors.Wrap(err, "failed to store drop"))
			return
		}

		logger.Debug("storing drop confirmation")
		if err := h.keeper.AddDropConfirmation(ctx, dr.Pulse, dr.JetID, dr.Split); err != nil {
			logger.Error(errors.Wrapf(err, "failed to add drop confirmation jet=%v", dr.JetID.DebugString()))
			return
		}

		logger.Debug("update jets")
		err = h.jets.Update(ctx, dr.Pulse, true, dr.JetID)
		if err != nil {
			logger.Error(errors.Wrapf(err, "failed to update jet %s", dr.JetID.DebugString()))
			return
		}

		logger.Debug("finalize pulse")
		FinalizePulse(ctx, h.pulseCalculator, h.backuper, h.keeper, dr.Pulse)
	}

	for {
		select {
		case data, ok := <-h.syncWaitingData:
			if !ok {
				return
			}
			work(data)
		case <-h.done:
			inslogger.FromContext(ctx).Info("heavy replicator stopped")
			return
		}
	}
}

func storeIndexes(
	ctx context.Context,
	mod object.IndexModifier,
	indexes []record.Index,
	pn insolar.PulseNumber,
) error {
	for _, idx := range indexes {
		err := mod.SetIndex(ctx, pn, idx)
		if err != nil {
			return err
		}
	}
	return nil
}

func storeDrop(
	ctx context.Context,
	drops drop.Modifier,
	rawDrop []byte,
) (*drop.Drop, error) {
	d, err := drop.Decode(rawDrop)
	if err != nil {
		inslogger.FromContext(ctx).Error(err)
		return nil, err
	}
	err = drops.Set(ctx, *d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func storeRecords(
	ctx context.Context,
	recordStorage object.RecordModifier,
	pcs insolar.PlatformCryptographyScheme,
	pn insolar.PulseNumber,
	records []record.Material,
) error {
	for _, rec := range records {
		hash := record.HashVirtual(pcs.ReferenceHasher(), rec.Virtual)
		id := *insolar.NewID(pn, hash)
		if rec.ID != id {
			return fmt.Errorf(
				"record id does not match (calculated: %s, received: %s)",
				id.DebugString(),
				rec.ID.DebugString(),
			)
		}
	}
	if err := recordStorage.BatchSet(ctx, records); err != nil {
		return errors.Wrap(err, "set method failed")
	}
	return nil
}
