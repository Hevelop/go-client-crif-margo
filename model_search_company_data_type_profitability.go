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

// checks if the SearchCompanyDataTypeProfitability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeProfitability{}

// SearchCompanyDataTypeProfitability struct for SearchCompanyDataTypeProfitability
type SearchCompanyDataTypeProfitability struct {
	Ros *float32 `json:"ros,omitempty"`
	Roe *float32 `json:"roe,omitempty"`
	IncidenceOfExtraFeaturesManagement *float32 `json:"incidenceOfExtraFeaturesManagement,omitempty"`
	Roi *float32 `json:"roi,omitempty"`
	RoaMonetary *float32 `json:"roaMonetary,omitempty"`
}

// NewSearchCompanyDataTypeProfitability instantiates a new SearchCompanyDataTypeProfitability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeProfitability() *SearchCompanyDataTypeProfitability {
	this := SearchCompanyDataTypeProfitability{}
	return &this
}

// NewSearchCompanyDataTypeProfitabilityWithDefaults instantiates a new SearchCompanyDataTypeProfitability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeProfitabilityWithDefaults() *SearchCompanyDataTypeProfitability {
	this := SearchCompanyDataTypeProfitability{}
	return &this
}

// GetRos returns the Ros field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeProfitability) GetRos() float32 {
	if o == nil || IsNil(o.Ros) {
		var ret float32
		return ret
	}
	return *o.Ros
}

// GetRosOk returns a tuple with the Ros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeProfitability) GetRosOk() (*float32, bool) {
	if o == nil || IsNil(o.Ros) {
		return nil, false
	}
	return o.Ros, true
}

// HasRos returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeProfitability) HasRos() bool {
	if o != nil && !IsNil(o.Ros) {
		return true
	}

	return false
}

// SetRos gets a reference to the given float32 and assigns it to the Ros field.
func (o *SearchCompanyDataTypeProfitability) SetRos(v float32) {
	o.Ros = &v
}

// GetRoe returns the Roe field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeProfitability) GetRoe() float32 {
	if o == nil || IsNil(o.Roe) {
		var ret float32
		return ret
	}
	return *o.Roe
}

// GetRoeOk returns a tuple with the Roe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeProfitability) GetRoeOk() (*float32, bool) {
	if o == nil || IsNil(o.Roe) {
		return nil, false
	}
	return o.Roe, true
}

// HasRoe returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeProfitability) HasRoe() bool {
	if o != nil && !IsNil(o.Roe) {
		return true
	}

	return false
}

// SetRoe gets a reference to the given float32 and assigns it to the Roe field.
func (o *SearchCompanyDataTypeProfitability) SetRoe(v float32) {
	o.Roe = &v
}

// GetIncidenceOfExtraFeaturesManagement returns the IncidenceOfExtraFeaturesManagement field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeProfitability) GetIncidenceOfExtraFeaturesManagement() float32 {
	if o == nil || IsNil(o.IncidenceOfExtraFeaturesManagement) {
		var ret float32
		return ret
	}
	return *o.IncidenceOfExtraFeaturesManagement
}

// GetIncidenceOfExtraFeaturesManagementOk returns a tuple with the IncidenceOfExtraFeaturesManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeProfitability) GetIncidenceOfExtraFeaturesManagementOk() (*float32, bool) {
	if o == nil || IsNil(o.IncidenceOfExtraFeaturesManagement) {
		return nil, false
	}
	return o.IncidenceOfExtraFeaturesManagement, true
}

// HasIncidenceOfExtraFeaturesManagement returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeProfitability) HasIncidenceOfExtraFeaturesManagement() bool {
	if o != nil && !IsNil(o.IncidenceOfExtraFeaturesManagement) {
		return true
	}

	return false
}

// SetIncidenceOfExtraFeaturesManagement gets a reference to the given float32 and assigns it to the IncidenceOfExtraFeaturesManagement field.
func (o *SearchCompanyDataTypeProfitability) SetIncidenceOfExtraFeaturesManagement(v float32) {
	o.IncidenceOfExtraFeaturesManagement = &v
}

// GetRoi returns the Roi field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeProfitability) GetRoi() float32 {
	if o == nil || IsNil(o.Roi) {
		var ret float32
		return ret
	}
	return *o.Roi
}

// GetRoiOk returns a tuple with the Roi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeProfitability) GetRoiOk() (*float32, bool) {
	if o == nil || IsNil(o.Roi) {
		return nil, false
	}
	return o.Roi, true
}

// HasRoi returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeProfitability) HasRoi() bool {
	if o != nil && !IsNil(o.Roi) {
		return true
	}

	return false
}

// SetRoi gets a reference to the given float32 and assigns it to the Roi field.
func (o *SearchCompanyDataTypeProfitability) SetRoi(v float32) {
	o.Roi = &v
}

// GetRoaMonetary returns the RoaMonetary field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeProfitability) GetRoaMonetary() float32 {
	if o == nil || IsNil(o.RoaMonetary) {
		var ret float32
		return ret
	}
	return *o.RoaMonetary
}

// GetRoaMonetaryOk returns a tuple with the RoaMonetary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeProfitability) GetRoaMonetaryOk() (*float32, bool) {
	if o == nil || IsNil(o.RoaMonetary) {
		return nil, false
	}
	return o.RoaMonetary, true
}

// HasRoaMonetary returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeProfitability) HasRoaMonetary() bool {
	if o != nil && !IsNil(o.RoaMonetary) {
		return true
	}

	return false
}

// SetRoaMonetary gets a reference to the given float32 and assigns it to the RoaMonetary field.
func (o *SearchCompanyDataTypeProfitability) SetRoaMonetary(v float32) {
	o.RoaMonetary = &v
}

func (o SearchCompanyDataTypeProfitability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeProfitability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ros) {
		toSerialize["ros"] = o.Ros
	}
	if !IsNil(o.Roe) {
		toSerialize["roe"] = o.Roe
	}
	if !IsNil(o.IncidenceOfExtraFeaturesManagement) {
		toSerialize["incidenceOfExtraFeaturesManagement"] = o.IncidenceOfExtraFeaturesManagement
	}
	if !IsNil(o.Roi) {
		toSerialize["roi"] = o.Roi
	}
	if !IsNil(o.RoaMonetary) {
		toSerialize["roaMonetary"] = o.RoaMonetary
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeProfitability struct {
	value *SearchCompanyDataTypeProfitability
	isSet bool
}

func (v NullableSearchCompanyDataTypeProfitability) Get() *SearchCompanyDataTypeProfitability {
	return v.value
}

func (v *NullableSearchCompanyDataTypeProfitability) Set(val *SearchCompanyDataTypeProfitability) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeProfitability) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeProfitability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeProfitability(val *SearchCompanyDataTypeProfitability) *NullableSearchCompanyDataTypeProfitability {
	return &NullableSearchCompanyDataTypeProfitability{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeProfitability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeProfitability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


