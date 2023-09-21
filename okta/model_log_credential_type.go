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

// LogCredentialType the model 'LogCredentialType'
type LogCredentialType string

// List of LogCredentialType
const (
	LOGCREDENTIALTYPE_ASSERTION  LogCredentialType = "ASSERTION"
	LOGCREDENTIALTYPE_EMAIL      LogCredentialType = "EMAIL"
	LOGCREDENTIALTYPE_IWA        LogCredentialType = "IWA"
	LOGCREDENTIALTYPE_JWT        LogCredentialType = "JWT"
	LOGCREDENTIALTYPE_O_AUTH_2_0 LogCredentialType = "OAuth 2.0"
	LOGCREDENTIALTYPE_OTP        LogCredentialType = "OTP"
	LOGCREDENTIALTYPE_PASSWORD   LogCredentialType = "PASSWORD"
	LOGCREDENTIALTYPE_SMS        LogCredentialType = "SMS"
)

// All allowed values of LogCredentialType enum
var AllowedLogCredentialTypeEnumValues = []LogCredentialType{
	"ASSERTION",
	"EMAIL",
	"IWA",
	"JWT",
	"OAuth 2.0",
	"OTP",
	"PASSWORD",
	"SMS",
}

func (v *LogCredentialType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogCredentialType(value)
	for _, existing := range AllowedLogCredentialTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogCredentialType", value)
}

// NewLogCredentialTypeFromValue returns a pointer to a valid LogCredentialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogCredentialTypeFromValue(v string) (*LogCredentialType, error) {
	ev := LogCredentialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogCredentialType: valid values are %v", v, AllowedLogCredentialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogCredentialType) IsValid() bool {
	for _, existing := range AllowedLogCredentialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogCredentialType value
func (v LogCredentialType) Ptr() *LogCredentialType {
	return &v
}

type NullableLogCredentialType struct {
	value *LogCredentialType
	isSet bool
}

func (v NullableLogCredentialType) Get() *LogCredentialType {
	return v.value
}

func (v *NullableLogCredentialType) Set(val *LogCredentialType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogCredentialType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogCredentialType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogCredentialType(val *LogCredentialType) *NullableLogCredentialType {
	return &NullableLogCredentialType{value: val, isSet: true}
}

func (v NullableLogCredentialType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogCredentialType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}