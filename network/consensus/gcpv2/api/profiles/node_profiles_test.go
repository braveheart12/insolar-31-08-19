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

package profiles

import (
	"testing"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/network/consensus/common/cryptkit"
	"github.com/insolar/insolar/network/consensus/common/endpoints"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/member"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/proofs"

	"github.com/stretchr/testify/require"
)

func TestNewMembershipProfile(t *testing.T) {
	nsh := proofs.NewNodeStateHashEvidenceMock(t)
	nas := proofs.NewMemberAnnouncementSignatureMock(t)
	index := uint16(1)
	power := member.Power(2)
	ep := member.Power(3)
	mp := NewMembershipProfile(member.ModeNormal, power, index, nsh, nas, ep)
	require.Equal(t, index, mp.Index)

	require.Equal(t, power, mp.Power)

	require.Equal(t, ep, mp.RequestedPower)

	require.Equal(t, nsh, mp.StateEvidence)

	require.Equal(t, nas, mp.AnnounceSignature)
}

func TestNewMembershipProfileByNode(t *testing.T) {
	np := NewActiveNodeMock(t)
	index := 1
	np.GetIndexMock.Set(func() int { return index })
	power := member.Power(2)
	np.GetDeclaredPowerMock.Set(func() member.Power { return power })
	np.GetOpModeMock.Set(func() (r member.OpMode) {
		return member.ModeNormal
	})

	nsh := proofs.NewNodeStateHashEvidenceMock(t)
	nas := proofs.NewMemberAnnouncementSignatureMock(t)
	ep := member.Power(3)
	mp := NewMembershipProfileByNode(np, nsh, nas, ep)
	require.Equal(t, uint16(index), mp.Index)

	require.Equal(t, power, mp.Power)

	require.Equal(t, ep, mp.RequestedPower)

	require.Equal(t, nsh, mp.StateEvidence)

	require.Equal(t, nas, mp.AnnounceSignature)
}

func TestIsEmpty(t *testing.T) {
	mp := MembershipProfile{}
	require.True(t, mp.IsEmpty())

	se := proofs.NewNodeStateHashEvidenceMock(t)
	mp.StateEvidence = se
	require.True(t, mp.IsEmpty())

	mp.StateEvidence = nil
	mp.AnnounceSignature = proofs.NewMemberAnnouncementSignatureMock(t)
	require.True(t, mp.IsEmpty())

	mp.StateEvidence = se
	require.False(t, mp.IsEmpty())
}

func TestEquals(t *testing.T) {
	mp1 := MembershipProfile{}
	mp2 := MembershipProfile{}
	require.False(t, mp1.Equals(mp2))

	mp1.Index = uint16(1)
	mp1.Power = member.Power(2)
	mp1.RequestedPower = member.Power(3)
	she1 := proofs.NewNodeStateHashEvidenceMock(t)
	mas1 := proofs.NewMemberAnnouncementSignatureMock(t)
	mp1.StateEvidence = she1
	mp1.AnnounceSignature = mas1

	mp2.Index = uint16(2)
	mp2.Power = mp1.Power
	mp2.RequestedPower = mp1.RequestedPower
	mp2.StateEvidence = mp1.StateEvidence
	mp2.AnnounceSignature = mp1.AnnounceSignature

	require.False(t, mp1.Equals(mp2))

	mp2.Index = mp1.Index
	mp2.Power = member.Power(3)
	require.False(t, mp1.Equals(mp2))

	mp2.Power = mp1.Power
	mp2.StateEvidence = nil
	require.False(t, mp1.Equals(mp2))

	mp2.StateEvidence = mp1.StateEvidence
	mp2.AnnounceSignature = nil
	require.False(t, mp1.Equals(mp2))

	mp2.AnnounceSignature = mp1.AnnounceSignature
	mp1.StateEvidence = nil
	require.False(t, mp1.Equals(mp2))

	mp1.StateEvidence = mp2.StateEvidence
	mp1.AnnounceSignature = nil
	require.False(t, mp1.Equals(mp2))

	mp1.AnnounceSignature = mp2.AnnounceSignature
	mp2.RequestedPower = member.Power(4)
	require.False(t, mp1.Equals(mp2))

	mp2.RequestedPower = mp1.RequestedPower
	she2 := proofs.NewNodeStateHashEvidenceMock(t)
	mp2.StateEvidence = she2
	nsh := proofs.NewNodeStateHashMock(t)
	she1.GetNodeStateHashMock.Set(func() proofs.NodeStateHash { return nsh })
	she2.GetNodeStateHashMock.Set(func() proofs.NodeStateHash { return nsh })
	nsh.EqualsMock.Set(func(cryptkit.DigestHolder) bool { return false })
	require.False(t, mp1.Equals(mp2))

	nsh.EqualsMock.Set(func(cryptkit.DigestHolder) bool { return true })
	sh := cryptkit.NewSignatureHolderMock(t)
	sh.EqualsMock.Set(func(cryptkit.SignatureHolder) bool { return false })
	she1.GetGlobulaNodeStateSignatureMock.Set(func() cryptkit.SignatureHolder { return sh })
	she2.GetGlobulaNodeStateSignatureMock.Set(func() cryptkit.SignatureHolder { return sh })
	require.False(t, mp1.Equals(mp2))

	sh.EqualsMock.Set(func(cryptkit.SignatureHolder) bool { return true })
	require.True(t, mp1.Equals(mp2))

	mp2.StateEvidence = she1
	mas2 := proofs.NewMemberAnnouncementSignatureMock(t)
	mp2.AnnounceSignature = mas2
	mas1.EqualsMock.Set(func(cryptkit.SignatureHolder) bool { return false })
	require.False(t, mp1.Equals(mp2))

	mas1.EqualsMock.Set(func(cryptkit.SignatureHolder) bool { return true })
	require.True(t, mp1.Equals(mp2))
}

