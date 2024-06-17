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

// checks if the SearchCompanyDataTypeTelcoAnalytics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeTelcoAnalytics{}

// SearchCompanyDataTypeTelcoAnalytics struct for SearchCompanyDataTypeTelcoAnalytics
type SearchCompanyDataTypeTelcoAnalytics struct {
	LandLineScore *SearchCompanyDataTypeTelcoAnalyticsLandLineScore `json:"landLineScore,omitempty"`
	MobileScore *SearchCompanyDataTypeTelcoAnalyticsMobileScore `json:"mobileScore,omitempty"`
}

// NewSearchCompanyDataTypeTelcoAnalytics instantiates a new SearchCompanyDataTypeTelcoAnalytics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeTelcoAnalytics() *SearchCompanyDataTypeTelcoAnalytics {
	this := SearchCompanyDataTypeTelcoAnalytics{}
	return &this
}

// NewSearchCompanyDataTypeTelcoAnalyticsWithDefaults instantiates a new SearchCompanyDataTypeTelcoAnalytics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeTelcoAnalyticsWithDefaults() *SearchCompanyDataTypeTelcoAnalytics {
	this := SearchCompanyDataTypeTelcoAnalytics{}
	return &this
}

// GetLandLineScore returns the LandLineScore field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeTelcoAnalytics) GetLandLineScore() SearchCompanyDataTypeTelcoAnalyticsLandLineScore {
	if o == nil || IsNil(o.LandLineScore) {
		var ret SearchCompanyDataTypeTelcoAnalyticsLandLineScore
		return ret
	}
	return *o.LandLineScore
}

// GetLandLineScoreOk returns a tuple with the LandLineScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeTelcoAnalytics) GetLandLineScoreOk() (*SearchCompanyDataTypeTelcoAnalyticsLandLineScore, bool) {
	if o == nil || IsNil(o.LandLineScore) {
		return nil, false
	}
	return o.LandLineScore, true
}

// HasLandLineScore returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeTelcoAnalytics) HasLandLineScore() bool {
	if o != nil && !IsNil(o.LandLineScore) {
		return true
	}

	return false
}

// SetLandLineScore gets a reference to the given SearchCompanyDataTypeTelcoAnalyticsLandLineScore and assigns it to the LandLineScore field.
func (o *SearchCompanyDataTypeTelcoAnalytics) SetLandLineScore(v SearchCompanyDataTypeTelcoAnalyticsLandLineScore) {
	o.LandLineScore = &v
}

// GetMobileScore returns the MobileScore field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeTelcoAnalytics) GetMobileScore() SearchCompanyDataTypeTelcoAnalyticsMobileScore {
	if o == nil || IsNil(o.MobileScore) {
		var ret SearchCompanyDataTypeTelcoAnalyticsMobileScore
		return ret
	}
	return *o.MobileScore
}

// GetMobileScoreOk returns a tuple with the MobileScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeTelcoAnalytics) GetMobileScoreOk() (*SearchCompanyDataTypeTelcoAnalyticsMobileScore, bool) {
	if o == nil || IsNil(o.MobileScore) {
		return nil, false
	}
	return o.MobileScore, true
}

// HasMobileScore returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeTelcoAnalytics) HasMobileScore() bool {
	if o != nil && !IsNil(o.MobileScore) {
		return true
	}

	return false
}

// SetMobileScore gets a reference to the given SearchCompanyDataTypeTelcoAnalyticsMobileScore and assigns it to the MobileScore field.
func (o *SearchCompanyDataTypeTelcoAnalytics) SetMobileScore(v SearchCompanyDataTypeTelcoAnalyticsMobileScore) {
	o.MobileScore = &v
}

func (o SearchCompanyDataTypeTelcoAnalytics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeTelcoAnalytics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LandLineScore) {
		toSerialize["landLineScore"] = o.LandLineScore
	}
	if !IsNil(o.MobileScore) {
		toSerialize["mobileScore"] = o.MobileScore
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeTelcoAnalytics struct {
	value *SearchCompanyDataTypeTelcoAnalytics
	isSet bool
}

func (v NullableSearchCompanyDataTypeTelcoAnalytics) Get() *SearchCompanyDataTypeTelcoAnalytics {
	return v.value
}

func (v *NullableSearchCompanyDataTypeTelcoAnalytics) Set(val *SearchCompanyDataTypeTelcoAnalytics) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeTelcoAnalytics) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeTelcoAnalytics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeTelcoAnalytics(val *SearchCompanyDataTypeTelcoAnalytics) *NullableSearchCompanyDataTypeTelcoAnalytics {
	return &NullableSearchCompanyDataTypeTelcoAnalytics{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeTelcoAnalytics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeTelcoAnalytics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

