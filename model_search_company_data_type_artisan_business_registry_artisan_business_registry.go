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
	"time"
)

// checks if the SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry{}

// SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry struct for SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry
type SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry struct {
	RegistrationDate *time.Time `json:"registrationDate,omitempty"`
	RegistrationNumber *float32 `json:"registrationNumber,omitempty"`
}

// NewSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry instantiates a new SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry() *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry {
	this := SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry{}
	return &this
}

// NewSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistryWithDefaults instantiates a new SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistryWithDefaults() *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry {
	this := SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry{}
	return &this
}

// GetRegistrationDate returns the RegistrationDate field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) GetRegistrationDate() time.Time {
	if o == nil || IsNil(o.RegistrationDate) {
		var ret time.Time
		return ret
	}
	return *o.RegistrationDate
}

// GetRegistrationDateOk returns a tuple with the RegistrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) GetRegistrationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RegistrationDate) {
		return nil, false
	}
	return o.RegistrationDate, true
}

// HasRegistrationDate returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) HasRegistrationDate() bool {
	if o != nil && !IsNil(o.RegistrationDate) {
		return true
	}

	return false
}

// SetRegistrationDate gets a reference to the given time.Time and assigns it to the RegistrationDate field.
func (o *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) SetRegistrationDate(v time.Time) {
	o.RegistrationDate = &v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) GetRegistrationNumber() float32 {
	if o == nil || IsNil(o.RegistrationNumber) {
		var ret float32
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) GetRegistrationNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.RegistrationNumber) {
		return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) HasRegistrationNumber() bool {
	if o != nil && !IsNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given float32 and assigns it to the RegistrationNumber field.
func (o *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) SetRegistrationNumber(v float32) {
	o.RegistrationNumber = &v
}

func (o SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegistrationDate) {
		toSerialize["registrationDate"] = o.RegistrationDate
	}
	if !IsNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry struct {
	value *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry
	isSet bool
}

func (v NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) Get() *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry {
	return v.value
}

func (v *NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) Set(val *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry(val *SearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) *NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry {
	return &NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeArtisanBusinessRegistryArtisanBusinessRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

