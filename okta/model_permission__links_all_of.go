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

// PermissionLinksAllOf struct for PermissionLinksAllOf
type PermissionLinksAllOf struct {
	Role                 *HrefObject `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PermissionLinksAllOf PermissionLinksAllOf

// NewPermissionLinksAllOf instantiates a new PermissionLinksAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionLinksAllOf() *PermissionLinksAllOf {
	this := PermissionLinksAllOf{}
	return &this
}

// NewPermissionLinksAllOfWithDefaults instantiates a new PermissionLinksAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionLinksAllOfWithDefaults() *PermissionLinksAllOf {
	this := PermissionLinksAllOf{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PermissionLinksAllOf) GetRole() HrefObject {
	if o == nil || o.Role == nil {
		var ret HrefObject
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionLinksAllOf) GetRoleOk() (*HrefObject, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PermissionLinksAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given HrefObject and assigns it to the Role field.
func (o *PermissionLinksAllOf) SetRole(v HrefObject) {
	o.Role = &v
}

func (o PermissionLinksAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PermissionLinksAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPermissionLinksAllOf := _PermissionLinksAllOf{}

	err = json.Unmarshal(bytes, &varPermissionLinksAllOf)
	if err == nil {
		*o = PermissionLinksAllOf(varPermissionLinksAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePermissionLinksAllOf struct {
	value *PermissionLinksAllOf
	isSet bool
}

func (v NullablePermissionLinksAllOf) Get() *PermissionLinksAllOf {
	return v.value
}

func (v *NullablePermissionLinksAllOf) Set(val *PermissionLinksAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionLinksAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionLinksAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionLinksAllOf(val *PermissionLinksAllOf) *NullablePermissionLinksAllOf {
	return &NullablePermissionLinksAllOf{value: val, isSet: true}
}

func (v NullablePermissionLinksAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionLinksAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}