func TestStringParts(t *testing.T) {
	mp := MembershipProfile{}
	require.True(t, len(mp.StringParts()) > 0)
	mp.Power = member.Power(1)
	require.True(t, len(mp.StringParts()) > 0)
}

func TestMembershipProfileString(t *testing.T) {
	mp := MembershipProfile{}
	require.True(t, len(mp.String()) > 0)
}

func TestEqualIntroProfiles(t *testing.T) {
	require.False(t, EqualIntroProfiles(nil, nil))
	p := NewNodeIntroProfileMock(t)
	require.False(t, EqualIntroProfiles(p, nil))

	require.False(t, EqualIntroProfiles(nil, p))

	require.True(t, EqualIntroProfiles(p, p))

	snID1 := insolar.ShortNodeID(1)
	p.GetShortNodeIDMock.Set(func() insolar.ShortNodeID { return *(&snID1) })
	primaryRole1 := member.PrimaryRoleNeutral
	p.GetPrimaryRoleMock.Set(func() member.PrimaryRole { return *(&primaryRole1) })
	specialRole1 := member.SpecialRoleDiscovery
	p.GetSpecialRolesMock.Set(func() member.SpecialRole { return *(&specialRole1) })
	power1 := member.Power(1)
	p.GetStartPowerMock.Set(func() member.Power { return *(&power1) })
	skh := cryptkit.NewSignatureKeyHolderMock(t)
	signHoldEq := true
	skh.EqualsMock.Set(func(cryptkit.SignatureKeyHolder) bool { return *(&signHoldEq) })
	p.GetNodePublicKeyMock.Set(func() cryptkit.SignatureKeyHolder { return skh })

	o := NewNodeIntroProfileMock(t)
	snID2 := insolar.ShortNodeID(2)
	o.GetShortNodeIDMock.Set(func() insolar.ShortNodeID { return *(&snID2) })
	primaryRole2 := primaryRole1
	o.GetPrimaryRoleMock.Set(func() member.PrimaryRole { return *(&primaryRole2) })
	specialRole2 := specialRole1
	o.GetSpecialRolesMock.Set(func() member.SpecialRole { return *(&specialRole2) })
	power2 := power1
	o.GetStartPowerMock.Set(func() member.Power { return *(&power2) })
	o.GetNodePublicKeyMock.Set(func() cryptkit.SignatureKeyHolder { return skh })
	require.False(t, EqualIntroProfiles(p, o))

	snID2 = snID1
	primaryRole2 = member.PrimaryRoleHeavyMaterial
	require.False(t, EqualIntroProfiles(p, o))

	primaryRole2 = primaryRole1
	specialRole2 = member.SpecialRoleNone
	require.False(t, EqualIntroProfiles(p, o))

	specialRole2 = specialRole1
	power2 = member.Power(2)
	require.False(t, EqualIntroProfiles(p, o))

	power1 = power2
	signHoldEq = false
	require.False(t, EqualIntroProfiles(p, o))

	signHoldEq = true
	ne1 := endpoints.NewOutboundMock(t)
	ne1.GetEndpointTypeMock.Set(func() endpoints.NodeEndpointType { return endpoints.NameEndpoint })
	ne1.GetNameAddressMock.Set(func() endpoints.Name { return endpoints.Name("test1") })
	p.GetDefaultEndpointMock.Set(func() endpoints.Outbound { return ne1 })
	ne2 := endpoints.NewOutboundMock(t)
	ne2.GetEndpointTypeMock.Set(func() endpoints.NodeEndpointType { return endpoints.NameEndpoint })
	ne2.GetNameAddressMock.Set(func() endpoints.Name { return endpoints.Name("test2") })
	o.GetDefaultEndpointMock.Set(func() endpoints.Outbound { return ne2 })
	require.False(t, EqualIntroProfiles(p, o))

	o.GetDefaultEndpointMock.Set(func() endpoints.Outbound { return ne1 })
	require.True(t, EqualIntroProfiles(p, o))
}
