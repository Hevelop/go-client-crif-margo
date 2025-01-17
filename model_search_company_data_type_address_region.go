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

// checks if the SearchCompanyDataTypeAddressRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeAddressRegion{}

// SearchCompanyDataTypeAddressRegion struct for SearchCompanyDataTypeAddressRegion
type SearchCompanyDataTypeAddressRegion struct {
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewSearchCompanyDataTypeAddressRegion instantiates a new SearchCompanyDataTypeAddressRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeAddressRegion() *SearchCompanyDataTypeAddressRegion {
	this := SearchCompanyDataTypeAddressRegion{}
	return &this
}

// NewSearchCompanyDataTypeAddressRegionWithDefaults instantiates a new SearchCompanyDataTypeAddressRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeAddressRegionWithDefaults() *SearchCompanyDataTypeAddressRegion {
	this := SearchCompanyDataTypeAddressRegion{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddressRegion) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddressRegion) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddressRegion) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SearchCompanyDataTypeAddressRegion) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddressRegion) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddressRegion) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddressRegion) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchCompanyDataTypeAddressRegion) SetDescription(v string) {
	o.Description = &v
}

func (o SearchCompanyDataTypeAddressRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeAddressRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeAddressRegion struct {
	value *SearchCompanyDataTypeAddressRegion
	isSet bool
}

func (v NullableSearchCompanyDataTypeAddressRegion) Get() *SearchCompanyDataTypeAddressRegion {
	return v.value
}

func (v *NullableSearchCompanyDataTypeAddressRegion) Set(val *SearchCompanyDataTypeAddressRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeAddressRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeAddressRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeAddressRegion(val *SearchCompanyDataTypeAddressRegion) *NullableSearchCompanyDataTypeAddressRegion {
	return &NullableSearchCompanyDataTypeAddressRegion{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeAddressRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeAddressRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


