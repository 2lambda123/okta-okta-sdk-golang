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
)

// CustomHotpUserFactorProfile struct for CustomHotpUserFactorProfile
type CustomHotpUserFactorProfile struct {
	SharedSecret         *string `json:"sharedSecret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomHotpUserFactorProfile CustomHotpUserFactorProfile

// NewCustomHotpUserFactorProfile instantiates a new CustomHotpUserFactorProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomHotpUserFactorProfile() *CustomHotpUserFactorProfile {
	this := CustomHotpUserFactorProfile{}
	return &this
}

// NewCustomHotpUserFactorProfileWithDefaults instantiates a new CustomHotpUserFactorProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomHotpUserFactorProfileWithDefaults() *CustomHotpUserFactorProfile {
	this := CustomHotpUserFactorProfile{}
	return &this
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *CustomHotpUserFactorProfile) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomHotpUserFactorProfile) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *CustomHotpUserFactorProfile) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *CustomHotpUserFactorProfile) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

func (o CustomHotpUserFactorProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SharedSecret != nil {
		toSerialize["sharedSecret"] = o.SharedSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomHotpUserFactorProfile) UnmarshalJSON(bytes []byte) (err error) {
	varCustomHotpUserFactorProfile := _CustomHotpUserFactorProfile{}

	err = json.Unmarshal(bytes, &varCustomHotpUserFactorProfile)
	if err == nil {
		*o = CustomHotpUserFactorProfile(varCustomHotpUserFactorProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "sharedSecret")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCustomHotpUserFactorProfile struct {
	value *CustomHotpUserFactorProfile
	isSet bool
}

func (v NullableCustomHotpUserFactorProfile) Get() *CustomHotpUserFactorProfile {
	return v.value
}

func (v *NullableCustomHotpUserFactorProfile) Set(val *CustomHotpUserFactorProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomHotpUserFactorProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomHotpUserFactorProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomHotpUserFactorProfile(val *CustomHotpUserFactorProfile) *NullableCustomHotpUserFactorProfile {
	return &NullableCustomHotpUserFactorProfile{value: val, isSet: true}
}

func (v NullableCustomHotpUserFactorProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomHotpUserFactorProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
