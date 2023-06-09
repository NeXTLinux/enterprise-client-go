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
)

// ServiceVersionEngine struct for ServiceVersionEngine
type ServiceVersionEngine struct {
	// Version of the installed engine library
	Version *string `json:"version,omitempty"`
	// Version of the installed engine db schema
	Db *string `json:"db,omitempty"`
}

// NewServiceVersionEngine instantiates a new ServiceVersionEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVersionEngine() *ServiceVersionEngine {
	this := ServiceVersionEngine{}
	return &this
}

// NewServiceVersionEngineWithDefaults instantiates a new ServiceVersionEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVersionEngineWithDefaults() *ServiceVersionEngine {
	this := ServiceVersionEngine{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceVersionEngine) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionEngine) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceVersionEngine) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ServiceVersionEngine) SetVersion(v string) {
	o.Version = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *ServiceVersionEngine) GetDb() string {
	if o == nil || o.Db == nil {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionEngine) GetDbOk() (*string, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *ServiceVersionEngine) HasDb() bool {
	if o != nil && o.Db != nil {
		return true
	}

	return false
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *ServiceVersionEngine) SetDb(v string) {
	o.Db = &v
}

func (o ServiceVersionEngine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Db != nil {
		toSerialize["db"] = o.Db
	}
	return json.Marshal(toSerialize)
}

type NullableServiceVersionEngine struct {
	value *ServiceVersionEngine
	isSet bool
}

func (v NullableServiceVersionEngine) Get() *ServiceVersionEngine {
	return v.value
}

func (v *NullableServiceVersionEngine) Set(val *ServiceVersionEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVersionEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVersionEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVersionEngine(val *ServiceVersionEngine) *NullableServiceVersionEngine {
	return &NullableServiceVersionEngine{value: val, isSet: true}
}

func (v NullableServiceVersionEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVersionEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


