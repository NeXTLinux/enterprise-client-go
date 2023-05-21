# SBOMVulnerabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SbomId** | Pointer to **string** |  | [optional] 
**VulnerabilityType** | Pointer to **string** |  | [optional] 
**Vulnerabilities** | Pointer to [**[]Vulnerability**](Vulnerability.md) | List of Vulnerability objects | [optional] 

## Methods

### NewSBOMVulnerabilitiesResponse

`func NewSBOMVulnerabilitiesResponse() *SBOMVulnerabilitiesResponse`

NewSBOMVulnerabilitiesResponse instantiates a new SBOMVulnerabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSBOMVulnerabilitiesResponseWithDefaults

`func NewSBOMVulnerabilitiesResponseWithDefaults() *SBOMVulnerabilitiesResponse`

NewSBOMVulnerabilitiesResponseWithDefaults instantiates a new SBOMVulnerabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSbomId

`func (o *SBOMVulnerabilitiesResponse) GetSbomId() string`

GetSbomId returns the SbomId field if non-nil, zero value otherwise.

### GetSbomIdOk

`func (o *SBOMVulnerabilitiesResponse) GetSbomIdOk() (*string, bool)`

GetSbomIdOk returns a tuple with the SbomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbomId

`func (o *SBOMVulnerabilitiesResponse) SetSbomId(v string)`

SetSbomId sets SbomId field to given value.

### HasSbomId

`func (o *SBOMVulnerabilitiesResponse) HasSbomId() bool`

HasSbomId returns a boolean if a field has been set.

### GetVulnerabilityType

`func (o *SBOMVulnerabilitiesResponse) GetVulnerabilityType() string`

GetVulnerabilityType returns the VulnerabilityType field if non-nil, zero value otherwise.

### GetVulnerabilityTypeOk

`func (o *SBOMVulnerabilitiesResponse) GetVulnerabilityTypeOk() (*string, bool)`

GetVulnerabilityTypeOk returns a tuple with the VulnerabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityType

`func (o *SBOMVulnerabilitiesResponse) SetVulnerabilityType(v string)`

SetVulnerabilityType sets VulnerabilityType field to given value.

### HasVulnerabilityType

`func (o *SBOMVulnerabilitiesResponse) HasVulnerabilityType() bool`

HasVulnerabilityType returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *SBOMVulnerabilitiesResponse) GetVulnerabilities() []Vulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *SBOMVulnerabilitiesResponse) GetVulnerabilitiesOk() (*[]Vulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *SBOMVulnerabilitiesResponse) SetVulnerabilities(v []Vulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *SBOMVulnerabilitiesResponse) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


