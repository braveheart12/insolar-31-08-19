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

package smoketests

import (
	"testing"

	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
)

// Information api

// TODO method has been moved to observer
// func TestGetInfo(t *testing.T) {
// 	response := apihelper.GetInfo(t)
// 	require.NotEmpty(t, response.RootDomain)
// 	require.NotEmpty(t, response.RootMember)
// 	require.NotEmpty(t, response.NodeDomain)
// 	require.NotEmpty(t, response.TraceID)
// }

func TestGetSeed(t *testing.T) {
	seed := apihelper.GetSeed(t)
	require.NotEmpty(t, seed)
}

// Member api

func TestCreateMember(t *testing.T) {
	member := apihelper.CreateMember(t)
	require.NotEmpty(t, member.MemberReference, "MemberReference")
}

func TestMemberTransfer(t *testing.T) {
	member1 := apihelper.CreateMember(t)
	member2 := apihelper.CreateMember(t)
	transfer := member1.Transfer(t, member2.MemberReference, "1")
	require.NotEmpty(t, transfer.Result.CallResult.Fee, "Fee")
}

func TestGetMember(t *testing.T) {
	member1 := apihelper.CreateMember(t)
	resp := member1.GetMember(t)
	require.Equal(t, member1.MemberReference, resp.Result.CallResult.Reference, "Reference")
	require.Empty(t, resp.Result.CallResult.MigrationAddress, "MigrationAddress")
}

// Migration api

func TestDepositTransfer(t *testing.T) {
	response := apihelper.DepositTransfer(t)
	require.NotEmpty(t, response.Result.CallResult)
}

func TestMemberMigrationCreate(t *testing.T) {
	var member = apihelper.MemberMigrationCreate(t)
	require.NotEmpty(t, member)
	require.NotEmpty(t, member.MemberResponseResult)
}
