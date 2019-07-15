//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package core

import (
	"fmt"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/network/consensus/common/cryptkit"
	"github.com/insolar/insolar/network/consensus/common/endpoints"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/member"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/phases"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/profiles"
	"testing"

	"github.com/insolar/insolar/network/consensus/gcpv2/core/errors"

	gcommon "github.com/insolar/insolar/network/consensus/gcpv2/common"

	"github.com/insolar/insolar/network/consensus/gcpv2/packets"
	"github.com/stretchr/testify/require"
)

func TestNewNodeAppearanceAsSelf(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.Equal(t, profiles.NodeStateLocalActive, r.state)

	require.Equal(t, member.SelfTrust, r.trust)

	require.Equal(t, lp, r.profile)

	require.Equal(t, callback, r.callback)

}

func TestInit(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.Panics(t, func() { r.init(nil, callback, 0) })

	r.init(lp, callback, 0)
	require.Equal(t, profiles.NodeStateLocalActive, r.state)

	require.Equal(t, member.SelfTrust, r.trust)

	require.Equal(t, lp, r.profile)

	require.Equal(t, callback, r.callback)
}

func TestString(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.Equal(t, fmt.Sprintf("node:{%v}", lp), r.String())
}

func TestLessByNeighbourWeightForNodeAppearance(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r1 := NewNodeAppearanceAsSelf(lp, callback)
	r2 := NewNodeAppearanceAsSelf(lp, callback)
	r1.neighbourWeight = 0
	r2.neighbourWeight = 1
	require.True(t, LessByNeighbourWeightForNodeAppearance(r1, r2))

	require.False(t, LessByNeighbourWeightForNodeAppearance(r2, r1))

	r2.neighbourWeight = 0
	require.False(t, LessByNeighbourWeightForNodeAppearance(r2, r1))
}

func TestCopySelfTo(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}

	source := NewNodeAppearanceAsSelf(lp, callback)
	source.stateEvidence = profiles.NewNodeStateHashEvidenceMock(t)
	source.announceSignature = gcommon.NewMemberAnnouncementSignatureMock(t)
	source.requestedPower = 1
	source.state = profiles.NodeStateLocalActive
	source.trust = member.TrustBySome

	target := NewNodeAppearanceAsSelf(lp, callback)
	target.stateEvidence = profiles.NewNodeStateHashEvidenceMock(t)
	target.announceSignature = gcommon.NewMemberAnnouncementSignatureMock(t)
	target.requestedPower = 2
	target.state = profiles.NodeStateReceivedPhases
	target.trust = member.TrustByNeighbors

	target.copySelfTo(source)

	require.Equal(t, target.stateEvidence, source.stateEvidence)

	require.Equal(t, target.announceSignature, source.announceSignature)

	require.Equal(t, target.requestedPower, source.requestedPower)

	require.Equal(t, target.state, source.state)

	require.Equal(t, target.trust, source.trust)
}

func TestIsJoiner(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	lp.IsJoinerMock.Set(func() (r bool) {
		return true
	})
	callback := &nodeContext{}

	r := NewNodeAppearanceAsSelf(lp, callback)
	require.True(t, r.IsJoiner())
}

func TestGetIndex(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	index := 1
	lp.GetIndexMock.Set(func() int { return index })
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.Equal(t, index, r.GetIndex())
}

func TestGetShortNodeID(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	lp.GetShortNodeIDMock.Set(func() insolar.ShortNodeID { return insolar.AbsentShortNodeID })
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.Equal(t, insolar.AbsentShortNodeID, r.GetShortNodeID())
}

func TestGetTrustLevel(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	r.trust = member.TrustBySome
	require.Equal(t, member.TrustBySome, r.GetTrustLevel())
}

func TestGetProfile(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.Equal(t, lp, r.GetProfile())
}

