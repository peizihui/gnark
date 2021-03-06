// +build bls377 !bn256,!bls381

/*
Copyright © 2020 ConsenSys

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mimc

import (
	"github.com/consensys/gnark/cs/internal/curve"
)

const mimcNbRounds = 91

// plain execution of a mimc run
// m: message
// k: encryption key
func (h MiMC) encrypt(m, k curve.Element) curve.Element {

	for _, cons := range h.Params {
		// m = (m+k+c)**-1
		m.Add(&m, &k).Add(&m, &cons).Inverse(&m)
	}
	m.Add(&m, &k)
	return m

}
