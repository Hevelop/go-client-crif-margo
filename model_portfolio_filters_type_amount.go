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

// checks if the PortfolioFiltersTypeAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioFiltersTypeAmount{}

// PortfolioFiltersTypeAmount struct for PortfolioFiltersTypeAmount
type PortfolioFiltersTypeAmount struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewPortfolioFiltersTypeAmount instantiates a new PortfolioFiltersTypeAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioFiltersTypeAmount() *PortfolioFiltersTypeAmount {
	this := PortfolioFiltersTypeAmount{}
	return &this
}

// NewPortfolioFiltersTypeAmountWithDefaults instantiates a new PortfolioFiltersTypeAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioFiltersTypeAmountWithDefaults() *PortfolioFiltersTypeAmount {
	this := PortfolioFiltersTypeAmount{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *PortfolioFiltersTypeAmount) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioFiltersTypeAmount) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *PortfolioFiltersTypeAmount) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *PortfolioFiltersTypeAmount) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *PortfolioFiltersTypeAmount) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioFiltersTypeAmount) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *PortfolioFiltersTypeAmount) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *PortfolioFiltersTypeAmount) SetMax(v int32) {
	o.Max = &v
}

func (o PortfolioFiltersTypeAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioFiltersTypeAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullablePortfolioFiltersTypeAmount struct {
	value *PortfolioFiltersTypeAmount
	isSet bool
}

func (v NullablePortfolioFiltersTypeAmount) Get() *PortfolioFiltersTypeAmount {
	return v.value
}

func (v *NullablePortfolioFiltersTypeAmount) Set(val *PortfolioFiltersTypeAmount) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioFiltersTypeAmount) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioFiltersTypeAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioFiltersTypeAmount(val *PortfolioFiltersTypeAmount) *NullablePortfolioFiltersTypeAmount {
	return &NullablePortfolioFiltersTypeAmount{value: val, isSet: true}
}

func (v NullablePortfolioFiltersTypeAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioFiltersTypeAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


