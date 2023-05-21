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

// EventResponseEventSource struct for EventResponseEventSource
type EventResponseEventSource struct {
	Servicename *string `json:"servicename,omitempty"`
	Hostid *string `json:"hostid,omitempty"`
	BaseUrl *string `json:"base_url,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
}

// NewEventResponseEventSource instantiates a new EventResponseEventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponseEventSource() *EventResponseEventSource {
	this := EventResponseEventSource{}
	return &this
}

// NewEventResponseEventSourceWithDefaults instantiates a new EventResponseEventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseEventSourceWithDefaults() *EventResponseEventSource {
	this := EventResponseEventSource{}
	return &this
}

// GetServicename returns the Servicename field value if set, zero value otherwise.
func (o *EventResponseEventSource) GetServicename() string {
	if o == nil || o.Servicename == nil {
		var ret string
		return ret
	}
	return *o.Servicename
}

// GetServicenameOk returns a tuple with the Servicename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEventSource) GetServicenameOk() (*string, bool) {
	if o == nil || o.Servicename == nil {
		return nil, false
	}
	return o.Servicename, true
}

// HasServicename returns a boolean if a field has been set.
func (o *EventResponseEventSource) HasServicename() bool {
	if o != nil && o.Servicename != nil {
		return true
	}

	return false
}

// SetServicename gets a reference to the given string and assigns it to the Servicename field.
func (o *EventResponseEventSource) SetServicename(v string) {
	o.Servicename = &v
}

// GetHostid returns the Hostid field value if set, zero value otherwise.
func (o *EventResponseEventSource) GetHostid() string {
	if o == nil || o.Hostid == nil {
		var ret string
		return ret
	}
	return *o.Hostid
}

// GetHostidOk returns a tuple with the Hostid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEventSource) GetHostidOk() (*string, bool) {
	if o == nil || o.Hostid == nil {
		return nil, false
	}
	return o.Hostid, true
}

// HasHostid returns a boolean if a field has been set.
func (o *EventResponseEventSource) HasHostid() bool {
	if o != nil && o.Hostid != nil {
		return true
	}

	return false
}

// SetHostid gets a reference to the given string and assigns it to the Hostid field.
func (o *EventResponseEventSource) SetHostid(v string) {
	o.Hostid = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *EventResponseEventSource) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEventSource) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *EventResponseEventSource) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *EventResponseEventSource) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *EventResponseEventSource) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEventSource) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *EventResponseEventSource) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *EventResponseEventSource) SetRequestId(v string) {
	o.RequestId = &v
}

func (o EventResponseEventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Servicename != nil {
		toSerialize["servicename"] = o.Servicename
	}
	if o.Hostid != nil {
		toSerialize["hostid"] = o.Hostid
	}
	if o.BaseUrl != nil {
		toSerialize["base_url"] = o.BaseUrl
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableEventResponseEventSource struct {
	value *EventResponseEventSource
	isSet bool
}

func (v NullableEventResponseEventSource) Get() *EventResponseEventSource {
	return v.value
}

func (v *NullableEventResponseEventSource) Set(val *EventResponseEventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponseEventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponseEventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponseEventSource(val *EventResponseEventSource) *NullableEventResponseEventSource {
	return &NullableEventResponseEventSource{value: val, isSet: true}
}

func (v NullableEventResponseEventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponseEventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


