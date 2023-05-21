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
)

// SubscriptionUpdate A modification to a subscription entry to change its status or value
type SubscriptionUpdate struct {
	// The new subscription value, e.g. the new tag to be subscribed to
	SubscriptionValue NullableString `json:"subscription_value,omitempty"`
	// Toggle the subscription processing on or off
	Active *bool `json:"active,omitempty"`
}

// NewSubscriptionUpdate instantiates a new SubscriptionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUpdate() *SubscriptionUpdate {
	this := SubscriptionUpdate{}
	return &this
}

// NewSubscriptionUpdateWithDefaults instantiates a new SubscriptionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUpdateWithDefaults() *SubscriptionUpdate {
	this := SubscriptionUpdate{}
	return &this
}

// GetSubscriptionValue returns the SubscriptionValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionUpdate) GetSubscriptionValue() string {
	if o == nil || o.SubscriptionValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionValue.Get()
}

// GetSubscriptionValueOk returns a tuple with the SubscriptionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionUpdate) GetSubscriptionValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionValue.Get(), o.SubscriptionValue.IsSet()
}

// HasSubscriptionValue returns a boolean if a field has been set.
func (o *SubscriptionUpdate) HasSubscriptionValue() bool {
	if o != nil && o.SubscriptionValue.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionValue gets a reference to the given NullableString and assigns it to the SubscriptionValue field.
func (o *SubscriptionUpdate) SetSubscriptionValue(v string) {
	o.SubscriptionValue.Set(&v)
}
// SetSubscriptionValueNil sets the value for SubscriptionValue to be an explicit nil
func (o *SubscriptionUpdate) SetSubscriptionValueNil() {
	o.SubscriptionValue.Set(nil)
}

// UnsetSubscriptionValue ensures that no value is present for SubscriptionValue, not even an explicit nil
func (o *SubscriptionUpdate) UnsetSubscriptionValue() {
	o.SubscriptionValue.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SubscriptionUpdate) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdate) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SubscriptionUpdate) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SubscriptionUpdate) SetActive(v bool) {
	o.Active = &v
}

func (o SubscriptionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscriptionValue.IsSet() {
		toSerialize["subscription_value"] = o.SubscriptionValue.Get()
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionUpdate struct {
	value *SubscriptionUpdate
	isSet bool
}

func (v NullableSubscriptionUpdate) Get() *SubscriptionUpdate {
	return v.value
}

func (v *NullableSubscriptionUpdate) Set(val *SubscriptionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUpdate(val *SubscriptionUpdate) *NullableSubscriptionUpdate {
	return &NullableSubscriptionUpdate{value: val, isSet: true}
}

func (v NullableSubscriptionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


