# PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Gate** | **string** |  | 
**Trigger** | **string** |  | 
**Action** | **string** |  | 
**Params** | Pointer to [**[]PolicyRuleParams**](PolicyRuleParams.md) |  | [optional] 
**Recommendation** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyRule

`func NewPolicyRule(gate string, trigger string, action string, ) *PolicyRule`

NewPolicyRule instantiates a new PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRuleWithDefaults

`func NewPolicyRuleWithDefaults() *PolicyRule`

NewPolicyRuleWithDefaults instantiates a new PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGate

`func (o *PolicyRule) GetGate() string`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *PolicyRule) GetGateOk() (*string, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *PolicyRule) SetGate(v string)`

SetGate sets Gate field to given value.


### GetTrigger

`func (o *PolicyRule) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PolicyRule) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PolicyRule) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetAction

`func (o *PolicyRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetParams

`func (o *PolicyRule) GetParams() []PolicyRuleParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *PolicyRule) GetParamsOk() (*[]PolicyRuleParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *PolicyRule) SetParams(v []PolicyRuleParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *PolicyRule) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetRecommendation

`func (o *PolicyRule) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *PolicyRule) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *PolicyRule) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *PolicyRule) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


