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

// THIS CODE IS AUTOGENERATED

package {{ .Package }}

import (
	"github.com/pkg/errors"
{{ range $contract := .Contracts }}
    {{ $contract.ImportName }} "{{ $contract.ImportPath }}"
{{- end }}

    XXX_insolar "github.com/insolar/insolar/insolar"
    XXX_artifacts "github.com/insolar/insolar/logicrunner/artifacts"
    XXX_rootdomain "github.com/insolar/insolar/insolar/rootdomain"
)

func InitializeContractMethods() map[string]XXX_insolar.ContractWrapper {
    return map[string]XXX_insolar.ContractWrapper{
{{- range $contract := .Contracts }}
        "{{ $contract.Name }}": {{ $contract.ImportName }}.Initialize(),
{{- end }}
    }
}

func shouldLoadRef(strRef string) XXX_insolar.Reference {
    ref, err := XXX_insolar.NewReferenceFromBase58(strRef)
    if err != nil {
        panic(errors.Wrap(err, "Unexpected error, bailing out"))
    }
    return *ref
}

func InitializeCodeRefs() map[XXX_insolar.Reference]string {
    rv := make(map[XXX_insolar.Reference]string, 0)

    {{ range $contract := .Contracts -}}
    rv[shouldLoadRef("{{ $contract.CodeReference }}")] = "{{ $contract.Name }}"
    {{ end }}

    return rv
}

func InitializeCodeDescriptors() []XXX_artifacts.CodeDescriptor {
    rv := make([]XXX_artifacts.CodeDescriptor, 0)

    {{ range $contract := .Contracts -}}
    // {{ $contract.Name }}
    rv = append(rv, XXX_artifacts.NewCodeDescriptor(
        /* code:        */ nil,
        /* machineType: */ XXX_insolar.MachineTypeBuiltin,
        /* ref:         */ shouldLoadRef("{{ $contract.CodeReference }}"),
    ))
    {{ end }}
    return rv
}

func InitializePrototypeDescriptors() []XXX_artifacts.ObjectDescriptor {
    rv := make([]XXX_artifacts.ObjectDescriptor, 0)

    {{ range $contract := .Contracts }}
    { // {{ $contract.Name }}
        pRef := shouldLoadRef("{{ $contract.PrototypeReference }}")
        cRef := shouldLoadRef("{{ $contract.CodeReference }}")
        rv = append(rv, XXX_artifacts.NewObjectDescriptor(
            /* head:         */ pRef,
            /* state:        */ *pRef.GetLocal(),
            /* prototype:    */ &cRef,
            /* isPrototype:  */ true,
            /* childPointer: */ nil,
            /* memory:       */ nil,
            /* parent:       */ XXX_rootdomain.RootDomain.Ref(),
        ))
    }
    {{ end }}
    return rv
}
