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

// checks if the DownloadCompanyDataTypePortfolioInformationAreaManager type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadCompanyDataTypePortfolioInformationAreaManager{}

// DownloadCompanyDataTypePortfolioInformationAreaManager struct for DownloadCompanyDataTypePortfolioInformationAreaManager
type DownloadCompanyDataTypePortfolioInformationAreaManager struct {
	Id *string `json:"id,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewDownloadCompanyDataTypePortfolioInformationAreaManager instantiates a new DownloadCompanyDataTypePortfolioInformationAreaManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadCompanyDataTypePortfolioInformationAreaManager() *DownloadCompanyDataTypePortfolioInformationAreaManager {
	this := DownloadCompanyDataTypePortfolioInformationAreaManager{}
	return &this
}

// NewDownloadCompanyDataTypePortfolioInformationAreaManagerWithDefaults instantiates a new DownloadCompanyDataTypePortfolioInformationAreaManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadCompanyDataTypePortfolioInformationAreaManagerWithDefaults() *DownloadCompanyDataTypePortfolioInformationAreaManager {
	this := DownloadCompanyDataTypePortfolioInformationAreaManager{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DownloadCompanyDataTypePortfolioInformationAreaManager) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadCompanyDataTypePortfolioInformationAreaManager) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DownloadCompanyDataTypePortfolioInformationAreaManager) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DownloadCompanyDataTypePortfolioInformationAreaManager) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DownloadCompanyDataTypePortfolioInformationAreaManager) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadCompanyDataTypePortfolioInformationAreaManager) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DownloadCompanyDataTypePortfolioInformationAreaManager) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DownloadCompanyDataTypePortfolioInformationAreaManager) SetDescription(v string) {
	o.Description = &v
}

func (o DownloadCompanyDataTypePortfolioInformationAreaManager) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadCompanyDataTypePortfolioInformationAreaManager) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableDownloadCompanyDataTypePortfolioInformationAreaManager struct {
	value *DownloadCompanyDataTypePortfolioInformationAreaManager
	isSet bool
}

func (v NullableDownloadCompanyDataTypePortfolioInformationAreaManager) Get() *DownloadCompanyDataTypePortfolioInformationAreaManager {
	return v.value
}

func (v *NullableDownloadCompanyDataTypePortfolioInformationAreaManager) Set(val *DownloadCompanyDataTypePortfolioInformationAreaManager) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadCompanyDataTypePortfolioInformationAreaManager) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadCompanyDataTypePortfolioInformationAreaManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadCompanyDataTypePortfolioInformationAreaManager(val *DownloadCompanyDataTypePortfolioInformationAreaManager) *NullableDownloadCompanyDataTypePortfolioInformationAreaManager {
	return &NullableDownloadCompanyDataTypePortfolioInformationAreaManager{value: val, isSet: true}
}

func (v NullableDownloadCompanyDataTypePortfolioInformationAreaManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadCompanyDataTypePortfolioInformationAreaManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

