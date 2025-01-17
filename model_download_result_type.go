/*
Global Marketing Solution

## Overview

API version: 1.0.0
Contact: PM_Margo@crif.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crifmargo

import (
	"encoding/json"
)

// checks if the DownloadResultType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadResultType{}

// DownloadResultType companies list with the Marketing List or datapackets requested
type DownloadResultType struct {
	// Total number of elements the API should return without pagination
	TotalElements *int32 `json:"totalElements,omitempty"`
	// Total number of pages available `totalPages = ceil (totalElements / size)`
	TotalPages *int64 `json:"totalPages,omitempty"`
	// maximum number of elements inside a page (like input)
	Size *int64 `json:"size,omitempty"`
	// page number (like input)
	Page *int64 `json:"page,omitempty"`
	// number of elements in this page `numberOfElements <= size`
	NumberOfElements *int64 `json:"numberOfElements,omitempty"`
	Content []DownloadCompanyDataType `json:"content,omitempty"`
}

// NewDownloadResultType instantiates a new DownloadResultType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadResultType() *DownloadResultType {
	this := DownloadResultType{}
	return &this
}

// NewDownloadResultTypeWithDefaults instantiates a new DownloadResultType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadResultTypeWithDefaults() *DownloadResultType {
	this := DownloadResultType{}
	return &this
}

// GetTotalElements returns the TotalElements field value if set, zero value otherwise.
func (o *DownloadResultType) GetTotalElements() int32 {
	if o == nil || IsNil(o.TotalElements) {
		var ret int32
		return ret
	}
	return *o.TotalElements
}

// GetTotalElementsOk returns a tuple with the TotalElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadResultType) GetTotalElementsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalElements) {
		return nil, false
	}
	return o.TotalElements, true
}

// HasTotalElements returns a boolean if a field has been set.
func (o *DownloadResultType) HasTotalElements() bool {
	if o != nil && !IsNil(o.TotalElements) {
		return true
	}

	return false
}

// SetTotalElements gets a reference to the given int32 and assigns it to the TotalElements field.
func (o *DownloadResultType) SetTotalElements(v int32) {
	o.TotalElements = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *DownloadResultType) GetTotalPages() int64 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int64
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadResultType) GetTotalPagesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *DownloadResultType) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int64 and assigns it to the TotalPages field.
func (o *DownloadResultType) SetTotalPages(v int64) {
	o.TotalPages = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DownloadResultType) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadResultType) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DownloadResultType) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *DownloadResultType) SetSize(v int64) {
	o.Size = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DownloadResultType) GetPage() int64 {
	if o == nil || IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadResultType) GetPageOk() (*int64, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DownloadResultType) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *DownloadResultType) SetPage(v int64) {
	o.Page = &v
}

// GetNumberOfElements returns the NumberOfElements field value if set, zero value otherwise.
func (o *DownloadResultType) GetNumberOfElements() int64 {
	if o == nil || IsNil(o.NumberOfElements) {
		var ret int64
		return ret
	}
	return *o.NumberOfElements
}

// GetNumberOfElementsOk returns a tuple with the NumberOfElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadResultType) GetNumberOfElementsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfElements) {
		return nil, false
	}
	return o.NumberOfElements, true
}

// HasNumberOfElements returns a boolean if a field has been set.
func (o *DownloadResultType) HasNumberOfElements() bool {
	if o != nil && !IsNil(o.NumberOfElements) {
		return true
	}

	return false
}

// SetNumberOfElements gets a reference to the given int64 and assigns it to the NumberOfElements field.
func (o *DownloadResultType) SetNumberOfElements(v int64) {
	o.NumberOfElements = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DownloadResultType) GetContent() []DownloadCompanyDataType {
	if o == nil || IsNil(o.Content) {
		var ret []DownloadCompanyDataType
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadResultType) GetContentOk() ([]DownloadCompanyDataType, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DownloadResultType) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []DownloadCompanyDataType and assigns it to the Content field.
func (o *DownloadResultType) SetContent(v []DownloadCompanyDataType) {
	o.Content = v
}

func (o DownloadResultType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadResultType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalElements) {
		toSerialize["totalElements"] = o.TotalElements
	}
	if !IsNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.NumberOfElements) {
		toSerialize["numberOfElements"] = o.NumberOfElements
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableDownloadResultType struct {
	value *DownloadResultType
	isSet bool
}

func (v NullableDownloadResultType) Get() *DownloadResultType {
	return v.value
}

func (v *NullableDownloadResultType) Set(val *DownloadResultType) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadResultType) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadResultType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadResultType(val *DownloadResultType) *NullableDownloadResultType {
	return &NullableDownloadResultType{value: val, isSet: true}
}

func (v NullableDownloadResultType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadResultType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


