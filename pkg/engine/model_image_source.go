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

// ImageSource A set of analysis source types. Only one may be set in any given request.
type ImageSource struct {
	Tag NullableRegistryTagSource `json:"tag,omitempty"`
	Digest NullableRegistryDigestSource `json:"digest,omitempty"`
	Archive NullableAnalysisArchiveSource `json:"archive,omitempty"`
	Import NullableImageImportManifest `json:"import,omitempty"`
}

// NewImageSource instantiates a new ImageSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageSource() *ImageSource {
	this := ImageSource{}
	return &this
}

// NewImageSourceWithDefaults instantiates a new ImageSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageSourceWithDefaults() *ImageSource {
	this := ImageSource{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageSource) GetTag() RegistryTagSource {
	if o == nil || o.Tag.Get() == nil {
		var ret RegistryTagSource
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageSource) GetTagOk() (*RegistryTagSource, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *ImageSource) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableRegistryTagSource and assigns it to the Tag field.
func (o *ImageSource) SetTag(v RegistryTagSource) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *ImageSource) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *ImageSource) UnsetTag() {
	o.Tag.Unset()
}

// GetDigest returns the Digest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageSource) GetDigest() RegistryDigestSource {
	if o == nil || o.Digest.Get() == nil {
		var ret RegistryDigestSource
		return ret
	}
	return *o.Digest.Get()
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageSource) GetDigestOk() (*RegistryDigestSource, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Digest.Get(), o.Digest.IsSet()
}

// HasDigest returns a boolean if a field has been set.
func (o *ImageSource) HasDigest() bool {
	if o != nil && o.Digest.IsSet() {
		return true
	}

	return false
}

// SetDigest gets a reference to the given NullableRegistryDigestSource and assigns it to the Digest field.
func (o *ImageSource) SetDigest(v RegistryDigestSource) {
	o.Digest.Set(&v)
}
// SetDigestNil sets the value for Digest to be an explicit nil
func (o *ImageSource) SetDigestNil() {
	o.Digest.Set(nil)
}

// UnsetDigest ensures that no value is present for Digest, not even an explicit nil
func (o *ImageSource) UnsetDigest() {
	o.Digest.Unset()
}

// GetArchive returns the Archive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageSource) GetArchive() AnalysisArchiveSource {
	if o == nil || o.Archive.Get() == nil {
		var ret AnalysisArchiveSource
		return ret
	}
	return *o.Archive.Get()
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageSource) GetArchiveOk() (*AnalysisArchiveSource, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Archive.Get(), o.Archive.IsSet()
}

// HasArchive returns a boolean if a field has been set.
func (o *ImageSource) HasArchive() bool {
	if o != nil && o.Archive.IsSet() {
		return true
	}

	return false
}

// SetArchive gets a reference to the given NullableAnalysisArchiveSource and assigns it to the Archive field.
func (o *ImageSource) SetArchive(v AnalysisArchiveSource) {
	o.Archive.Set(&v)
}
// SetArchiveNil sets the value for Archive to be an explicit nil
func (o *ImageSource) SetArchiveNil() {
	o.Archive.Set(nil)
}

// UnsetArchive ensures that no value is present for Archive, not even an explicit nil
func (o *ImageSource) UnsetArchive() {
	o.Archive.Unset()
}

// GetImport returns the Import field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageSource) GetImport() ImageImportManifest {
	if o == nil || o.Import.Get() == nil {
		var ret ImageImportManifest
		return ret
	}
	return *o.Import.Get()
}

// GetImportOk returns a tuple with the Import field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageSource) GetImportOk() (*ImageImportManifest, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Import.Get(), o.Import.IsSet()
}

// HasImport returns a boolean if a field has been set.
func (o *ImageSource) HasImport() bool {
	if o != nil && o.Import.IsSet() {
		return true
	}

	return false
}

// SetImport gets a reference to the given NullableImageImportManifest and assigns it to the Import field.
func (o *ImageSource) SetImport(v ImageImportManifest) {
	o.Import.Set(&v)
}
// SetImportNil sets the value for Import to be an explicit nil
func (o *ImageSource) SetImportNil() {
	o.Import.Set(nil)
}

// UnsetImport ensures that no value is present for Import, not even an explicit nil
func (o *ImageSource) UnsetImport() {
	o.Import.Unset()
}

func (o ImageSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	if o.Digest.IsSet() {
		toSerialize["digest"] = o.Digest.Get()
	}
	if o.Archive.IsSet() {
		toSerialize["archive"] = o.Archive.Get()
	}
	if o.Import.IsSet() {
		toSerialize["import"] = o.Import.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableImageSource struct {
	value *ImageSource
	isSet bool
}

func (v NullableImageSource) Get() *ImageSource {
	return v.value
}

func (v *NullableImageSource) Set(val *ImageSource) {
	v.value = val
	v.isSet = true
}

func (v NullableImageSource) IsSet() bool {
	return v.isSet
}

func (v *NullableImageSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageSource(val *ImageSource) *NullableImageSource {
	return &NullableImageSource{value: val, isSet: true}
}

func (v NullableImageSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


