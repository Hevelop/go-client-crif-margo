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

// checks if the SearchCompanyDataTypeDevelopment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeDevelopment{}

// SearchCompanyDataTypeDevelopment struct for SearchCompanyDataTypeDevelopment
type SearchCompanyDataTypeDevelopment struct {
	EbitVariation *float32 `json:"ebitVariation,omitempty"`
	GrossFinancialDebt *float32 `json:"grossFinancialDebt,omitempty"`
	TotalAssets *float32 `json:"totalAssets,omitempty"`
	Mol *float32 `json:"mol,omitempty"`
	AddedValue *float32 `json:"addedValue,omitempty"`
	ProductionValue *float32 `json:"productionValue,omitempty"`
}

// NewSearchCompanyDataTypeDevelopment instantiates a new SearchCompanyDataTypeDevelopment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeDevelopment() *SearchCompanyDataTypeDevelopment {
	this := SearchCompanyDataTypeDevelopment{}
	return &this
}

// NewSearchCompanyDataTypeDevelopmentWithDefaults instantiates a new SearchCompanyDataTypeDevelopment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeDevelopmentWithDefaults() *SearchCompanyDataTypeDevelopment {
	this := SearchCompanyDataTypeDevelopment{}
	return &this
}

// GetEbitVariation returns the EbitVariation field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeDevelopment) GetEbitVariation() float32 {
	if o == nil || IsNil(o.EbitVariation) {
		var ret float32
		return ret
	}
	return *o.EbitVariation
}

// GetEbitVariationOk returns a tuple with the EbitVariation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeDevelopment) GetEbitVariationOk() (*float32, bool) {
	if o == nil || IsNil(o.EbitVariation) {
		return nil, false
	}
	return o.EbitVariation, true
}

// HasEbitVariation returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeDevelopment) HasEbitVariation() bool {
	if o != nil && !IsNil(o.EbitVariation) {
		return true
	}

	return false
}

// SetEbitVariation gets a reference to the given float32 and assigns it to the EbitVariation field.
func (o *SearchCompanyDataTypeDevelopment) SetEbitVariation(v float32) {
	o.EbitVariation = &v
}

// GetGrossFinancialDebt returns the GrossFinancialDebt field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeDevelopment) GetGrossFinancialDebt() float32 {
	if o == nil || IsNil(o.GrossFinancialDebt) {
		var ret float32
		return ret
	}
	return *o.GrossFinancialDebt
}

// GetGrossFinancialDebtOk returns a tuple with the GrossFinancialDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeDevelopment) GetGrossFinancialDebtOk() (*float32, bool) {
	if o == nil || IsNil(o.GrossFinancialDebt) {
		return nil, false
	}
	return o.GrossFinancialDebt, true
}

// HasGrossFinancialDebt returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeDevelopment) HasGrossFinancialDebt() bool {
	if o != nil && !IsNil(o.GrossFinancialDebt) {
		return true
	}

	return false
}

// SetGrossFinancialDebt gets a reference to the given float32 and assigns it to the GrossFinancialDebt field.
func (o *SearchCompanyDataTypeDevelopment) SetGrossFinancialDebt(v float32) {
	o.GrossFinancialDebt = &v
}

// GetTotalAssets returns the TotalAssets field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeDevelopment) GetTotalAssets() float32 {
	if o == nil || IsNil(o.TotalAssets) {
		var ret float32
		return ret
	}
	return *o.TotalAssets
}

// GetTotalAssetsOk returns a tuple with the TotalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeDevelopment) GetTotalAssetsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalAssets) {
		return nil, false
	}
	return o.TotalAssets, true
}

// HasTotalAssets returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeDevelopment) HasTotalAssets() bool {
	if o != nil && !IsNil(o.TotalAssets) {
		return true
	}

	return false
}

// SetTotalAssets gets a reference to the given float32 and assigns it to the TotalAssets field.
func (o *SearchCompanyDataTypeDevelopment) SetTotalAssets(v float32) {
	o.TotalAssets = &v
}

