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

// checks if the SearchCompanyDataTypeManagersInnerGender type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeManagersInnerGender{}

// SearchCompanyDataTypeManagersInnerGender struct for SearchCompanyDataTypeManagersInnerGender
type SearchCompanyDataTypeManagersInnerGender struct {
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewSearchCompanyDataTypeManagersInnerGender instantiates a new SearchCompanyDataTypeManagersInnerGender object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeManagersInnerGender() *SearchCompanyDataTypeManagersInnerGender {
	this := SearchCompanyDataTypeManagersInnerGender{}
	return &this
}

// NewSearchCompanyDataTypeManagersInnerGenderWithDefaults instantiates a new SearchCompanyDataTypeManagersInnerGender object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeManagersInnerGenderWithDefaults() *SearchCompanyDataTypeManagersInnerGender {
	this := SearchCompanyDataTypeManagersInnerGender{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInnerGender) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInnerGender) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInnerGender) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SearchCompanyDataTypeManagersInnerGender) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInnerGender) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInnerGender) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInnerGender) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchCompanyDataTypeManagersInnerGender) SetDescription(v string) {
	o.Description = &v
}

func (o SearchCompanyDataTypeManagersInnerGender) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeManagersInnerGender) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeManagersInnerGender struct {
	value *SearchCompanyDataTypeManagersInnerGender
	isSet bool
}

func (v NullableSearchCompanyDataTypeManagersInnerGender) Get() *SearchCompanyDataTypeManagersInnerGender {
	return v.value
}

func (v *NullableSearchCompanyDataTypeManagersInnerGender) Set(val *SearchCompanyDataTypeManagersInnerGender) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeManagersInnerGender) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeManagersInnerGender) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeManagersInnerGender(val *SearchCompanyDataTypeManagersInnerGender) *NullableSearchCompanyDataTypeManagersInnerGender {
	return &NullableSearchCompanyDataTypeManagersInnerGender{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeManagersInnerGender) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeManagersInnerGender) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

