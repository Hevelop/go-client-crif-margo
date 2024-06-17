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

// checks if the SearchCompanyDataTypeStateOwned type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeStateOwned{}

// SearchCompanyDataTypeStateOwned struct for SearchCompanyDataTypeStateOwned
type SearchCompanyDataTypeStateOwned struct {
	IsStateOwned *bool `json:"isStateOwned,omitempty"`
	StateOwnedType *SearchCompanyDataTypeStateOwnedStateOwnedType `json:"stateOwnedType,omitempty"`
}

// NewSearchCompanyDataTypeStateOwned instantiates a new SearchCompanyDataTypeStateOwned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeStateOwned() *SearchCompanyDataTypeStateOwned {
	this := SearchCompanyDataTypeStateOwned{}
	return &this
}

// NewSearchCompanyDataTypeStateOwnedWithDefaults instantiates a new SearchCompanyDataTypeStateOwned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeStateOwnedWithDefaults() *SearchCompanyDataTypeStateOwned {
	this := SearchCompanyDataTypeStateOwned{}
	return &this
}

// GetIsStateOwned returns the IsStateOwned field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeStateOwned) GetIsStateOwned() bool {
	if o == nil || IsNil(o.IsStateOwned) {
		var ret bool
		return ret
	}
	return *o.IsStateOwned
}

// GetIsStateOwnedOk returns a tuple with the IsStateOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeStateOwned) GetIsStateOwnedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStateOwned) {
		return nil, false
	}
	return o.IsStateOwned, true
}

// HasIsStateOwned returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeStateOwned) HasIsStateOwned() bool {
	if o != nil && !IsNil(o.IsStateOwned) {
		return true
	}

	return false
}

// SetIsStateOwned gets a reference to the given bool and assigns it to the IsStateOwned field.
func (o *SearchCompanyDataTypeStateOwned) SetIsStateOwned(v bool) {
	o.IsStateOwned = &v
}

// GetStateOwnedType returns the StateOwnedType field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeStateOwned) GetStateOwnedType() SearchCompanyDataTypeStateOwnedStateOwnedType {
	if o == nil || IsNil(o.StateOwnedType) {
		var ret SearchCompanyDataTypeStateOwnedStateOwnedType
		return ret
	}
	return *o.StateOwnedType
}

// GetStateOwnedTypeOk returns a tuple with the StateOwnedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeStateOwned) GetStateOwnedTypeOk() (*SearchCompanyDataTypeStateOwnedStateOwnedType, bool) {
	if o == nil || IsNil(o.StateOwnedType) {
		return nil, false
	}
	return o.StateOwnedType, true
}

// HasStateOwnedType returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeStateOwned) HasStateOwnedType() bool {
	if o != nil && !IsNil(o.StateOwnedType) {
		return true
	}

	return false
}

// SetStateOwnedType gets a reference to the given SearchCompanyDataTypeStateOwnedStateOwnedType and assigns it to the StateOwnedType field.
func (o *SearchCompanyDataTypeStateOwned) SetStateOwnedType(v SearchCompanyDataTypeStateOwnedStateOwnedType) {
	o.StateOwnedType = &v
}

func (o SearchCompanyDataTypeStateOwned) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeStateOwned) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsStateOwned) {
		toSerialize["isStateOwned"] = o.IsStateOwned
	}
	if !IsNil(o.StateOwnedType) {
		toSerialize["stateOwnedType"] = o.StateOwnedType
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeStateOwned struct {
	value *SearchCompanyDataTypeStateOwned
	isSet bool
}

func (v NullableSearchCompanyDataTypeStateOwned) Get() *SearchCompanyDataTypeStateOwned {
	return v.value
}

func (v *NullableSearchCompanyDataTypeStateOwned) Set(val *SearchCompanyDataTypeStateOwned) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeStateOwned) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeStateOwned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeStateOwned(val *SearchCompanyDataTypeStateOwned) *NullableSearchCompanyDataTypeStateOwned {
	return &NullableSearchCompanyDataTypeStateOwned{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeStateOwned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeStateOwned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

