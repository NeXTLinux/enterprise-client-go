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

// VulnDiffResult The results of the comparing two vulnerability records during an update
type VulnDiffResult struct {
	Added *[]interface{} `json:"added,omitempty"`
	Updated *[]interface{} `json:"updated,omitempty"`
	Removed *[]interface{} `json:"removed,omitempty"`
}

// NewVulnDiffResult instantiates a new VulnDiffResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnDiffResult() *VulnDiffResult {
	this := VulnDiffResult{}
	return &this
}

// NewVulnDiffResultWithDefaults instantiates a new VulnDiffResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnDiffResultWithDefaults() *VulnDiffResult {
	this := VulnDiffResult{}
	return &this
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *VulnDiffResult) GetAdded() []interface{} {
	if o == nil || o.Added == nil {
		var ret []interface{}
		return ret
	}
	return *o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnDiffResult) GetAddedOk() (*[]interface{}, bool) {
	if o == nil || o.Added == nil {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *VulnDiffResult) HasAdded() bool {
	if o != nil && o.Added != nil {
		return true
	}

	return false
}

// SetAdded gets a reference to the given []interface{} and assigns it to the Added field.
func (o *VulnDiffResult) SetAdded(v []interface{}) {
	o.Added = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *VulnDiffResult) GetUpdated() []interface{} {
	if o == nil || o.Updated == nil {
		var ret []interface{}
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnDiffResult) GetUpdatedOk() (*[]interface{}, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *VulnDiffResult) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given []interface{} and assigns it to the Updated field.
func (o *VulnDiffResult) SetUpdated(v []interface{}) {
	o.Updated = &v
}

// GetRemoved returns the Removed field value if set, zero value otherwise.
func (o *VulnDiffResult) GetRemoved() []interface{} {
	if o == nil || o.Removed == nil {
		var ret []interface{}
		return ret
	}
	return *o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnDiffResult) GetRemovedOk() (*[]interface{}, bool) {
	if o == nil || o.Removed == nil {
		return nil, false
	}
	return o.Removed, true
}

// HasRemoved returns a boolean if a field has been set.
func (o *VulnDiffResult) HasRemoved() bool {
	if o != nil && o.Removed != nil {
		return true
	}

	return false
}

// SetRemoved gets a reference to the given []interface{} and assigns it to the Removed field.
func (o *VulnDiffResult) SetRemoved(v []interface{}) {
	o.Removed = &v
}

func (o VulnDiffResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Added != nil {
		toSerialize["added"] = o.Added
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Removed != nil {
		toSerialize["removed"] = o.Removed
	}
	return json.Marshal(toSerialize)
}

type NullableVulnDiffResult struct {
	value *VulnDiffResult
	isSet bool
}

func (v NullableVulnDiffResult) Get() *VulnDiffResult {
	return v.value
}

func (v *NullableVulnDiffResult) Set(val *VulnDiffResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnDiffResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnDiffResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnDiffResult(val *VulnDiffResult) *NullableVulnDiffResult {
	return &NullableVulnDiffResult{value: val, isSet: true}
}

func (v NullableVulnDiffResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnDiffResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


