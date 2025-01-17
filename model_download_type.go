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
	"bytes"
	"fmt"
)

// checks if the DownloadType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadType{}

// DownloadType It defines the ouput requested in term of what is the marketing list  or data packets.
type DownloadType struct {
	// Zero-based number of the page to obtain
	Size *float32 `json:"size,omitempty"`
	// It manages the number of calls and it governs pagination. continueToken is provided after first call by Margo and after it must be insert in the body. The continueToken changes after each call and it must refresh in the body.
	ContinueToken *string `json:"continueToken,omitempty"`
	Filters *FiltersType `json:"filters,omitempty"`
	PortfolioFilters *PortfolioFiltersType `json:"portfolioFilters,omitempty"`
	Content DownloadTypeContent `json:"content"`
}

type _DownloadType DownloadType

// NewDownloadType instantiates a new DownloadType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadType(content DownloadTypeContent) *DownloadType {
	this := DownloadType{}
	this.Content = content
	return &this
}

// NewDownloadTypeWithDefaults instantiates a new DownloadType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadTypeWithDefaults() *DownloadType {
	this := DownloadType{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DownloadType) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadType) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DownloadType) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *DownloadType) SetSize(v float32) {
	o.Size = &v
}

// GetContinueToken returns the ContinueToken field value if set, zero value otherwise.
func (o *DownloadType) GetContinueToken() string {
	if o == nil || IsNil(o.ContinueToken) {
		var ret string
		return ret
	}
	return *o.ContinueToken
}

// GetContinueTokenOk returns a tuple with the ContinueToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadType) GetContinueTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinueToken) {
		return nil, false
	}
	return o.ContinueToken, true
}

// HasContinueToken returns a boolean if a field has been set.
func (o *DownloadType) HasContinueToken() bool {
	if o != nil && !IsNil(o.ContinueToken) {
		return true
	}

	return false
}

// SetContinueToken gets a reference to the given string and assigns it to the ContinueToken field.
func (o *DownloadType) SetContinueToken(v string) {
	o.ContinueToken = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *DownloadType) GetFilters() FiltersType {
	if o == nil || IsNil(o.Filters) {
		var ret FiltersType
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadType) GetFiltersOk() (*FiltersType, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *DownloadType) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersType and assigns it to the Filters field.
func (o *DownloadType) SetFilters(v FiltersType) {
	o.Filters = &v
}

// GetPortfolioFilters returns the PortfolioFilters field value if set, zero value otherwise.
func (o *DownloadType) GetPortfolioFilters() PortfolioFiltersType {
	if o == nil || IsNil(o.PortfolioFilters) {
		var ret PortfolioFiltersType
		return ret
	}
	return *o.PortfolioFilters
}

// GetPortfolioFiltersOk returns a tuple with the PortfolioFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadType) GetPortfolioFiltersOk() (*PortfolioFiltersType, bool) {
	if o == nil || IsNil(o.PortfolioFilters) {
		return nil, false
	}
	return o.PortfolioFilters, true
}

// HasPortfolioFilters returns a boolean if a field has been set.
func (o *DownloadType) HasPortfolioFilters() bool {
	if o != nil && !IsNil(o.PortfolioFilters) {
		return true
	}

	return false
}

// SetPortfolioFilters gets a reference to the given PortfolioFiltersType and assigns it to the PortfolioFilters field.
func (o *DownloadType) SetPortfolioFilters(v PortfolioFiltersType) {
	o.PortfolioFilters = &v
}

// GetContent returns the Content field value
func (o *DownloadType) GetContent() DownloadTypeContent {
	if o == nil {
		var ret DownloadTypeContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *DownloadType) GetContentOk() (*DownloadTypeContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *DownloadType) SetContent(v DownloadTypeContent) {
	o.Content = v
}

func (o DownloadType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.ContinueToken) {
		toSerialize["continueToken"] = o.ContinueToken
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.PortfolioFilters) {
		toSerialize["portfolioFilters"] = o.PortfolioFilters
	}
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

func (o *DownloadType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDownloadType := _DownloadType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDownloadType)

	if err != nil {
		return err
	}

	*o = DownloadType(varDownloadType)

	return err
}

type NullableDownloadType struct {
	value *DownloadType
	isSet bool
}

func (v NullableDownloadType) Get() *DownloadType {
	return v.value
}

func (v *NullableDownloadType) Set(val *DownloadType) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadType) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadType(val *DownloadType) *NullableDownloadType {
	return &NullableDownloadType{value: val, isSet: true}
}

func (v NullableDownloadType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


