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
///
// +build apitests

package insolarapi

import (
	"sync"
	"testing"

	"github.com/insolar/insolar/apitests/tests"
	"github.com/stretchr/testify/require"
)

func TestMemberMinusTransfer(t *testing.T) {
	member1 := CreateMember(t)
	member2 := CreateMember(t)
	input := []string{"-1", //INS-2183
		"+0",
		"0", //INS-2184
		"-0"}
	for _, v := range input {
		transfer := member1.Transfer(t, member2.MemberReference, v)
		expErr := tests.TestError{-32000, "[ makeCall ] Error in called method: amount must be larger then zero"}
		require.Equal(t, expErr.Code, transfer.Error.Code)
		require.Equal(t, expErr.Message, transfer.Error.Message)
	}
}

func TestMemberTransferByIncorrectRef(t *testing.T) {
	member1 := CreateMember(t)
	transfer := member1.Transfer(t, "5gFY3nZ5uDPCCU2MwQbFSQ17XA2b1eUo9xp3p8AkdAB.11111111111111111111111111111111", "100")
	error := tests.TestError{-32000, "member not found"} //https://insolar.atlassian.net/browse/INS-3309
	require.Equal(t, error.Code, transfer.Error.Code)
	require.Equal(t, error.Message, transfer.Error.Message)
}

// The identifier is to be incremented in every request and each response will contain a corresponding one.
// The transfer request sends an amount of funds to member identified by a reference:
func TestCreateTransferGetBalance(t *testing.T) {
	var Wg sync.WaitGroup

	member1 := CreateMember(t)
	member2 := CreateMember(t)

	count := 5
	Wg.Add(count * 2)

	for i := 0; i < count; i++ {
		go member1.Transfer(t, member2.MemberReference, "1")
		go member2.Transfer(t, member1.MemberReference, "1")
	}
	Wg.Wait()
}
