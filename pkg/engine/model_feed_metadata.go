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

// FeedMetadata Metadata on a single feed based on what the engine finds from querying the endpoints
type FeedMetadata struct {
	// name of the feed
	Name *string `json:"name,omitempty"`
	// Date the metadata record was created in engine (first seen on source)
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date the metadata was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Groups *[]FeedGroupMetadata `json:"groups,omitempty"`
	LastFullSync *time.Time `json:"last_full_sync,omitempty"`
	// If feed is enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewFeedMetadata instantiates a new FeedMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedMetadata() *FeedMetadata {
	this := FeedMetadata{}
	return &this
}

// NewFeedMetadataWithDefaults instantiates a new FeedMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedMetadataWithDefaults() *FeedMetadata {
	this := FeedMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeedMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeedMetadata) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeedMetadata) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FeedMetadata) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FeedMetadata) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FeedMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FeedMetadata) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FeedMetadata) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FeedMetadata) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *FeedMetadata) GetGroups() []FeedGroupMetadata {
	if o == nil || o.Groups == nil {
		var ret []FeedGroupMetadata
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetGroupsOk() (*[]FeedGroupMetadata, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *FeedMetadata) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []FeedGroupMetadata and assigns it to the Groups field.
func (o *FeedMetadata) SetGroups(v []FeedGroupMetadata) {
	o.Groups = &v
}

// GetLastFullSync returns the LastFullSync field value if set, zero value otherwise.
func (o *FeedMetadata) GetLastFullSync() time.Time {
	if o == nil || o.LastFullSync == nil {
		var ret time.Time
		return ret
	}
	return *o.LastFullSync
}

// GetLastFullSyncOk returns a tuple with the LastFullSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetLastFullSyncOk() (*time.Time, bool) {
	if o == nil || o.LastFullSync == nil {
		return nil, false
	}
	return o.LastFullSync, true
}

// HasLastFullSync returns a boolean if a field has been set.
func (o *FeedMetadata) HasLastFullSync() bool {
	if o != nil && o.LastFullSync != nil {
		return true
	}

	return false
}

// SetLastFullSync gets a reference to the given time.Time and assigns it to the LastFullSync field.
func (o *FeedMetadata) SetLastFullSync(v time.Time) {
	o.LastFullSync = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *FeedMetadata) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *FeedMetadata) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *FeedMetadata) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o FeedMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.LastFullSync != nil {
		toSerialize["last_full_sync"] = o.LastFullSync
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableFeedMetadata struct {
	value *FeedMetadata
	isSet bool
}

func (v NullableFeedMetadata) Get() *FeedMetadata {
	return v.value
}

func (v *NullableFeedMetadata) Set(val *FeedMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedMetadata(val *FeedMetadata) *NullableFeedMetadata {
	return &NullableFeedMetadata{value: val, isSet: true}
}

func (v NullableFeedMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


