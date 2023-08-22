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

// WebUserFactorProfile struct for WebUserFactorProfile
type WebUserFactorProfile struct {
	CredentialId         *string `json:"credentialId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebUserFactorProfile WebUserFactorProfile

// NewWebUserFactorProfile instantiates a new WebUserFactorProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebUserFactorProfile() *WebUserFactorProfile {
	this := WebUserFactorProfile{}
	return &this
}

// NewWebUserFactorProfileWithDefaults instantiates a new WebUserFactorProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebUserFactorProfileWithDefaults() *WebUserFactorProfile {
	this := WebUserFactorProfile{}
	return &this
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *WebUserFactorProfile) GetCredentialId() string {
	if o == nil || o.CredentialId == nil {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebUserFactorProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || o.CredentialId == nil {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *WebUserFactorProfile) HasCredentialId() bool {
	if o != nil && o.CredentialId != nil {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *WebUserFactorProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

func (o WebUserFactorProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredentialId != nil {
		toSerialize["credentialId"] = o.CredentialId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WebUserFactorProfile) UnmarshalJSON(bytes []byte) (err error) {
	varWebUserFactorProfile := _WebUserFactorProfile{}

	err = json.Unmarshal(bytes, &varWebUserFactorProfile)
	if err == nil {
		*o = WebUserFactorProfile(varWebUserFactorProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentialId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWebUserFactorProfile struct {
	value *WebUserFactorProfile
	isSet bool
}

func (v NullableWebUserFactorProfile) Get() *WebUserFactorProfile {
	return v.value
}

func (v *NullableWebUserFactorProfile) Set(val *WebUserFactorProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableWebUserFactorProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableWebUserFactorProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebUserFactorProfile(val *WebUserFactorProfile) *NullableWebUserFactorProfile {
	return &NullableWebUserFactorProfile{value: val, isSet: true}
}

func (v NullableWebUserFactorProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebUserFactorProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
