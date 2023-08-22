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

// ErrorPageTouchPointVariant the model 'ErrorPageTouchPointVariant'
type ErrorPageTouchPointVariant string

// List of ErrorPageTouchPointVariant
const (
	ERRORPAGETOUCHPOINTVARIANT_BACKGROUND_IMAGE           ErrorPageTouchPointVariant = "BACKGROUND_IMAGE"
	ERRORPAGETOUCHPOINTVARIANT_BACKGROUND_SECONDARY_COLOR ErrorPageTouchPointVariant = "BACKGROUND_SECONDARY_COLOR"
	ERRORPAGETOUCHPOINTVARIANT_OKTA_DEFAULT               ErrorPageTouchPointVariant = "OKTA_DEFAULT"
)

// All allowed values of ErrorPageTouchPointVariant enum
var AllowedErrorPageTouchPointVariantEnumValues = []ErrorPageTouchPointVariant{
	"BACKGROUND_IMAGE",
	"BACKGROUND_SECONDARY_COLOR",
	"OKTA_DEFAULT",
}

func (v *ErrorPageTouchPointVariant) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorPageTouchPointVariant(value)
	for _, existing := range AllowedErrorPageTouchPointVariantEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorPageTouchPointVariant", value)
}

// NewErrorPageTouchPointVariantFromValue returns a pointer to a valid ErrorPageTouchPointVariant
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorPageTouchPointVariantFromValue(v string) (*ErrorPageTouchPointVariant, error) {
	ev := ErrorPageTouchPointVariant(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorPageTouchPointVariant: valid values are %v", v, AllowedErrorPageTouchPointVariantEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorPageTouchPointVariant) IsValid() bool {
	for _, existing := range AllowedErrorPageTouchPointVariantEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorPageTouchPointVariant value
func (v ErrorPageTouchPointVariant) Ptr() *ErrorPageTouchPointVariant {
	return &v
}

type NullableErrorPageTouchPointVariant struct {
	value *ErrorPageTouchPointVariant
	isSet bool
}

func (v NullableErrorPageTouchPointVariant) Get() *ErrorPageTouchPointVariant {
	return v.value
}

func (v *NullableErrorPageTouchPointVariant) Set(val *ErrorPageTouchPointVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorPageTouchPointVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorPageTouchPointVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorPageTouchPointVariant(val *ErrorPageTouchPointVariant) *NullableErrorPageTouchPointVariant {
	return &NullableErrorPageTouchPointVariant{value: val, isSet: true}
}

func (v NullableErrorPageTouchPointVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorPageTouchPointVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