// GetMol returns the Mol field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeDevelopment) GetMol() float32 {
	if o == nil || IsNil(o.Mol) {
		var ret float32
		return ret
	}
	return *o.Mol
}

// GetMolOk returns a tuple with the Mol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeDevelopment) GetMolOk() (*float32, bool) {
	if o == nil || IsNil(o.Mol) {
		return nil, false
	}
	return o.Mol, true
}

// HasMol returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeDevelopment) HasMol() bool {
	if o != nil && !IsNil(o.Mol) {
		return true
	}

	return false
}

// SetMol gets a reference to the given float32 and assigns it to the Mol field.
func (o *SearchCompanyDataTypeDevelopment) SetMol(v float32) {
	o.Mol = &v
}

// GetAddedValue returns the AddedValue field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeDevelopment) GetAddedValue() float32 {
	if o == nil || IsNil(o.AddedValue) {
		var ret float32
		return ret
	}
	return *o.AddedValue
}

// GetAddedValueOk returns a tuple with the AddedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeDevelopment) GetAddedValueOk() (*float32, bool) {
	if o == nil || IsNil(o.AddedValue) {
		return nil, false
	}
	return o.AddedValue, true
}

// HasAddedValue returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeDevelopment) HasAddedValue() bool {
	if o != nil && !IsNil(o.AddedValue) {
		return true
	}

	return false
}

// SetAddedValue gets a reference to the given float32 and assigns it to the AddedValue field.
func (o *SearchCompanyDataTypeDevelopment) SetAddedValue(v float32) {
	o.AddedValue = &v
}

// GetProductionValue returns the ProductionValue field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeDevelopment) GetProductionValue() float32 {
	if o == nil || IsNil(o.ProductionValue) {
		var ret float32
		return ret
	}
	return *o.ProductionValue
}

// GetProductionValueOk returns a tuple with the ProductionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeDevelopment) GetProductionValueOk() (*float32, bool) {
	if o == nil || IsNil(o.ProductionValue) {
		return nil, false
	}
	return o.ProductionValue, true
}

// HasProductionValue returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeDevelopment) HasProductionValue() bool {
	if o != nil && !IsNil(o.ProductionValue) {
		return true
	}

	return false
}

// SetProductionValue gets a reference to the given float32 and assigns it to the ProductionValue field.
func (o *SearchCompanyDataTypeDevelopment) SetProductionValue(v float32) {
	o.ProductionValue = &v
}

func (o SearchCompanyDataTypeDevelopment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeDevelopment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EbitVariation) {
		toSerialize["ebitVariation"] = o.EbitVariation
	}
	if !IsNil(o.GrossFinancialDebt) {
		toSerialize["grossFinancialDebt"] = o.GrossFinancialDebt
	}
	if !IsNil(o.TotalAssets) {
		toSerialize["totalAssets"] = o.TotalAssets
	}
	if !IsNil(o.Mol) {
		toSerialize["mol"] = o.Mol
	}
	if !IsNil(o.AddedValue) {
		toSerialize["addedValue"] = o.AddedValue
	}
	if !IsNil(o.ProductionValue) {
		toSerialize["productionValue"] = o.ProductionValue
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeDevelopment struct {
	value *SearchCompanyDataTypeDevelopment
	isSet bool
}

func (v NullableSearchCompanyDataTypeDevelopment) Get() *SearchCompanyDataTypeDevelopment {
	return v.value
}

func (v *NullableSearchCompanyDataTypeDevelopment) Set(val *SearchCompanyDataTypeDevelopment) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeDevelopment) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeDevelopment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeDevelopment(val *SearchCompanyDataTypeDevelopment) *NullableSearchCompanyDataTypeDevelopment {
	return &NullableSearchCompanyDataTypeDevelopment{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeDevelopment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeDevelopment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


