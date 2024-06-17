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

// checks if the SearchCompanyDataTypeFinancialBurden type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeFinancialBurden{}

// SearchCompanyDataTypeFinancialBurden struct for SearchCompanyDataTypeFinancialBurden
type SearchCompanyDataTypeFinancialBurden struct {
	Rod *float32 `json:"rod,omitempty"`
	RodFinanziario *float32 `json:"rodFinanziario,omitempty"`
	BurdenIndex *float32 `json:"burdenIndex,omitempty"`
}

// NewSearchCompanyDataTypeFinancialBurden instantiates a new SearchCompanyDataTypeFinancialBurden object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeFinancialBurden() *SearchCompanyDataTypeFinancialBurden {
	this := SearchCompanyDataTypeFinancialBurden{}
	return &this
}

// NewSearchCompanyDataTypeFinancialBurdenWithDefaults instantiates a new SearchCompanyDataTypeFinancialBurden object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeFinancialBurdenWithDefaults() *SearchCompanyDataTypeFinancialBurden {
	this := SearchCompanyDataTypeFinancialBurden{}
	return &this
}

// GetRod returns the Rod field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeFinancialBurden) GetRod() float32 {
	if o == nil || IsNil(o.Rod) {
		var ret float32
		return ret
	}
	return *o.Rod
}

// GetRodOk returns a tuple with the Rod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeFinancialBurden) GetRodOk() (*float32, bool) {
	if o == nil || IsNil(o.Rod) {
		return nil, false
	}
	return o.Rod, true
}

// HasRod returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeFinancialBurden) HasRod() bool {
	if o != nil && !IsNil(o.Rod) {
		return true
	}

	return false
}

// SetRod gets a reference to the given float32 and assigns it to the Rod field.
func (o *SearchCompanyDataTypeFinancialBurden) SetRod(v float32) {
	o.Rod = &v
}

// GetRodFinanziario returns the RodFinanziario field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeFinancialBurden) GetRodFinanziario() float32 {
	if o == nil || IsNil(o.RodFinanziario) {
		var ret float32
		return ret
	}
	return *o.RodFinanziario
}

// GetRodFinanziarioOk returns a tuple with the RodFinanziario field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeFinancialBurden) GetRodFinanziarioOk() (*float32, bool) {
	if o == nil || IsNil(o.RodFinanziario) {
		return nil, false
	}
	return o.RodFinanziario, true
}

// HasRodFinanziario returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeFinancialBurden) HasRodFinanziario() bool {
	if o != nil && !IsNil(o.RodFinanziario) {
		return true
	}

	return false
}

// SetRodFinanziario gets a reference to the given float32 and assigns it to the RodFinanziario field.
func (o *SearchCompanyDataTypeFinancialBurden) SetRodFinanziario(v float32) {
	o.RodFinanziario = &v
}

// GetBurdenIndex returns the BurdenIndex field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeFinancialBurden) GetBurdenIndex() float32 {
	if o == nil || IsNil(o.BurdenIndex) {
		var ret float32
		return ret
	}
	return *o.BurdenIndex
}

// GetBurdenIndexOk returns a tuple with the BurdenIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeFinancialBurden) GetBurdenIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.BurdenIndex) {
		return nil, false
	}
	return o.BurdenIndex, true
}

// HasBurdenIndex returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeFinancialBurden) HasBurdenIndex() bool {
	if o != nil && !IsNil(o.BurdenIndex) {
		return true
	}

	return false
}

// SetBurdenIndex gets a reference to the given float32 and assigns it to the BurdenIndex field.
func (o *SearchCompanyDataTypeFinancialBurden) SetBurdenIndex(v float32) {
	o.BurdenIndex = &v
}

func (o SearchCompanyDataTypeFinancialBurden) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeFinancialBurden) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rod) {
		toSerialize["rod"] = o.Rod
	}
	if !IsNil(o.RodFinanziario) {
		toSerialize["rodFinanziario"] = o.RodFinanziario
	}
	if !IsNil(o.BurdenIndex) {
		toSerialize["burdenIndex"] = o.BurdenIndex
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeFinancialBurden struct {
	value *SearchCompanyDataTypeFinancialBurden
	isSet bool
}

func (v NullableSearchCompanyDataTypeFinancialBurden) Get() *SearchCompanyDataTypeFinancialBurden {
	return v.value
}

func (v *NullableSearchCompanyDataTypeFinancialBurden) Set(val *SearchCompanyDataTypeFinancialBurden) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeFinancialBurden) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeFinancialBurden) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeFinancialBurden(val *SearchCompanyDataTypeFinancialBurden) *NullableSearchCompanyDataTypeFinancialBurden {
	return &NullableSearchCompanyDataTypeFinancialBurden{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeFinancialBurden) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeFinancialBurden) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


