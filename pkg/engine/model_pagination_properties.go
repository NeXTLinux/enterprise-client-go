/*
Nextlinux Engine API Server

This is the Nextlinux Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// PaginationProperties Properties for common pagination handling to be included in any wrapping object that needs pagination elements
type PaginationProperties struct {
	// The page number returned (should match the requested page query string param)
	Page *string `json:"page,omitempty"`
	// True if additional pages exist (page + 1) or False if this is the last page
	NextPage *string `json:"next_page,omitempty"`
	// The number of items sent in this response
	ReturnedCount *int32 `json:"returned_count,omitempty"`
}

// NewPaginationProperties instantiates a new PaginationProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationProperties() *PaginationProperties {
	this := PaginationProperties{}
	return &this
}

// NewPaginationPropertiesWithDefaults instantiates a new PaginationProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationPropertiesWithDefaults() *PaginationProperties {
	this := PaginationProperties{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginationProperties) GetPage() string {
	if o == nil || o.Page == nil {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationProperties) GetPageOk() (*string, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginationProperties) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *PaginationProperties) SetPage(v string) {
	o.Page = &v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *PaginationProperties) GetNextPage() string {
	if o == nil || o.NextPage == nil {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationProperties) GetNextPageOk() (*string, bool) {
	if o == nil || o.NextPage == nil {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *PaginationProperties) HasNextPage() bool {
	if o != nil && o.NextPage != nil {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *PaginationProperties) SetNextPage(v string) {
	o.NextPage = &v
}

// GetReturnedCount returns the ReturnedCount field value if set, zero value otherwise.
func (o *PaginationProperties) GetReturnedCount() int32 {
	if o == nil || o.ReturnedCount == nil {
		var ret int32
		return ret
	}
	return *o.ReturnedCount
}

// GetReturnedCountOk returns a tuple with the ReturnedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationProperties) GetReturnedCountOk() (*int32, bool) {
	if o == nil || o.ReturnedCount == nil {
		return nil, false
	}
	return o.ReturnedCount, true
}

// HasReturnedCount returns a boolean if a field has been set.
func (o *PaginationProperties) HasReturnedCount() bool {
	if o != nil && o.ReturnedCount != nil {
		return true
	}

	return false
}

// SetReturnedCount gets a reference to the given int32 and assigns it to the ReturnedCount field.
func (o *PaginationProperties) SetReturnedCount(v int32) {
	o.ReturnedCount = &v
}

func (o PaginationProperties) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullablePaginationProperties struct {
	value *PaginationProperties
	isSet bool
}

func (v NullablePaginationProperties) Get() *PaginationProperties {
	return v.value
}

func (v *NullablePaginationProperties) Set(val *PaginationProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationProperties(val *PaginationProperties) *NullablePaginationProperties {
	return &NullablePaginationProperties{value: val, isSet: true}
}

func (v NullablePaginationProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


