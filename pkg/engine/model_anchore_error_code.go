/*
Nextlinux Engine API Server

This is the Nextlinux Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// NextlinuxErrorCode A description of an nextlinux error code (name, description)
type NextlinuxErrorCode struct {
	// Error code name
	Name *string `json:"name,omitempty"`
	// Description of the error code
	Description *string `json:"description,omitempty"`
}

// NewNextlinuxErrorCode instantiates a new NextlinuxErrorCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextlinuxErrorCode() *NextlinuxErrorCode {
	this := NextlinuxErrorCode{}
	return &this
}

// NewNextlinuxErrorCodeWithDefaults instantiates a new NextlinuxErrorCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextlinuxErrorCodeWithDefaults() *NextlinuxErrorCode {
	this := NextlinuxErrorCode{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NextlinuxErrorCode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextlinuxErrorCode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NextlinuxErrorCode) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NextlinuxErrorCode) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NextlinuxErrorCode) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextlinuxErrorCode) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NextlinuxErrorCode) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NextlinuxErrorCode) SetDescription(v string) {
	o.Description = &v
}

func (o NextlinuxErrorCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableNextlinuxErrorCode struct {
	value *NextlinuxErrorCode
	isSet bool
}

func (v NullableNextlinuxErrorCode) Get() *NextlinuxErrorCode {
	return v.value
}

func (v *NullableNextlinuxErrorCode) Set(val *NextlinuxErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableNextlinuxErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableNextlinuxErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextlinuxErrorCode(val *NextlinuxErrorCode) *NullableNextlinuxErrorCode {
	return &NullableNextlinuxErrorCode{value: val, isSet: true}
}

func (v NullableNextlinuxErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextlinuxErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


