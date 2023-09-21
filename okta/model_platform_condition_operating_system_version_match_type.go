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

// PlatformConditionOperatingSystemVersionMatchType the model 'PlatformConditionOperatingSystemVersionMatchType'
type PlatformConditionOperatingSystemVersionMatchType string

// List of PlatformConditionOperatingSystemVersionMatchType
const (
	PLATFORMCONDITIONOPERATINGSYSTEMVERSIONMATCHTYPE_EXPRESSION PlatformConditionOperatingSystemVersionMatchType = "EXPRESSION"
	PLATFORMCONDITIONOPERATINGSYSTEMVERSIONMATCHTYPE_SEMVER     PlatformConditionOperatingSystemVersionMatchType = "SEMVER"
)

// All allowed values of PlatformConditionOperatingSystemVersionMatchType enum
var AllowedPlatformConditionOperatingSystemVersionMatchTypeEnumValues = []PlatformConditionOperatingSystemVersionMatchType{
	"EXPRESSION",
	"SEMVER",
}

func (v *PlatformConditionOperatingSystemVersionMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlatformConditionOperatingSystemVersionMatchType(value)
	for _, existing := range AllowedPlatformConditionOperatingSystemVersionMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlatformConditionOperatingSystemVersionMatchType", value)
}

// NewPlatformConditionOperatingSystemVersionMatchTypeFromValue returns a pointer to a valid PlatformConditionOperatingSystemVersionMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlatformConditionOperatingSystemVersionMatchTypeFromValue(v string) (*PlatformConditionOperatingSystemVersionMatchType, error) {
	ev := PlatformConditionOperatingSystemVersionMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlatformConditionOperatingSystemVersionMatchType: valid values are %v", v, AllowedPlatformConditionOperatingSystemVersionMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlatformConditionOperatingSystemVersionMatchType) IsValid() bool {
	for _, existing := range AllowedPlatformConditionOperatingSystemVersionMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlatformConditionOperatingSystemVersionMatchType value
func (v PlatformConditionOperatingSystemVersionMatchType) Ptr() *PlatformConditionOperatingSystemVersionMatchType {
	return &v
}

type NullablePlatformConditionOperatingSystemVersionMatchType struct {
	value *PlatformConditionOperatingSystemVersionMatchType
	isSet bool
}

func (v NullablePlatformConditionOperatingSystemVersionMatchType) Get() *PlatformConditionOperatingSystemVersionMatchType {
	return v.value
}

func (v *NullablePlatformConditionOperatingSystemVersionMatchType) Set(val *PlatformConditionOperatingSystemVersionMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformConditionOperatingSystemVersionMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformConditionOperatingSystemVersionMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformConditionOperatingSystemVersionMatchType(val *PlatformConditionOperatingSystemVersionMatchType) *NullablePlatformConditionOperatingSystemVersionMatchType {
	return &NullablePlatformConditionOperatingSystemVersionMatchType{value: val, isSet: true}
}

func (v NullablePlatformConditionOperatingSystemVersionMatchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformConditionOperatingSystemVersionMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}