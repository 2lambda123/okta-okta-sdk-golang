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

// OAuth2ScopeConsentType the model 'OAuth2ScopeConsentType'
type OAuth2ScopeConsentType string

// List of OAuth2ScopeConsentType
const (
	OAUTH2SCOPECONSENTTYPE_ADMIN    OAuth2ScopeConsentType = "ADMIN"
	OAUTH2SCOPECONSENTTYPE_IMPLICIT OAuth2ScopeConsentType = "IMPLICIT"
	OAUTH2SCOPECONSENTTYPE_REQUIRED OAuth2ScopeConsentType = "REQUIRED"
)

// All allowed values of OAuth2ScopeConsentType enum
var AllowedOAuth2ScopeConsentTypeEnumValues = []OAuth2ScopeConsentType{
	"ADMIN",
	"IMPLICIT",
	"REQUIRED",
}

func (v *OAuth2ScopeConsentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OAuth2ScopeConsentType(value)
	for _, existing := range AllowedOAuth2ScopeConsentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OAuth2ScopeConsentType", value)
}

// NewOAuth2ScopeConsentTypeFromValue returns a pointer to a valid OAuth2ScopeConsentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOAuth2ScopeConsentTypeFromValue(v string) (*OAuth2ScopeConsentType, error) {
	ev := OAuth2ScopeConsentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OAuth2ScopeConsentType: valid values are %v", v, AllowedOAuth2ScopeConsentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OAuth2ScopeConsentType) IsValid() bool {
	for _, existing := range AllowedOAuth2ScopeConsentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OAuth2ScopeConsentType value
func (v OAuth2ScopeConsentType) Ptr() *OAuth2ScopeConsentType {
	return &v
}

type NullableOAuth2ScopeConsentType struct {
	value *OAuth2ScopeConsentType
	isSet bool
}

func (v NullableOAuth2ScopeConsentType) Get() *OAuth2ScopeConsentType {
	return v.value
}

func (v *NullableOAuth2ScopeConsentType) Set(val *OAuth2ScopeConsentType) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ScopeConsentType) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ScopeConsentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ScopeConsentType(val *OAuth2ScopeConsentType) *NullableOAuth2ScopeConsentType {
	return &NullableOAuth2ScopeConsentType{value: val, isSet: true}
}

func (v NullableOAuth2ScopeConsentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ScopeConsentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
