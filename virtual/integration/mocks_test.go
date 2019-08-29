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

package integration

import (
	"context"
	"crypto"
	"crypto/rand"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/network"
	networknode "github.com/insolar/insolar/network/node"
)

var (
	light = nodeMock{
		ref:     gen.Reference(),
		shortID: 1,
		role:    insolar.StaticRoleLightMaterial,
	}
	virtual = nodeMock{
		ref:     gen.Reference(),
		shortID: 3,
		role:    insolar.StaticRoleVirtual,
	}
)

type nodeMock struct {
	ref     insolar.Reference
	shortID insolar.ShortNodeID
	role    insolar.StaticRole
}

func (n *nodeMock) ID() insolar.Reference {
	return n.ref
}

func (n *nodeMock) ShortID() insolar.ShortNodeID {
	return n.shortID
}

func (n *nodeMock) Role() insolar.StaticRole {
	return n.role
}

func (n *nodeMock) PublicKey() crypto.PublicKey {
	panic("implement me")
}

func (n *nodeMock) Address() string {
	return ""
}

func (n *nodeMock) GetGlobuleID() insolar.GlobuleID {
	panic("implement me")
}

func (n *nodeMock) Version() string {
	panic("implement me")
}

func (n *nodeMock) LeavingETA() insolar.PulseNumber {
	panic("implement me")
}

func (n *nodeMock) GetState() insolar.NodeState {
	return insolar.NodeReady
}

func (n *nodeMock) GetPower() insolar.Power {
	return 1
}

type nodeNetMock struct {
	me insolar.NetworkNode
}

func (n *nodeNetMock) GetAccessor(insolar.PulseNumber) network.Accessor {
	return networknode.NewAccessor(networknode.NewSnapshot(insolar.GenesisPulse.PulseNumber, []insolar.NetworkNode{&virtual, &light}))
}

func newNodeNetMock(me insolar.NetworkNode) *nodeNetMock {
	return &nodeNetMock{me: me}
}

func (n *nodeNetMock) GetOrigin() insolar.NetworkNode {
	return n.me
}

func SendMessage(
	ctx context.Context, s *Server, msg payload.Payload,
) payload.Payload {
	reps, done := s.SendToSelf(ctx, msg)
	defer done()

	rep, ok := <-reps
	if !ok {
		panic("no reply")
	}
	pl, err := payload.UnmarshalFromMeta(rep.Payload)
	checkError(ctx, err, "unmarshal payload")
	return pl
}

func MakeSetIncomingRequest(
	objectID insolar.ID,
	reasonID insolar.ID,
	reasonObjectID insolar.ID,
	noWait bool,
	isAPI bool,
) (payload.SetIncomingRequest, record.Virtual) {
	args := make([]byte, 100)
	_, err := rand.Read(args)
	panicIfErr(err)

	req := record.IncomingRequest{
		Arguments: args,
		Reason:    *insolar.NewReference(reasonID),
	}

	if noWait {
		req.ReturnMode = record.ReturnNoWait
	} else {
		req.Object = insolar.NewReference(objectID)
	}

	if isAPI {
		req.APINode = gen.Reference()
	} else {
		req.Caller = *insolar.NewReference(reasonObjectID)
	}

	rec := record.Wrap(&req)
	pl := payload.SetIncomingRequest{
		Request: rec,
	}
	return pl, rec
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
