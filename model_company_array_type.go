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

// checks if the CompanyArrayType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyArrayType{}

// CompanyArrayType Array of Companies
type CompanyArrayType struct {
	Companies []CompanyType `json:"companies,omitempty"`
}

// NewCompanyArrayType instantiates a new CompanyArrayType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyArrayType() *CompanyArrayType {
	this := CompanyArrayType{}
	return &this
}

// NewCompanyArrayTypeWithDefaults instantiates a new CompanyArrayType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyArrayTypeWithDefaults() *CompanyArrayType {
	this := CompanyArrayType{}
	return &this
}

// GetCompanies returns the Companies field value if set, zero value otherwise.
func (o *CompanyArrayType) GetCompanies() []CompanyType {
	if o == nil || IsNil(o.Companies) {
		var ret []CompanyType
		return ret
	}
	return o.Companies
}

// GetCompaniesOk returns a tuple with the Companies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyArrayType) GetCompaniesOk() ([]CompanyType, bool) {
	if o == nil || IsNil(o.Companies) {
		return nil, false
	}
	return o.Companies, true
}

// HasCompanies returns a boolean if a field has been set.
func (o *CompanyArrayType) HasCompanies() bool {
	if o != nil && !IsNil(o.Companies) {
		return true
	}

	return false
}

// SetCompanies gets a reference to the given []CompanyType and assigns it to the Companies field.
func (o *CompanyArrayType) SetCompanies(v []CompanyType) {
	o.Companies = v
}

func (o CompanyArrayType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyArrayType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Companies) {
		toSerialize["companies"] = o.Companies
	}
	return toSerialize, nil
}

type NullableCompanyArrayType struct {
	value *CompanyArrayType
	isSet bool
}

func (v NullableCompanyArrayType) Get() *CompanyArrayType {
	return v.value
}

func (v *NullableCompanyArrayType) Set(val *CompanyArrayType) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyArrayType) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyArrayType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyArrayType(val *CompanyArrayType) *NullableCompanyArrayType {
	return &NullableCompanyArrayType{value: val, isSet: true}
}

func (v NullableCompanyArrayType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyArrayType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


