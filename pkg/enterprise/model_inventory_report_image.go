/*
Nextlinux Enterprise API Server

This is the Nextlinux Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
Contact: dev@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// InventoryReportImage defines an image that may be provided for image inventory
type InventoryReportImage struct {
	// The tag name of the image. Must have more than one character.
	Tag string `json:"tag"`
	// the image digest
	RepoDigest *string `json:"repoDigest,omitempty"`
}

// NewInventoryReportImage instantiates a new InventoryReportImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReportImage(tag string) *InventoryReportImage {
	this := InventoryReportImage{}
	this.Tag = tag
	return &this
}

// NewInventoryReportImageWithDefaults instantiates a new InventoryReportImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReportImageWithDefaults() *InventoryReportImage {
	this := InventoryReportImage{}
	return &this
}

// GetTag returns the Tag field value
func (o *InventoryReportImage) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *InventoryReportImage) GetTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *InventoryReportImage) SetTag(v string) {
	o.Tag = v
}

// GetRepoDigest returns the RepoDigest field value if set, zero value otherwise.
func (o *InventoryReportImage) GetRepoDigest() string {
	if o == nil || o.RepoDigest == nil {
		var ret string
		return ret
	}
	return *o.RepoDigest
}

// GetRepoDigestOk returns a tuple with the RepoDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReportImage) GetRepoDigestOk() (*string, bool) {
	if o == nil || o.RepoDigest == nil {
		return nil, false
	}
	return o.RepoDigest, true
}

// HasRepoDigest returns a boolean if a field has been set.
func (o *InventoryReportImage) HasRepoDigest() bool {
	if o != nil && o.RepoDigest != nil {
		return true
	}

	return false
}

// SetRepoDigest gets a reference to the given string and assigns it to the RepoDigest field.
func (o *InventoryReportImage) SetRepoDigest(v string) {
	o.RepoDigest = &v
}

func (o InventoryReportImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if o.RepoDigest != nil {
		toSerialize["repoDigest"] = o.RepoDigest
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryReportImage struct {
	value *InventoryReportImage
	isSet bool
}

func (v NullableInventoryReportImage) Get() *InventoryReportImage {
	return v.value
}

func (v *NullableInventoryReportImage) Set(val *InventoryReportImage) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReportImage) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReportImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReportImage(val *InventoryReportImage) *NullableInventoryReportImage {
	return &NullableInventoryReportImage{value: val, isSet: true}
}

func (v NullableInventoryReportImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReportImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


