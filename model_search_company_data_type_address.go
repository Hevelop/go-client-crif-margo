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

// checks if the SearchCompanyDataTypeAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeAddress{}

// SearchCompanyDataTypeAddress struct for SearchCompanyDataTypeAddress
type SearchCompanyDataTypeAddress struct {
	StreetName *string `json:"streetName,omitempty"`
	Hamlet *string `json:"hamlet,omitempty"`
	ZipCode *string `json:"zipCode,omitempty"`
	Town *string `json:"town,omitempty"`
	Province *SearchCompanyDataTypeAddressProvince `json:"province,omitempty"`
	Region *SearchCompanyDataTypeAddressRegion `json:"region,omitempty"`
	Country *SearchCompanyDataTypeAddressCountry `json:"country,omitempty"`
}

// NewSearchCompanyDataTypeAddress instantiates a new SearchCompanyDataTypeAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeAddress() *SearchCompanyDataTypeAddress {
	this := SearchCompanyDataTypeAddress{}
	return &this
}

// NewSearchCompanyDataTypeAddressWithDefaults instantiates a new SearchCompanyDataTypeAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeAddressWithDefaults() *SearchCompanyDataTypeAddress {
	this := SearchCompanyDataTypeAddress{}
	return &this
}

// GetStreetName returns the StreetName field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddress) GetStreetName() string {
	if o == nil || IsNil(o.StreetName) {
		var ret string
		return ret
	}
	return *o.StreetName
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddress) GetStreetNameOk() (*string, bool) {
	if o == nil || IsNil(o.StreetName) {
		return nil, false
	}
	return o.StreetName, true
}

// HasStreetName returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddress) HasStreetName() bool {
	if o != nil && !IsNil(o.StreetName) {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given string and assigns it to the StreetName field.
func (o *SearchCompanyDataTypeAddress) SetStreetName(v string) {
	o.StreetName = &v
}

// GetHamlet returns the Hamlet field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddress) GetHamlet() string {
	if o == nil || IsNil(o.Hamlet) {
		var ret string
		return ret
	}
	return *o.Hamlet
}

// GetHamletOk returns a tuple with the Hamlet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddress) GetHamletOk() (*string, bool) {
	if o == nil || IsNil(o.Hamlet) {
		return nil, false
	}
	return o.Hamlet, true
}

// HasHamlet returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddress) HasHamlet() bool {
	if o != nil && !IsNil(o.Hamlet) {
		return true
	}

	return false
}

// SetHamlet gets a reference to the given string and assigns it to the Hamlet field.
func (o *SearchCompanyDataTypeAddress) SetHamlet(v string) {
	o.Hamlet = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddress) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddress) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddress) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *SearchCompanyDataTypeAddress) SetZipCode(v string) {
	o.ZipCode = &v
}

// GetTown returns the Town field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddress) GetTown() string {
	if o == nil || IsNil(o.Town) {
		var ret string
		return ret
	}
	return *o.Town
}

// GetTownOk returns a tuple with the Town field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddress) GetTownOk() (*string, bool) {
	if o == nil || IsNil(o.Town) {
		return nil, false
	}
	return o.Town, true
}

// HasTown returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddress) HasTown() bool {
	if o != nil && !IsNil(o.Town) {
		return true
	}

	return false
}

// SetTown gets a reference to the given string and assigns it to the Town field.
func (o *SearchCompanyDataTypeAddress) SetTown(v string) {
	o.Town = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddress) GetProvince() SearchCompanyDataTypeAddressProvince {
	if o == nil || IsNil(o.Province) {
		var ret SearchCompanyDataTypeAddressProvince
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddress) GetProvinceOk() (*SearchCompanyDataTypeAddressProvince, bool) {
	if o == nil || IsNil(o.Province) {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddress) HasProvince() bool {
	if o != nil && !IsNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given SearchCompanyDataTypeAddressProvince and assigns it to the Province field.
func (o *SearchCompanyDataTypeAddress) SetProvince(v SearchCompanyDataTypeAddressProvince) {
	o.Province = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddress) GetRegion() SearchCompanyDataTypeAddressRegion {
	if o == nil || IsNil(o.Region) {
		var ret SearchCompanyDataTypeAddressRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddress) GetRegionOk() (*SearchCompanyDataTypeAddressRegion, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddress) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given SearchCompanyDataTypeAddressRegion and assigns it to the Region field.
func (o *SearchCompanyDataTypeAddress) SetRegion(v SearchCompanyDataTypeAddressRegion) {
	o.Region = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeAddress) GetCountry() SearchCompanyDataTypeAddressCountry {
	if o == nil || IsNil(o.Country) {
		var ret SearchCompanyDataTypeAddressCountry
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeAddress) GetCountryOk() (*SearchCompanyDataTypeAddressCountry, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given SearchCompanyDataTypeAddressCountry and assigns it to the Country field.
func (o *SearchCompanyDataTypeAddress) SetCountry(v SearchCompanyDataTypeAddressCountry) {
	o.Country = &v
}

func (o SearchCompanyDataTypeAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StreetName) {
		toSerialize["streetName"] = o.StreetName
	}
	if !IsNil(o.Hamlet) {
		toSerialize["hamlet"] = o.Hamlet
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zipCode"] = o.ZipCode
	}
	if !IsNil(o.Town) {
		toSerialize["town"] = o.Town
	}
	if !IsNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeAddress struct {
	value *SearchCompanyDataTypeAddress
	isSet bool
}

func (v NullableSearchCompanyDataTypeAddress) Get() *SearchCompanyDataTypeAddress {
	return v.value
}

func (v *NullableSearchCompanyDataTypeAddress) Set(val *SearchCompanyDataTypeAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeAddress(val *SearchCompanyDataTypeAddress) *NullableSearchCompanyDataTypeAddress {
	return &NullableSearchCompanyDataTypeAddress{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

