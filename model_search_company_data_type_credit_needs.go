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

// checks if the SearchCompanyDataTypeCreditNeeds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeCreditNeeds{}

// SearchCompanyDataTypeCreditNeeds struct for SearchCompanyDataTypeCreditNeeds
type SearchCompanyDataTypeCreditNeeds struct {
	ShortTermScore *SearchCompanyDataTypeCreditNeedsShortTermScore `json:"shortTermScore,omitempty"`
	MediumTermScore *SearchCompanyDataTypeCreditNeedsShortTermScore `json:"mediumTermScore,omitempty"`
	CarLeaseScore *SearchCompanyDataTypeCreditNeedsShortTermScore `json:"carLeaseScore,omitempty"`
	EquipmentLeaseScore *SearchCompanyDataTypeCreditNeedsEquipmentLeaseScore `json:"equipmentLeaseScore,omitempty"`
	RealEstateLeaseScore *SearchCompanyDataTypeCreditNeedsRealEstateLeaseScore `json:"realEstateLeaseScore,omitempty"`
	FactoringScore *SearchCompanyDataTypeCreditNeedsFactoringScore `json:"factoringScore,omitempty"`
}

// NewSearchCompanyDataTypeCreditNeeds instantiates a new SearchCompanyDataTypeCreditNeeds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeCreditNeeds() *SearchCompanyDataTypeCreditNeeds {
	this := SearchCompanyDataTypeCreditNeeds{}
	return &this
}

// NewSearchCompanyDataTypeCreditNeedsWithDefaults instantiates a new SearchCompanyDataTypeCreditNeeds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeCreditNeedsWithDefaults() *SearchCompanyDataTypeCreditNeeds {
	this := SearchCompanyDataTypeCreditNeeds{}
	return &this
}

// GetShortTermScore returns the ShortTermScore field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCreditNeeds) GetShortTermScore() SearchCompanyDataTypeCreditNeedsShortTermScore {
	if o == nil || IsNil(o.ShortTermScore) {
		var ret SearchCompanyDataTypeCreditNeedsShortTermScore
		return ret
	}
	return *o.ShortTermScore
}

// GetShortTermScoreOk returns a tuple with the ShortTermScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCreditNeeds) GetShortTermScoreOk() (*SearchCompanyDataTypeCreditNeedsShortTermScore, bool) {
	if o == nil || IsNil(o.ShortTermScore) {
		return nil, false
	}
	return o.ShortTermScore, true
}

// HasShortTermScore returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCreditNeeds) HasShortTermScore() bool {
	if o != nil && !IsNil(o.ShortTermScore) {
		return true
	}

	return false
}

// SetShortTermScore gets a reference to the given SearchCompanyDataTypeCreditNeedsShortTermScore and assigns it to the ShortTermScore field.
func (o *SearchCompanyDataTypeCreditNeeds) SetShortTermScore(v SearchCompanyDataTypeCreditNeedsShortTermScore) {
	o.ShortTermScore = &v
}

// GetMediumTermScore returns the MediumTermScore field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCreditNeeds) GetMediumTermScore() SearchCompanyDataTypeCreditNeedsShortTermScore {
	if o == nil || IsNil(o.MediumTermScore) {
		var ret SearchCompanyDataTypeCreditNeedsShortTermScore
		return ret
	}
	return *o.MediumTermScore
}

// GetMediumTermScoreOk returns a tuple with the MediumTermScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCreditNeeds) GetMediumTermScoreOk() (*SearchCompanyDataTypeCreditNeedsShortTermScore, bool) {
	if o == nil || IsNil(o.MediumTermScore) {
		return nil, false
	}
	return o.MediumTermScore, true
}

// HasMediumTermScore returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCreditNeeds) HasMediumTermScore() bool {
	if o != nil && !IsNil(o.MediumTermScore) {
		return true
	}

	return false
}

// SetMediumTermScore gets a reference to the given SearchCompanyDataTypeCreditNeedsShortTermScore and assigns it to the MediumTermScore field.
func (o *SearchCompanyDataTypeCreditNeeds) SetMediumTermScore(v SearchCompanyDataTypeCreditNeedsShortTermScore) {
	o.MediumTermScore = &v
}

// GetCarLeaseScore returns the CarLeaseScore field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCreditNeeds) GetCarLeaseScore() SearchCompanyDataTypeCreditNeedsShortTermScore {
	if o == nil || IsNil(o.CarLeaseScore) {
		var ret SearchCompanyDataTypeCreditNeedsShortTermScore
		return ret
	}
	return *o.CarLeaseScore
}

// GetCarLeaseScoreOk returns a tuple with the CarLeaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCreditNeeds) GetCarLeaseScoreOk() (*SearchCompanyDataTypeCreditNeedsShortTermScore, bool) {
	if o == nil || IsNil(o.CarLeaseScore) {
		return nil, false
	}
	return o.CarLeaseScore, true
}

// HasCarLeaseScore returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCreditNeeds) HasCarLeaseScore() bool {
	if o != nil && !IsNil(o.CarLeaseScore) {
		return true
	}

	return false
}

// SetCarLeaseScore gets a reference to the given SearchCompanyDataTypeCreditNeedsShortTermScore and assigns it to the CarLeaseScore field.
func (o *SearchCompanyDataTypeCreditNeeds) SetCarLeaseScore(v SearchCompanyDataTypeCreditNeedsShortTermScore) {
	o.CarLeaseScore = &v
}

