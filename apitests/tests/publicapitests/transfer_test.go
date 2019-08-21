////////////////////////////////////////////////////////////////////////////////
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
////////////////////////////////////////////////////////////////////////////////

package publicapitests

import (
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemberMinusTransfer(t *testing.T) {
	member1 := apihelper.CreateMember(t)
	member2 := apihelper.CreateMember(t)
	//transfer := member1.Transfer(t, member2.MemberReference, "-1")//INS-2183
	transfer := member1.Transfer(t, member2.MemberReference, "0") //INS-2184
	require.Equal(t, transfer.Error.Code, -32000)
	require.Equal(t, transfer.Error.Message, "[ makeCall ] Error in called method: amount must be larger then zero")
}

func TestMemberTransferToBadMember(t *testing.T) {
	member1 := apihelper.CreateMember(t)
	transfer := member1.Transfer(t, "5gFY3nZ5uDPCCU2MwQbFSQ17XA2b1eUo9xp3p8AkdAB.11111111111111111111111111111111", "100") //INS-2185
	require.Equal(t, transfer.Error.Code, -32000)
	require.Equal(t, transfer.Error.Message, "[ makeCall ] Error in called method: amount must be larger then zero")

}
