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

// PolicyEvaluationResult struct for PolicyEvaluationResult
type PolicyEvaluationResult struct {
	AccountId *string `json:"account_id,omitempty"`
	PolicyId *string `json:"policy_id,omitempty"`
	EvaluationId *string `json:"evaluation_id,omitempty"`
	SourceId *string `json:"source_id,omitempty"`
	VcsHost *string `json:"vcs_host,omitempty"`
	RepositoryName *string `json:"repository_name,omitempty"`
	FinalAction *string `json:"final_action,omitempty"`
	EvaluationUrl *string `json:"evaluation_url,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	Result *interface{} `json:"result,omitempty"`
}

// NewPolicyEvaluationResult instantiates a new PolicyEvaluationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationResult() *PolicyEvaluationResult {
	this := PolicyEvaluationResult{}
	return &this
}

// NewPolicyEvaluationResultWithDefaults instantiates a new PolicyEvaluationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationResultWithDefaults() *PolicyEvaluationResult {
	this := PolicyEvaluationResult{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *PolicyEvaluationResult) SetAccountId(v string) {
	o.AccountId = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *PolicyEvaluationResult) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetEvaluationId returns the EvaluationId field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetEvaluationId() string {
	if o == nil || o.EvaluationId == nil {
		var ret string
		return ret
	}
	return *o.EvaluationId
}

// GetEvaluationIdOk returns a tuple with the EvaluationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetEvaluationIdOk() (*string, bool) {
	if o == nil || o.EvaluationId == nil {
		return nil, false
	}
	return o.EvaluationId, true
}

// HasEvaluationId returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasEvaluationId() bool {
	if o != nil && o.EvaluationId != nil {
		return true
	}

	return false
}

// SetEvaluationId gets a reference to the given string and assigns it to the EvaluationId field.
func (o *PolicyEvaluationResult) SetEvaluationId(v string) {
	o.EvaluationId = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *PolicyEvaluationResult) SetSourceId(v string) {
	o.SourceId = &v
}

// GetVcsHost returns the VcsHost field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetVcsHost() string {
	if o == nil || o.VcsHost == nil {
		var ret string
		return ret
	}
	return *o.VcsHost
}

// GetVcsHostOk returns a tuple with the VcsHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetVcsHostOk() (*string, bool) {
	if o == nil || o.VcsHost == nil {
		return nil, false
	}
	return o.VcsHost, true
}

// HasVcsHost returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasVcsHost() bool {
	if o != nil && o.VcsHost != nil {
		return true
	}

	return false
}

// SetVcsHost gets a reference to the given string and assigns it to the VcsHost field.
func (o *PolicyEvaluationResult) SetVcsHost(v string) {
	o.VcsHost = &v
}

// GetRepositoryName returns the RepositoryName field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetRepositoryName() string {
	if o == nil || o.RepositoryName == nil {
		var ret string
		return ret
	}
	return *o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetRepositoryNameOk() (*string, bool) {
	if o == nil || o.RepositoryName == nil {
		return nil, false
	}
	return o.RepositoryName, true
}

// HasRepositoryName returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasRepositoryName() bool {
	if o != nil && o.RepositoryName != nil {
		return true
	}

	return false
}

// SetRepositoryName gets a reference to the given string and assigns it to the RepositoryName field.
func (o *PolicyEvaluationResult) SetRepositoryName(v string) {
	o.RepositoryName = &v
}

// GetFinalAction returns the FinalAction field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetFinalAction() string {
	if o == nil || o.FinalAction == nil {
		var ret string
		return ret
	}
	return *o.FinalAction
}

// GetFinalActionOk returns a tuple with the FinalAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetFinalActionOk() (*string, bool) {
	if o == nil || o.FinalAction == nil {
		return nil, false
	}
	return o.FinalAction, true
}

// HasFinalAction returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasFinalAction() bool {
	if o != nil && o.FinalAction != nil {
		return true
	}

	return false
}

// SetFinalAction gets a reference to the given string and assigns it to the FinalAction field.
func (o *PolicyEvaluationResult) SetFinalAction(v string) {
	o.FinalAction = &v
}

// GetEvaluationUrl returns the EvaluationUrl field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetEvaluationUrl() string {
	if o == nil || o.EvaluationUrl == nil {
		var ret string
		return ret
	}
	return *o.EvaluationUrl
}

// GetEvaluationUrlOk returns a tuple with the EvaluationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetEvaluationUrlOk() (*string, bool) {
	if o == nil || o.EvaluationUrl == nil {
		return nil, false
	}
	return o.EvaluationUrl, true
}

// HasEvaluationUrl returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasEvaluationUrl() bool {
	if o != nil && o.EvaluationUrl != nil {
		return true
	}

	return false
}

// SetEvaluationUrl gets a reference to the given string and assigns it to the EvaluationUrl field.
func (o *PolicyEvaluationResult) SetEvaluationUrl(v string) {
	o.EvaluationUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PolicyEvaluationResult) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PolicyEvaluationResult) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetResult() interface{} {
	if o == nil || o.Result == nil {
		var ret interface{}
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetResultOk() (*interface{}, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given interface{} and assigns it to the Result field.
func (o *PolicyEvaluationResult) SetResult(v interface{}) {
	o.Result = &v
}

func (o PolicyEvaluationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.EvaluationId != nil {
		toSerialize["evaluation_id"] = o.EvaluationId
	}
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	if o.VcsHost != nil {
		toSerialize["vcs_host"] = o.VcsHost
	}
	if o.RepositoryName != nil {
		toSerialize["repository_name"] = o.RepositoryName
	}
	if o.FinalAction != nil {
		toSerialize["final_action"] = o.FinalAction
	}
	if o.EvaluationUrl != nil {
		toSerialize["evaluation_url"] = o.EvaluationUrl
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEvaluationResult struct {
	value *PolicyEvaluationResult
	isSet bool
}

func (v NullablePolicyEvaluationResult) Get() *PolicyEvaluationResult {
	return v.value
}

func (v *NullablePolicyEvaluationResult) Set(val *PolicyEvaluationResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationResult(val *PolicyEvaluationResult) *NullablePolicyEvaluationResult {
	return &NullablePolicyEvaluationResult{value: val, isSet: true}
}

func (v NullablePolicyEvaluationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


