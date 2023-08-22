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
)

// U2fUserFactor struct for U2fUserFactor
type U2fUserFactor struct {
	UserFactor
	Profile              *U2fUserFactorProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _U2fUserFactor U2fUserFactor

// NewU2fUserFactor instantiates a new U2fUserFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewU2fUserFactor() *U2fUserFactor {
	this := U2fUserFactor{}
	return &this
}

// NewU2fUserFactorWithDefaults instantiates a new U2fUserFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewU2fUserFactorWithDefaults() *U2fUserFactor {
	this := U2fUserFactor{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *U2fUserFactor) GetProfile() U2fUserFactorProfile {
	if o == nil || o.Profile == nil {
		var ret U2fUserFactorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *U2fUserFactor) GetProfileOk() (*U2fUserFactorProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *U2fUserFactor) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given U2fUserFactorProfile and assigns it to the Profile field.
func (o *U2fUserFactor) SetProfile(v U2fUserFactorProfile) {
	o.Profile = &v
}

func (o U2fUserFactor) MarshalJSON() ([]byte, error) {
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

func (o *U2fUserFactor) UnmarshalJSON(bytes []byte) (err error) {
	type U2fUserFactorWithoutEmbeddedStruct struct {
		Profile *U2fUserFactorProfile `json:"profile,omitempty"`
	}

	varU2fUserFactorWithoutEmbeddedStruct := U2fUserFactorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varU2fUserFactorWithoutEmbeddedStruct)
	if err == nil {
		varU2fUserFactor := _U2fUserFactor{}
		varU2fUserFactor.Profile = varU2fUserFactorWithoutEmbeddedStruct.Profile
		*o = U2fUserFactor(varU2fUserFactor)
	} else {
		return err
	}

	varU2fUserFactor := _U2fUserFactor{}

	err = json.Unmarshal(bytes, &varU2fUserFactor)
	if err == nil {
		o.UserFactor = varU2fUserFactor.UserFactor
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

type NullableU2fUserFactor struct {
	value *U2fUserFactor
	isSet bool
}

func (v NullableU2fUserFactor) Get() *U2fUserFactor {
	return v.value
}

func (v *NullableU2fUserFactor) Set(val *U2fUserFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableU2fUserFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableU2fUserFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableU2fUserFactor(val *U2fUserFactor) *NullableU2fUserFactor {
	return &NullableU2fUserFactor{value: val, isSet: true}
}

func (v NullableU2fUserFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableU2fUserFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
