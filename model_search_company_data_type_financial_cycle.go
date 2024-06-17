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

// checks if the SearchCompanyDataTypeFinancialCycle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeFinancialCycle{}

// SearchCompanyDataTypeFinancialCycle struct for SearchCompanyDataTypeFinancialCycle
type SearchCompanyDataTypeFinancialCycle struct {
	AccountsReceivableDuration *float32 `json:"accountsReceivableDuration,omitempty"`
	StockDuration *float32 `json:"stockDuration,omitempty"`
	DebtsToSuppliersDuration *float32 `json:"debtsToSuppliersDuration,omitempty"`
	FinancialCycleDuration *float32 `json:"financialCycleDuration,omitempty"`
}

// NewSearchCompanyDataTypeFinancialCycle instantiates a new SearchCompanyDataTypeFinancialCycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeFinancialCycle() *SearchCompanyDataTypeFinancialCycle {
	this := SearchCompanyDataTypeFinancialCycle{}
	return &this
}

// NewSearchCompanyDataTypeFinancialCycleWithDefaults instantiates a new SearchCompanyDataTypeFinancialCycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeFinancialCycleWithDefaults() *SearchCompanyDataTypeFinancialCycle {
	this := SearchCompanyDataTypeFinancialCycle{}
	return &this
}

// GetAccountsReceivableDuration returns the AccountsReceivableDuration field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeFinancialCycle) GetAccountsReceivableDuration() float32 {
	if o == nil || IsNil(o.AccountsReceivableDuration) {
		var ret float32
		return ret
	}
	return *o.AccountsReceivableDuration
}

// GetAccountsReceivableDurationOk returns a tuple with the AccountsReceivableDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeFinancialCycle) GetAccountsReceivableDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.AccountsReceivableDuration) {
		return nil, false
	}
	return o.AccountsReceivableDuration, true
}

// HasAccountsReceivableDuration returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeFinancialCycle) HasAccountsReceivableDuration() bool {
	if o != nil && !IsNil(o.AccountsReceivableDuration) {
		return true
	}

	return false
}

// SetAccountsReceivableDuration gets a reference to the given float32 and assigns it to the AccountsReceivableDuration field.
func (o *SearchCompanyDataTypeFinancialCycle) SetAccountsReceivableDuration(v float32) {
	o.AccountsReceivableDuration = &v
}

// GetStockDuration returns the StockDuration field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeFinancialCycle) GetStockDuration() float32 {
	if o == nil || IsNil(o.StockDuration) {
		var ret float32
		return ret
	}
	return *o.StockDuration
}

// GetStockDurationOk returns a tuple with the StockDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeFinancialCycle) GetStockDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.StockDuration) {
		return nil, false
	}
	return o.StockDuration, true
}

// HasStockDuration returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeFinancialCycle) HasStockDuration() bool {
	if o != nil && !IsNil(o.StockDuration) {
		return true
	}

	return false
}

// SetStockDuration gets a reference to the given float32 and assigns it to the StockDuration field.
func (o *SearchCompanyDataTypeFinancialCycle) SetStockDuration(v float32) {
	o.StockDuration = &v
}

// GetDebtsToSuppliersDuration returns the DebtsToSuppliersDuration field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeFinancialCycle) GetDebtsToSuppliersDuration() float32 {
	if o == nil || IsNil(o.DebtsToSuppliersDuration) {
		var ret float32
		return ret
	}
	return *o.DebtsToSuppliersDuration
}

// GetDebtsToSuppliersDurationOk returns a tuple with the DebtsToSuppliersDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeFinancialCycle) GetDebtsToSuppliersDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.DebtsToSuppliersDuration) {
		return nil, false
	}
	return o.DebtsToSuppliersDuration, true
}

// HasDebtsToSuppliersDuration returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeFinancialCycle) HasDebtsToSuppliersDuration() bool {
	if o != nil && !IsNil(o.DebtsToSuppliersDuration) {
		return true
	}

	return false
}

// SetDebtsToSuppliersDuration gets a reference to the given float32 and assigns it to the DebtsToSuppliersDuration field.
func (o *SearchCompanyDataTypeFinancialCycle) SetDebtsToSuppliersDuration(v float32) {
	o.DebtsToSuppliersDuration = &v
}

// GetFinancialCycleDuration returns the FinancialCycleDuration field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeFinancialCycle) GetFinancialCycleDuration() float32 {
	if o == nil || IsNil(o.FinancialCycleDuration) {
		var ret float32
		return ret
	}
	return *o.FinancialCycleDuration
}

// GetFinancialCycleDurationOk returns a tuple with the FinancialCycleDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeFinancialCycle) GetFinancialCycleDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.FinancialCycleDuration) {
		return nil, false
	}
	return o.FinancialCycleDuration, true
}

// HasFinancialCycleDuration returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeFinancialCycle) HasFinancialCycleDuration() bool {
	if o != nil && !IsNil(o.FinancialCycleDuration) {
		return true
	}

	return false
}

// SetFinancialCycleDuration gets a reference to the given float32 and assigns it to the FinancialCycleDuration field.
func (o *SearchCompanyDataTypeFinancialCycle) SetFinancialCycleDuration(v float32) {
	o.FinancialCycleDuration = &v
}

func (o SearchCompanyDataTypeFinancialCycle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeFinancialCycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountsReceivableDuration) {
		toSerialize["accountsReceivableDuration"] = o.AccountsReceivableDuration
	}
	if !IsNil(o.StockDuration) {
		toSerialize["stockDuration"] = o.StockDuration
	}
	if !IsNil(o.DebtsToSuppliersDuration) {
		toSerialize["debtsToSuppliersDuration"] = o.DebtsToSuppliersDuration
	}
	if !IsNil(o.FinancialCycleDuration) {
		toSerialize["financialCycleDuration"] = o.FinancialCycleDuration
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeFinancialCycle struct {
	value *SearchCompanyDataTypeFinancialCycle
	isSet bool
}

func (v NullableSearchCompanyDataTypeFinancialCycle) Get() *SearchCompanyDataTypeFinancialCycle {
	return v.value
}

func (v *NullableSearchCompanyDataTypeFinancialCycle) Set(val *SearchCompanyDataTypeFinancialCycle) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeFinancialCycle) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeFinancialCycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeFinancialCycle(val *SearchCompanyDataTypeFinancialCycle) *NullableSearchCompanyDataTypeFinancialCycle {
	return &NullableSearchCompanyDataTypeFinancialCycle{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeFinancialCycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeFinancialCycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

