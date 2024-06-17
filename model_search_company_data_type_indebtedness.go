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

// checks if the SearchCompanyDataTypeIndebtedness type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeIndebtedness{}

// SearchCompanyDataTypeIndebtedness struct for SearchCompanyDataTypeIndebtedness
type SearchCompanyDataTypeIndebtedness struct {
	BankDebtTotalAssets *float32 `json:"bankDebtTotalAssets,omitempty"`
	GrossFinancialDebtTotalAssets *float32 `json:"grossFinancialDebtTotalAssets,omitempty"`
	CapitalizationDegree *float32 `json:"capitalizationDegree,omitempty"`
	Leverage *float32 `json:"leverage,omitempty"`
	BankDebtRatio *float32 `json:"bankDebtRatio,omitempty"`
	DebtRatio *float32 `json:"debtRatio,omitempty"`
}

// NewSearchCompanyDataTypeIndebtedness instantiates a new SearchCompanyDataTypeIndebtedness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeIndebtedness() *SearchCompanyDataTypeIndebtedness {
	this := SearchCompanyDataTypeIndebtedness{}
	return &this
}

// NewSearchCompanyDataTypeIndebtednessWithDefaults instantiates a new SearchCompanyDataTypeIndebtedness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeIndebtednessWithDefaults() *SearchCompanyDataTypeIndebtedness {
	this := SearchCompanyDataTypeIndebtedness{}
	return &this
}

// GetBankDebtTotalAssets returns the BankDebtTotalAssets field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeIndebtedness) GetBankDebtTotalAssets() float32 {
	if o == nil || IsNil(o.BankDebtTotalAssets) {
		var ret float32
		return ret
	}
	return *o.BankDebtTotalAssets
}

// GetBankDebtTotalAssetsOk returns a tuple with the BankDebtTotalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeIndebtedness) GetBankDebtTotalAssetsOk() (*float32, bool) {
	if o == nil || IsNil(o.BankDebtTotalAssets) {
		return nil, false
	}
	return o.BankDebtTotalAssets, true
}

// HasBankDebtTotalAssets returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeIndebtedness) HasBankDebtTotalAssets() bool {
	if o != nil && !IsNil(o.BankDebtTotalAssets) {
		return true
	}

	return false
}

// SetBankDebtTotalAssets gets a reference to the given float32 and assigns it to the BankDebtTotalAssets field.
func (o *SearchCompanyDataTypeIndebtedness) SetBankDebtTotalAssets(v float32) {
	o.BankDebtTotalAssets = &v
}

// GetGrossFinancialDebtTotalAssets returns the GrossFinancialDebtTotalAssets field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeIndebtedness) GetGrossFinancialDebtTotalAssets() float32 {
	if o == nil || IsNil(o.GrossFinancialDebtTotalAssets) {
		var ret float32
		return ret
	}
	return *o.GrossFinancialDebtTotalAssets
}

// GetGrossFinancialDebtTotalAssetsOk returns a tuple with the GrossFinancialDebtTotalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeIndebtedness) GetGrossFinancialDebtTotalAssetsOk() (*float32, bool) {
	if o == nil || IsNil(o.GrossFinancialDebtTotalAssets) {
		return nil, false
	}
	return o.GrossFinancialDebtTotalAssets, true
}

// HasGrossFinancialDebtTotalAssets returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeIndebtedness) HasGrossFinancialDebtTotalAssets() bool {
	if o != nil && !IsNil(o.GrossFinancialDebtTotalAssets) {
		return true
	}

	return false
}

// SetGrossFinancialDebtTotalAssets gets a reference to the given float32 and assigns it to the GrossFinancialDebtTotalAssets field.
func (o *SearchCompanyDataTypeIndebtedness) SetGrossFinancialDebtTotalAssets(v float32) {
	o.GrossFinancialDebtTotalAssets = &v
}

// GetCapitalizationDegree returns the CapitalizationDegree field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeIndebtedness) GetCapitalizationDegree() float32 {
	if o == nil || IsNil(o.CapitalizationDegree) {
		var ret float32
		return ret
	}
	return *o.CapitalizationDegree
}

