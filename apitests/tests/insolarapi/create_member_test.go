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
	"testing"
	"time"

	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/insolar/insolar/apitests/apihelper/apilogger"
	"github.com/insolar/insolar/apitests/tests"
	"github.com/insolar/insolar/testutils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

var err error

func TestCreateMemberWithBadSeed(t *testing.T) {
	randomString := testutils.RandomString()
	uuids, error := uuid.NewV4()
	require.Nil(t, error)
	data := []tests.Cases{
		{"", tests.TestError{-32000, "[ checkSeed ] Bad seed param"}},
		{" ", tests.TestError{-32000, "[ checkSeed ] Failed to decode seed from string"}},
		{"node.getInfo", tests.TestError{-32000, "[ checkSeed ] Failed to decode seed from string"}},
		{randomString, tests.TestError{-32000, "[ checkSeed ] Failed to decode seed from string"}},
		{"1111", tests.TestError{-32000, "[ checkSeed ] Bad seed param"}},
		{uuids.String(), tests.TestError{-32000, "[ checkSeed ] Failed to decode seed from string"}},
	}
	for _, tc := range data {
		ms, _ := apihelper.NewMemberSignature()
		request := insolar_api.MemberCreateRequest{
			Jsonrpc: JSONRPCVersion,
			Id:      apihelper.GetRequestId(),
			Method:  ContractCall,
			Params: insolar_api.MemberCreateRequestParams{
				Seed:      tc.Input,
				CallSite:  MemberCreateMethod,
				PublicKey: string(ms.PemPublicKey),
			},
		}
		responseErrCreateMember(request, ms, t, tc.ExpectedError)
	}
}

func TestCreateMemberWithOldSeed(t *testing.T) {
	seed := GetSeed(t)

	time.Sleep(5 * time.Second)

	ms, _ := apihelper.NewMemberSignature()
	request := insolar_api.MemberCreateRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      apihelper.GetRequestId(),
		Method:  ContractCall,
		Params: insolar_api.MemberCreateRequestParams{
			Seed:      seed,
			CallSite:  MemberCreateMethod,
			PublicKey: string(ms.PemPublicKey),
		},
	}
	responseErrCreateMember(request, ms, t, tests.TestError{-32000, "[ checkSeed ] Incorrect seed"})
}

func TestCreateMemberWithEmptyPK(t *testing.T) {
	ms, _ := apihelper.NewMemberSignature()
	request := insolar_api.MemberCreateRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      apihelper.GetRequestId(),
		Method:  ContractCall,
		Params: insolar_api.MemberCreateRequestParams{
			Seed:      GetSeed(t),
			CallSite:  MemberCreateMethod,
			PublicKey: "",
		},
	}
	responseErrCreateMember(request, ms, t, tests.TestError{-32000, "[ makeCall ] Error in called method: error while verify signature: problems with decoding. Key - "})
}

func TestCreateMemberWithOtherPK(t *testing.T) {
	ms, _ := apihelper.NewMemberSignature()
	ms2, _ := apihelper.NewMemberSignature()
	request := insolar_api.MemberCreateRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      apihelper.GetRequestId(),
		Method:  ContractCall,
		Params: insolar_api.MemberCreateRequestParams{
			Seed:      GetSeed(t),
			CallSite:  MemberCreateMethod,
			PublicKey: string(ms2.PemPublicKey),
		},
	}
	responseErrCreateMember(request, ms, t, tests.TestError{-32000, "[ makeCall ] Error in called method: error while verify signature: invalid signature"})
}

func responseErrCreateMember(request insolar_api.MemberCreateRequest, ms apihelper.MemberSignature, t *testing.T, error tests.TestError) {
	response := logCreateMember(request, ms, t)
	require.NotEmpty(t, response.Error, "Error")
	require.Equal(t, error.Code, response.Error.Code, "ErrorCode")
	require.Equal(t, error.Message, response.Error.Message, "ErrorMessage")
}

func logCreateMember(request insolar_api.MemberCreateRequest, ms apihelper.MemberSignature, t *testing.T) insolar_api.MemberCreateResponse200 {
	d, s, m := apihelper.Sign(request, ms.PrivateKey)
	apilogger.LogApiRequest(MemberCreateMethod, request, m)
	response, http, err := GetClient().MemberApi.MemberCreate(nil, d, s, request)
	require.Nil(t, err)
	apilogger.LogApiResponse(http, response)
	return response
}