func TestVerifyPacketAuthenticity(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	var isAcceptable bool
	lp.IsAcceptableHostMock.Set(func(endpoints.Inbound) bool { return *(&isAcceptable) })
	sv := cryptkit.NewSignatureVerifierMock(t)
	var isSignOfSignatureMethodSupported bool
	sv.IsSignOfSignatureMethodSupportedMock.Set(func(cryptkit.SignatureMethod) bool { return *(&isSignOfSignatureMethodSupported) })
	var isValidDigestSignature bool
	sv.IsValidDigestSignatureMock.Set(func(cryptkit.DigestHolder, cryptkit.SignatureHolder) bool {
		return *(&isValidDigestSignature)
	})
	lp.GetSignatureVerifierMock.Set(func() cryptkit.SignatureVerifier { return sv })
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	packet := packets.NewPacketParserMock(t)
	packet.GetPacketSignatureMock.Set(func() cryptkit.SignedDigest { return cryptkit.SignedDigest{} })
	from := endpoints.NewHostIdentityHolderMock(t)
	strictFrom := true
	isAcceptable = false
	require.NotEqual(t, nil, r.VerifyPacketAuthenticity(packet, from, strictFrom))

	strictFrom = false
	isSignOfSignatureMethodSupported = false
	require.NotEqual(t, nil, r.VerifyPacketAuthenticity(packet, from, strictFrom))

	isSignOfSignatureMethodSupported = true
	isValidDigestSignature = false
	require.NotEqual(t, nil, r.VerifyPacketAuthenticity(packet, from, strictFrom))

	isValidDigestSignature = true
	require.Equal(t, nil, r.VerifyPacketAuthenticity(packet, from, strictFrom))
}

func TestSetReceivedPhase(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.True(t, r.SetReceivedPhase(phases.Phase1))

	require.False(t, r.SetReceivedPhase(phases.Phase1))
}

func TestSetReceivedByPacketType(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.True(t, r.SetPacketReceived(phases.PacketPhase1))

	require.False(t, r.SetPacketReceived(phases.PacketPhase1))

	require.False(t, r.SetPacketReceived(phases.PacketTypeCount))
}

func TestSetSentPhase(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.True(t, r.SetSentPhase(phases.Phase1))

	require.False(t, r.SetSentPhase(phases.Phase1))
}

func TestSetSentByPacketType(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.True(t, r.SetPacketSent(phases.PacketPhase1))

	require.True(t, r.SetPacketSent(phases.PacketPhase1))

	require.False(t, r.SetPacketSent(phases.PacketTypeCount))
}

func TestSetReceivedWithDupCheck(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	require.Equal(t, r.SetPacketReceivedWithDupError(phases.PacketPhase1), nil)

	require.Equal(t, r.SetPacketReceivedWithDupError(phases.PacketPhase1), errors.ErrPacketLimitExceeded)

	require.Equal(t, r.SetPacketReceivedWithDupError(phases.PacketTypeCount), errors.ErrPacketLimitExceeded)
}

func TestGetSignatureVerifier(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	sv1 := cryptkit.NewSignatureVerifierMock(t)
	lp.GetSignatureVerifierMock.Set(func() cryptkit.SignatureVerifier { return sv1 })
	lp.GetNodePublicKeyStoreMock.Set(func() cryptkit.PublicKeyStore { return nil })
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)
	svf := cryptkit.NewSignatureVerifierFactoryMock(t)
	sv2 := cryptkit.NewSignatureVerifierMock(t)
	svf.GetSignatureVerifierWithPKSMock.Set(func(cryptkit.PublicKeyStore) cryptkit.SignatureVerifier { return sv2 })
	require.Equal(t, r.GetSignatureVerifier(svf), sv1)

	lp.GetSignatureVerifierMock.Set(func() cryptkit.SignatureVerifier { return nil })
	require.Equal(t, sv2, r.GetSignatureVerifier(svf))
}

func TestCreateSignatureVerifier(t *testing.T) {
	lp := gcommon.NewLocalNodeProfileMock(t)
	lp.LocalNodeProfileMock.Set(func() {})
	lp.GetNodePublicKeyStoreMock.Set(func() cryptkit.PublicKeyStore { return nil })
	callback := &nodeContext{}
	r := NewNodeAppearanceAsSelf(lp, callback)

	svf := cryptkit.NewSignatureVerifierFactoryMock(t)
	sv := cryptkit.NewSignatureVerifierMock(t)
	svf.GetSignatureVerifierWithPKSMock.Set(func(cryptkit.PublicKeyStore) cryptkit.SignatureVerifier { return sv })
	require.Equal(t, sv, r.CreateSignatureVerifier(svf))
}