// GetCapitalizationDegreeOk returns a tuple with the CapitalizationDegree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeIndebtedness) GetCapitalizationDegreeOk() (*float32, bool) {
	if o == nil || IsNil(o.CapitalizationDegree) {
		return nil, false
	}
	return o.CapitalizationDegree, true
}

// HasCapitalizationDegree returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeIndebtedness) HasCapitalizationDegree() bool {
	if o != nil && !IsNil(o.CapitalizationDegree) {
		return true
	}

	return false
}

// SetCapitalizationDegree gets a reference to the given float32 and assigns it to the CapitalizationDegree field.
func (o *SearchCompanyDataTypeIndebtedness) SetCapitalizationDegree(v float32) {
	o.CapitalizationDegree = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeIndebtedness) GetLeverage() float32 {
	if o == nil || IsNil(o.Leverage) {
		var ret float32
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeIndebtedness) GetLeverageOk() (*float32, bool) {
	if o == nil || IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeIndebtedness) HasLeverage() bool {
	if o != nil && !IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given float32 and assigns it to the Leverage field.
func (o *SearchCompanyDataTypeIndebtedness) SetLeverage(v float32) {
	o.Leverage = &v
}

// GetBankDebtRatio returns the BankDebtRatio field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeIndebtedness) GetBankDebtRatio() float32 {
	if o == nil || IsNil(o.BankDebtRatio) {
		var ret float32
		return ret
	}
	return *o.BankDebtRatio
}

// GetBankDebtRatioOk returns a tuple with the BankDebtRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeIndebtedness) GetBankDebtRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.BankDebtRatio) {
		return nil, false
	}
	return o.BankDebtRatio, true
}

// HasBankDebtRatio returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeIndebtedness) HasBankDebtRatio() bool {
	if o != nil && !IsNil(o.BankDebtRatio) {
		return true
	}

	return false
}

// SetBankDebtRatio gets a reference to the given float32 and assigns it to the BankDebtRatio field.
func (o *SearchCompanyDataTypeIndebtedness) SetBankDebtRatio(v float32) {
	o.BankDebtRatio = &v
}

// GetDebtRatio returns the DebtRatio field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeIndebtedness) GetDebtRatio() float32 {
	if o == nil || IsNil(o.DebtRatio) {
		var ret float32
		return ret
	}
	return *o.DebtRatio
}

// GetDebtRatioOk returns a tuple with the DebtRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeIndebtedness) GetDebtRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.DebtRatio) {
		return nil, false
	}
	return o.DebtRatio, true
}

// HasDebtRatio returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeIndebtedness) HasDebtRatio() bool {
	if o != nil && !IsNil(o.DebtRatio) {
		return true
	}

	return false
}

// SetDebtRatio gets a reference to the given float32 and assigns it to the DebtRatio field.
func (o *SearchCompanyDataTypeIndebtedness) SetDebtRatio(v float32) {
	o.DebtRatio = &v
}

func (o SearchCompanyDataTypeIndebtedness) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeIndebtedness) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BankDebtTotalAssets) {
		toSerialize["bankDebtTotalAssets"] = o.BankDebtTotalAssets
	}
	if !IsNil(o.GrossFinancialDebtTotalAssets) {
		toSerialize["grossFinancialDebtTotalAssets"] = o.GrossFinancialDebtTotalAssets
	}
	if !IsNil(o.CapitalizationDegree) {
		toSerialize["capitalizationDegree"] = o.CapitalizationDegree
	}
	if !IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !IsNil(o.BankDebtRatio) {
		toSerialize["bankDebtRatio"] = o.BankDebtRatio
	}
	if !IsNil(o.DebtRatio) {
		toSerialize["debtRatio"] = o.DebtRatio
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeIndebtedness struct {
	value *SearchCompanyDataTypeIndebtedness
	isSet bool
}

func (v NullableSearchCompanyDataTypeIndebtedness) Get() *SearchCompanyDataTypeIndebtedness {
	return v.value
}

func (v *NullableSearchCompanyDataTypeIndebtedness) Set(val *SearchCompanyDataTypeIndebtedness) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeIndebtedness) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeIndebtedness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeIndebtedness(val *SearchCompanyDataTypeIndebtedness) *NullableSearchCompanyDataTypeIndebtedness {
	return &NullableSearchCompanyDataTypeIndebtedness{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeIndebtedness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeIndebtedness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

