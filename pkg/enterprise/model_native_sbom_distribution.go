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

// NativeSBOMDistribution struct for NativeSBOMDistribution
type NativeSBOMDistribution struct {
	Name NullableString `json:"name,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Version NullableString `json:"version,omitempty"`
	VersionID NullableString `json:"versionID,omitempty"`
	IdLike *interface{} `json:"idLike,omitempty"`
}

// NewNativeSBOMDistribution instantiates a new NativeSBOMDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMDistribution() *NativeSBOMDistribution {
	this := NativeSBOMDistribution{}
	return &this
}

// NewNativeSBOMDistributionWithDefaults instantiates a new NativeSBOMDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMDistributionWithDefaults() *NativeSBOMDistribution {
	this := NativeSBOMDistribution{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NativeSBOMDistribution) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NativeSBOMDistribution) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NativeSBOMDistribution) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NativeSBOMDistribution) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NativeSBOMDistribution) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NativeSBOMDistribution) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NativeSBOMDistribution) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NativeSBOMDistribution) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *NativeSBOMDistribution) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *NativeSBOMDistribution) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *NativeSBOMDistribution) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *NativeSBOMDistribution) UnsetId() {
	o.Id.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NativeSBOMDistribution) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NativeSBOMDistribution) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *NativeSBOMDistribution) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *NativeSBOMDistribution) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *NativeSBOMDistribution) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *NativeSBOMDistribution) UnsetVersion() {
	o.Version.Unset()
}

// GetVersionID returns the VersionID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NativeSBOMDistribution) GetVersionID() string {
	if o == nil || o.VersionID.Get() == nil {
		var ret string
		return ret
	}
	return *o.VersionID.Get()
}

// GetVersionIDOk returns a tuple with the VersionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NativeSBOMDistribution) GetVersionIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VersionID.Get(), o.VersionID.IsSet()
}

// HasVersionID returns a boolean if a field has been set.
func (o *NativeSBOMDistribution) HasVersionID() bool {
	if o != nil && o.VersionID.IsSet() {
		return true
	}

	return false
}

// SetVersionID gets a reference to the given NullableString and assigns it to the VersionID field.
func (o *NativeSBOMDistribution) SetVersionID(v string) {
	o.VersionID.Set(&v)
}
// SetVersionIDNil sets the value for VersionID to be an explicit nil
func (o *NativeSBOMDistribution) SetVersionIDNil() {
	o.VersionID.Set(nil)
}

// UnsetVersionID ensures that no value is present for VersionID, not even an explicit nil
func (o *NativeSBOMDistribution) UnsetVersionID() {
	o.VersionID.Unset()
}

// GetIdLike returns the IdLike field value if set, zero value otherwise.
func (o *NativeSBOMDistribution) GetIdLike() interface{} {
	if o == nil || o.IdLike == nil {
		var ret interface{}
		return ret
	}
	return *o.IdLike
}

// GetIdLikeOk returns a tuple with the IdLike field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMDistribution) GetIdLikeOk() (*interface{}, bool) {
	if o == nil || o.IdLike == nil {
		return nil, false
	}
	return o.IdLike, true
}

// HasIdLike returns a boolean if a field has been set.
func (o *NativeSBOMDistribution) HasIdLike() bool {
	if o != nil && o.IdLike != nil {
		return true
	}

	return false
}

// SetIdLike gets a reference to the given interface{} and assigns it to the IdLike field.
func (o *NativeSBOMDistribution) SetIdLike(v interface{}) {
	o.IdLike = &v
}

func (o NativeSBOMDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.VersionID.IsSet() {
		toSerialize["versionID"] = o.VersionID.Get()
	}
	if o.IdLike != nil {
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


