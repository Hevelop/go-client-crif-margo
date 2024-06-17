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

// checks if the SearchCompanyDataTypeInnovativeSmeAndSu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeInnovativeSmeAndSu{}

// SearchCompanyDataTypeInnovativeSmeAndSu struct for SearchCompanyDataTypeInnovativeSmeAndSu
type SearchCompanyDataTypeInnovativeSmeAndSu struct {
	IsInnovativeStartUp *bool `json:"isInnovativeStartUp,omitempty"`
	IsInnovativeSme *bool `json:"isInnovativeSme,omitempty"`
}

// NewSearchCompanyDataTypeInnovativeSmeAndSu instantiates a new SearchCompanyDataTypeInnovativeSmeAndSu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeInnovativeSmeAndSu() *SearchCompanyDataTypeInnovativeSmeAndSu {
	this := SearchCompanyDataTypeInnovativeSmeAndSu{}
	return &this
}

// NewSearchCompanyDataTypeInnovativeSmeAndSuWithDefaults instantiates a new SearchCompanyDataTypeInnovativeSmeAndSu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeInnovativeSmeAndSuWithDefaults() *SearchCompanyDataTypeInnovativeSmeAndSu {
	this := SearchCompanyDataTypeInnovativeSmeAndSu{}
	return &this
}

// GetIsInnovativeStartUp returns the IsInnovativeStartUp field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeInnovativeSmeAndSu) GetIsInnovativeStartUp() bool {
	if o == nil || IsNil(o.IsInnovativeStartUp) {
		var ret bool
		return ret
	}
	return *o.IsInnovativeStartUp
}

// GetIsInnovativeStartUpOk returns a tuple with the IsInnovativeStartUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeInnovativeSmeAndSu) GetIsInnovativeStartUpOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInnovativeStartUp) {
		return nil, false
	}
	return o.IsInnovativeStartUp, true
}

// HasIsInnovativeStartUp returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeInnovativeSmeAndSu) HasIsInnovativeStartUp() bool {
	if o != nil && !IsNil(o.IsInnovativeStartUp) {
		return true
	}

	return false
}

// SetIsInnovativeStartUp gets a reference to the given bool and assigns it to the IsInnovativeStartUp field.
func (o *SearchCompanyDataTypeInnovativeSmeAndSu) SetIsInnovativeStartUp(v bool) {
	o.IsInnovativeStartUp = &v
}

// GetIsInnovativeSme returns the IsInnovativeSme field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeInnovativeSmeAndSu) GetIsInnovativeSme() bool {
	if o == nil || IsNil(o.IsInnovativeSme) {
		var ret bool
		return ret
	}
	return *o.IsInnovativeSme
}

// GetIsInnovativeSmeOk returns a tuple with the IsInnovativeSme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeInnovativeSmeAndSu) GetIsInnovativeSmeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInnovativeSme) {
		return nil, false
	}
	return o.IsInnovativeSme, true
}

// HasIsInnovativeSme returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeInnovativeSmeAndSu) HasIsInnovativeSme() bool {
	if o != nil && !IsNil(o.IsInnovativeSme) {
		return true
	}

	return false
}

// SetIsInnovativeSme gets a reference to the given bool and assigns it to the IsInnovativeSme field.
func (o *SearchCompanyDataTypeInnovativeSmeAndSu) SetIsInnovativeSme(v bool) {
	o.IsInnovativeSme = &v
}

func (o SearchCompanyDataTypeInnovativeSmeAndSu) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeInnovativeSmeAndSu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsInnovativeStartUp) {
		toSerialize["isInnovativeStartUp"] = o.IsInnovativeStartUp
	}
	if !IsNil(o.IsInnovativeSme) {
		toSerialize["isInnovativeSme"] = o.IsInnovativeSme
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeInnovativeSmeAndSu struct {
	value *SearchCompanyDataTypeInnovativeSmeAndSu
	isSet bool
}

func (v NullableSearchCompanyDataTypeInnovativeSmeAndSu) Get() *SearchCompanyDataTypeInnovativeSmeAndSu {
	return v.value
}

func (v *NullableSearchCompanyDataTypeInnovativeSmeAndSu) Set(val *SearchCompanyDataTypeInnovativeSmeAndSu) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeInnovativeSmeAndSu) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeInnovativeSmeAndSu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeInnovativeSmeAndSu(val *SearchCompanyDataTypeInnovativeSmeAndSu) *NullableSearchCompanyDataTypeInnovativeSmeAndSu {
	return &NullableSearchCompanyDataTypeInnovativeSmeAndSu{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeInnovativeSmeAndSu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeInnovativeSmeAndSu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


