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

// SourceManifestArtifacts Sbom content reference
type SourceManifestArtifacts struct {
	ContentType *string `json:"content_type,omitempty"`
	Digest *string `json:"digest,omitempty"`
	Bucket *string `json:"bucket,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewSourceManifestArtifacts instantiates a new SourceManifestArtifacts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceManifestArtifacts() *SourceManifestArtifacts {
	this := SourceManifestArtifacts{}
	return &this
}

// NewSourceManifestArtifactsWithDefaults instantiates a new SourceManifestArtifacts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceManifestArtifactsWithDefaults() *SourceManifestArtifacts {
	this := SourceManifestArtifacts{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SourceManifestArtifacts) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestArtifacts) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SourceManifestArtifacts) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SourceManifestArtifacts) SetContentType(v string) {
	o.ContentType = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *SourceManifestArtifacts) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestArtifacts) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *SourceManifestArtifacts) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *SourceManifestArtifacts) SetDigest(v string) {
	o.Digest = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *SourceManifestArtifacts) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestArtifacts) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *SourceManifestArtifacts) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *SourceManifestArtifacts) SetBucket(v string) {
	o.Bucket = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SourceManifestArtifacts) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestArtifacts) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SourceManifestArtifacts) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SourceManifestArtifacts) SetKey(v string) {
	o.Key = &v
}

func (o SourceManifestArtifacts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableSourceManifestArtifacts struct {
	value *SourceManifestArtifacts
	isSet bool
}

func (v NullableSourceManifestArtifacts) Get() *SourceManifestArtifacts {
	return v.value
}

func (v *NullableSourceManifestArtifacts) Set(val *SourceManifestArtifacts) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceManifestArtifacts) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceManifestArtifacts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceManifestArtifacts(val *SourceManifestArtifacts) *NullableSourceManifestArtifacts {
	return &NullableSourceManifestArtifacts{value: val, isSet: true}
}

func (v NullableSourceManifestArtifacts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceManifestArtifacts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


