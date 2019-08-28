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
	"github.com/insolar/insolar/apitests/apihelper/apilogger"
	"testing"
	"time"

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

	time.Sleep(5000)
	get := apihelper.Member(t, member.MemberReference)
	require.Empty(t, get.Error)
	require.NotEmpty(t, get.Balance)
	require.NotEmpty(t, get.MigrationAddress)

	balance := apihelper.Balance(t, member.MemberReference)
	require.Empty(t, balance.Error)
	require.NotEmpty(t, balance.Balance)
}

func TestMember(t *testing.T) {
	member := apihelper.CreateMember(t)
	require.NotEmpty(t, member.MemberReference, "MemberReference")
	response := apihelper.Member(t, member.MemberReference)
	require.Empty(t, response.Error)
	require.NotEmpty(t, response.Balance)
	require.NotEmpty(t, response.MigrationAddress)
	//require.NotEmpty(t, response.Deposits)
}

func TestTransaction(t *testing.T) {
	amount := "1"
	member1 := apihelper.CreateMember(t)
	member2 := apihelper.CreateMember(t)
	transfer := member1.Transfer(t, member2.MemberReference, amount)
	apihelper.CheckResponseHasNoError(t, transfer)
	apilogger.Println("Transfer OK. Fee: " + transfer.Result.CallResult.Fee)
	require.NotEmpty(t, transfer.Result.CallResult.Fee, "Fee")

	time.Sleep(60 * time.Second)
	response := apihelper.Transaction(t, transfer.Result.RequestReference)

	require.Equal(t, amount, response.Amount)
	require.Equal(t, transfer.Result.CallResult.Fee, response.Fee)
	require.Equal(t, member1.MemberReference, response.FromMemberReference)
	require.Equal(t, member2.MemberReference, response.ToMemberReference)
	require.Equal(t, "SUCCESS", response.Status)

	require.Empty(t, response.Error)
}

func TestTransactionList(t *testing.T) {
	member1 := apihelper.CreateMember(t)
	member2 := apihelper.CreateMember(t)
	transfer := member1.Transfer(t, member2.MemberReference, "1")
	apihelper.CheckResponseHasNoError(t, transfer)
	apilogger.Println("Transfer OK. Fee: " + transfer.Result.CallResult.Fee)
	require.NotEmpty(t, transfer.Result.CallResult.Fee, "Fee")

	transactions := apihelper.TransactionList(t, member1.MemberReference)
	require.NotEmpty(t, transactions)
}
