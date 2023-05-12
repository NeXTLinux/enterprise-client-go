# ContentFilesResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** |  | [optional] 
**Gid** | Pointer to **int32** |  | [optional] 
**Linkdest** | Pointer to **NullableString** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Sha256** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **int32** |  | [optional] 

## Methods

### NewContentFilesResponseContent

`func NewContentFilesResponseContent() *ContentFilesResponseContent`

NewContentFilesResponseContent instantiates a new ContentFilesResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentFilesResponseContentWithDefaults

`func NewContentFilesResponseContentWithDefaults() *ContentFilesResponseContent`

NewContentFilesResponseContentWithDefaults instantiates a new ContentFilesResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *ContentFilesResponseContent) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ContentFilesResponseContent) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ContentFilesResponseContent) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ContentFilesResponseContent) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetGid

`func (o *ContentFilesResponseContent) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *ContentFilesResponseContent) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *ContentFilesResponseContent) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *ContentFilesResponseContent) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetLinkdest

`func (o *ContentFilesResponseContent) GetLinkdest() string`

GetLinkdest returns the Linkdest field if non-nil, zero value otherwise.

### GetLinkdestOk

`func (o *ContentFilesResponseContent) GetLinkdestOk() (*string, bool)`

GetLinkdestOk returns a tuple with the Linkdest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkdest

`func (o *ContentFilesResponseContent) SetLinkdest(v string)`

SetLinkdest sets Linkdest field to given value.

### HasLinkdest

`func (o *ContentFilesResponseContent) HasLinkdest() bool`

HasLinkdest returns a boolean if a field has been set.

### SetLinkdestNil

`func (o *ContentFilesResponseContent) SetLinkdestNil(b bool)`

 SetLinkdestNil sets the value for Linkdest to be an explicit nil

### UnsetLinkdest
`func (o *ContentFilesResponseContent) UnsetLinkdest()`

UnsetLinkdest ensures that no value is present for Linkdest, not even an explicit nil
### GetMode

`func (o *ContentFilesResponseContent) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ContentFilesResponseContent) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ContentFilesResponseContent) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ContentFilesResponseContent) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSha256

`func (o *ContentFilesResponseContent) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ContentFilesResponseContent) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ContentFilesResponseContent) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ContentFilesResponseContent) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *ContentFilesResponseContent) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *ContentFilesResponseContent) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetSize

`func (o *ContentFilesResponseContent) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentFilesResponseContent) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentFilesResponseContent) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentFilesResponseContent) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ContentFilesResponseContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentFilesResponseContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentFilesResponseContent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentFilesResponseContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *ContentFilesResponseContent) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ContentFilesResponseContent) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ContentFilesResponseContent) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ContentFilesResponseContent) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


