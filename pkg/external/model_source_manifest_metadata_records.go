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
	"time"
)

// SourceManifestMetadataRecords Metadata associated with a source upload
type SourceManifestMetadataRecords struct {
	Uuid *string `json:"uuid,omitempty"`
	CiWorkflowName *string `json:"ci_workflow_name,omitempty"`
	CiWorkflowExecutionTime *time.Time `json:"ci_workflow_execution_time,omitempty"`
	BranchName *string `json:"branch_name,omitempty"`
	ChangeAuthor *string `json:"change_author,omitempty"`
	VcsType *string `json:"vcs_type,omitempty"`
}

// NewSourceManifestMetadataRecords instantiates a new SourceManifestMetadataRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceManifestMetadataRecords() *SourceManifestMetadataRecords {
	this := SourceManifestMetadataRecords{}
	return &this
}

// NewSourceManifestMetadataRecordsWithDefaults instantiates a new SourceManifestMetadataRecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceManifestMetadataRecordsWithDefaults() *SourceManifestMetadataRecords {
	this := SourceManifestMetadataRecords{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SourceManifestMetadataRecords) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestMetadataRecords) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecords) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SourceManifestMetadataRecords) SetUuid(v string) {
	o.Uuid = &v
}

// GetCiWorkflowName returns the CiWorkflowName field value if set, zero value otherwise.
func (o *SourceManifestMetadataRecords) GetCiWorkflowName() string {
	if o == nil || o.CiWorkflowName == nil {
		var ret string
		return ret
	}
	return *o.CiWorkflowName
}

// GetCiWorkflowNameOk returns a tuple with the CiWorkflowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestMetadataRecords) GetCiWorkflowNameOk() (*string, bool) {
	if o == nil || o.CiWorkflowName == nil {
		return nil, false
	}
	return o.CiWorkflowName, true
}

// HasCiWorkflowName returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecords) HasCiWorkflowName() bool {
	if o != nil && o.CiWorkflowName != nil {
		return true
	}

	return false
}

// SetCiWorkflowName gets a reference to the given string and assigns it to the CiWorkflowName field.
func (o *SourceManifestMetadataRecords) SetCiWorkflowName(v string) {
	o.CiWorkflowName = &v
}

// GetCiWorkflowExecutionTime returns the CiWorkflowExecutionTime field value if set, zero value otherwise.
func (o *SourceManifestMetadataRecords) GetCiWorkflowExecutionTime() time.Time {
	if o == nil || o.CiWorkflowExecutionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CiWorkflowExecutionTime
}

// GetCiWorkflowExecutionTimeOk returns a tuple with the CiWorkflowExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestMetadataRecords) GetCiWorkflowExecutionTimeOk() (*time.Time, bool) {
	if o == nil || o.CiWorkflowExecutionTime == nil {
		return nil, false
	}
	return o.CiWorkflowExecutionTime, true
}

// HasCiWorkflowExecutionTime returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecords) HasCiWorkflowExecutionTime() bool {
	if o != nil && o.CiWorkflowExecutionTime != nil {
		return true
	}

	return false
}

// SetCiWorkflowExecutionTime gets a reference to the given time.Time and assigns it to the CiWorkflowExecutionTime field.
func (o *SourceManifestMetadataRecords) SetCiWorkflowExecutionTime(v time.Time) {
	o.CiWorkflowExecutionTime = &v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *SourceManifestMetadataRecords) GetBranchName() string {
	if o == nil || o.BranchName == nil {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestMetadataRecords) GetBranchNameOk() (*string, bool) {
	if o == nil || o.BranchName == nil {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecords) HasBranchName() bool {
	if o != nil && o.BranchName != nil {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *SourceManifestMetadataRecords) SetBranchName(v string) {
	o.BranchName = &v
}

// GetChangeAuthor returns the ChangeAuthor field value if set, zero value otherwise.
func (o *SourceManifestMetadataRecords) GetChangeAuthor() string {
	if o == nil || o.ChangeAuthor == nil {
		var ret string
		return ret
	}
	return *o.ChangeAuthor
}

// GetChangeAuthorOk returns a tuple with the ChangeAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestMetadataRecords) GetChangeAuthorOk() (*string, bool) {
	if o == nil || o.ChangeAuthor == nil {
		return nil, false
	}
	return o.ChangeAuthor, true
}

// HasChangeAuthor returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecords) HasChangeAuthor() bool {
	if o != nil && o.ChangeAuthor != nil {
		return true
	}

	return false
}

// SetChangeAuthor gets a reference to the given string and assigns it to the ChangeAuthor field.
func (o *SourceManifestMetadataRecords) SetChangeAuthor(v string) {
	o.ChangeAuthor = &v
}

// GetVcsType returns the VcsType field value if set, zero value otherwise.
func (o *SourceManifestMetadataRecords) GetVcsType() string {
	if o == nil || o.VcsType == nil {
		var ret string
		return ret
	}
	return *o.VcsType
}

// GetVcsTypeOk returns a tuple with the VcsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestMetadataRecords) GetVcsTypeOk() (*string, bool) {
	if o == nil || o.VcsType == nil {
		return nil, false
	}
	return o.VcsType, true
}

// HasVcsType returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecords) HasVcsType() bool {
	if o != nil && o.VcsType != nil {
		return true
	}

	return false
}

// SetVcsType gets a reference to the given string and assigns it to the VcsType field.
func (o *SourceManifestMetadataRecords) SetVcsType(v string) {
	o.VcsType = &v
}

func (o SourceManifestMetadataRecords) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.CiWorkflowName != nil {
		toSerialize["ci_workflow_name"] = o.CiWorkflowName
	}
	if o.CiWorkflowExecutionTime != nil {
		toSerialize["ci_workflow_execution_time"] = o.CiWorkflowExecutionTime
	}
	if o.BranchName != nil {
		toSerialize["branch_name"] = o.BranchName
	}
	if o.ChangeAuthor != nil {
		toSerialize["change_author"] = o.ChangeAuthor
	}
	if o.VcsType != nil {
		toSerialize["vcs_type"] = o.VcsType
	}
	return json.Marshal(toSerialize)
}

type NullableSourceManifestMetadataRecords struct {
	value *SourceManifestMetadataRecords
	isSet bool
}

func (v NullableSourceManifestMetadataRecords) Get() *SourceManifestMetadataRecords {
	return v.value
}

func (v *NullableSourceManifestMetadataRecords) Set(val *SourceManifestMetadataRecords) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceManifestMetadataRecords) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceManifestMetadataRecords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceManifestMetadataRecords(val *SourceManifestMetadataRecords) *NullableSourceManifestMetadataRecords {
	return &NullableSourceManifestMetadataRecords{value: val, isSet: true}
}

func (v NullableSourceManifestMetadataRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceManifestMetadataRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

