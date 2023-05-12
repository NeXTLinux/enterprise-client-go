# NextlinuxImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageContent** | Pointer to **interface{}** | A metadata content record for a specific image, containing different content type entries | [optional] 
**ImageDetail** | Pointer to [**[]ImageDetail**](ImageDetail.md) | Details specific to an image reference and type such as tag and image source | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**ParentDigest** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **interface{}** |  | [optional] 
**ImageStatus** | Pointer to **string** | State of the image | [optional] 
**AnalysisStatus** | Pointer to **string** | A state value for the current status of the analysis progress of the image | [optional] 
**RecordVersion** | Pointer to **string** | The version of the record, used for internal schema updates and data migrations. | [optional] 
**AnalysisStatusDetail** | Pointer to [**[]AnalysisStatusDetail**](AnalysisStatusDetail.md) |  | [optional] 

## Methods

### NewNextlinuxImage

`func NewNextlinuxImage() *NextlinuxImage`

NewNextlinuxImage instantiates a new NextlinuxImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextlinuxImageWithDefaults

`func NewNextlinuxImageWithDefaults() *NextlinuxImage`

NewNextlinuxImageWithDefaults instantiates a new NextlinuxImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageContent

`func (o *NextlinuxImage) GetImageContent() interface{}`

GetImageContent returns the ImageContent field if non-nil, zero value otherwise.

### GetImageContentOk

`func (o *NextlinuxImage) GetImageContentOk() (*interface{}, bool)`

GetImageContentOk returns a tuple with the ImageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContent

`func (o *NextlinuxImage) SetImageContent(v interface{})`

SetImageContent sets ImageContent field to given value.

### HasImageContent

`func (o *NextlinuxImage) HasImageContent() bool`

HasImageContent returns a boolean if a field has been set.

### GetImageDetail

`func (o *NextlinuxImage) GetImageDetail() []ImageDetail`

GetImageDetail returns the ImageDetail field if non-nil, zero value otherwise.

### GetImageDetailOk

`func (o *NextlinuxImage) GetImageDetailOk() (*[]ImageDetail, bool)`

GetImageDetailOk returns a tuple with the ImageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDetail

`func (o *NextlinuxImage) SetImageDetail(v []ImageDetail)`

SetImageDetail sets ImageDetail field to given value.

### HasImageDetail

`func (o *NextlinuxImage) HasImageDetail() bool`

HasImageDetail returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NextlinuxImage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NextlinuxImage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NextlinuxImage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NextlinuxImage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NextlinuxImage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NextlinuxImage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NextlinuxImage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NextlinuxImage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetImageDigest

`func (o *NextlinuxImage) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *NextlinuxImage) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *NextlinuxImage) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *NextlinuxImage) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetParentDigest

`func (o *NextlinuxImage) GetParentDigest() string`

GetParentDigest returns the ParentDigest field if non-nil, zero value otherwise.

### GetParentDigestOk

`func (o *NextlinuxImage) GetParentDigestOk() (*string, bool)`

GetParentDigestOk returns a tuple with the ParentDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDigest

`func (o *NextlinuxImage) SetParentDigest(v string)`

SetParentDigest sets ParentDigest field to given value.

### HasParentDigest

`func (o *NextlinuxImage) HasParentDigest() bool`

HasParentDigest returns a boolean if a field has been set.

### GetUserId

`func (o *NextlinuxImage) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *NextlinuxImage) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *NextlinuxImage) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *NextlinuxImage) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAnnotations

`func (o *NextlinuxImage) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *NextlinuxImage) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *NextlinuxImage) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *NextlinuxImage) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetImageStatus

`func (o *NextlinuxImage) GetImageStatus() string`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *NextlinuxImage) GetImageStatusOk() (*string, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *NextlinuxImage) SetImageStatus(v string)`

SetImageStatus sets ImageStatus field to given value.

### HasImageStatus

`func (o *NextlinuxImage) HasImageStatus() bool`

HasImageStatus returns a boolean if a field has been set.

### GetAnalysisStatus

`func (o *NextlinuxImage) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *NextlinuxImage) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *NextlinuxImage) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *NextlinuxImage) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NextlinuxImage) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NextlinuxImage) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NextlinuxImage) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NextlinuxImage) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetAnalysisStatusDetail

`func (o *NextlinuxImage) GetAnalysisStatusDetail() []AnalysisStatusDetail`

GetAnalysisStatusDetail returns the AnalysisStatusDetail field if non-nil, zero value otherwise.

### GetAnalysisStatusDetailOk

`func (o *NextlinuxImage) GetAnalysisStatusDetailOk() (*[]AnalysisStatusDetail, bool)`

GetAnalysisStatusDetailOk returns a tuple with the AnalysisStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatusDetail

`func (o *NextlinuxImage) SetAnalysisStatusDetail(v []AnalysisStatusDetail)`

SetAnalysisStatusDetail sets AnalysisStatusDetail field to given value.

### HasAnalysisStatusDetail

`func (o *NextlinuxImage) HasAnalysisStatusDetail() bool`

HasAnalysisStatusDetail returns a boolean if a field has been set.

### SetAnalysisStatusDetailNil

`func (o *NextlinuxImage) SetAnalysisStatusDetailNil(b bool)`

 SetAnalysisStatusDetailNil sets the value for AnalysisStatusDetail to be an explicit nil

### UnsetAnalysisStatusDetail
`func (o *NextlinuxImage) UnsetAnalysisStatusDetail()`

UnsetAnalysisStatusDetail ensures that no value is present for AnalysisStatusDetail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


