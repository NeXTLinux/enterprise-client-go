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

// PaginatedVulnerableImageList Pagination wrapped list of images with vulnerabilties that match some filter
type PaginatedVulnerableImageList struct {
	// The page number returned (should match the requested page query string param)
	Page *string `json:"page,omitempty"`
	// True if additional pages exist (page + 1) or False if this is the last page
	NextPage *string `json:"next_page,omitempty"`
	// The number of items sent in this response
	ReturnedCount *int32 `json:"returned_count,omitempty"`
	Images *[]VulnerableImage `json:"images,omitempty"`
}

// NewPaginatedVulnerableImageList instantiates a new PaginatedVulnerableImageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedVulnerableImageList() *PaginatedVulnerableImageList {
	this := PaginatedVulnerableImageList{}
	return &this
}

// NewPaginatedVulnerableImageListWithDefaults instantiates a new PaginatedVulnerableImageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedVulnerableImageListWithDefaults() *PaginatedVulnerableImageList {
	this := PaginatedVulnerableImageList{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedVulnerableImageList) GetPage() string {
	if o == nil || o.Page == nil {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedVulnerableImageList) GetPageOk() (*string, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedVulnerableImageList) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *PaginatedVulnerableImageList) SetPage(v string) {
	o.Page = &v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *PaginatedVulnerableImageList) GetNextPage() string {
	if o == nil || o.NextPage == nil {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedVulnerableImageList) GetNextPageOk() (*string, bool) {
	if o == nil || o.NextPage == nil {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *PaginatedVulnerableImageList) HasNextPage() bool {
	if o != nil && o.NextPage != nil {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *PaginatedVulnerableImageList) SetNextPage(v string) {
	o.NextPage = &v
}

// GetReturnedCount returns the ReturnedCount field value if set, zero value otherwise.
func (o *PaginatedVulnerableImageList) GetReturnedCount() int32 {
	if o == nil || o.ReturnedCount == nil {
		var ret int32
		return ret
	}
	return *o.ReturnedCount
}

// GetReturnedCountOk returns a tuple with the ReturnedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedVulnerableImageList) GetReturnedCountOk() (*int32, bool) {
	if o == nil || o.ReturnedCount == nil {
		return nil, false
	}
	return o.ReturnedCount, true
}

// HasReturnedCount returns a boolean if a field has been set.
func (o *PaginatedVulnerableImageList) HasReturnedCount() bool {
	if o != nil && o.ReturnedCount != nil {
		return true
	}

	return false
}

// SetReturnedCount gets a reference to the given int32 and assigns it to the ReturnedCount field.
func (o *PaginatedVulnerableImageList) SetReturnedCount(v int32) {
	o.ReturnedCount = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *PaginatedVulnerableImageList) GetImages() []VulnerableImage {
	if o == nil || o.Images == nil {
		var ret []VulnerableImage
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedVulnerableImageList) GetImagesOk() (*[]VulnerableImage, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *PaginatedVulnerableImageList) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []VulnerableImage and assigns it to the Images field.
func (o *PaginatedVulnerableImageList) SetImages(v []VulnerableImage) {
	o.Images = &v
}

func (o PaginatedVulnerableImageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.NextPage != nil {
		toSerialize["next_page"] = o.NextPage
	}
	if o.ReturnedCount != nil {
		toSerialize["returned_count"] = o.ReturnedCount
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedVulnerableImageList struct {
	value *PaginatedVulnerableImageList
	isSet bool
}

func (v NullablePaginatedVulnerableImageList) Get() *PaginatedVulnerableImageList {
	return v.value
}

func (v *NullablePaginatedVulnerableImageList) Set(val *PaginatedVulnerableImageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedVulnerableImageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedVulnerableImageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedVulnerableImageList(val *PaginatedVulnerableImageList) *NullablePaginatedVulnerableImageList {
	return &NullablePaginatedVulnerableImageList{value: val, isSet: true}
}

func (v NullablePaginatedVulnerableImageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedVulnerableImageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


