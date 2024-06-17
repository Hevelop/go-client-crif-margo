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

// checks if the PortfolioSummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioSummaryType{}

// PortfolioSummaryType struct for PortfolioSummaryType
type PortfolioSummaryType struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	UserCode *string `json:"userCode,omitempty"`
	TotalCompanies *int32 `json:"totalCompanies,omitempty"`
	IsMonitored *bool `json:"isMonitored,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewPortfolioSummaryType instantiates a new PortfolioSummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioSummaryType() *PortfolioSummaryType {
	this := PortfolioSummaryType{}
	return &this
}

// NewPortfolioSummaryTypeWithDefaults instantiates a new PortfolioSummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioSummaryTypeWithDefaults() *PortfolioSummaryType {
	this := PortfolioSummaryType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortfolioSummaryType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSummaryType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortfolioSummaryType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortfolioSummaryType) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PortfolioSummaryType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSummaryType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PortfolioSummaryType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PortfolioSummaryType) SetName(v string) {
	o.Name = &v
}

// GetUserCode returns the UserCode field value if set, zero value otherwise.
func (o *PortfolioSummaryType) GetUserCode() string {
	if o == nil || IsNil(o.UserCode) {
		var ret string
		return ret
	}
	return *o.UserCode
}

// GetUserCodeOk returns a tuple with the UserCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSummaryType) GetUserCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UserCode) {
		return nil, false
	}
	return o.UserCode, true
}

// HasUserCode returns a boolean if a field has been set.
func (o *PortfolioSummaryType) HasUserCode() bool {
	if o != nil && !IsNil(o.UserCode) {
		return true
	}

	return false
}

// SetUserCode gets a reference to the given string and assigns it to the UserCode field.
func (o *PortfolioSummaryType) SetUserCode(v string) {
	o.UserCode = &v
}

// GetTotalCompanies returns the TotalCompanies field value if set, zero value otherwise.
func (o *PortfolioSummaryType) GetTotalCompanies() int32 {
	if o == nil || IsNil(o.TotalCompanies) {
		var ret int32
		return ret
	}
	return *o.TotalCompanies
}

// GetTotalCompaniesOk returns a tuple with the TotalCompanies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSummaryType) GetTotalCompaniesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCompanies) {
		return nil, false
	}
	return o.TotalCompanies, true
}

// HasTotalCompanies returns a boolean if a field has been set.
func (o *PortfolioSummaryType) HasTotalCompanies() bool {
	if o != nil && !IsNil(o.TotalCompanies) {
		return true
	}

	return false
}

// SetTotalCompanies gets a reference to the given int32 and assigns it to the TotalCompanies field.
func (o *PortfolioSummaryType) SetTotalCompanies(v int32) {
	o.TotalCompanies = &v
}

// GetIsMonitored returns the IsMonitored field value if set, zero value otherwise.
func (o *PortfolioSummaryType) GetIsMonitored() bool {
	if o == nil || IsNil(o.IsMonitored) {
		var ret bool
		return ret
	}
	return *o.IsMonitored
}

// GetIsMonitoredOk returns a tuple with the IsMonitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSummaryType) GetIsMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMonitored) {
		return nil, false
	}
	return o.IsMonitored, true
}

// HasIsMonitored returns a boolean if a field has been set.
func (o *PortfolioSummaryType) HasIsMonitored() bool {
	if o != nil && !IsNil(o.IsMonitored) {
		return true
	}

	return false
}

// SetIsMonitored gets a reference to the given bool and assigns it to the IsMonitored field.
func (o *PortfolioSummaryType) SetIsMonitored(v bool) {
	o.IsMonitored = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PortfolioSummaryType) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSummaryType) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PortfolioSummaryType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PortfolioSummaryType) SetStatus(v string) {
	o.Status = &v
}

func (o PortfolioSummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioSummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UserCode) {
		toSerialize["userCode"] = o.UserCode
	}
	if !IsNil(o.TotalCompanies) {
		toSerialize["totalCompanies"] = o.TotalCompanies
	}
	if !IsNil(o.IsMonitored) {
		toSerialize["isMonitored"] = o.IsMonitored
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePortfolioSummaryType struct {
	value *PortfolioSummaryType
	isSet bool
}

func (v NullablePortfolioSummaryType) Get() *PortfolioSummaryType {
	return v.value
}

func (v *NullablePortfolioSummaryType) Set(val *PortfolioSummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioSummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioSummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioSummaryType(val *PortfolioSummaryType) *NullablePortfolioSummaryType {
	return &NullablePortfolioSummaryType{value: val, isSet: true}
}

func (v NullablePortfolioSummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioSummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


