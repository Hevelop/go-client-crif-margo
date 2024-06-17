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

// checks if the SearchCompanyDataTypeEfficiency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeEfficiency{}

// SearchCompanyDataTypeEfficiency struct for SearchCompanyDataTypeEfficiency
type SearchCompanyDataTypeEfficiency struct {
	TurnoverIndex *float32 `json:"turnoverIndex,omitempty"`
	InventoryRotation *float32 `json:"inventoryRotation,omitempty"`
	CurrentAssetsRotation *float32 `json:"currentAssetsRotation,omitempty"`
	AccountsReceivableRotation *float32 `json:"accountsReceivableRotation,omitempty"`
}

// NewSearchCompanyDataTypeEfficiency instantiates a new SearchCompanyDataTypeEfficiency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeEfficiency() *SearchCompanyDataTypeEfficiency {
	this := SearchCompanyDataTypeEfficiency{}
	return &this
}

// NewSearchCompanyDataTypeEfficiencyWithDefaults instantiates a new SearchCompanyDataTypeEfficiency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeEfficiencyWithDefaults() *SearchCompanyDataTypeEfficiency {
	this := SearchCompanyDataTypeEfficiency{}
	return &this
}

// GetTurnoverIndex returns the TurnoverIndex field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeEfficiency) GetTurnoverIndex() float32 {
	if o == nil || IsNil(o.TurnoverIndex) {
		var ret float32
		return ret
	}
	return *o.TurnoverIndex
}

// GetTurnoverIndexOk returns a tuple with the TurnoverIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeEfficiency) GetTurnoverIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.TurnoverIndex) {
		return nil, false
	}
	return o.TurnoverIndex, true
}

// HasTurnoverIndex returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeEfficiency) HasTurnoverIndex() bool {
	if o != nil && !IsNil(o.TurnoverIndex) {
		return true
	}

	return false
}

// SetTurnoverIndex gets a reference to the given float32 and assigns it to the TurnoverIndex field.
func (o *SearchCompanyDataTypeEfficiency) SetTurnoverIndex(v float32) {
	o.TurnoverIndex = &v
}

// GetInventoryRotation returns the InventoryRotation field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeEfficiency) GetInventoryRotation() float32 {
	if o == nil || IsNil(o.InventoryRotation) {
		var ret float32
		return ret
	}
	return *o.InventoryRotation
}

// GetInventoryRotationOk returns a tuple with the InventoryRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeEfficiency) GetInventoryRotationOk() (*float32, bool) {
	if o == nil || IsNil(o.InventoryRotation) {
		return nil, false
	}
	return o.InventoryRotation, true
}

// HasInventoryRotation returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeEfficiency) HasInventoryRotation() bool {
	if o != nil && !IsNil(o.InventoryRotation) {
		return true
	}

	return false
}

// SetInventoryRotation gets a reference to the given float32 and assigns it to the InventoryRotation field.
func (o *SearchCompanyDataTypeEfficiency) SetInventoryRotation(v float32) {
	o.InventoryRotation = &v
}

// GetCurrentAssetsRotation returns the CurrentAssetsRotation field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeEfficiency) GetCurrentAssetsRotation() float32 {
	if o == nil || IsNil(o.CurrentAssetsRotation) {
		var ret float32
		return ret
	}
	return *o.CurrentAssetsRotation
}

// GetCurrentAssetsRotationOk returns a tuple with the CurrentAssetsRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeEfficiency) GetCurrentAssetsRotationOk() (*float32, bool) {
	if o == nil || IsNil(o.CurrentAssetsRotation) {
		return nil, false
	}
	return o.CurrentAssetsRotation, true
}

// HasCurrentAssetsRotation returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeEfficiency) HasCurrentAssetsRotation() bool {
	if o != nil && !IsNil(o.CurrentAssetsRotation) {
		return true
	}

	return false
}

// SetCurrentAssetsRotation gets a reference to the given float32 and assigns it to the CurrentAssetsRotation field.
func (o *SearchCompanyDataTypeEfficiency) SetCurrentAssetsRotation(v float32) {
	o.CurrentAssetsRotation = &v
}

// GetAccountsReceivableRotation returns the AccountsReceivableRotation field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeEfficiency) GetAccountsReceivableRotation() float32 {
	if o == nil || IsNil(o.AccountsReceivableRotation) {
		var ret float32
		return ret
	}
	return *o.AccountsReceivableRotation
}

// GetAccountsReceivableRotationOk returns a tuple with the AccountsReceivableRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeEfficiency) GetAccountsReceivableRotationOk() (*float32, bool) {
	if o == nil || IsNil(o.AccountsReceivableRotation) {
		return nil, false
	}
	return o.AccountsReceivableRotation, true
}

// HasAccountsReceivableRotation returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeEfficiency) HasAccountsReceivableRotation() bool {
	if o != nil && !IsNil(o.AccountsReceivableRotation) {
		return true
	}

	return false
}

// SetAccountsReceivableRotation gets a reference to the given float32 and assigns it to the AccountsReceivableRotation field.
func (o *SearchCompanyDataTypeEfficiency) SetAccountsReceivableRotation(v float32) {
	o.AccountsReceivableRotation = &v
}

func (o SearchCompanyDataTypeEfficiency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeEfficiency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TurnoverIndex) {
		toSerialize["turnoverIndex"] = o.TurnoverIndex
	}
	if !IsNil(o.InventoryRotation) {
		toSerialize["inventoryRotation"] = o.InventoryRotation
	}
	if !IsNil(o.CurrentAssetsRotation) {
		toSerialize["currentAssetsRotation"] = o.CurrentAssetsRotation
	}
	if !IsNil(o.AccountsReceivableRotation) {
		toSerialize["accountsReceivableRotation"] = o.AccountsReceivableRotation
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeEfficiency struct {
	value *SearchCompanyDataTypeEfficiency
	isSet bool
}

func (v NullableSearchCompanyDataTypeEfficiency) Get() *SearchCompanyDataTypeEfficiency {
	return v.value
}

func (v *NullableSearchCompanyDataTypeEfficiency) Set(val *SearchCompanyDataTypeEfficiency) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeEfficiency) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeEfficiency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeEfficiency(val *SearchCompanyDataTypeEfficiency) *NullableSearchCompanyDataTypeEfficiency {
	return &NullableSearchCompanyDataTypeEfficiency{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeEfficiency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeEfficiency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

