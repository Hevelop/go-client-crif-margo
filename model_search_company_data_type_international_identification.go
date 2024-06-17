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

// checks if the SearchCompanyDataTypeInternationalIdentification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeInternationalIdentification{}

// SearchCompanyDataTypeInternationalIdentification struct for SearchCompanyDataTypeInternationalIdentification
type SearchCompanyDataTypeInternationalIdentification struct {
	HoldingDunsNumber *string `json:"holdingDunsNumber,omitempty"`
	NationalParentDunsNumber *string `json:"nationalParentDunsNumber,omitempty"`
}

// NewSearchCompanyDataTypeInternationalIdentification instantiates a new SearchCompanyDataTypeInternationalIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeInternationalIdentification() *SearchCompanyDataTypeInternationalIdentification {
	this := SearchCompanyDataTypeInternationalIdentification{}
	return &this
}

// NewSearchCompanyDataTypeInternationalIdentificationWithDefaults instantiates a new SearchCompanyDataTypeInternationalIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeInternationalIdentificationWithDefaults() *SearchCompanyDataTypeInternationalIdentification {
	this := SearchCompanyDataTypeInternationalIdentification{}
	return &this
}

// GetHoldingDunsNumber returns the HoldingDunsNumber field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeInternationalIdentification) GetHoldingDunsNumber() string {
	if o == nil || IsNil(o.HoldingDunsNumber) {
		var ret string
		return ret
	}
	return *o.HoldingDunsNumber
}

// GetHoldingDunsNumberOk returns a tuple with the HoldingDunsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeInternationalIdentification) GetHoldingDunsNumberOk() (*string, bool) {
	if o == nil || IsNil(o.HoldingDunsNumber) {
		return nil, false
	}
	return o.HoldingDunsNumber, true
}

// HasHoldingDunsNumber returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeInternationalIdentification) HasHoldingDunsNumber() bool {
	if o != nil && !IsNil(o.HoldingDunsNumber) {
		return true
	}

	return false
}

// SetHoldingDunsNumber gets a reference to the given string and assigns it to the HoldingDunsNumber field.
func (o *SearchCompanyDataTypeInternationalIdentification) SetHoldingDunsNumber(v string) {
	o.HoldingDunsNumber = &v
}

// GetNationalParentDunsNumber returns the NationalParentDunsNumber field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeInternationalIdentification) GetNationalParentDunsNumber() string {
	if o == nil || IsNil(o.NationalParentDunsNumber) {
		var ret string
		return ret
	}
	return *o.NationalParentDunsNumber
}

// GetNationalParentDunsNumberOk returns a tuple with the NationalParentDunsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeInternationalIdentification) GetNationalParentDunsNumberOk() (*string, bool) {
	if o == nil || IsNil(o.NationalParentDunsNumber) {
		return nil, false
	}
	return o.NationalParentDunsNumber, true
}

// HasNationalParentDunsNumber returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeInternationalIdentification) HasNationalParentDunsNumber() bool {
	if o != nil && !IsNil(o.NationalParentDunsNumber) {
		return true
	}

	return false
}

// SetNationalParentDunsNumber gets a reference to the given string and assigns it to the NationalParentDunsNumber field.
func (o *SearchCompanyDataTypeInternationalIdentification) SetNationalParentDunsNumber(v string) {
	o.NationalParentDunsNumber = &v
}

func (o SearchCompanyDataTypeInternationalIdentification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeInternationalIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HoldingDunsNumber) {
		toSerialize["holdingDunsNumber"] = o.HoldingDunsNumber
	}
	if !IsNil(o.NationalParentDunsNumber) {
		toSerialize["nationalParentDunsNumber"] = o.NationalParentDunsNumber
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeInternationalIdentification struct {
	value *SearchCompanyDataTypeInternationalIdentification
	isSet bool
}

func (v NullableSearchCompanyDataTypeInternationalIdentification) Get() *SearchCompanyDataTypeInternationalIdentification {
	return v.value
}

func (v *NullableSearchCompanyDataTypeInternationalIdentification) Set(val *SearchCompanyDataTypeInternationalIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeInternationalIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeInternationalIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeInternationalIdentification(val *SearchCompanyDataTypeInternationalIdentification) *NullableSearchCompanyDataTypeInternationalIdentification {
	return &NullableSearchCompanyDataTypeInternationalIdentification{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeInternationalIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeInternationalIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


