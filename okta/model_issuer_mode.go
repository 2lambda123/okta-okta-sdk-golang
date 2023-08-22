/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// IssuerMode the model 'IssuerMode'
type IssuerMode string

// List of IssuerMode
const (
	ISSUERMODE_CUSTOM_URL IssuerMode = "CUSTOM_URL"
	ISSUERMODE_DYNAMIC    IssuerMode = "DYNAMIC"
	ISSUERMODE_ORG_URL    IssuerMode = "ORG_URL"
)

// All allowed values of IssuerMode enum
var AllowedIssuerModeEnumValues = []IssuerMode{
	"CUSTOM_URL",
	"DYNAMIC",
	"ORG_URL",
}

func (v *IssuerMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IssuerMode(value)
	for _, existing := range AllowedIssuerModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IssuerMode", value)
}

// NewIssuerModeFromValue returns a pointer to a valid IssuerMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIssuerModeFromValue(v string) (*IssuerMode, error) {
	ev := IssuerMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IssuerMode: valid values are %v", v, AllowedIssuerModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IssuerMode) IsValid() bool {
	for _, existing := range AllowedIssuerModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuerMode value
func (v IssuerMode) Ptr() *IssuerMode {
	return &v
}

type NullableIssuerMode struct {
	value *IssuerMode
	isSet bool
}

func (v NullableIssuerMode) Get() *IssuerMode {
	return v.value
}

func (v *NullableIssuerMode) Set(val *IssuerMode) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuerMode) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuerMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuerMode(val *IssuerMode) *NullableIssuerMode {
	return &NullableIssuerMode{value: val, isSet: true}
}

func (v NullableIssuerMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuerMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
