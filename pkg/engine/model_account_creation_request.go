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

// AccountCreationRequest An account to create/add to the system. If already exists will return 400.
type AccountCreationRequest struct {
	// The account name to use. This will identify the account and must be globally unique in the system.
	Name string `json:"name"`
	// An optional email to associate with the account for contact purposes
	Email *string `json:"email,omitempty"`
}

// NewAccountCreationRequest instantiates a new AccountCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreationRequest(name string) *AccountCreationRequest {
	this := AccountCreationRequest{}
	this.Name = name
	return &this
}

// NewAccountCreationRequestWithDefaults instantiates a new AccountCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreationRequestWithDefaults() *AccountCreationRequest {
	this := AccountCreationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AccountCreationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountCreationRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountCreationRequest) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AccountCreationRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreationRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AccountCreationRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AccountCreationRequest) SetEmail(v string) {
	o.Email = &v
}

func (o AccountCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCreationRequest struct {
	value *AccountCreationRequest
	isSet bool
}

func (v NullableAccountCreationRequest) Get() *AccountCreationRequest {
	return v.value
}

func (v *NullableAccountCreationRequest) Set(val *AccountCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreationRequest(val *AccountCreationRequest) *NullableAccountCreationRequest {
	return &NullableAccountCreationRequest{value: val, isSet: true}
}

func (v NullableAccountCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


