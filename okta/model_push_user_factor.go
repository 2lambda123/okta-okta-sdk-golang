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
	"reflect"
	"strings"
	"time"
)

// PushUserFactor struct for PushUserFactor
type PushUserFactor struct {
	UserFactor
	ExpiresAt            *time.Time             `json:"expiresAt,omitempty"`
	FactorResult         *FactorResultType      `json:"factorResult,omitempty"`
	Profile              *PushUserFactorProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PushUserFactor PushUserFactor

// NewPushUserFactor instantiates a new PushUserFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushUserFactor() *PushUserFactor {
	this := PushUserFactor{}
	return &this
}

// NewPushUserFactorWithDefaults instantiates a new PushUserFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushUserFactorWithDefaults() *PushUserFactor {
	this := PushUserFactor{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PushUserFactor) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushUserFactor) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PushUserFactor) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *PushUserFactor) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFactorResult returns the FactorResult field value if set, zero value otherwise.
func (o *PushUserFactor) GetFactorResult() FactorResultType {
	if o == nil || o.FactorResult == nil {
		var ret FactorResultType
		return ret
	}
	return *o.FactorResult
}

// GetFactorResultOk returns a tuple with the FactorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushUserFactor) GetFactorResultOk() (*FactorResultType, bool) {
	if o == nil || o.FactorResult == nil {
		return nil, false
	}
	return o.FactorResult, true
}

// HasFactorResult returns a boolean if a field has been set.
func (o *PushUserFactor) HasFactorResult() bool {
	if o != nil && o.FactorResult != nil {
		return true
	}

	return false
}

// SetFactorResult gets a reference to the given FactorResultType and assigns it to the FactorResult field.
func (o *PushUserFactor) SetFactorResult(v FactorResultType) {
	o.FactorResult = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PushUserFactor) GetProfile() PushUserFactorProfile {
	if o == nil || o.Profile == nil {
		var ret PushUserFactorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushUserFactor) GetProfileOk() (*PushUserFactorProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PushUserFactor) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given PushUserFactorProfile and assigns it to the Profile field.
func (o *PushUserFactor) SetProfile(v PushUserFactorProfile) {
	o.Profile = &v
}

func (o PushUserFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.FactorResult != nil {
		toSerialize["factorResult"] = o.FactorResult
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PushUserFactor) UnmarshalJSON(bytes []byte) (err error) {
	type PushUserFactorWithoutEmbeddedStruct struct {
		ExpiresAt    *time.Time             `json:"expiresAt,omitempty"`
		FactorResult *FactorResultType      `json:"factorResult,omitempty"`
		Profile      *PushUserFactorProfile `json:"profile,omitempty"`
	}

	varPushUserFactorWithoutEmbeddedStruct := PushUserFactorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPushUserFactorWithoutEmbeddedStruct)
	if err == nil {
		varPushUserFactor := _PushUserFactor{}
		varPushUserFactor.ExpiresAt = varPushUserFactorWithoutEmbeddedStruct.ExpiresAt
		varPushUserFactor.FactorResult = varPushUserFactorWithoutEmbeddedStruct.FactorResult
		varPushUserFactor.Profile = varPushUserFactorWithoutEmbeddedStruct.Profile
		*o = PushUserFactor(varPushUserFactor)
	} else {
		return err
	}

	varPushUserFactor := _PushUserFactor{}

	err = json.Unmarshal(bytes, &varPushUserFactor)
	if err == nil {
		o.UserFactor = varPushUserFactor.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "factorResult")
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

type NullablePushUserFactor struct {
	value *PushUserFactor
	isSet bool
}

func (v NullablePushUserFactor) Get() *PushUserFactor {
	return v.value
}

func (v *NullablePushUserFactor) Set(val *PushUserFactor) {
	v.value = val
	v.isSet = true
}

func (v NullablePushUserFactor) IsSet() bool {
	return v.isSet
}

func (v *NullablePushUserFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushUserFactor(val *PushUserFactor) *NullablePushUserFactor {
	return &NullablePushUserFactor{value: val, isSet: true}
}

func (v NullablePushUserFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushUserFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
