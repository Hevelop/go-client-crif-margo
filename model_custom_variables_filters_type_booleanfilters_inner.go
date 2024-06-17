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

// checks if the CustomVariablesFiltersTypeBooleanfiltersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomVariablesFiltersTypeBooleanfiltersInner{}

// CustomVariablesFiltersTypeBooleanfiltersInner struct for CustomVariablesFiltersTypeBooleanfiltersInner
type CustomVariablesFiltersTypeBooleanfiltersInner struct {
	VariableIndex *int32 `json:"variableIndex,omitempty"`
	Value *bool `json:"value,omitempty"`
}

// NewCustomVariablesFiltersTypeBooleanfiltersInner instantiates a new CustomVariablesFiltersTypeBooleanfiltersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomVariablesFiltersTypeBooleanfiltersInner() *CustomVariablesFiltersTypeBooleanfiltersInner {
	this := CustomVariablesFiltersTypeBooleanfiltersInner{}
	return &this
}

// NewCustomVariablesFiltersTypeBooleanfiltersInnerWithDefaults instantiates a new CustomVariablesFiltersTypeBooleanfiltersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomVariablesFiltersTypeBooleanfiltersInnerWithDefaults() *CustomVariablesFiltersTypeBooleanfiltersInner {
	this := CustomVariablesFiltersTypeBooleanfiltersInner{}
	return &this
}

// GetVariableIndex returns the VariableIndex field value if set, zero value otherwise.
func (o *CustomVariablesFiltersTypeBooleanfiltersInner) GetVariableIndex() int32 {
	if o == nil || IsNil(o.VariableIndex) {
		var ret int32
		return ret
	}
	return *o.VariableIndex
}

// GetVariableIndexOk returns a tuple with the VariableIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomVariablesFiltersTypeBooleanfiltersInner) GetVariableIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.VariableIndex) {
		return nil, false
	}
	return o.VariableIndex, true
}

// HasVariableIndex returns a boolean if a field has been set.
func (o *CustomVariablesFiltersTypeBooleanfiltersInner) HasVariableIndex() bool {
	if o != nil && !IsNil(o.VariableIndex) {
		return true
	}

	return false
}

// SetVariableIndex gets a reference to the given int32 and assigns it to the VariableIndex field.
func (o *CustomVariablesFiltersTypeBooleanfiltersInner) SetVariableIndex(v int32) {
	o.VariableIndex = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomVariablesFiltersTypeBooleanfiltersInner) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomVariablesFiltersTypeBooleanfiltersInner) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomVariablesFiltersTypeBooleanfiltersInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *CustomVariablesFiltersTypeBooleanfiltersInner) SetValue(v bool) {
	o.Value = &v
}

func (o CustomVariablesFiltersTypeBooleanfiltersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomVariablesFiltersTypeBooleanfiltersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VariableIndex) {
		toSerialize["variableIndex"] = o.VariableIndex
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCustomVariablesFiltersTypeBooleanfiltersInner struct {
	value *CustomVariablesFiltersTypeBooleanfiltersInner
	isSet bool
}

func (v NullableCustomVariablesFiltersTypeBooleanfiltersInner) Get() *CustomVariablesFiltersTypeBooleanfiltersInner {
	return v.value
}

func (v *NullableCustomVariablesFiltersTypeBooleanfiltersInner) Set(val *CustomVariablesFiltersTypeBooleanfiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomVariablesFiltersTypeBooleanfiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomVariablesFiltersTypeBooleanfiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomVariablesFiltersTypeBooleanfiltersInner(val *CustomVariablesFiltersTypeBooleanfiltersInner) *NullableCustomVariablesFiltersTypeBooleanfiltersInner {
	return &NullableCustomVariablesFiltersTypeBooleanfiltersInner{value: val, isSet: true}
}

func (v NullableCustomVariablesFiltersTypeBooleanfiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomVariablesFiltersTypeBooleanfiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


