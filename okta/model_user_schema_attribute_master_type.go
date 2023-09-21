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

// UserSchemaAttributeMasterType the model 'UserSchemaAttributeMasterType'
type UserSchemaAttributeMasterType string

// List of UserSchemaAttributeMasterType
const (
	USERSCHEMAATTRIBUTEMASTERTYPE_OKTA           UserSchemaAttributeMasterType = "OKTA"
	USERSCHEMAATTRIBUTEMASTERTYPE_OVERRIDE       UserSchemaAttributeMasterType = "OVERRIDE"
	USERSCHEMAATTRIBUTEMASTERTYPE_PROFILE_MASTER UserSchemaAttributeMasterType = "PROFILE_MASTER"
)

// All allowed values of UserSchemaAttributeMasterType enum
var AllowedUserSchemaAttributeMasterTypeEnumValues = []UserSchemaAttributeMasterType{
	"OKTA",
	"OVERRIDE",
	"PROFILE_MASTER",
}

func (v *UserSchemaAttributeMasterType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserSchemaAttributeMasterType(value)
	for _, existing := range AllowedUserSchemaAttributeMasterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserSchemaAttributeMasterType", value)
}

// NewUserSchemaAttributeMasterTypeFromValue returns a pointer to a valid UserSchemaAttributeMasterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserSchemaAttributeMasterTypeFromValue(v string) (*UserSchemaAttributeMasterType, error) {
	ev := UserSchemaAttributeMasterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserSchemaAttributeMasterType: valid values are %v", v, AllowedUserSchemaAttributeMasterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserSchemaAttributeMasterType) IsValid() bool {
	for _, existing := range AllowedUserSchemaAttributeMasterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserSchemaAttributeMasterType value
func (v UserSchemaAttributeMasterType) Ptr() *UserSchemaAttributeMasterType {
	return &v
}

type NullableUserSchemaAttributeMasterType struct {
	value *UserSchemaAttributeMasterType
	isSet bool
}

func (v NullableUserSchemaAttributeMasterType) Get() *UserSchemaAttributeMasterType {
	return v.value
}

func (v *NullableUserSchemaAttributeMasterType) Set(val *UserSchemaAttributeMasterType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaAttributeMasterType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaAttributeMasterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaAttributeMasterType(val *UserSchemaAttributeMasterType) *NullableUserSchemaAttributeMasterType {
	return &NullableUserSchemaAttributeMasterType{value: val, isSet: true}
}

func (v NullableUserSchemaAttributeMasterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaAttributeMasterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}