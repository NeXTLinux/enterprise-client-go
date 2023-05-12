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
	"time"
)

// ImageAnalysisRequest A request to add an image to be watched and analyzed by the engine. Optionally include the dockerfile content. Either source, digest or tag must be present.
type ImageAnalysisRequest struct {
	// Base64 encoded content of the dockerfile for the image, if available. Deprecated in favor of the 'source' field.
	Dockerfile *string `json:"dockerfile,omitempty"`
	// A digest string for an image, maybe a pull string or just a digest. e.g. nginx@sha256:123 or sha256:abc123. If a pull string, it must have same regisry/repo as the tag field. Deprecated in favor of the 'source' field
	Digest *string `json:"digest,omitempty"`
	// Full pullable tag reference for image. e.g. docker.io/nginx:latest. Deprecated in favor of the 'source' field
	Tag *string `json:"tag,omitempty"`
	// Optional override of the image creation time, only honored when both tag and digest are also supplied  e.g. 2018-10-17T18:14:00Z. Deprecated in favor of the 'source' field
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Optional. The type of image this is adding, defaults to \"docker\". This can be ommitted until multiple image types are supported.
	ImageType *string `json:"image_type,omitempty"`
	// Annotations to be associated with the added image in key/value form
	Annotations *interface{} `json:"annotations,omitempty"`
	Source *ImageSource `json:"source,omitempty"`
}

// NewImageAnalysisRequest instantiates a new ImageAnalysisRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAnalysisRequest() *ImageAnalysisRequest {
	this := ImageAnalysisRequest{}
	return &this
}

// NewImageAnalysisRequestWithDefaults instantiates a new ImageAnalysisRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAnalysisRequestWithDefaults() *ImageAnalysisRequest {
	this := ImageAnalysisRequest{}
	return &this
}

// GetDockerfile returns the Dockerfile field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetDockerfile() string {
	if o == nil || o.Dockerfile == nil {
		var ret string
		return ret
	}
	return *o.Dockerfile
}

// GetDockerfileOk returns a tuple with the Dockerfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetDockerfileOk() (*string, bool) {
	if o == nil || o.Dockerfile == nil {
		return nil, false
	}
	return o.Dockerfile, true
}

// HasDockerfile returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasDockerfile() bool {
	if o != nil && o.Dockerfile != nil {
		return true
	}

	return false
}

// SetDockerfile gets a reference to the given string and assigns it to the Dockerfile field.
func (o *ImageAnalysisRequest) SetDockerfile(v string) {
	o.Dockerfile = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *ImageAnalysisRequest) SetDigest(v string) {
	o.Digest = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ImageAnalysisRequest) SetTag(v string) {
	o.Tag = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ImageAnalysisRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetImageType() string {
	if o == nil || o.ImageType == nil {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetImageTypeOk() (*string, bool) {
	if o == nil || o.ImageType == nil {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasImageType() bool {
	if o != nil && o.ImageType != nil {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *ImageAnalysisRequest) SetImageType(v string) {
	o.ImageType = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetAnnotations() interface{} {
	if o == nil || o.Annotations == nil {
		var ret interface{}
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetAnnotationsOk() (*interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *ImageAnalysisRequest) SetAnnotations(v interface{}) {
	o.Annotations = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetSource() ImageSource {
	if o == nil || o.Source == nil {
		var ret ImageSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetSourceOk() (*ImageSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given ImageSource and assigns it to the Source field.
func (o *ImageAnalysisRequest) SetSource(v ImageSource) {
	o.Source = &v
}

func (o ImageAnalysisRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dockerfile != nil {
		toSerialize["dockerfile"] = o.Dockerfile
	}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ImageType != nil {
		toSerialize["image_type"] = o.ImageType
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableImageAnalysisRequest struct {
	value *ImageAnalysisRequest
	isSet bool
}

func (v NullableImageAnalysisRequest) Get() *ImageAnalysisRequest {
	return v.value
}

func (v *NullableImageAnalysisRequest) Set(val *ImageAnalysisRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAnalysisRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAnalysisRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAnalysisRequest(val *ImageAnalysisRequest) *NullableImageAnalysisRequest {
	return &NullableImageAnalysisRequest{value: val, isSet: true}
}

func (v NullableImageAnalysisRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAnalysisRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


