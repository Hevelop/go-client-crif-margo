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

// checks if the CustomVariablesFiltersType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomVariablesFiltersType{}

// CustomVariablesFiltersType One portfolio's feature are the custom variables and their data type could be: numeric, date, boolean, string and domain.  Like `filters` object, also `customVariablesFilters` object contains a key for each data type.
type CustomVariablesFiltersType struct {
	// The array contains the custom variable with data type **numeric**.  The object must have the following structure: * variableIndex: it identifies the index associated at custom variable * value: this object have to contain the range of value to be filtered.
	Numericfilters []CustomVariablesFiltersTypeNumericfiltersInner `json:"numericfilters,omitempty"`
	// The array contains the custom variable with data type **date**.  The object must have the following structure: * variableIndex: it identifies the index associated at custom variable * value: this object have to contain the range of value to be filtered.
	Datefilters []CustomVariablesFiltersTypeDatefiltersInner `json:"datefilters,omitempty"`
	// The array contains all fields with data type ***boolean***.  The object must have the following structure: * variableIndex: it identifies the index associated at custom variable * value: the field have to contain a boolean value (true or false).
	Booleanfilters []CustomVariablesFiltersTypeBooleanfiltersInner `json:"booleanfilters,omitempty"`
	// The array contains all fields with data type ***string***.  The object must have the following structure: * variableIndex: it identifies the index associated at custom variable * value: the field have to contain a label.
	Stringfilters []CustomVariablesFiltersTypeStringfiltersInner `json:"stringfilters,omitempty"`
}

// NewCustomVariablesFiltersType instantiates a new CustomVariablesFiltersType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomVariablesFiltersType() *CustomVariablesFiltersType {
	this := CustomVariablesFiltersType{}
	return &this
}

// NewCustomVariablesFiltersTypeWithDefaults instantiates a new CustomVariablesFiltersType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomVariablesFiltersTypeWithDefaults() *CustomVariablesFiltersType {
	this := CustomVariablesFiltersType{}
	return &this
}

// GetNumericfilters returns the Numericfilters field value if set, zero value otherwise.
func (o *CustomVariablesFiltersType) GetNumericfilters() []CustomVariablesFiltersTypeNumericfiltersInner {
	if o == nil || IsNil(o.Numericfilters) {
		var ret []CustomVariablesFiltersTypeNumericfiltersInner
		return ret
	}
	return o.Numericfilters
}

// GetNumericfiltersOk returns a tuple with the Numericfilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomVariablesFiltersType) GetNumericfiltersOk() ([]CustomVariablesFiltersTypeNumericfiltersInner, bool) {
	if o == nil || IsNil(o.Numericfilters) {
		return nil, false
	}
	return o.Numericfilters, true
}

// HasNumericfilters returns a boolean if a field has been set.
func (o *CustomVariablesFiltersType) HasNumericfilters() bool {
	if o != nil && !IsNil(o.Numericfilters) {
		return true
	}

	return false
}

// SetNumericfilters gets a reference to the given []CustomVariablesFiltersTypeNumericfiltersInner and assigns it to the Numericfilters field.
func (o *CustomVariablesFiltersType) SetNumericfilters(v []CustomVariablesFiltersTypeNumericfiltersInner) {
	o.Numericfilters = v
}

// GetDatefilters returns the Datefilters field value if set, zero value otherwise.
func (o *CustomVariablesFiltersType) GetDatefilters() []CustomVariablesFiltersTypeDatefiltersInner {
	if o == nil || IsNil(o.Datefilters) {
		var ret []CustomVariablesFiltersTypeDatefiltersInner
		return ret
	}
	return o.Datefilters
}

// GetDatefiltersOk returns a tuple with the Datefilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomVariablesFiltersType) GetDatefiltersOk() ([]CustomVariablesFiltersTypeDatefiltersInner, bool) {
	if o == nil || IsNil(o.Datefilters) {
		return nil, false
	}
	return o.Datefilters, true
}

// HasDatefilters returns a boolean if a field has been set.
func (o *CustomVariablesFiltersType) HasDatefilters() bool {
	if o != nil && !IsNil(o.Datefilters) {
		return true
	}

	return false
}

