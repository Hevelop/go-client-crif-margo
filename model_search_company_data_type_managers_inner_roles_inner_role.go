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

// checks if the SearchCompanyDataTypeManagersInnerRolesInnerRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeManagersInnerRolesInnerRole{}

// SearchCompanyDataTypeManagersInnerRolesInnerRole struct for SearchCompanyDataTypeManagersInnerRolesInnerRole
type SearchCompanyDataTypeManagersInnerRolesInnerRole struct {
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewSearchCompanyDataTypeManagersInnerRolesInnerRole instantiates a new SearchCompanyDataTypeManagersInnerRolesInnerRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeManagersInnerRolesInnerRole() *SearchCompanyDataTypeManagersInnerRolesInnerRole {
	this := SearchCompanyDataTypeManagersInnerRolesInnerRole{}
	return &this
}

// NewSearchCompanyDataTypeManagersInnerRolesInnerRoleWithDefaults instantiates a new SearchCompanyDataTypeManagersInnerRolesInnerRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeManagersInnerRolesInnerRoleWithDefaults() *SearchCompanyDataTypeManagersInnerRolesInnerRole {
	this := SearchCompanyDataTypeManagersInnerRolesInnerRole{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInnerRolesInnerRole) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInnerRolesInnerRole) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInnerRolesInnerRole) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SearchCompanyDataTypeManagersInnerRolesInnerRole) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInnerRolesInnerRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInnerRolesInnerRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInnerRolesInnerRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchCompanyDataTypeManagersInnerRolesInnerRole) SetDescription(v string) {
	o.Description = &v
}

func (o SearchCompanyDataTypeManagersInnerRolesInnerRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeManagersInnerRolesInnerRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeManagersInnerRolesInnerRole struct {
	value *SearchCompanyDataTypeManagersInnerRolesInnerRole
	isSet bool
}

func (v NullableSearchCompanyDataTypeManagersInnerRolesInnerRole) Get() *SearchCompanyDataTypeManagersInnerRolesInnerRole {
	return v.value
}

func (v *NullableSearchCompanyDataTypeManagersInnerRolesInnerRole) Set(val *SearchCompanyDataTypeManagersInnerRolesInnerRole) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeManagersInnerRolesInnerRole) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeManagersInnerRolesInnerRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeManagersInnerRolesInnerRole(val *SearchCompanyDataTypeManagersInnerRolesInnerRole) *NullableSearchCompanyDataTypeManagersInnerRolesInnerRole {
	return &NullableSearchCompanyDataTypeManagersInnerRolesInnerRole{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeManagersInnerRolesInnerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeManagersInnerRolesInnerRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


