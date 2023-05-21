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

// TagUpdateNotificationPayload struct for TagUpdateNotificationPayload
type TagUpdateNotificationPayload struct {
	UserId *string `json:"userId,omitempty"`
	SubscriptionKey *string `json:"subscription_key,omitempty"`
	SubscriptionType *string `json:"subscription_type,omitempty"`
	NotificationId *string `json:"notificationId,omitempty"`
	// A list containing the current image digest
	CurrEval *[]interface{} `json:"curr_eval,omitempty"`
	// A list containing the previous image digests
	LastEval *[]interface{} `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations interface{} `json:"annotations,omitempty"`
}

// NewTagUpdateNotificationPayload instantiates a new TagUpdateNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdateNotificationPayload() *TagUpdateNotificationPayload {
	this := TagUpdateNotificationPayload{}
	return &this
}

// NewTagUpdateNotificationPayloadWithDefaults instantiates a new TagUpdateNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateNotificationPayloadWithDefaults() *TagUpdateNotificationPayload {
	this := TagUpdateNotificationPayload{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *TagUpdateNotificationPayload) SetUserId(v string) {
	o.UserId = &v
}

// GetSubscriptionKey returns the SubscriptionKey field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetSubscriptionKey() string {
	if o == nil || o.SubscriptionKey == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil || o.SubscriptionKey == nil {
		return nil, false
	}
	return o.SubscriptionKey, true
}

// HasSubscriptionKey returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasSubscriptionKey() bool {
	if o != nil && o.SubscriptionKey != nil {
		return true
	}

	return false
}

// SetSubscriptionKey gets a reference to the given string and assigns it to the SubscriptionKey field.
func (o *TagUpdateNotificationPayload) SetSubscriptionKey(v string) {
	o.SubscriptionKey = &v
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetSubscriptionType() string {
	if o == nil || o.SubscriptionType == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || o.SubscriptionType == nil {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasSubscriptionType() bool {
	if o != nil && o.SubscriptionType != nil {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *TagUpdateNotificationPayload) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetNotificationId() string {
	if o == nil || o.NotificationId == nil {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetNotificationIdOk() (*string, bool) {
	if o == nil || o.NotificationId == nil {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasNotificationId() bool {
	if o != nil && o.NotificationId != nil {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *TagUpdateNotificationPayload) SetNotificationId(v string) {
	o.NotificationId = &v
}

// GetCurrEval returns the CurrEval field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetCurrEval() []interface{} {
	if o == nil || o.CurrEval == nil {
		var ret []interface{}
		return ret
	}
	return *o.CurrEval
}

// GetCurrEvalOk returns a tuple with the CurrEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetCurrEvalOk() (*[]interface{}, bool) {
	if o == nil || o.CurrEval == nil {
		return nil, false
	}
	return o.CurrEval, true
}

// HasCurrEval returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasCurrEval() bool {
	if o != nil && o.CurrEval != nil {
		return true
	}

	return false
}

// SetCurrEval gets a reference to the given []interface{} and assigns it to the CurrEval field.
func (o *TagUpdateNotificationPayload) SetCurrEval(v []interface{}) {
	o.CurrEval = &v
}

// GetLastEval returns the LastEval field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetLastEval() []interface{} {
	if o == nil || o.LastEval == nil {
		var ret []interface{}
		return ret
	}
	return *o.LastEval
}

// GetLastEvalOk returns a tuple with the LastEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetLastEvalOk() (*[]interface{}, bool) {
	if o == nil || o.LastEval == nil {
		return nil, false
	}
	return o.LastEval, true
}

// HasLastEval returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasLastEval() bool {
	if o != nil && o.LastEval != nil {
		return true
	}

	return false
}

// SetLastEval gets a reference to the given []interface{} and assigns it to the LastEval field.
func (o *TagUpdateNotificationPayload) SetLastEval(v []interface{}) {
	o.LastEval = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagUpdateNotificationPayload) GetAnnotations() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagUpdateNotificationPayload) GetAnnotationsOk() (*interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *TagUpdateNotificationPayload) SetAnnotations(v interface{}) {
	o.Annotations = v
}

func (o TagUpdateNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.SubscriptionKey != nil {
		toSerialize["subscription_key"] = o.SubscriptionKey
	}
	if o.SubscriptionType != nil {
		toSerialize["subscription_type"] = o.SubscriptionType
	}
	if o.NotificationId != nil {
		toSerialize["notificationId"] = o.NotificationId
	}
	if o.CurrEval != nil {
		toSerialize["curr_eval"] = o.CurrEval
	}
	if o.LastEval != nil {
		toSerialize["last_eval"] = o.LastEval
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableTagUpdateNotificationPayload struct {
	value *TagUpdateNotificationPayload
	isSet bool
}

func (v NullableTagUpdateNotificationPayload) Get() *TagUpdateNotificationPayload {
	return v.value
}

func (v *NullableTagUpdateNotificationPayload) Set(val *TagUpdateNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdateNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdateNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdateNotificationPayload(val *TagUpdateNotificationPayload) *NullableTagUpdateNotificationPayload {
	return &NullableTagUpdateNotificationPayload{value: val, isSet: true}
}

func (v NullableTagUpdateNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdateNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


