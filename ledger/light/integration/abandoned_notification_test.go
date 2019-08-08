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

package integration_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/insolar/pulse"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/drop"
	"github.com/stretchr/testify/require"
)

func HeavyResponse(pl payload.Payload) []payload.Payload {
	switch p := pl.(type) {
	case *payload.Replication, *payload.GotHotConfirmation:
		return nil
	case *payload.GetFilament:
		virtual := record.Wrap(&record.PendingFilament{
			RecordID:       p.ObjectID,
			PreviousRecord: nil,
		})

		return []payload.Payload{&payload.FilamentSegment{
			ObjectID: p.ObjectID,
			Records: []record.CompositeFilamentRecord{
				{
					RecordID: p.ObjectID,
					MetaID:   p.StartFrom,
					Meta:     record.Material{Virtual: virtual},
					Record:   record.Material{Virtual: record.Wrap(&record.IncomingRequest{})},
				},
			},
		}}
	case *payload.GetLightInitialState:
		return []payload.Payload{&payload.LightInitialState{
			NetworkStart: true,
			JetIDs:       []insolar.JetID{insolar.ZeroJetID},
			Pulse: pulse.PulseProto{
				PulseNumber: insolar.FirstPulseNumber,
			},
			Drops: [][]byte{
				drop.MustEncode(&drop.Drop{JetID: insolar.ZeroJetID, Pulse: insolar.FirstPulseNumber}),
			},
		}}
	}

	panic(fmt.Sprintf("unexpected message to heavy %T", pl))
}

func Test_AbandonedNotification(t *testing.T) {
	t.Parallel()

	ctx := inslogger.TestContext(t)
	cfg := DefaultLightConfig()
	cfg.Ledger.LightChainLimit = 5
	cfg.Bus.ReplyTimeout = 3 * time.Minute

	received := make(chan payload.AbandonedRequestsNotification)
	receivedConfirmations := make(chan payload.GotHotConfirmation)
	s, err := NewServer(ctx, cfg, func(meta payload.Meta, pl payload.Payload) []payload.Payload {
		if notification, ok := pl.(*payload.AbandonedRequestsNotification); ok {
			received <- *notification
		}
		if confirmation, ok := pl.(*payload.GotHotConfirmation); ok {
			receivedConfirmations <- *confirmation
		}
		if meta.Receiver == NodeHeavy() {
			return HeavyResponse(pl)
		}
		return nil
	})
	require.NoError(t, err)
	defer s.Stop()

	// First pulse goes in storage then interrupts.
	s.SetPulse(ctx)
	// Second pulse goes in storage and starts processing, including pulse change in flow dispatcher.
	s.SetPulse(ctx)

	t.Run("abandoned notification", func(t *testing.T) {
		msg, _ := MakeSetIncomingRequest(gen.ID(), gen.IDWithPulse(s.Pulse()), true, true)
		rep := SendMessage(ctx, s, &msg)
		RequireNotError(rep)
		reqInfo := rep.(*payload.RequestInfo)
		objectID := reqInfo.ObjectID

		<-receivedConfirmations
		s.SetPulse(ctx)
		<-receivedConfirmations
		s.SetPulse(ctx)
		<-receivedConfirmations

		for i := 0; i < 100; i++ {
			s.SetPulse(ctx)
			<-receivedConfirmations

			notification := <-received
			require.Equal(t, objectID, notification.ObjectID)
		}

		requestID := reqInfo.RequestID

		// Set result.
		resMsg, _ := MakeSetResult(objectID, requestID)
		rep = SendMessage(ctx, s, &resMsg)
		RequireNotError(rep)

		for j := 0; j < 10; j++ {
			s.SetPulse(ctx)
			<-receivedConfirmations

			select {
			case _ = <-received:
				t.Error("unexpected abandoned notifications reply")
			default:
				// Do nothing. It's ok.
			}
		}
	})
}
