/*
Nextlinux Enterprise API Server

This is the Nextlinux Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
Contact: dev@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ComplianceViolationAlert Alert raised by the system on a compliance check failure
type ComplianceViolationAlert struct {
	// Identifier for the alert
	Uuid *string `json:"uuid,omitempty"`
	// Type of alert generated
	Type *string `json:"type,omitempty"`
	// Current state of the alert
	State *string `json:"state,omitempty"`
	Resource *ComplianceResource `json:"resource,omitempty"`
	// Account that closed the alert
	ClosedBy *string `json:"closed_by,omitempty"`
	// Reason for closing the alert
	ClosedReason *string `json:"closed_reason,omitempty"`
	// RFC 3339 formatted UTC timestamp when the alert was generated
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the alert was last modified
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Reason for compliance check status. Compliance check could fail due to policy evaluation or blacklisting or errors evaluating compliance
	ComplianceStatusReason *string `json:"compliance_status_reason,omitempty"`
	// Number of STOP action results in the compliance check report
	ViolationsCount *int32 `json:"violations_count,omitempty"`
}

// NewComplianceViolationAlert instantiates a new ComplianceViolationAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplianceViolationAlert() *ComplianceViolationAlert {
	this := ComplianceViolationAlert{}
	return &this
}

// NewComplianceViolationAlertWithDefaults instantiates a new ComplianceViolationAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplianceViolationAlertWithDefaults() *ComplianceViolationAlert {
	this := ComplianceViolationAlert{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ComplianceViolationAlert) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ComplianceViolationAlert) SetType(v string) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ComplianceViolationAlert) SetState(v string) {
	o.State = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetResource() ComplianceResource {
	if o == nil || o.Resource == nil {
		var ret ComplianceResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetResourceOk() (*ComplianceResource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given ComplianceResource and assigns it to the Resource field.
func (o *ComplianceViolationAlert) SetResource(v ComplianceResource) {
	o.Resource = &v
}

// GetClosedBy returns the ClosedBy field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetClosedBy() string {
	if o == nil || o.ClosedBy == nil {
		var ret string
		return ret
	}
	return *o.ClosedBy
}

// GetClosedByOk returns a tuple with the ClosedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetClosedByOk() (*string, bool) {
	if o == nil || o.ClosedBy == nil {
		return nil, false
	}
	return o.ClosedBy, true
}

// HasClosedBy returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasClosedBy() bool {
	if o != nil && o.ClosedBy != nil {
		return true
	}

	return false
}

// SetClosedBy gets a reference to the given string and assigns it to the ClosedBy field.
func (o *ComplianceViolationAlert) SetClosedBy(v string) {
	o.ClosedBy = &v
}

// GetClosedReason returns the ClosedReason field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetClosedReason() string {
	if o == nil || o.ClosedReason == nil {
		var ret string
		return ret
	}
	return *o.ClosedReason
}

// GetClosedReasonOk returns a tuple with the ClosedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetClosedReasonOk() (*string, bool) {
	if o == nil || o.ClosedReason == nil {
		return nil, false
	}
	return o.ClosedReason, true
}

// HasClosedReason returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasClosedReason() bool {
	if o != nil && o.ClosedReason != nil {
		return true
	}

	return false
}

// SetClosedReason gets a reference to the given string and assigns it to the ClosedReason field.
func (o *ComplianceViolationAlert) SetClosedReason(v string) {
	o.ClosedReason = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ComplianceViolationAlert) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ComplianceViolationAlert) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetComplianceStatusReason returns the ComplianceStatusReason field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetComplianceStatusReason() string {
	if o == nil || o.ComplianceStatusReason == nil {
		var ret string
		return ret
	}
	return *o.ComplianceStatusReason
}

// GetComplianceStatusReasonOk returns a tuple with the ComplianceStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetComplianceStatusReasonOk() (*string, bool) {
	if o == nil || o.ComplianceStatusReason == nil {
		return nil, false
	}
	return o.ComplianceStatusReason, true
}

// HasComplianceStatusReason returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasComplianceStatusReason() bool {
	if o != nil && o.ComplianceStatusReason != nil {
		return true
	}

	return false
}

// SetComplianceStatusReason gets a reference to the given string and assigns it to the ComplianceStatusReason field.
func (o *ComplianceViolationAlert) SetComplianceStatusReason(v string) {
	o.ComplianceStatusReason = &v
}

// GetViolationsCount returns the ViolationsCount field value if set, zero value otherwise.
func (o *ComplianceViolationAlert) GetViolationsCount() int32 {
	if o == nil || o.ViolationsCount == nil {
		var ret int32
		return ret
	}
	return *o.ViolationsCount
}

// GetViolationsCountOk returns a tuple with the ViolationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlert) GetViolationsCountOk() (*int32, bool) {
	if o == nil || o.ViolationsCount == nil {
		return nil, false
	}
	return o.ViolationsCount, true
}

// HasViolationsCount returns a boolean if a field has been set.
func (o *ComplianceViolationAlert) HasViolationsCount() bool {
	if o != nil && o.ViolationsCount != nil {
		return true
	}

	return false
}

// SetViolationsCount gets a reference to the given int32 and assigns it to the ViolationsCount field.
func (o *ComplianceViolationAlert) SetViolationsCount(v int32) {
	o.ViolationsCount = &v
}

func (o ComplianceViolationAlert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.ClosedBy != nil {
		toSerialize["closed_by"] = o.ClosedBy
	}
	if o.ClosedReason != nil {
		toSerialize["closed_reason"] = o.ClosedReason
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.ComplianceStatusReason != nil {
		toSerialize["compliance_status_reason"] = o.ComplianceStatusReason
	}
	if o.ViolationsCount != nil {
		toSerialize["violations_count"] = o.ViolationsCount
	}
	return json.Marshal(toSerialize)
}

type NullableComplianceViolationAlert struct {
	value *ComplianceViolationAlert
	isSet bool
}

func (v NullableComplianceViolationAlert) Get() *ComplianceViolationAlert {
	return v.value
}

func (v *NullableComplianceViolationAlert) Set(val *ComplianceViolationAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableComplianceViolationAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableComplianceViolationAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplianceViolationAlert(val *ComplianceViolationAlert) *NullableComplianceViolationAlert {
	return &NullableComplianceViolationAlert{value: val, isSet: true}
}

func (v NullableComplianceViolationAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplianceViolationAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


