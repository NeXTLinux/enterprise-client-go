/*
Nextlinux Enterprise API Server

This is the Nextlinux Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.2.1
Contact: dev@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external

import (
	"encoding/json"
)

// ResourceLabel Label on the resource in the key value format
type ResourceLabel struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewResourceLabel instantiates a new ResourceLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceLabel() *ResourceLabel {
	this := ResourceLabel{}
	return &this
}

// NewResourceLabelWithDefaults instantiates a new ResourceLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceLabelWithDefaults() *ResourceLabel {
	this := ResourceLabel{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ResourceLabel) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ResourceLabel) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ResourceLabel) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResourceLabel) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResourceLabel) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResourceLabel) SetValue(v string) {
	o.Value = &v
}

func (o ResourceLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableResourceLabel struct {
	value *ResourceLabel
	isSet bool
}

func (v NullableResourceLabel) Get() *ResourceLabel {
	return v.value
}

func (v *NullableResourceLabel) Set(val *ResourceLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLabel(val *ResourceLabel) *NullableResourceLabel {
	return &NullableResourceLabel{value: val, isSet: true}
}

func (v NullableResourceLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

