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

// checks if the SearchCompanyDataTypeEmployees type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeEmployees{}

// SearchCompanyDataTypeEmployees struct for SearchCompanyDataTypeEmployees
type SearchCompanyDataTypeEmployees struct {
	EmployeeRange *SearchCompanyDataTypeEmployeesEmployeeRange `json:"employeeRange,omitempty"`
	Employee *float32 `json:"employee,omitempty"`
	EmployeeTrend *float32 `json:"employeeTrend,omitempty"`
}

// NewSearchCompanyDataTypeEmployees instantiates a new SearchCompanyDataTypeEmployees object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeEmployees() *SearchCompanyDataTypeEmployees {
	this := SearchCompanyDataTypeEmployees{}
	return &this
}

// NewSearchCompanyDataTypeEmployeesWithDefaults instantiates a new SearchCompanyDataTypeEmployees object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeEmployeesWithDefaults() *SearchCompanyDataTypeEmployees {
	this := SearchCompanyDataTypeEmployees{}
	return &this
}

// GetEmployeeRange returns the EmployeeRange field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeEmployees) GetEmployeeRange() SearchCompanyDataTypeEmployeesEmployeeRange {
	if o == nil || IsNil(o.EmployeeRange) {
		var ret SearchCompanyDataTypeEmployeesEmployeeRange
		return ret
	}
	return *o.EmployeeRange
}

// GetEmployeeRangeOk returns a tuple with the EmployeeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeEmployees) GetEmployeeRangeOk() (*SearchCompanyDataTypeEmployeesEmployeeRange, bool) {
	if o == nil || IsNil(o.EmployeeRange) {
		return nil, false
	}
	return o.EmployeeRange, true
}

// HasEmployeeRange returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeEmployees) HasEmployeeRange() bool {
	if o != nil && !IsNil(o.EmployeeRange) {
		return true
	}

	return false
}

// SetEmployeeRange gets a reference to the given SearchCompanyDataTypeEmployeesEmployeeRange and assigns it to the EmployeeRange field.
func (o *SearchCompanyDataTypeEmployees) SetEmployeeRange(v SearchCompanyDataTypeEmployeesEmployeeRange) {
	o.EmployeeRange = &v
}

// GetEmployee returns the Employee field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeEmployees) GetEmployee() float32 {
	if o == nil || IsNil(o.Employee) {
		var ret float32
		return ret
	}
	return *o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeEmployees) GetEmployeeOk() (*float32, bool) {
	if o == nil || IsNil(o.Employee) {
		return nil, false
	}
	return o.Employee, true
}

// HasEmployee returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeEmployees) HasEmployee() bool {
	if o != nil && !IsNil(o.Employee) {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given float32 and assigns it to the Employee field.
func (o *SearchCompanyDataTypeEmployees) SetEmployee(v float32) {
	o.Employee = &v
}

// GetEmployeeTrend returns the EmployeeTrend field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeEmployees) GetEmployeeTrend() float32 {
	if o == nil || IsNil(o.EmployeeTrend) {
		var ret float32
		return ret
	}
	return *o.EmployeeTrend
}

// GetEmployeeTrendOk returns a tuple with the EmployeeTrend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeEmployees) GetEmployeeTrendOk() (*float32, bool) {
	if o == nil || IsNil(o.EmployeeTrend) {
		return nil, false
	}
	return o.EmployeeTrend, true
}

// HasEmployeeTrend returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeEmployees) HasEmployeeTrend() bool {
	if o != nil && !IsNil(o.EmployeeTrend) {
		return true
	}

	return false
}

// SetEmployeeTrend gets a reference to the given float32 and assigns it to the EmployeeTrend field.
func (o *SearchCompanyDataTypeEmployees) SetEmployeeTrend(v float32) {
	o.EmployeeTrend = &v
}

func (o SearchCompanyDataTypeEmployees) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeEmployees) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmployeeRange) {
		toSerialize["employeeRange"] = o.EmployeeRange
	}
	if !IsNil(o.Employee) {
		toSerialize["employee"] = o.Employee
	}
	if !IsNil(o.EmployeeTrend) {
		toSerialize["employeeTrend"] = o.EmployeeTrend
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeEmployees struct {
	value *SearchCompanyDataTypeEmployees
	isSet bool
}

func (v NullableSearchCompanyDataTypeEmployees) Get() *SearchCompanyDataTypeEmployees {
	return v.value
}

func (v *NullableSearchCompanyDataTypeEmployees) Set(val *SearchCompanyDataTypeEmployees) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeEmployees) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeEmployees) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeEmployees(val *SearchCompanyDataTypeEmployees) *NullableSearchCompanyDataTypeEmployees {
	return &NullableSearchCompanyDataTypeEmployees{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeEmployees) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeEmployees) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


