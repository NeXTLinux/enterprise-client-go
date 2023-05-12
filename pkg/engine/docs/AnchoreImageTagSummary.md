# NextlinuxImageTagSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | Pointer to **string** |  | [optional] 
**ParentDigest** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**AnalysisStatus** | Pointer to **string** |  | [optional] 
**Fulltag** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**AnalyzedAt** | Pointer to **int32** |  | [optional] 
**TagDetectedAt** | Pointer to **int32** |  | [optional] 
**ImageStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewNextlinuxImageTagSummary

`func NewNextlinuxImageTagSummary() *NextlinuxImageTagSummary`

NewNextlinuxImageTagSummary instantiates a new NextlinuxImageTagSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextlinuxImageTagSummaryWithDefaults

`func NewNextlinuxImageTagSummaryWithDefaults() *NextlinuxImageTagSummary`

NewNextlinuxImageTagSummaryWithDefaults instantiates a new NextlinuxImageTagSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDigest

`func (o *NextlinuxImageTagSummary) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *NextlinuxImageTagSummary) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *NextlinuxImageTagSummary) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *NextlinuxImageTagSummary) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetParentDigest

`func (o *NextlinuxImageTagSummary) GetParentDigest() string`

GetParentDigest returns the ParentDigest field if non-nil, zero value otherwise.

### GetParentDigestOk

`func (o *NextlinuxImageTagSummary) GetParentDigestOk() (*string, bool)`

GetParentDigestOk returns a tuple with the ParentDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDigest

`func (o *NextlinuxImageTagSummary) SetParentDigest(v string)`

SetParentDigest sets ParentDigest field to given value.

### HasParentDigest

`func (o *NextlinuxImageTagSummary) HasParentDigest() bool`

HasParentDigest returns a boolean if a field has been set.

### GetImageId

`func (o *NextlinuxImageTagSummary) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *NextlinuxImageTagSummary) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *NextlinuxImageTagSummary) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *NextlinuxImageTagSummary) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetAnalysisStatus

`func (o *NextlinuxImageTagSummary) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *NextlinuxImageTagSummary) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *NextlinuxImageTagSummary) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *NextlinuxImageTagSummary) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetFulltag

`func (o *NextlinuxImageTagSummary) GetFulltag() string`

GetFulltag returns the Fulltag field if non-nil, zero value otherwise.

### GetFulltagOk

`func (o *NextlinuxImageTagSummary) GetFulltagOk() (*string, bool)`

GetFulltagOk returns a tuple with the Fulltag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltag

`func (o *NextlinuxImageTagSummary) SetFulltag(v string)`

SetFulltag sets Fulltag field to given value.

### HasFulltag

`func (o *NextlinuxImageTagSummary) HasFulltag() bool`

HasFulltag returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NextlinuxImageTagSummary) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NextlinuxImageTagSummary) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NextlinuxImageTagSummary) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NextlinuxImageTagSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAnalyzedAt

`func (o *NextlinuxImageTagSummary) GetAnalyzedAt() int32`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *NextlinuxImageTagSummary) GetAnalyzedAtOk() (*int32, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *NextlinuxImageTagSummary) SetAnalyzedAt(v int32)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *NextlinuxImageTagSummary) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### GetTagDetectedAt

`func (o *NextlinuxImageTagSummary) GetTagDetectedAt() int32`

GetTagDetectedAt returns the TagDetectedAt field if non-nil, zero value otherwise.

### GetTagDetectedAtOk

`func (o *NextlinuxImageTagSummary) GetTagDetectedAtOk() (*int32, bool)`

GetTagDetectedAtOk returns a tuple with the TagDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagDetectedAt

`func (o *NextlinuxImageTagSummary) SetTagDetectedAt(v int32)`

SetTagDetectedAt sets TagDetectedAt field to given value.

### HasTagDetectedAt

`func (o *NextlinuxImageTagSummary) HasTagDetectedAt() bool`

HasTagDetectedAt returns a boolean if a field has been set.

### GetImageStatus

`func (o *NextlinuxImageTagSummary) GetImageStatus() string`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *NextlinuxImageTagSummary) GetImageStatusOk() (*string, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *NextlinuxImageTagSummary) SetImageStatus(v string)`

SetImageStatus sets ImageStatus field to given value.

### HasImageStatus

`func (o *NextlinuxImageTagSummary) HasImageStatus() bool`

HasImageStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


