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

// DeviceAssuranceChromeOSPlatformAllOf struct for DeviceAssuranceChromeOSPlatformAllOf
type DeviceAssuranceChromeOSPlatformAllOf struct {
	ThirdPartySignalProviders *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _DeviceAssuranceChromeOSPlatformAllOf DeviceAssuranceChromeOSPlatformAllOf

// NewDeviceAssuranceChromeOSPlatformAllOf instantiates a new DeviceAssuranceChromeOSPlatformAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceChromeOSPlatformAllOf() *DeviceAssuranceChromeOSPlatformAllOf {
	this := DeviceAssuranceChromeOSPlatformAllOf{}
	return &this
}

// NewDeviceAssuranceChromeOSPlatformAllOfWithDefaults instantiates a new DeviceAssuranceChromeOSPlatformAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceChromeOSPlatformAllOfWithDefaults() *DeviceAssuranceChromeOSPlatformAllOf {
	this := DeviceAssuranceChromeOSPlatformAllOf{}
	return &this
}

// GetThirdPartySignalProviders returns the ThirdPartySignalProviders field value if set, zero value otherwise.
func (o *DeviceAssuranceChromeOSPlatformAllOf) GetThirdPartySignalProviders() DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders {
	if o == nil || o.ThirdPartySignalProviders == nil {
		var ret DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders
		return ret
	}
	return *o.ThirdPartySignalProviders
}

// GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceChromeOSPlatformAllOf) GetThirdPartySignalProvidersOk() (*DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders, bool) {
	if o == nil || o.ThirdPartySignalProviders == nil {
		return nil, false
	}
	return o.ThirdPartySignalProviders, true
}

// HasThirdPartySignalProviders returns a boolean if a field has been set.
func (o *DeviceAssuranceChromeOSPlatformAllOf) HasThirdPartySignalProviders() bool {
	if o != nil && o.ThirdPartySignalProviders != nil {
		return true
	}

	return false
}

// SetThirdPartySignalProviders gets a reference to the given DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders and assigns it to the ThirdPartySignalProviders field.
func (o *DeviceAssuranceChromeOSPlatformAllOf) SetThirdPartySignalProviders(v DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) {
	o.ThirdPartySignalProviders = &v
}

func (o DeviceAssuranceChromeOSPlatformAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ThirdPartySignalProviders != nil {
		toSerialize["thirdPartySignalProviders"] = o.ThirdPartySignalProviders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceChromeOSPlatformAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceAssuranceChromeOSPlatformAllOf := _DeviceAssuranceChromeOSPlatformAllOf{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceChromeOSPlatformAllOf)
	if err == nil {
		*o = DeviceAssuranceChromeOSPlatformAllOf(varDeviceAssuranceChromeOSPlatformAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "thirdPartySignalProviders")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceAssuranceChromeOSPlatformAllOf struct {
	value *DeviceAssuranceChromeOSPlatformAllOf
	isSet bool
}

func (v NullableDeviceAssuranceChromeOSPlatformAllOf) Get() *DeviceAssuranceChromeOSPlatformAllOf {
	return v.value
}

func (v *NullableDeviceAssuranceChromeOSPlatformAllOf) Set(val *DeviceAssuranceChromeOSPlatformAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceChromeOSPlatformAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceChromeOSPlatformAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceChromeOSPlatformAllOf(val *DeviceAssuranceChromeOSPlatformAllOf) *NullableDeviceAssuranceChromeOSPlatformAllOf {
	return &NullableDeviceAssuranceChromeOSPlatformAllOf{value: val, isSet: true}
}

func (v NullableDeviceAssuranceChromeOSPlatformAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceChromeOSPlatformAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
