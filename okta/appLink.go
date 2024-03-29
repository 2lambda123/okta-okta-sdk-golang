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

type AppLink struct {
	AppAssignmentId  string `json:"appAssignmentId,omitempty"`
	AppInstanceId    string `json:"appInstanceId,omitempty"`
	AppName          string `json:"appName,omitempty"`
	CredentialsSetup *bool  `json:"credentialsSetup,omitempty"`
	Hidden           *bool  `json:"hidden,omitempty"`
	Id               string `json:"id,omitempty"`
	Label            string `json:"label,omitempty"`
	LinkUrl          string `json:"linkUrl,omitempty"`
	LogoUrl          string `json:"logoUrl,omitempty"`
	SortOrder        int64  `json:"-"`
	SortOrderPtr     *int64 `json:"sortOrder,omitempty"`
}

func (a *AppLink) MarshalJSON() ([]byte, error) {
	type Alias AppLink
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.SortOrder != 0 {
		result.SortOrderPtr = Int64Ptr(a.SortOrder)
	}
	return json.Marshal(&result)
}

func (a *AppLink) UnmarshalJSON(data []byte) error {
	type Alias AppLink

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.SortOrderPtr != nil {
		a.SortOrder = *result.SortOrderPtr
		a.SortOrderPtr = result.SortOrderPtr
	}
	return nil
}
