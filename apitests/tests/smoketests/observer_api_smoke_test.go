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
// +build smoke

package smoketests

import (
	"testing"

	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
)

func TestNotification(t *testing.T) {
	response := apihelper.Notification(t)
	require.NotEmpty(t, response.Notification)
}

func TestBalance(t *testing.T) {
	member := apihelper.CreateMember(t)
	require.NotEmpty(t, member.MemberReference, "MemberReference")

	member.GetMember(t)

	response := apihelper.Balance(t, member.MemberReference)
	require.Empty(t, response.Error)
	require.NotEmpty(t, response.Balance)
}

func TestMember(t *testing.T) {
	member := apihelper.CreateMember(t)
	require.NotEmpty(t, member.MemberReference, "MemberReference")
	response := apihelper.Member(t, member.MemberReference)
	require.Empty(t, response.Error)
	require.NotEmpty(t, response.Balance)
	require.NotEmpty(t, response.MigrationAddress)
	require.NotEmpty(t, response.Deposits)
}

func TestTransaction(t *testing.T) {
	response := apihelper.Transaction(t, "")
	require.NotEmpty(t, response.Amount)
	require.NotEmpty(t, response.Fee)
	require.NotEmpty(t, response.FromMemberReference)
	require.Empty(t, response.Error)
}

func TestTransactionList(t *testing.T) {
	response := apihelper.TransactionList(t, "")
	require.NotEmpty(t, response)
}
