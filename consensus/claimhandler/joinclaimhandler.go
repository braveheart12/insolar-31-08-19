/*
 *    Copyright 2019 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package claimhandler

import (
	"github.com/insolar/insolar/consensus/packets"
	"github.com/insolar/insolar/core"
)

type joinClaimHandler struct {
	next        ClaimHandler
	queue       Queue
	ref         core.RecordRef
	activeCount int
}

func NewJoinClaimHandler(activeNodesCount int, claims []*packets.NodeJoinClaim, pulse *core.Pulse, next ClaimHandler) ClaimHandler {
	handler := &joinClaimHandler{activeCount: activeNodesCount, next: next}
	for _, claim := range claims {
		handler.queue.PushClaim(claim, getPriority(claim.NodeRef, pulse.Entropy))
	}
	return handler
}

func (jch *joinClaimHandler) HandleClaim(claim packets.ReferendumClaim) packets.ReferendumClaim {
	_, ok := claim.(*packets.NodeJoinClaim)
	if !ok {
		if jch.next == nil {
			return claim
		}
		jch.next.HandleClaim(claim)
	}
	return jch.handle(claim)
}

func (jch *joinClaimHandler) handle(claim packets.ReferendumClaim) packets.ReferendumClaim {
	return jch.queue.PopClaim()
}

func getPriority(ref core.RecordRef, entropy core.Entropy) []byte {
	res := make([]byte, len(ref))
	for i := 0; i < len(ref); i++ {
		res[i] = ref[i] ^ entropy[i]
	}
	return res
}
