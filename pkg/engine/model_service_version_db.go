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

// ServiceVersionDb struct for ServiceVersionDb
type ServiceVersionDb struct {
	// Semantic version of the db schema
	SchemaVersion *string `json:"schema_version,omitempty"`
}

// NewServiceVersionDb instantiates a new ServiceVersionDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVersionDb() *ServiceVersionDb {
	this := ServiceVersionDb{}
	return &this
}

// NewServiceVersionDbWithDefaults instantiates a new ServiceVersionDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVersionDbWithDefaults() *ServiceVersionDb {
	this := ServiceVersionDb{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *ServiceVersionDb) GetSchemaVersion() string {
	if o == nil || o.SchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDb) GetSchemaVersionOk() (*string, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *ServiceVersionDb) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *ServiceVersionDb) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

func (o ServiceVersionDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SchemaVersion != nil {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return json.Marshal(toSerialize)
}

type NullableServiceVersionDb struct {
	value *ServiceVersionDb
	isSet bool
}

func (v NullableServiceVersionDb) Get() *ServiceVersionDb {
	return v.value
}

func (v *NullableServiceVersionDb) Set(val *ServiceVersionDb) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVersionDb) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVersionDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVersionDb(val *ServiceVersionDb) *NullableServiceVersionDb {
	return &NullableServiceVersionDb{value: val, isSet: true}
}

func (v NullableServiceVersionDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVersionDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


