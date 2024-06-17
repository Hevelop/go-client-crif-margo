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

// checks if the SearchCompanyDataTypeAtecoClassification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeAtecoClassification{}

// SearchCompanyDataTypeAtecoClassification struct for SearchCompanyDataTypeAtecoClassification
type SearchCompanyDataTypeAtecoClassification struct {
	Ateco *SearchCompanyDataTypeAteco `json:"ateco,omitempty"`
	SecondaryAteco *string `json:"secondaryAteco,omitempty"`
}

// NewSearchCompanyDataTypeAtecoClassification instantiates a new SearchCompanyDataTypeAtecoClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeAtecoClassification() *SearchCompanyDataTypeAtecoClassification {
	this := SearchCompanyDataTypeAtecoClassification{}
	return &this
}

// NewSearchCompanyDataTypeAtecoClassificationWithDefaults instantiates a new SearchCompanyDataTypeAtecoClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeAtecoClassificationWithDefaults() *SearchCompanyDataTypeAtecoClassification {
	this := SearchCompanyDataTypeAtecoClassification{}
	return &this
}

// GetAteco returns the Ateco field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAtecoClassification) GetAteco() SearchCompanyDataTypeAteco {
	if o == nil || IsNil(o.Ateco) {
		var ret SearchCompanyDataTypeAteco
		return ret
	}
	return *o.Ateco
}

// GetAtecoOk returns a tuple with the Ateco field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAtecoClassification) GetAtecoOk() (*SearchCompanyDataTypeAteco, bool) {
	if o == nil || IsNil(o.Ateco) {
		return nil, false
	}
	return o.Ateco, true
}

// HasAteco returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAtecoClassification) HasAteco() bool {
	if o != nil && !IsNil(o.Ateco) {
		return true
	}

	return false
}

// SetAteco gets a reference to the given SearchCompanyDataTypeAteco and assigns it to the Ateco field.
func (o *SearchCompanyDataTypeAtecoClassification) SetAteco(v SearchCompanyDataTypeAteco) {
	o.Ateco = &v
}

// GetSecondaryAteco returns the SecondaryAteco field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAtecoClassification) GetSecondaryAteco() string {
	if o == nil || IsNil(o.SecondaryAteco) {
		var ret string
		return ret
	}
	return *o.SecondaryAteco
}

// GetSecondaryAtecoOk returns a tuple with the SecondaryAteco field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAtecoClassification) GetSecondaryAtecoOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryAteco) {
		return nil, false
	}
	return o.SecondaryAteco, true
}

// HasSecondaryAteco returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAtecoClassification) HasSecondaryAteco() bool {
	if o != nil && !IsNil(o.SecondaryAteco) {
		return true
	}

	return false
}

// SetSecondaryAteco gets a reference to the given string and assigns it to the SecondaryAteco field.
func (o *SearchCompanyDataTypeAtecoClassification) SetSecondaryAteco(v string) {
	o.SecondaryAteco = &v
}

func (o SearchCompanyDataTypeAtecoClassification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeAtecoClassification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ateco) {
		toSerialize["ateco"] = o.Ateco
	}
	if !IsNil(o.SecondaryAteco) {
		toSerialize["secondaryAteco"] = o.SecondaryAteco
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeAtecoClassification struct {
	value *SearchCompanyDataTypeAtecoClassification
	isSet bool
}

func (v NullableSearchCompanyDataTypeAtecoClassification) Get() *SearchCompanyDataTypeAtecoClassification {
	return v.value
}

func (v *NullableSearchCompanyDataTypeAtecoClassification) Set(val *SearchCompanyDataTypeAtecoClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeAtecoClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeAtecoClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeAtecoClassification(val *SearchCompanyDataTypeAtecoClassification) *NullableSearchCompanyDataTypeAtecoClassification {
	return &NullableSearchCompanyDataTypeAtecoClassification{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeAtecoClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeAtecoClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

