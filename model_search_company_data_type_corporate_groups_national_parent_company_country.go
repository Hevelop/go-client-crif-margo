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

// checks if the SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry{}

// SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry struct for SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry
type SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry struct {
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry instantiates a new SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry() *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry {
	this := SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry{}
	return &this
}

// NewSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountryWithDefaults instantiates a new SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountryWithDefaults() *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry {
	this := SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) SetDescription(v string) {
	o.Description = &v
}

func (o SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry struct {
	value *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry
	isSet bool
}

func (v NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) Get() *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry {
	return v.value
}

func (v *NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) Set(val *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry(val *SearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) *NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry {
	return &NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeCorporateGroupsNationalParentCompanyCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


