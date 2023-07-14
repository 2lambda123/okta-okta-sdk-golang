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
)

// PasswordPolicyRuleAllOf struct for PasswordPolicyRuleAllOf
type PasswordPolicyRuleAllOf struct {
	Actions              *PasswordPolicyRuleActions    `json:"actions,omitempty"`
	Conditions           *PasswordPolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRuleAllOf PasswordPolicyRuleAllOf

// NewPasswordPolicyRuleAllOf instantiates a new PasswordPolicyRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRuleAllOf() *PasswordPolicyRuleAllOf {
	this := PasswordPolicyRuleAllOf{}
	return &this
}

// NewPasswordPolicyRuleAllOfWithDefaults instantiates a new PasswordPolicyRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRuleAllOfWithDefaults() *PasswordPolicyRuleAllOf {
	this := PasswordPolicyRuleAllOf{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *PasswordPolicyRuleAllOf) GetActions() PasswordPolicyRuleActions {
	if o == nil || o.Actions == nil {
		var ret PasswordPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRuleAllOf) GetActionsOk() (*PasswordPolicyRuleActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *PasswordPolicyRuleAllOf) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given PasswordPolicyRuleActions and assigns it to the Actions field.
func (o *PasswordPolicyRuleAllOf) SetActions(v PasswordPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PasswordPolicyRuleAllOf) GetConditions() PasswordPolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PasswordPolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRuleAllOf) GetConditionsOk() (*PasswordPolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PasswordPolicyRuleAllOf) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PasswordPolicyRuleConditions and assigns it to the Conditions field.
func (o *PasswordPolicyRuleAllOf) SetConditions(v PasswordPolicyRuleConditions) {
	o.Conditions = &v
}

func (o PasswordPolicyRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyRuleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyRuleAllOf := _PasswordPolicyRuleAllOf{}

	err = json.Unmarshal(bytes, &varPasswordPolicyRuleAllOf)
	if err == nil {
		*o = PasswordPolicyRuleAllOf(varPasswordPolicyRuleAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyRuleAllOf struct {
	value *PasswordPolicyRuleAllOf
	isSet bool
}

func (v NullablePasswordPolicyRuleAllOf) Get() *PasswordPolicyRuleAllOf {
	return v.value
}

func (v *NullablePasswordPolicyRuleAllOf) Set(val *PasswordPolicyRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRuleAllOf(val *PasswordPolicyRuleAllOf) *NullablePasswordPolicyRuleAllOf {
	return &NullablePasswordPolicyRuleAllOf{value: val, isSet: true}
}

func (v NullablePasswordPolicyRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
