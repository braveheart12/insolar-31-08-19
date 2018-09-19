/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package pulsar

import (
	"bytes"
	"testing"

	"github.com/insolar/insolar/core"
)

var mockEntropy = [64]byte{1, 2, 3, 4, 5, 6, 7, 8}

type MockEntropyGenerator struct {
}

func (generator *MockEntropyGenerator) GenerateEntropy() core.Entropy {
	return mockEntropy
}

func TestNewPulse(t *testing.T) {
	generator := &MockEntropyGenerator{}
	previousPulse := core.PulseNumber(876)
	expectedPulse := previousPulse + 1

	result := NewPulse(previousPulse, generator)

	if !bytes.Equal(result.Entropy[:], mockEntropy[:]) {
		t.Errorf("Expeced and actual entropies are different, got: %v, want: %v", result.Entropy, mockEntropy)
	}

	if result.PulseNumber != core.PulseNumber(expectedPulse) {
		t.Errorf("Expeced and actual pulse numbers are different, got: %v, want: %v", result.PulseNumber, expectedPulse)
	}
}
