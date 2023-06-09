/*
Nextlinux Enterprise RBAC API

Enterprise API for managing roles, permissions, and user mappings

API version: 0.1.0
Contact: dev@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac

import (
	"encoding/json"
	"time"
)

// RoleMembership Membership for a role in an account
type RoleMembership struct {
	// The name of the role the user has permissions for
	Role *string `json:"role,omitempty"`
	ForAccount *[]string `json:"for_account,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewRoleMembership instantiates a new RoleMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMembership() *RoleMembership {
	this := RoleMembership{}
	return &this
}

// NewRoleMembershipWithDefaults instantiates a new RoleMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMembershipWithDefaults() *RoleMembership {
	this := RoleMembership{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleMembership) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMembership) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleMembership) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *RoleMembership) SetRole(v string) {
	o.Role = &v
}

// GetForAccount returns the ForAccount field value if set, zero value otherwise.
func (o *RoleMembership) GetForAccount() []string {
	if o == nil || o.ForAccount == nil {
		var ret []string
		return ret
	}
	return *o.ForAccount
}

// GetForAccountOk returns a tuple with the ForAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMembership) GetForAccountOk() (*[]string, bool) {
	if o == nil || o.ForAccount == nil {
		return nil, false
	}
	return o.ForAccount, true
}

// HasForAccount returns a boolean if a field has been set.
func (o *RoleMembership) HasForAccount() bool {
	if o != nil && o.ForAccount != nil {
		return true
	}

	return false
}

// SetForAccount gets a reference to the given []string and assigns it to the ForAccount field.
func (o *RoleMembership) SetForAccount(v []string) {
	o.ForAccount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RoleMembership) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMembership) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RoleMembership) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RoleMembership) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o RoleMembership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ForAccount != nil {
		toSerialize["for_account"] = o.ForAccount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRoleMembership struct {
	value *RoleMembership
	isSet bool
}

func (v NullableRoleMembership) Get() *RoleMembership {
	return v.value
}

func (v *NullableRoleMembership) Set(val *RoleMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMembership(val *RoleMembership) *NullableRoleMembership {
	return &NullableRoleMembership{value: val, isSet: true}
}

func (v NullableRoleMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


