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

package apihelper

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
)

type MemberSignature struct {
	PublicKey     ecdsa.PublicKey
	PrivateKey    *ecdsa.PrivateKey
	X509PublicKey []byte
	PemPublicKey  []byte
}

type MemberObject struct {
	MemberReference      string
	Signature            MemberSignature
	MemberResponseResult insolar_api.MemberCreateResponse200
}

type ecdsaSignature struct {
	R, S *big.Int
}