// GetEquipmentLeaseScore returns the EquipmentLeaseScore field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCreditNeeds) GetEquipmentLeaseScore() SearchCompanyDataTypeCreditNeedsEquipmentLeaseScore {
	if o == nil || IsNil(o.EquipmentLeaseScore) {
		var ret SearchCompanyDataTypeCreditNeedsEquipmentLeaseScore
		return ret
	}
	return *o.EquipmentLeaseScore
}

// GetEquipmentLeaseScoreOk returns a tuple with the EquipmentLeaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCreditNeeds) GetEquipmentLeaseScoreOk() (*SearchCompanyDataTypeCreditNeedsEquipmentLeaseScore, bool) {
	if o == nil || IsNil(o.EquipmentLeaseScore) {
		return nil, false
	}
	return o.EquipmentLeaseScore, true
}

// HasEquipmentLeaseScore returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCreditNeeds) HasEquipmentLeaseScore() bool {
	if o != nil && !IsNil(o.EquipmentLeaseScore) {
		return true
	}

	return false
}

// SetEquipmentLeaseScore gets a reference to the given SearchCompanyDataTypeCreditNeedsEquipmentLeaseScore and assigns it to the EquipmentLeaseScore field.
func (o *SearchCompanyDataTypeCreditNeeds) SetEquipmentLeaseScore(v SearchCompanyDataTypeCreditNeedsEquipmentLeaseScore) {
	o.EquipmentLeaseScore = &v
}

// GetRealEstateLeaseScore returns the RealEstateLeaseScore field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCreditNeeds) GetRealEstateLeaseScore() SearchCompanyDataTypeCreditNeedsRealEstateLeaseScore {
	if o == nil || IsNil(o.RealEstateLeaseScore) {
		var ret SearchCompanyDataTypeCreditNeedsRealEstateLeaseScore
		return ret
	}
	return *o.RealEstateLeaseScore
}

// GetRealEstateLeaseScoreOk returns a tuple with the RealEstateLeaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCreditNeeds) GetRealEstateLeaseScoreOk() (*SearchCompanyDataTypeCreditNeedsRealEstateLeaseScore, bool) {
	if o == nil || IsNil(o.RealEstateLeaseScore) {
		return nil, false
	}
	return o.RealEstateLeaseScore, true
}

// HasRealEstateLeaseScore returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCreditNeeds) HasRealEstateLeaseScore() bool {
	if o != nil && !IsNil(o.RealEstateLeaseScore) {
		return true
	}

	return false
}

// SetRealEstateLeaseScore gets a reference to the given SearchCompanyDataTypeCreditNeedsRealEstateLeaseScore and assigns it to the RealEstateLeaseScore field.
func (o *SearchCompanyDataTypeCreditNeeds) SetRealEstateLeaseScore(v SearchCompanyDataTypeCreditNeedsRealEstateLeaseScore) {
	o.RealEstateLeaseScore = &v
}

// GetFactoringScore returns the FactoringScore field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCreditNeeds) GetFactoringScore() SearchCompanyDataTypeCreditNeedsFactoringScore {
	if o == nil || IsNil(o.FactoringScore) {
		var ret SearchCompanyDataTypeCreditNeedsFactoringScore
		return ret
	}
	return *o.FactoringScore
}

// GetFactoringScoreOk returns a tuple with the FactoringScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCreditNeeds) GetFactoringScoreOk() (*SearchCompanyDataTypeCreditNeedsFactoringScore, bool) {
	if o == nil || IsNil(o.FactoringScore) {
		return nil, false
	}
	return o.FactoringScore, true
}

// HasFactoringScore returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCreditNeeds) HasFactoringScore() bool {
	if o != nil && !IsNil(o.FactoringScore) {
		return true
	}

	return false
}

// SetFactoringScore gets a reference to the given SearchCompanyDataTypeCreditNeedsFactoringScore and assigns it to the FactoringScore field.
func (o *SearchCompanyDataTypeCreditNeeds) SetFactoringScore(v SearchCompanyDataTypeCreditNeedsFactoringScore) {
	o.FactoringScore = &v
}

func (o SearchCompanyDataTypeCreditNeeds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeCreditNeeds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShortTermScore) {
		toSerialize["shortTermScore"] = o.ShortTermScore
	}
	if !IsNil(o.MediumTermScore) {
		toSerialize["mediumTermScore"] = o.MediumTermScore
	}
	if !IsNil(o.CarLeaseScore) {
		toSerialize["carLeaseScore"] = o.CarLeaseScore
	}
	if !IsNil(o.EquipmentLeaseScore) {
		toSerialize["equipmentLeaseScore"] = o.EquipmentLeaseScore
	}
	if !IsNil(o.RealEstateLeaseScore) {
		toSerialize["realEstateLeaseScore"] = o.RealEstateLeaseScore
	}
	if !IsNil(o.FactoringScore) {
		toSerialize["factoringScore"] = o.FactoringScore
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeCreditNeeds struct {
	value *SearchCompanyDataTypeCreditNeeds
	isSet bool
}

func (v NullableSearchCompanyDataTypeCreditNeeds) Get() *SearchCompanyDataTypeCreditNeeds {
	return v.value
}

func (v *NullableSearchCompanyDataTypeCreditNeeds) Set(val *SearchCompanyDataTypeCreditNeeds) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeCreditNeeds) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeCreditNeeds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeCreditNeeds(val *SearchCompanyDataTypeCreditNeeds) *NullableSearchCompanyDataTypeCreditNeeds {
	return &NullableSearchCompanyDataTypeCreditNeeds{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeCreditNeeds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeCreditNeeds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


