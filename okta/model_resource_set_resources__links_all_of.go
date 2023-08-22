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

// ResourceSetResourcesLinksAllOf struct for ResourceSetResourcesLinksAllOf
type ResourceSetResourcesLinksAllOf struct {
	ResourceSet          *HrefObject `json:"resource-set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetResourcesLinksAllOf ResourceSetResourcesLinksAllOf

// NewResourceSetResourcesLinksAllOf instantiates a new ResourceSetResourcesLinksAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetResourcesLinksAllOf() *ResourceSetResourcesLinksAllOf {
	this := ResourceSetResourcesLinksAllOf{}
	return &this
}

// NewResourceSetResourcesLinksAllOfWithDefaults instantiates a new ResourceSetResourcesLinksAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetResourcesLinksAllOfWithDefaults() *ResourceSetResourcesLinksAllOf {
	this := ResourceSetResourcesLinksAllOf{}
	return &this
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *ResourceSetResourcesLinksAllOf) GetResourceSet() HrefObject {
	if o == nil || o.ResourceSet == nil {
		var ret HrefObject
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResourcesLinksAllOf) GetResourceSetOk() (*HrefObject, bool) {
	if o == nil || o.ResourceSet == nil {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *ResourceSetResourcesLinksAllOf) HasResourceSet() bool {
	if o != nil && o.ResourceSet != nil {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given HrefObject and assigns it to the ResourceSet field.
func (o *ResourceSetResourcesLinksAllOf) SetResourceSet(v HrefObject) {
	o.ResourceSet = &v
}

func (o ResourceSetResourcesLinksAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceSet != nil {
		toSerialize["resource-set"] = o.ResourceSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetResourcesLinksAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetResourcesLinksAllOf := _ResourceSetResourcesLinksAllOf{}

	err = json.Unmarshal(bytes, &varResourceSetResourcesLinksAllOf)
	if err == nil {
		*o = ResourceSetResourcesLinksAllOf(varResourceSetResourcesLinksAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resource-set")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetResourcesLinksAllOf struct {
	value *ResourceSetResourcesLinksAllOf
	isSet bool
}

func (v NullableResourceSetResourcesLinksAllOf) Get() *ResourceSetResourcesLinksAllOf {
	return v.value
}

func (v *NullableResourceSetResourcesLinksAllOf) Set(val *ResourceSetResourcesLinksAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetResourcesLinksAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetResourcesLinksAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetResourcesLinksAllOf(val *ResourceSetResourcesLinksAllOf) *NullableResourceSetResourcesLinksAllOf {
	return &NullableResourceSetResourcesLinksAllOf{value: val, isSet: true}
}

func (v NullableResourceSetResourcesLinksAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetResourcesLinksAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
