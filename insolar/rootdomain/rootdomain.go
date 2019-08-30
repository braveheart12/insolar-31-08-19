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

package rootdomain

import (
	"sync"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/genesisrefs"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/logicrunner/builtin/proxy/member"
	"github.com/insolar/insolar/logicrunner/builtin/proxy/migrationshard"
	"github.com/insolar/insolar/logicrunner/builtin/proxy/pkshard"
	"github.com/insolar/insolar/platformpolicy"
)

const (
	GenesisPrototypeSuffix = "_proto"
)

func init() {
	for _, el := range insolar.GenesisNameMigrationDaemonMembers {
		genesisrefs.PredefinedPrototypes[el+GenesisPrototypeSuffix] = *member.PrototypeReference
	}

	for _, el := range insolar.GenesisNamePublicKeyShards {
		genesisrefs.PredefinedPrototypes[el+GenesisPrototypeSuffix] = *pkshard.PrototypeReference
	}

	for _, el := range insolar.GenesisNameMigrationAddressShards {
		genesisrefs.PredefinedPrototypes[el+GenesisPrototypeSuffix] = *migrationshard.PrototypeReference
	}
}

var genesisPulse = insolar.GenesisPulse.PulseNumber

// Record provides methods to calculate root domain's identifiers.
type Record struct {
	once                sync.Once
	rootDomainID        insolar.ID
	rootDomainReference insolar.Reference
	PCS                 insolar.PlatformCryptographyScheme
}

// RootDomain is the root domain instance.
var RootDomain = &Record{
	PCS: platformpolicy.NewPlatformCryptographyScheme(),
}

func (r *Record) initialize() {
	rootRecord := record.IncomingRequest{
		CallType: record.CTGenesis,
		Method:   insolar.GenesisNameRootDomain,
	}
	virtualRec := record.Wrap(&rootRecord)
	hash := record.HashVirtual(r.PCS.ReferenceHasher(), virtualRec)

	r.rootDomainID = *insolar.NewID(genesisPulse, hash)
	r.rootDomainReference = *insolar.NewReference(r.rootDomainID)
}

// ID returns insolar.ID  to root domain object.
func (r *Record) ID() insolar.ID {
	r.once.Do(r.initialize)

	return r.rootDomainID
}

// Reference returns insolar.Reference to root domain object
func (r *Record) Reference() insolar.Reference {
	r.once.Do(r.initialize)

	return r.rootDomainReference
}
