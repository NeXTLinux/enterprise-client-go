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

// NativeSBOMDistribution struct for NativeSBOMDistribution
type NativeSBOMDistribution struct {
	Name string `json:"name"`
	Version string `json:"version"`
	IdLike string `json:"idLike"`
}

// NewNativeSBOMDistribution instantiates a new NativeSBOMDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMDistribution(name string, version string, idLike string) *NativeSBOMDistribution {
	this := NativeSBOMDistribution{}
	this.Name = name
	this.Version = version
	this.IdLike = idLike
	return &this
}

// NewNativeSBOMDistributionWithDefaults instantiates a new NativeSBOMDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMDistributionWithDefaults() *NativeSBOMDistribution {
	this := NativeSBOMDistribution{}
	return &this
}

// GetName returns the Name field value
func (o *NativeSBOMDistribution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMDistribution) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NativeSBOMDistribution) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *NativeSBOMDistribution) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMDistribution) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NativeSBOMDistribution) SetVersion(v string) {
	o.Version = v
}

// GetIdLike returns the IdLike field value
func (o *NativeSBOMDistribution) GetIdLike() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdLike
}

// GetIdLikeOk returns a tuple with the IdLike field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMDistribution) GetIdLikeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdLike, true
}

// SetIdLike sets field value
func (o *NativeSBOMDistribution) SetIdLike(v string) {
	o.IdLike = v
}

func (o NativeSBOMDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["idLike"] = o.IdLike
	}
	return json.Marshal(toSerialize)
}

type NullableNativeSBOMDistribution struct {
	value *NativeSBOMDistribution
	isSet bool
}

func (v NullableNativeSBOMDistribution) Get() *NativeSBOMDistribution {
	return v.value
}

func (v *NullableNativeSBOMDistribution) Set(val *NativeSBOMDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOMDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOMDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOMDistribution(val *NativeSBOMDistribution) *NullableNativeSBOMDistribution {
	return &NullableNativeSBOMDistribution{value: val, isSet: true}
}

func (v NullableNativeSBOMDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOMDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


