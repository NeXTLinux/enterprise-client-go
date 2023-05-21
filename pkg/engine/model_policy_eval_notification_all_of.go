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

// PolicyEvalNotificationAllOf The Notification Object definition for Policy Eval Notifications
type PolicyEvalNotificationAllOf struct {
	Data *PolicyEvalNotificationData `json:"data,omitempty"`
}

// NewPolicyEvalNotificationAllOf instantiates a new PolicyEvalNotificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvalNotificationAllOf() *PolicyEvalNotificationAllOf {
	this := PolicyEvalNotificationAllOf{}
	return &this
}

// NewPolicyEvalNotificationAllOfWithDefaults instantiates a new PolicyEvalNotificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvalNotificationAllOfWithDefaults() *PolicyEvalNotificationAllOf {
	this := PolicyEvalNotificationAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PolicyEvalNotificationAllOf) GetData() PolicyEvalNotificationData {
	if o == nil || o.Data == nil {
		var ret PolicyEvalNotificationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvalNotificationAllOf) GetDataOk() (*PolicyEvalNotificationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PolicyEvalNotificationAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PolicyEvalNotificationData and assigns it to the Data field.
func (o *PolicyEvalNotificationAllOf) SetData(v PolicyEvalNotificationData) {
	o.Data = &v
}

func (o PolicyEvalNotificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEvalNotificationAllOf struct {
	value *PolicyEvalNotificationAllOf
	isSet bool
}

func (v NullablePolicyEvalNotificationAllOf) Get() *PolicyEvalNotificationAllOf {
	return v.value
}

func (v *NullablePolicyEvalNotificationAllOf) Set(val *PolicyEvalNotificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvalNotificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvalNotificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvalNotificationAllOf(val *PolicyEvalNotificationAllOf) *NullablePolicyEvalNotificationAllOf {
	return &NullablePolicyEvalNotificationAllOf{value: val, isSet: true}
}

func (v NullablePolicyEvalNotificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvalNotificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


