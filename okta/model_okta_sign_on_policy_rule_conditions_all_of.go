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

// OktaSignOnPolicyRuleConditionsAllOf struct for OktaSignOnPolicyRuleConditionsAllOf
type OktaSignOnPolicyRuleConditionsAllOf struct {
	AuthContext          *PolicyRuleAuthContextCondition `json:"authContext,omitempty"`
	Network              *PolicyNetworkCondition         `json:"network,omitempty"`
	People               *PolicyPeopleCondition          `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicyRuleConditionsAllOf OktaSignOnPolicyRuleConditionsAllOf

// NewOktaSignOnPolicyRuleConditionsAllOf instantiates a new OktaSignOnPolicyRuleConditionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicyRuleConditionsAllOf() *OktaSignOnPolicyRuleConditionsAllOf {
	this := OktaSignOnPolicyRuleConditionsAllOf{}
	return &this
}

// NewOktaSignOnPolicyRuleConditionsAllOfWithDefaults instantiates a new OktaSignOnPolicyRuleConditionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyRuleConditionsAllOfWithDefaults() *OktaSignOnPolicyRuleConditionsAllOf {
	this := OktaSignOnPolicyRuleConditionsAllOf{}
	return &this
}

// GetAuthContext returns the AuthContext field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleConditionsAllOf) GetAuthContext() PolicyRuleAuthContextCondition {
	if o == nil || o.AuthContext == nil {
		var ret PolicyRuleAuthContextCondition
		return ret
	}
	return *o.AuthContext
}

// GetAuthContextOk returns a tuple with the AuthContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleConditionsAllOf) GetAuthContextOk() (*PolicyRuleAuthContextCondition, bool) {
	if o == nil || o.AuthContext == nil {
		return nil, false
	}
	return o.AuthContext, true
}

// HasAuthContext returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleConditionsAllOf) HasAuthContext() bool {
	if o != nil && o.AuthContext != nil {
		return true
	}

	return false
}

// SetAuthContext gets a reference to the given PolicyRuleAuthContextCondition and assigns it to the AuthContext field.
func (o *OktaSignOnPolicyRuleConditionsAllOf) SetAuthContext(v PolicyRuleAuthContextCondition) {
	o.AuthContext = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleConditionsAllOf) GetNetwork() PolicyNetworkCondition {
	if o == nil || o.Network == nil {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleConditionsAllOf) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleConditionsAllOf) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *OktaSignOnPolicyRuleConditionsAllOf) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleConditionsAllOf) GetPeople() PolicyPeopleCondition {
	if o == nil || o.People == nil {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleConditionsAllOf) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleConditionsAllOf) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *OktaSignOnPolicyRuleConditionsAllOf) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

func (o OktaSignOnPolicyRuleConditionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthContext != nil {
		toSerialize["authContext"] = o.AuthContext
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaSignOnPolicyRuleConditionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOktaSignOnPolicyRuleConditionsAllOf := _OktaSignOnPolicyRuleConditionsAllOf{}

	err = json.Unmarshal(bytes, &varOktaSignOnPolicyRuleConditionsAllOf)
	if err == nil {
		*o = OktaSignOnPolicyRuleConditionsAllOf(varOktaSignOnPolicyRuleConditionsAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authContext")
		delete(additionalProperties, "network")
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaSignOnPolicyRuleConditionsAllOf struct {
	value *OktaSignOnPolicyRuleConditionsAllOf
	isSet bool
}

func (v NullableOktaSignOnPolicyRuleConditionsAllOf) Get() *OktaSignOnPolicyRuleConditionsAllOf {
	return v.value
}

func (v *NullableOktaSignOnPolicyRuleConditionsAllOf) Set(val *OktaSignOnPolicyRuleConditionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyRuleConditionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyRuleConditionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyRuleConditionsAllOf(val *OktaSignOnPolicyRuleConditionsAllOf) *NullableOktaSignOnPolicyRuleConditionsAllOf {
	return &NullableOktaSignOnPolicyRuleConditionsAllOf{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyRuleConditionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyRuleConditionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
