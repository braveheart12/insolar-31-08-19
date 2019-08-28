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

package insolarapi

import (
	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/insolar/insolar/apitests/apihelper/apilogger"
	"github.com/insolar/insolar/apitests/tests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetNotExistMember(t *testing.T) {
	ms, _ := apihelper.NewMemberSignature()
	seed := GetSeed(t)
	request := insolar_api.MemberGetRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      apihelper.GetRequestId(),
		Method:  ContractCall,
		Params: insolar_api.MemberGetRequestParams{
			Seed:       seed,
			CallSite:   MemberGetMethod,
			CallParams: nil,
			PublicKey:  string(ms.PemPublicKey),
		},
	}
	d, s, m := apihelper.Sign(request, ms.PrivateKey)
	apilogger.LogApiRequest(request.Params.CallSite, request, m)
	response, http, err := GetClient().MemberApi.MemberGet(nil, d, s, request)
	apilogger.LogApiResponse(http, response)
	require.Nil(t, err)
	error := tests.TestError{-32000, "[ makeCall ] Error in called method: failed to get reference by public key: failed to get reference in shard: failed to find reference by key"}
	require.Equal(t, error.Code, response.Error.Code)
	require.Equal(t, error.Message, response.Error.Message)
}
