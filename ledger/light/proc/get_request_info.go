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

package proc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/bus"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/light/executor"
)

type GetRequestInfo struct {
	message   payload.Meta
	objectID  insolar.ID
	requestID insolar.ID

	dep struct {
		filament executor.FilamentCalculator
		sender   bus.Sender
	}
}

func NewGetRequestInfo(msg payload.Meta, objectID insolar.ID, requestID insolar.ID) *GetRequestInfo {
	return &GetRequestInfo{
		message:   msg,
		objectID:  objectID,
		requestID: requestID,
	}
}

func (p *GetRequestInfo) Dep(
	filament executor.FilamentCalculator,
	sender bus.Sender,
) {
	p.dep.filament = filament
	p.dep.sender = sender
}

func (p *GetRequestInfo) Proceed(ctx context.Context) error {

	if p.requestID.IsEmpty() {
		return errors.New("request id is empty")
	}

	objectID := p.objectID

	logger := inslogger.FromContext(ctx).WithFields(map[string]interface{}{
		"request_id": p.requestID.DebugString(),
		"object_id":  objectID.DebugString(),
	})

	// Searching for request
	{
		var (
			reqBuf []byte
			resBuf []byte
		)
		requestID := p.requestID
		foundRequest, foundResult, err := p.dep.filament.RequestInfo(ctx, objectID, requestID)
		if err != nil {
			return errors.Wrap(err, "failed to get request info")
		}
		if foundRequest != nil || foundResult != nil {
			if foundRequest != nil {
				reqBuf, err = foundRequest.Record.Marshal()
				if err != nil {
					return errors.Wrap(err, "failed to marshal request record")
				}
				requestID = foundRequest.RecordID
			}
			if foundResult != nil {
				resBuf, err = foundResult.Record.Marshal()
				if err != nil {
					return errors.Wrap(err, "failed to marshal result record")
				}
			}

			msg, err := payload.NewMessage(&payload.RequestInfo{
				ObjectID:  objectID,
				RequestID: requestID,
				Request:   reqBuf,
				Result:    resBuf,
			})
			if err != nil {
				return errors.Wrap(err, "failed to create reply")
			}
			p.dep.sender.Reply(ctx, p.message, msg)
			logger.WithFields(map[string]interface{}{
				"request":    foundRequest != nil,
				"has_result": foundResult != nil,
			}).Debug("result info found")
			return nil
		}
	}

	msg, err := payload.NewMessage(&payload.RequestInfo{
		ObjectID:  objectID,
		RequestID: p.requestID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create reply")
	}
	p.dep.sender.Reply(ctx, p.message, msg)
	return nil
}
