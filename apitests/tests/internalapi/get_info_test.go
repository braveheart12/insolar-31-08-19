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
package internalapi

import (
	"net/http"
	"testing"

	"github.com/insolar/insolar/apitests/apiclient/insolar_internal_api_new"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/insolar/insolar/apitests/apihelper/apilogger"
	"github.com/insolar/insolar/apitests/tests"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func TestGetInfoWithBadMethod(t *testing.T) {
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
		r := insolar_internal_api.NetworkGetInfoRequest{
			Jsonrpc: JSONRPCVersion,
			Id:      apihelper.GetRequestId(),
			Method:  NetworkGetInfo,
			Params:  nil,
		}
		getInfoWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetInfoWithBadJsonVersion(t *testing.T) {
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
		r := insolar_internal_api.NetworkGetInfoRequest{
			Jsonrpc: tc.Input,
			Id:      apihelper.GetRequestId(),
			Method:  NetworkGetInfo,
		}
		getInfoWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetInfoWithBadRequestId(t *testing.T) {
	data := []int32{0, -1, -2147483648, 2147483647}

	for _, v := range data {
		r := insolar_internal_api.NetworkGetInfoRequest{
			Jsonrpc: JSONRPCVersion,
			Id:      v,
			Method:  NetworkGetInfo,
		}
		getInfoRequest(t, r)
	}
}

func TestGetInfoWithTwoRequestId(t *testing.T) {
	r := insolar_internal_api.NetworkGetInfoRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      1,
		Method:  NetworkGetInfo,
	}
	getInfoRequest(t, r)
	getInfoRequest(t, r)
}
func getInfoWithBadRequest(t *testing.T, r insolar_internal_api.NetworkGetInfoRequest, error tests.TestError) {
	response, http := loggingGetInfoRequest(t, r)
	require.Equal(t, 200, http.StatusCode)
	require.Equal(t, error.Message, response.Error.Message)
	require.Equal(t, int32(error.Code), response.Error.Code)
	require.Empty(t, response.Result)
}
func getInfoRequest(t *testing.T, r insolar_internal_api.NetworkGetInfoRequest) insolar_internal_api.NetworkGetInfoResponse200Result {
	response, http := loggingGetInfoRequest(t, r)
	apihelper.CheckResponseHasNoError(t, response)
	require.Equal(t, 200, http.StatusCode)
	require.NotEmpty(t, response.Result.TraceID)
	require.NotEmpty(t, response.Result.RootMember)
	require.NotEmpty(t, response.Result.RootDomain)
	require.NotEmpty(t, response.Result.NodeDomain)
	require.Empty(t, response.Error)
	return response.Result
}

func loggingGetInfoRequest(t *testing.T, r insolar_internal_api.NetworkGetInfoRequest) (insolar_internal_api.NetworkGetInfoResponse200, *http.Response) {
	apilogger.LogApiRequest(r.Method, r, nil)
	response, http, _ := internalInformationApi.GetInfo(nil, r)
	apilogger.LogApiResponse(http, response)
	return response, http
}
