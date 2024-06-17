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

// checks if the SearchCompanyDataTypeTelcoAnalyticsLandLineScore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeTelcoAnalyticsLandLineScore{}

// SearchCompanyDataTypeTelcoAnalyticsLandLineScore struct for SearchCompanyDataTypeTelcoAnalyticsLandLineScore
type SearchCompanyDataTypeTelcoAnalyticsLandLineScore struct {
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewSearchCompanyDataTypeTelcoAnalyticsLandLineScore instantiates a new SearchCompanyDataTypeTelcoAnalyticsLandLineScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeTelcoAnalyticsLandLineScore() *SearchCompanyDataTypeTelcoAnalyticsLandLineScore {
	this := SearchCompanyDataTypeTelcoAnalyticsLandLineScore{}
	return &this
}

// NewSearchCompanyDataTypeTelcoAnalyticsLandLineScoreWithDefaults instantiates a new SearchCompanyDataTypeTelcoAnalyticsLandLineScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeTelcoAnalyticsLandLineScoreWithDefaults() *SearchCompanyDataTypeTelcoAnalyticsLandLineScore {
	this := SearchCompanyDataTypeTelcoAnalyticsLandLineScore{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) SetDescription(v string) {
	o.Description = &v
}

func (o SearchCompanyDataTypeTelcoAnalyticsLandLineScore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeTelcoAnalyticsLandLineScore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore struct {
	value *SearchCompanyDataTypeTelcoAnalyticsLandLineScore
	isSet bool
}

func (v NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore) Get() *SearchCompanyDataTypeTelcoAnalyticsLandLineScore {
	return v.value
}

func (v *NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore) Set(val *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore(val *SearchCompanyDataTypeTelcoAnalyticsLandLineScore) *NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore {
	return &NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeTelcoAnalyticsLandLineScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

