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

// ProvisioningConnection struct for ProvisioningConnection
type ProvisioningConnection struct {
	AuthScheme           ProvisioningConnectionAuthScheme `json:"authScheme"`
	Profile              *ProvisioningConnectionProfile   `json:"profile,omitempty"`
	Status               ProvisioningConnectionStatus     `json:"status"`
	Links                *LinksSelfAndLifecycle           `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnection ProvisioningConnection

// NewProvisioningConnection instantiates a new ProvisioningConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnection(authScheme ProvisioningConnectionAuthScheme, status ProvisioningConnectionStatus) *ProvisioningConnection {
	this := ProvisioningConnection{}
	this.AuthScheme = authScheme
	this.Status = status
	return &this
}

// NewProvisioningConnectionWithDefaults instantiates a new ProvisioningConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionWithDefaults() *ProvisioningConnection {
	this := ProvisioningConnection{}
	var status ProvisioningConnectionStatus = PROVISIONINGCONNECTIONSTATUS_DISABLED
	this.Status = status
	return &this
}

// GetAuthScheme returns the AuthScheme field value
func (o *ProvisioningConnection) GetAuthScheme() ProvisioningConnectionAuthScheme {
	if o == nil {
		var ret ProvisioningConnectionAuthScheme
		return ret
	}

	return o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnection) GetAuthSchemeOk() (*ProvisioningConnectionAuthScheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthScheme, true
}

// SetAuthScheme sets field value
func (o *ProvisioningConnection) SetAuthScheme(v ProvisioningConnectionAuthScheme) {
	o.AuthScheme = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ProvisioningConnection) GetProfile() ProvisioningConnectionProfile {
	if o == nil || o.Profile == nil {
		var ret ProvisioningConnectionProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnection) GetProfileOk() (*ProvisioningConnectionProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ProvisioningConnection) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ProvisioningConnectionProfile and assigns it to the Profile field.
func (o *ProvisioningConnection) SetProfile(v ProvisioningConnectionProfile) {
	o.Profile = &v
}

// GetStatus returns the Status field value
func (o *ProvisioningConnection) GetStatus() ProvisioningConnectionStatus {
	if o == nil {
		var ret ProvisioningConnectionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnection) GetStatusOk() (*ProvisioningConnectionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProvisioningConnection) SetStatus(v ProvisioningConnectionStatus) {
	o.Status = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProvisioningConnection) GetLinks() LinksSelfAndLifecycle {
	if o == nil || o.Links == nil {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnection) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProvisioningConnection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *ProvisioningConnection) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o ProvisioningConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authScheme"] = o.AuthScheme
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnection) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnection := _ProvisioningConnection{}

	err = json.Unmarshal(bytes, &varProvisioningConnection)
	if err == nil {
		*o = ProvisioningConnection(varProvisioningConnection)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioningConnection struct {
	value *ProvisioningConnection
	isSet bool
}

func (v NullableProvisioningConnection) Get() *ProvisioningConnection {
	return v.value
}

func (v *NullableProvisioningConnection) Set(val *ProvisioningConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnection(val *ProvisioningConnection) *NullableProvisioningConnection {
	return &NullableProvisioningConnection{value: val, isSet: true}
}

func (v NullableProvisioningConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}