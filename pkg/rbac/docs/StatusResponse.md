# StatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Busy** | Pointer to **bool** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewStatusResponse

`func NewStatusResponse() *StatusResponse`

NewStatusResponse instantiates a new StatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseWithDefaults

`func NewStatusResponseWithDefaults() *StatusResponse`

NewStatusResponseWithDefaults instantiates a new StatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusy

`func (o *StatusResponse) GetBusy() bool`

GetBusy returns the Busy field if non-nil, zero value otherwise.

### GetBusyOk

`func (o *StatusResponse) GetBusyOk() (*bool, bool)`

GetBusyOk returns a tuple with the Busy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusy

`func (o *StatusResponse) SetBusy(v bool)`

SetBusy sets Busy field to given value.

### HasBusy

`func (o *StatusResponse) HasBusy() bool`

HasBusy returns a boolean if a field has been set.

### GetUp

`func (o *StatusResponse) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *StatusResponse) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *StatusResponse) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *StatusResponse) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetMessage

`func (o *StatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


