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
	"reflect"
	"strings"
)

// SecurityQuestionUserFactor struct for SecurityQuestionUserFactor
type SecurityQuestionUserFactor struct {
	UserFactor
	Profile              *SecurityQuestionUserFactorProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityQuestionUserFactor SecurityQuestionUserFactor

// NewSecurityQuestionUserFactor instantiates a new SecurityQuestionUserFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityQuestionUserFactor() *SecurityQuestionUserFactor {
	this := SecurityQuestionUserFactor{}
	return &this
}

// NewSecurityQuestionUserFactorWithDefaults instantiates a new SecurityQuestionUserFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityQuestionUserFactorWithDefaults() *SecurityQuestionUserFactor {
	this := SecurityQuestionUserFactor{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SecurityQuestionUserFactor) GetProfile() SecurityQuestionUserFactorProfile {
	if o == nil || o.Profile == nil {
		var ret SecurityQuestionUserFactorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityQuestionUserFactor) GetProfileOk() (*SecurityQuestionUserFactorProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SecurityQuestionUserFactor) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given SecurityQuestionUserFactorProfile and assigns it to the Profile field.
func (o *SecurityQuestionUserFactor) SetProfile(v SecurityQuestionUserFactorProfile) {
	o.Profile = &v
}

func (o SecurityQuestionUserFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityQuestionUserFactor) UnmarshalJSON(bytes []byte) (err error) {
	type SecurityQuestionUserFactorWithoutEmbeddedStruct struct {
		Profile *SecurityQuestionUserFactorProfile `json:"profile,omitempty"`
	}

	varSecurityQuestionUserFactorWithoutEmbeddedStruct := SecurityQuestionUserFactorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSecurityQuestionUserFactorWithoutEmbeddedStruct)
	if err == nil {
		varSecurityQuestionUserFactor := _SecurityQuestionUserFactor{}
		varSecurityQuestionUserFactor.Profile = varSecurityQuestionUserFactorWithoutEmbeddedStruct.Profile
		*o = SecurityQuestionUserFactor(varSecurityQuestionUserFactor)
	} else {
		return err
	}

	varSecurityQuestionUserFactor := _SecurityQuestionUserFactor{}

	err = json.Unmarshal(bytes, &varSecurityQuestionUserFactor)
	if err == nil {
		o.UserFactor = varSecurityQuestionUserFactor.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")

		// remove fields from embedded structs
		reflectUserFactor := reflect.ValueOf(o.UserFactor)
		for i := 0; i < reflectUserFactor.Type().NumField(); i++ {
			t := reflectUserFactor.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityQuestionUserFactor struct {
	value *SecurityQuestionUserFactor
	isSet bool
}

func (v NullableSecurityQuestionUserFactor) Get() *SecurityQuestionUserFactor {
	return v.value
}

func (v *NullableSecurityQuestionUserFactor) Set(val *SecurityQuestionUserFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityQuestionUserFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityQuestionUserFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityQuestionUserFactor(val *SecurityQuestionUserFactor) *NullableSecurityQuestionUserFactor {
	return &NullableSecurityQuestionUserFactor{value: val, isSet: true}
}

func (v NullableSecurityQuestionUserFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityQuestionUserFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
