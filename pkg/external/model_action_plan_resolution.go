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

// ActionPlanResolution defines the trigger IDs and content of a resolution for an action plan
type ActionPlanResolution struct {
	TriggerIds []string `json:"trigger_ids,omitempty"`
	Content *string `json:"content,omitempty"`
}

// NewActionPlanResolution instantiates a new ActionPlanResolution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionPlanResolution() *ActionPlanResolution {
	this := ActionPlanResolution{}
	return &this
}

// NewActionPlanResolutionWithDefaults instantiates a new ActionPlanResolution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionPlanResolutionWithDefaults() *ActionPlanResolution {
	this := ActionPlanResolution{}
	return &this
}

// GetTriggerIds returns the TriggerIds field value if set, zero value otherwise.
func (o *ActionPlanResolution) GetTriggerIds() []string {
	if o == nil || o.TriggerIds == nil {
		var ret []string
		return ret
	}
	return o.TriggerIds
}

// GetTriggerIdsOk returns a tuple with the TriggerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlanResolution) GetTriggerIdsOk() ([]string, bool) {
	if o == nil || o.TriggerIds == nil {
		return nil, false
	}
	return o.TriggerIds, true
}

// HasTriggerIds returns a boolean if a field has been set.
func (o *ActionPlanResolution) HasTriggerIds() bool {
	if o != nil && o.TriggerIds != nil {
		return true
	}

	return false
}

// SetTriggerIds gets a reference to the given []string and assigns it to the TriggerIds field.
func (o *ActionPlanResolution) SetTriggerIds(v []string) {
	o.TriggerIds = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ActionPlanResolution) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlanResolution) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ActionPlanResolution) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ActionPlanResolution) SetContent(v string) {
	o.Content = &v
}

func (o ActionPlanResolution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TriggerIds != nil {
		toSerialize["trigger_ids"] = o.TriggerIds
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableActionPlanResolution struct {
	value *ActionPlanResolution
	isSet bool
}

func (v NullableActionPlanResolution) Get() *ActionPlanResolution {
	return v.value
}

func (v *NullableActionPlanResolution) Set(val *ActionPlanResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableActionPlanResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableActionPlanResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionPlanResolution(val *ActionPlanResolution) *NullableActionPlanResolution {
	return &NullableActionPlanResolution{value: val, isSet: true}
}

func (v NullableActionPlanResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionPlanResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


