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

// checks if the PortfolioCreationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioCreationType{}

// PortfolioCreationType struct for PortfolioCreationType
type PortfolioCreationType struct {
	Name *string `json:"name,omitempty"`
	CustomVariables []CustomVariableType `json:"customVariables,omitempty"`
}

// NewPortfolioCreationType instantiates a new PortfolioCreationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioCreationType() *PortfolioCreationType {
	this := PortfolioCreationType{}
	return &this
}

// NewPortfolioCreationTypeWithDefaults instantiates a new PortfolioCreationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioCreationTypeWithDefaults() *PortfolioCreationType {
	this := PortfolioCreationType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PortfolioCreationType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioCreationType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PortfolioCreationType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PortfolioCreationType) SetName(v string) {
	o.Name = &v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *PortfolioCreationType) GetCustomVariables() []CustomVariableType {
	if o == nil || IsNil(o.CustomVariables) {
		var ret []CustomVariableType
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioCreationType) GetCustomVariablesOk() ([]CustomVariableType, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return nil, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *PortfolioCreationType) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given []CustomVariableType and assigns it to the CustomVariables field.
func (o *PortfolioCreationType) SetCustomVariables(v []CustomVariableType) {
	o.CustomVariables = v
}

func (o PortfolioCreationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioCreationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	return toSerialize, nil
}

type NullablePortfolioCreationType struct {
	value *PortfolioCreationType
	isSet bool
}

func (v NullablePortfolioCreationType) Get() *PortfolioCreationType {
	return v.value
}

func (v *NullablePortfolioCreationType) Set(val *PortfolioCreationType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioCreationType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioCreationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioCreationType(val *PortfolioCreationType) *NullablePortfolioCreationType {
	return &NullablePortfolioCreationType{value: val, isSet: true}
}

func (v NullablePortfolioCreationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioCreationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

