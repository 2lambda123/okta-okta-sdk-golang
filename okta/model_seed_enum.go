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

API version: 4.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// SeedEnum the model 'SeedEnum'
type SeedEnum string

// List of SeedEnum
const (
	SEEDENUM_OKTA   SeedEnum = "OKTA"
	SEEDENUM_RANDOM SeedEnum = "RANDOM"
)

// All allowed values of SeedEnum enum
var AllowedSeedEnumEnumValues = []SeedEnum{
	"OKTA",
	"RANDOM",
}

func (v *SeedEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SeedEnum(value)
	for _, existing := range AllowedSeedEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SeedEnum", value)
}

// NewSeedEnumFromValue returns a pointer to a valid SeedEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSeedEnumFromValue(v string) (*SeedEnum, error) {
	ev := SeedEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SeedEnum: valid values are %v", v, AllowedSeedEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SeedEnum) IsValid() bool {
	for _, existing := range AllowedSeedEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeedEnum value
func (v SeedEnum) Ptr() *SeedEnum {
	return &v
}

type NullableSeedEnum struct {
	value *SeedEnum
	isSet bool
}

func (v NullableSeedEnum) Get() *SeedEnum {
	return v.value
}

func (v *NullableSeedEnum) Set(val *SeedEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSeedEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSeedEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeedEnum(val *SeedEnum) *NullableSeedEnum {
	return &NullableSeedEnum{value: val, isSet: true}
}

func (v NullableSeedEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeedEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
