/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

import "encoding/json"

type PasswordCredentialHash struct {
	Algorithm     string `json:"algorithm,omitempty"`
	Salt          string `json:"salt,omitempty"`
	SaltOrder     string `json:"saltOrder,omitempty"`
	Value         string `json:"value,omitempty"`
	WorkFactor    int64  `json:"-"`
	WorkFactorPtr *int64 `json:"workFactor,omitempty"`
}

func (a *PasswordCredentialHash) MarshalJSON() ([]byte, error) {
	type Alias PasswordCredentialHash
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.WorkFactor != 0 {
		result.WorkFactorPtr = Int64Ptr(a.WorkFactor)
	}
	return json.Marshal(&result)
}

func (a *PasswordCredentialHash) UnmarshalJSON(data []byte) error {
	type Alias PasswordCredentialHash

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.WorkFactorPtr != nil {
		a.WorkFactor = *result.WorkFactorPtr
		a.WorkFactorPtr = result.WorkFactorPtr
	}
	return nil
}