// SetDatefilters gets a reference to the given []CustomVariablesFiltersTypeDatefiltersInner and assigns it to the Datefilters field.
func (o *CustomVariablesFiltersType) SetDatefilters(v []CustomVariablesFiltersTypeDatefiltersInner) {
	o.Datefilters = v
}

// GetBooleanfilters returns the Booleanfilters field value if set, zero value otherwise.
func (o *CustomVariablesFiltersType) GetBooleanfilters() []CustomVariablesFiltersTypeBooleanfiltersInner {
	if o == nil || IsNil(o.Booleanfilters) {
		var ret []CustomVariablesFiltersTypeBooleanfiltersInner
		return ret
	}
	return o.Booleanfilters
}

// GetBooleanfiltersOk returns a tuple with the Booleanfilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomVariablesFiltersType) GetBooleanfiltersOk() ([]CustomVariablesFiltersTypeBooleanfiltersInner, bool) {
	if o == nil || IsNil(o.Booleanfilters) {
		return nil, false
	}
	return o.Booleanfilters, true
}

// HasBooleanfilters returns a boolean if a field has been set.
func (o *CustomVariablesFiltersType) HasBooleanfilters() bool {
	if o != nil && !IsNil(o.Booleanfilters) {
		return true
	}

	return false
}

// SetBooleanfilters gets a reference to the given []CustomVariablesFiltersTypeBooleanfiltersInner and assigns it to the Booleanfilters field.
func (o *CustomVariablesFiltersType) SetBooleanfilters(v []CustomVariablesFiltersTypeBooleanfiltersInner) {
	o.Booleanfilters = v
}

// GetStringfilters returns the Stringfilters field value if set, zero value otherwise.
func (o *CustomVariablesFiltersType) GetStringfilters() []CustomVariablesFiltersTypeStringfiltersInner {
	if o == nil || IsNil(o.Stringfilters) {
		var ret []CustomVariablesFiltersTypeStringfiltersInner
		return ret
	}
	return o.Stringfilters
}

// GetStringfiltersOk returns a tuple with the Stringfilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomVariablesFiltersType) GetStringfiltersOk() ([]CustomVariablesFiltersTypeStringfiltersInner, bool) {
	if o == nil || IsNil(o.Stringfilters) {
		return nil, false
	}
	return o.Stringfilters, true
}

// HasStringfilters returns a boolean if a field has been set.
func (o *CustomVariablesFiltersType) HasStringfilters() bool {
	if o != nil && !IsNil(o.Stringfilters) {
		return true
	}

	return false
}

// SetStringfilters gets a reference to the given []CustomVariablesFiltersTypeStringfiltersInner and assigns it to the Stringfilters field.
func (o *CustomVariablesFiltersType) SetStringfilters(v []CustomVariablesFiltersTypeStringfiltersInner) {
	o.Stringfilters = v
}

func (o CustomVariablesFiltersType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomVariablesFiltersType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Numericfilters) {
		toSerialize["numericfilters"] = o.Numericfilters
	}
	if !IsNil(o.Datefilters) {
		toSerialize["datefilters"] = o.Datefilters
	}
	if !IsNil(o.Booleanfilters) {
		toSerialize["booleanfilters"] = o.Booleanfilters
	}
	if !IsNil(o.Stringfilters) {
		toSerialize["stringfilters"] = o.Stringfilters
	}
	return toSerialize, nil
}

type NullableCustomVariablesFiltersType struct {
	value *CustomVariablesFiltersType
	isSet bool
}

func (v NullableCustomVariablesFiltersType) Get() *CustomVariablesFiltersType {
	return v.value
}

func (v *NullableCustomVariablesFiltersType) Set(val *CustomVariablesFiltersType) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomVariablesFiltersType) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomVariablesFiltersType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomVariablesFiltersType(val *CustomVariablesFiltersType) *NullableCustomVariablesFiltersType {
	return &NullableCustomVariablesFiltersType{value: val, isSet: true}
}

func (v NullableCustomVariablesFiltersType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomVariablesFiltersType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

