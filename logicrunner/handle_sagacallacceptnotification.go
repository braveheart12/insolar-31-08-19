package logicrunner

import (
	"context"

	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/flow/bus"
	"github.com/insolar/insolar/insolar/reply"
)

type HandleSagaCallAcceptNotification struct {
	dep *Dependencies

	Message bus.Message
}

func (h *HandleSagaCallAcceptNotification) Present(ctx context.Context, f flow.Flow) error {
	h.Message.ReplyTo <- bus.Reply{Reply: &reply.OK{}, Err: nil}
	return nil

	// AALEKSEEV TODO implement
	// msg, ok := parcel.Message().(*message.AbandonedRequestsNotification)

}
