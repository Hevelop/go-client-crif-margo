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

// checks if the SearchCompanyDataTypeDigitalAttitudeScore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeDigitalAttitudeScore{}

// SearchCompanyDataTypeDigitalAttitudeScore struct for SearchCompanyDataTypeDigitalAttitudeScore
type SearchCompanyDataTypeDigitalAttitudeScore struct {
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewSearchCompanyDataTypeDigitalAttitudeScore instantiates a new SearchCompanyDataTypeDigitalAttitudeScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeDigitalAttitudeScore() *SearchCompanyDataTypeDigitalAttitudeScore {
	this := SearchCompanyDataTypeDigitalAttitudeScore{}
	return &this
}

// NewSearchCompanyDataTypeDigitalAttitudeScoreWithDefaults instantiates a new SearchCompanyDataTypeDigitalAttitudeScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeDigitalAttitudeScoreWithDefaults() *SearchCompanyDataTypeDigitalAttitudeScore {
	this := SearchCompanyDataTypeDigitalAttitudeScore{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeDigitalAttitudeScore) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeDigitalAttitudeScore) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeDigitalAttitudeScore) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SearchCompanyDataTypeDigitalAttitudeScore) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeDigitalAttitudeScore) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeDigitalAttitudeScore) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeDigitalAttitudeScore) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchCompanyDataTypeDigitalAttitudeScore) SetDescription(v string) {
	o.Description = &v
}

func (o SearchCompanyDataTypeDigitalAttitudeScore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeDigitalAttitudeScore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeDigitalAttitudeScore struct {
	value *SearchCompanyDataTypeDigitalAttitudeScore
	isSet bool
}

func (v NullableSearchCompanyDataTypeDigitalAttitudeScore) Get() *SearchCompanyDataTypeDigitalAttitudeScore {
	return v.value
}

func (v *NullableSearchCompanyDataTypeDigitalAttitudeScore) Set(val *SearchCompanyDataTypeDigitalAttitudeScore) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeDigitalAttitudeScore) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeDigitalAttitudeScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeDigitalAttitudeScore(val *SearchCompanyDataTypeDigitalAttitudeScore) *NullableSearchCompanyDataTypeDigitalAttitudeScore {
	return &NullableSearchCompanyDataTypeDigitalAttitudeScore{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeDigitalAttitudeScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeDigitalAttitudeScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


