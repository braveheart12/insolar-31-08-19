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

package internalapi

import (
	"net/http"
	"testing"

	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/insolar/insolar/apitests/apihelper/apilogger"
	"github.com/insolar/insolar/apitests/tests"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func TestGetStatusWithBadMethod(t *testing.T) {
	randomString := testutils.RandomString()
	data := []tests.Cases{
		{"", tests.TestError{-32000, "rpc: service/method request ill-formed: \"\""}},
		{" ", tests.TestError{-32000, "rpc: service/method request ill-formed: \" \""}},
		{"node.getInfo", tests.TestError{-32000, "rpc: can't find method \"node.getInfo\""}},
		{randomString, tests.TestError{-32000, "rpc: service/method request ill-formed: \"" + randomString + "\""}},
		{"1111", tests.TestError{-32000, "rpc: service/method request ill-formed: \"1111\""}},
		{"node getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node getInfo\""}},
		{"node^getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node^getInfo\""}},
		{"node*getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node*getInfo\""}},
		{"node&getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node&getInfo\""}},
		{"node%getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node%getInfo\""}},
		{"getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"getInfo\""}},
	}
	for _, tc := range data {
		r := insolar_internal_api.NodeGetStatusRequest{
			Jsonrpc: JSONRPCVersion,
			Id:      apihelper.GetRequestId(),
			Method:  tc.Input,
		}
		getStatusWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetStatusWithBadJsonVersion(t *testing.T) {
	randomString := testutils.RandomString()
	data := []tests.Cases{
		{"1.0", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{"", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{" ", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{randomString, tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{"0", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{"-1", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{"3.0", tests.TestError{-32600, "jsonrpc must be 2.0"}},
	}
	for _, tc := range data {
		r := insolar_internal_api.NodeGetStatusRequest{
			Jsonrpc: tc.Input,
			Id:      apihelper.GetRequestId(),
			Method:  GetStatusMethod,
		}
		getStatusWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetStatusWithBadRequestId(t *testing.T) {
	data := []int32{0, -1, -2147483648, 2147483647}

	for _, v := range data {
		r := insolar_internal_api.NodeGetStatusRequest{
			Jsonrpc: JSONRPCVersion,
			Id:      v,
			Method:  GetStatusMethod,
		}
		GetStatusRequest(t, r)
	}
}

func TestGetStatusWithTwoRequestId(t *testing.T) {
	r := insolar_internal_api.NodeGetStatusRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      1,
		Method:  GetStatusMethod,
	}
	GetStatusRequest(t, r)
	GetStatusRequest(t, r)
}
func getStatusWithBadRequest(t *testing.T, r insolar_internal_api.NodeGetStatusRequest, error tests.TestError) {
	response, http := loggingGetStatusRequest(t, r)
	require.Equal(t, 200, http.StatusCode)
	require.Equal(t, error.Message, response.Error.Message)
	require.Equal(t, int32(error.Code), response.Error.Code)
	require.Empty(t, response.Result)
}

func GetStatusRequest(t *testing.T, r insolar_internal_api.NodeGetStatusRequest) insolar_internal_api.NodeGetStatusResponse200Result {
	response, http := loggingGetStatusRequest(t, r)
	require.Equal(t, 200, http.StatusCode)
	require.Equal(t, "CompleteNetworkState", response.Result.NetworkState)
	require.NotEmpty(t, response.Result.ActiveListSize)
	require.NotEmpty(t, response.Result.Entropy)
	for _, v := range response.Result.Nodes {
		require.Equal(t, true, v.IsWorking)
	}
	require.Equal(t, true, response.Result.Origin.IsWorking)
	require.NotEmpty(t, response.Result.PulseNumber)
	require.NotEmpty(t, response.Result.Version)
	require.Empty(t, response.Error)
	return response.Result
}

func loggingGetStatusRequest(t *testing.T, r insolar_internal_api.NodeGetStatusRequest) (insolar_internal_api.NodeGetStatusResponse200, *http.Response) {
	apilogger.LogApiRequest(r.Method, r, nil)
	response, http, err := GetInternalClient().InformationApi.GetStatus(nil, r)
	require.Nil(t, err)
	apilogger.LogApiResponse(http, response)
	return response, http
}
