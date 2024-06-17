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

// checks if the PortfolioType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioType{}

// PortfolioType struct for PortfolioType
type PortfolioType struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TotalCompanies *int32 `json:"totalCompanies,omitempty"`
	UserCode *string `json:"userCode,omitempty"`
	IsMonitored *bool `json:"isMonitored,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Status *string `json:"status,omitempty"`
	CompaniesPortfolioStatus *PortfolioTypeCompaniesPortfolioStatus `json:"companiesPortfolioStatus,omitempty"`
	CustomVariables []CustomVariableType `json:"customVariables,omitempty"`
}

// NewPortfolioType instantiates a new PortfolioType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioType() *PortfolioType {
	this := PortfolioType{}
	return &this
}

// NewPortfolioTypeWithDefaults instantiates a new PortfolioType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioTypeWithDefaults() *PortfolioType {
	this := PortfolioType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortfolioType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortfolioType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortfolioType) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PortfolioType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PortfolioType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PortfolioType) SetName(v string) {
	o.Name = &v
}

// GetTotalCompanies returns the TotalCompanies field value if set, zero value otherwise.
func (o *PortfolioType) GetTotalCompanies() int32 {
	if o == nil || IsNil(o.TotalCompanies) {
		var ret int32
		return ret
	}
	return *o.TotalCompanies
}

// GetTotalCompaniesOk returns a tuple with the TotalCompanies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetTotalCompaniesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCompanies) {
		return nil, false
	}
	return o.TotalCompanies, true
}

// HasTotalCompanies returns a boolean if a field has been set.
func (o *PortfolioType) HasTotalCompanies() bool {
	if o != nil && !IsNil(o.TotalCompanies) {
		return true
	}

	return false
}

// SetTotalCompanies gets a reference to the given int32 and assigns it to the TotalCompanies field.
func (o *PortfolioType) SetTotalCompanies(v int32) {
	o.TotalCompanies = &v
}

// GetUserCode returns the UserCode field value if set, zero value otherwise.
func (o *PortfolioType) GetUserCode() string {
	if o == nil || IsNil(o.UserCode) {
		var ret string
		return ret
	}
	return *o.UserCode
}

// GetUserCodeOk returns a tuple with the UserCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetUserCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UserCode) {
		return nil, false
	}
	return o.UserCode, true
}

// HasUserCode returns a boolean if a field has been set.
func (o *PortfolioType) HasUserCode() bool {
	if o != nil && !IsNil(o.UserCode) {
		return true
	}

	return false
}

// SetUserCode gets a reference to the given string and assigns it to the UserCode field.
func (o *PortfolioType) SetUserCode(v string) {
	o.UserCode = &v
}

// GetIsMonitored returns the IsMonitored field value if set, zero value otherwise.
func (o *PortfolioType) GetIsMonitored() bool {
	if o == nil || IsNil(o.IsMonitored) {
		var ret bool
		return ret
	}
	return *o.IsMonitored
}

// GetIsMonitoredOk returns a tuple with the IsMonitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetIsMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMonitored) {
		return nil, false
	}
	return o.IsMonitored, true
}

// HasIsMonitored returns a boolean if a field has been set.
func (o *PortfolioType) HasIsMonitored() bool {
	if o != nil && !IsNil(o.IsMonitored) {
		return true
	}

	return false
}

// SetIsMonitored gets a reference to the given bool and assigns it to the IsMonitored field.
func (o *PortfolioType) SetIsMonitored(v bool) {
	o.IsMonitored = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PortfolioType) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PortfolioType) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PortfolioType) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PortfolioType) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PortfolioType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PortfolioType) SetStatus(v string) {
	o.Status = &v
}

// GetCompaniesPortfolioStatus returns the CompaniesPortfolioStatus field value if set, zero value otherwise.
func (o *PortfolioType) GetCompaniesPortfolioStatus() PortfolioTypeCompaniesPortfolioStatus {
	if o == nil || IsNil(o.CompaniesPortfolioStatus) {
		var ret PortfolioTypeCompaniesPortfolioStatus
		return ret
	}
	return *o.CompaniesPortfolioStatus
}

// GetCompaniesPortfolioStatusOk returns a tuple with the CompaniesPortfolioStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetCompaniesPortfolioStatusOk() (*PortfolioTypeCompaniesPortfolioStatus, bool) {
	if o == nil || IsNil(o.CompaniesPortfolioStatus) {
		return nil, false
	}
	return o.CompaniesPortfolioStatus, true
}

// HasCompaniesPortfolioStatus returns a boolean if a field has been set.
func (o *PortfolioType) HasCompaniesPortfolioStatus() bool {
	if o != nil && !IsNil(o.CompaniesPortfolioStatus) {
		return true
	}

	return false
}

// SetCompaniesPortfolioStatus gets a reference to the given PortfolioTypeCompaniesPortfolioStatus and assigns it to the CompaniesPortfolioStatus field.
func (o *PortfolioType) SetCompaniesPortfolioStatus(v PortfolioTypeCompaniesPortfolioStatus) {
	o.CompaniesPortfolioStatus = &v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *PortfolioType) GetCustomVariables() []CustomVariableType {
	if o == nil || IsNil(o.CustomVariables) {
		var ret []CustomVariableType
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioType) GetCustomVariablesOk() ([]CustomVariableType, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return nil, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *PortfolioType) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given []CustomVariableType and assigns it to the CustomVariables field.
func (o *PortfolioType) SetCustomVariables(v []CustomVariableType) {
	o.CustomVariables = v
}

func (o PortfolioType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TotalCompanies) {
		toSerialize["totalCompanies"] = o.TotalCompanies
	}
	if !IsNil(o.UserCode) {
		toSerialize["userCode"] = o.UserCode
	}
	if !IsNil(o.IsMonitored) {
		toSerialize["isMonitored"] = o.IsMonitored
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CompaniesPortfolioStatus) {
		toSerialize["companiesPortfolioStatus"] = o.CompaniesPortfolioStatus
	}
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	return toSerialize, nil
}

type NullablePortfolioType struct {
	value *PortfolioType
	isSet bool
}

func (v NullablePortfolioType) Get() *PortfolioType {
	return v.value
}

func (v *NullablePortfolioType) Set(val *PortfolioType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioType(val *PortfolioType) *NullablePortfolioType {
	return &NullablePortfolioType{value: val, isSet: true}
}

func (v NullablePortfolioType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

