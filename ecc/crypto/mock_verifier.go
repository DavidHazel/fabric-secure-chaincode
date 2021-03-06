/*
* Copyright IBM Corp. 2018 All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package crypto

// MockVerifier implements Verifier interface!
type MockVerifier struct {
}

// Verify returns true if signature validation of enclave return is correct; other false
func (v *MockVerifier) Verify(args, responseData []byte, readset, writeset [][]byte, signature, enclavePk []byte) (bool, error) {
	return true, nil
}
