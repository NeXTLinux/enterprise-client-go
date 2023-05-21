/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.20
Contact: nurmi@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
	"time"
)

// AnalysisArchiveTransitionHistory A rule for auto-archiving image analysis by time and/or tag-history
type AnalysisArchiveTransitionHistory struct {
	// The task that created & updated this entry
	TransitionTaskId *string `json:"transition_task_id,omitempty"`
	RuleId *string `json:"rule_id,omitempty"`
	ImageDigest *string `json:"imageDigest,omitempty"`
	Transition *string `json:"transition,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewAnalysisArchiveTransitionHistory instantiates a new AnalysisArchiveTransitionHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisArchiveTransitionHistory() *AnalysisArchiveTransitionHistory {
	this := AnalysisArchiveTransitionHistory{}
	return &this
}

// NewAnalysisArchiveTransitionHistoryWithDefaults instantiates a new AnalysisArchiveTransitionHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisArchiveTransitionHistoryWithDefaults() *AnalysisArchiveTransitionHistory {
	this := AnalysisArchiveTransitionHistory{}
	return &this
}

// GetTransitionTaskId returns the TransitionTaskId field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionHistory) GetTransitionTaskId() string {
	if o == nil || o.TransitionTaskId == nil {
		var ret string
		return ret
	}
	return *o.TransitionTaskId
}

// GetTransitionTaskIdOk returns a tuple with the TransitionTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionHistory) GetTransitionTaskIdOk() (*string, bool) {
	if o == nil || o.TransitionTaskId == nil {
		return nil, false
	}
	return o.TransitionTaskId, true
}

// HasTransitionTaskId returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionHistory) HasTransitionTaskId() bool {
	if o != nil && o.TransitionTaskId != nil {
		return true
	}

	return false
}

// SetTransitionTaskId gets a reference to the given string and assigns it to the TransitionTaskId field.
func (o *AnalysisArchiveTransitionHistory) SetTransitionTaskId(v string) {
	o.TransitionTaskId = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionHistory) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionHistory) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionHistory) HasRuleId() bool {
	if o != nil && o.RuleId != nil {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *AnalysisArchiveTransitionHistory) SetRuleId(v string) {
	o.RuleId = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionHistory) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionHistory) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionHistory) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *AnalysisArchiveTransitionHistory) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetTransition returns the Transition field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionHistory) GetTransition() string {
	if o == nil || o.Transition == nil {
		var ret string
		return ret
	}
	return *o.Transition
}

// GetTransitionOk returns a tuple with the Transition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionHistory) GetTransitionOk() (*string, bool) {
	if o == nil || o.Transition == nil {
		return nil, false
	}
	return o.Transition, true
}

// HasTransition returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionHistory) HasTransition() bool {
	if o != nil && o.Transition != nil {
		return true
	}

	return false
}

// SetTransition gets a reference to the given string and assigns it to the Transition field.
func (o *AnalysisArchiveTransitionHistory) SetTransition(v string) {
	o.Transition = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionHistory) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionHistory) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AnalysisArchiveTransitionHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionHistory) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionHistory) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionHistory) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AnalysisArchiveTransitionHistory) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o AnalysisArchiveTransitionHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransitionTaskId != nil {
		toSerialize["transition_task_id"] = o.TransitionTaskId
	}
	if o.RuleId != nil {
		toSerialize["rule_id"] = o.RuleId
	}
	if o.ImageDigest != nil {
		toSerialize["imageDigest"] = o.ImageDigest
	}
	if o.Transition != nil {
		toSerialize["transition"] = o.Transition
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisArchiveTransitionHistory struct {
	value *AnalysisArchiveTransitionHistory
	isSet bool
}

func (v NullableAnalysisArchiveTransitionHistory) Get() *AnalysisArchiveTransitionHistory {
	return v.value
}

func (v *NullableAnalysisArchiveTransitionHistory) Set(val *AnalysisArchiveTransitionHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisArchiveTransitionHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisArchiveTransitionHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisArchiveTransitionHistory(val *AnalysisArchiveTransitionHistory) *NullableAnalysisArchiveTransitionHistory {
	return &NullableAnalysisArchiveTransitionHistory{value: val, isSet: true}
}

func (v NullableAnalysisArchiveTransitionHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisArchiveTransitionHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


