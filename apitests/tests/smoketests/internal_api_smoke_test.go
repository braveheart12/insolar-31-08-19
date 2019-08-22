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
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestMigrationAddAddresses(t *testing.T) {
	uuids, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	response := apihelper.AddMigrationAddresses(t, []string{uuids.String()})
	require.NotEmpty(t, response.Result)
	require.Empty(t, response.Error)
}

func TestMigrationDeposit(t *testing.T) {
	response := apihelper.MigrationDeposit(t)
	require.NotEmpty(t, response.Result)
	require.Empty(t, response.Error)
}

func TestObserverGetToken(t *testing.T) {
	response := apihelper.ObserverToken(t) //not worked
	require.NotEmpty(t, response)
}

func TestMemberGetBalance(t *testing.T) {
	member := apihelper.CreateMember(t)
	response := apihelper.GetBalance(t, member)
	require.NotEmpty(t, response.Result.CallResult.Deposits)
	require.NotEmpty(t, response.Result.CallResult.Balance)
}

/* "code": 217,
   "message": "[ makeCall ] Error in called method: unknown method: 'member.getBalance'"*/

func TestMigrationDeactivateDaemon(t *testing.T) {
	response := apihelper.MigrationDeactivateDaemon(t, "")
	require.NotEmpty(t, response.Result)
	require.Empty(t, response.Error)
}

func TestMigrationActivateDaemon(t *testing.T) {
	response := apihelper.MigrationActivateDaemon(t, "")
	require.NotEmpty(t, response.Result)
	require.Empty(t, response.Error)
}

func TestGetStatus(t *testing.T) {
	response := apihelper.GetStatus(t)
	require.Equal(t, "CompleteNetworkState", response.NetworkState)
	require.NotEmpty(t, response.ActiveListSize)
	require.NotEmpty(t, response.Entropy)
	for _, v := range response.Nodes {
		require.Equal(t, true, v.IsWorking)
	}
	require.Equal(t, true, response.Origin.IsWorking)
	require.NotEmpty(t, response.PulseNumber)
	require.NotEmpty(t, response.Version)
}

func TestGetInfo(t *testing.T) {
	response := apihelper.GetInfo(t)
	require.NotEmpty(t, response.RootDomain)
	require.NotEmpty(t, response.RootMember)
	require.NotEmpty(t, response.NodeDomain)
	require.NotEmpty(t, response.TraceID)
}
