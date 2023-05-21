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

// SourceMappingRule struct for SourceMappingRule
type SourceMappingRule struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	WhitelistIds *[]string `json:"whitelist_ids,omitempty"`
	// Optional single policy to evalute, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1.
	PolicyId *string `json:"policy_id,omitempty"`
	// List of policyIds to evaluate in order, to completion
	PolicyIds *[]string `json:"policy_ids,omitempty"`
	Host string `json:"host"`
	Repository string `json:"repository"`
}

// NewSourceMappingRule instantiates a new SourceMappingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceMappingRule(name string, host string, repository string) *SourceMappingRule {
	this := SourceMappingRule{}
	this.Name = name
	this.Host = host
	this.Repository = repository
	return &this
}

// NewSourceMappingRuleWithDefaults instantiates a new SourceMappingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceMappingRuleWithDefaults() *SourceMappingRule {
	this := SourceMappingRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceMappingRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceMappingRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceMappingRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SourceMappingRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceMappingRule) SetName(v string) {
	o.Name = v
}

// GetWhitelistIds returns the WhitelistIds field value if set, zero value otherwise.
func (o *SourceMappingRule) GetWhitelistIds() []string {
	if o == nil || o.WhitelistIds == nil {
		var ret []string
		return ret
	}
	return *o.WhitelistIds
}

// GetWhitelistIdsOk returns a tuple with the WhitelistIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetWhitelistIdsOk() (*[]string, bool) {
	if o == nil || o.WhitelistIds == nil {
		return nil, false
	}
	return o.WhitelistIds, true
}

// HasWhitelistIds returns a boolean if a field has been set.
func (o *SourceMappingRule) HasWhitelistIds() bool {
	if o != nil && o.WhitelistIds != nil {
		return true
	}

	return false
}

// SetWhitelistIds gets a reference to the given []string and assigns it to the WhitelistIds field.
func (o *SourceMappingRule) SetWhitelistIds(v []string) {
	o.WhitelistIds = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *SourceMappingRule) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *SourceMappingRule) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *SourceMappingRule) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyIds returns the PolicyIds field value if set, zero value otherwise.
func (o *SourceMappingRule) GetPolicyIds() []string {
	if o == nil || o.PolicyIds == nil {
		var ret []string
		return ret
	}
	return *o.PolicyIds
}

// GetPolicyIdsOk returns a tuple with the PolicyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetPolicyIdsOk() (*[]string, bool) {
	if o == nil || o.PolicyIds == nil {
		return nil, false
	}
	return o.PolicyIds, true
}

// HasPolicyIds returns a boolean if a field has been set.
func (o *SourceMappingRule) HasPolicyIds() bool {
	if o != nil && o.PolicyIds != nil {
		return true
	}

	return false
}

// SetPolicyIds gets a reference to the given []string and assigns it to the PolicyIds field.
func (o *SourceMappingRule) SetPolicyIds(v []string) {
	o.PolicyIds = &v
}

// GetHost returns the Host field value
func (o *SourceMappingRule) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *SourceMappingRule) SetHost(v string) {
	o.Host = v
}

// GetRepository returns the Repository field value
func (o *SourceMappingRule) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetRepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *SourceMappingRule) SetRepository(v string) {
	o.Repository = v
}

func (o SourceMappingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.WhitelistIds != nil {
		toSerialize["whitelist_ids"] = o.WhitelistIds
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.PolicyIds != nil {
		toSerialize["policy_ids"] = o.PolicyIds
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	return json.Marshal(toSerialize)
}

type NullableSourceMappingRule struct {
	value *SourceMappingRule
	isSet bool
}

func (v NullableSourceMappingRule) Get() *SourceMappingRule {
	return v.value
}

func (v *NullableSourceMappingRule) Set(val *SourceMappingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceMappingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceMappingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceMappingRule(val *SourceMappingRule) *NullableSourceMappingRule {
	return &NullableSourceMappingRule{value: val, isSet: true}
}

func (v NullableSourceMappingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceMappingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


