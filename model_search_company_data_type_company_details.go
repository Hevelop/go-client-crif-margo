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

// checks if the SearchCompanyDataTypeCompanyDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeCompanyDetails{}

// SearchCompanyDataTypeCompanyDetails struct for SearchCompanyDataTypeCompanyDetails
type SearchCompanyDataTypeCompanyDetails struct {
	CompanyName *string `json:"companyName,omitempty"`
	VatCode *string `json:"vatCode,omitempty"`
	TaxCode *string `json:"taxCode,omitempty"`
	LastUpdateDate *time.Time `json:"lastUpdateDate,omitempty"`
	Cciaa *string `json:"cciaa,omitempty"`
	ReaCode *string `json:"reaCode,omitempty"`
	CrifNumber *string `json:"crifNumber,omitempty"`
	OfficeType *SearchCompanyDataTypeOfficeType `json:"officeType,omitempty"`
}

// NewSearchCompanyDataTypeCompanyDetails instantiates a new SearchCompanyDataTypeCompanyDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeCompanyDetails() *SearchCompanyDataTypeCompanyDetails {
	this := SearchCompanyDataTypeCompanyDetails{}
	return &this
}

// NewSearchCompanyDataTypeCompanyDetailsWithDefaults instantiates a new SearchCompanyDataTypeCompanyDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeCompanyDetailsWithDefaults() *SearchCompanyDataTypeCompanyDetails {
	this := SearchCompanyDataTypeCompanyDetails{}
	return &this
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCompanyDetails) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCompanyDetails) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCompanyDetails) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *SearchCompanyDataTypeCompanyDetails) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetVatCode returns the VatCode field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCompanyDetails) GetVatCode() string {
	if o == nil || IsNil(o.VatCode) {
		var ret string
		return ret
	}
	return *o.VatCode
}

// GetVatCodeOk returns a tuple with the VatCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCompanyDetails) GetVatCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VatCode) {
		return nil, false
	}
	return o.VatCode, true
}

// HasVatCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCompanyDetails) HasVatCode() bool {
	if o != nil && !IsNil(o.VatCode) {
		return true
	}

	return false
}

// SetVatCode gets a reference to the given string and assigns it to the VatCode field.
func (o *SearchCompanyDataTypeCompanyDetails) SetVatCode(v string) {
	o.VatCode = &v
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCompanyDetails) GetTaxCode() string {
	if o == nil || IsNil(o.TaxCode) {
		var ret string
		return ret
	}
	return *o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCompanyDetails) GetTaxCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxCode) {
		return nil, false
	}
	return o.TaxCode, true
}

// HasTaxCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCompanyDetails) HasTaxCode() bool {
	if o != nil && !IsNil(o.TaxCode) {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given string and assigns it to the TaxCode field.
func (o *SearchCompanyDataTypeCompanyDetails) SetTaxCode(v string) {
	o.TaxCode = &v
}

// GetLastUpdateDate returns the LastUpdateDate field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCompanyDetails) GetLastUpdateDate() time.Time {
	if o == nil || IsNil(o.LastUpdateDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDate
}

// GetLastUpdateDateOk returns a tuple with the LastUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCompanyDetails) GetLastUpdateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDate) {
		return nil, false
	}
	return o.LastUpdateDate, true
}

// HasLastUpdateDate returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCompanyDetails) HasLastUpdateDate() bool {
	if o != nil && !IsNil(o.LastUpdateDate) {
		return true
	}

	return false
}

// SetLastUpdateDate gets a reference to the given time.Time and assigns it to the LastUpdateDate field.
func (o *SearchCompanyDataTypeCompanyDetails) SetLastUpdateDate(v time.Time) {
	o.LastUpdateDate = &v
}

// GetCciaa returns the Cciaa field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCompanyDetails) GetCciaa() string {
	if o == nil || IsNil(o.Cciaa) {
		var ret string
		return ret
	}
	return *o.Cciaa
}

// GetCciaaOk returns a tuple with the Cciaa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCompanyDetails) GetCciaaOk() (*string, bool) {
	if o == nil || IsNil(o.Cciaa) {
		return nil, false
	}
	return o.Cciaa, true
}

// HasCciaa returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCompanyDetails) HasCciaa() bool {
	if o != nil && !IsNil(o.Cciaa) {
		return true
	}

	return false
}

// SetCciaa gets a reference to the given string and assigns it to the Cciaa field.
func (o *SearchCompanyDataTypeCompanyDetails) SetCciaa(v string) {
	o.Cciaa = &v
}

// GetReaCode returns the ReaCode field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCompanyDetails) GetReaCode() string {
	if o == nil || IsNil(o.ReaCode) {
		var ret string
		return ret
	}
	return *o.ReaCode
}

// GetReaCodeOk returns a tuple with the ReaCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCompanyDetails) GetReaCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReaCode) {
		return nil, false
	}
	return o.ReaCode, true
}

// HasReaCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCompanyDetails) HasReaCode() bool {
	if o != nil && !IsNil(o.ReaCode) {
		return true
	}

	return false
}

// SetReaCode gets a reference to the given string and assigns it to the ReaCode field.
func (o *SearchCompanyDataTypeCompanyDetails) SetReaCode(v string) {
	o.ReaCode = &v
}

// GetCrifNumber returns the CrifNumber field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCompanyDetails) GetCrifNumber() string {
	if o == nil || IsNil(o.CrifNumber) {
		var ret string
		return ret
	}
	return *o.CrifNumber
}

// GetCrifNumberOk returns a tuple with the CrifNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCompanyDetails) GetCrifNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CrifNumber) {
		return nil, false
	}
	return o.CrifNumber, true
}

// HasCrifNumber returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCompanyDetails) HasCrifNumber() bool {
	if o != nil && !IsNil(o.CrifNumber) {
		return true
	}

	return false
}

// SetCrifNumber gets a reference to the given string and assigns it to the CrifNumber field.
func (o *SearchCompanyDataTypeCompanyDetails) SetCrifNumber(v string) {
	o.CrifNumber = &v
}

// GetOfficeType returns the OfficeType field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeCompanyDetails) GetOfficeType() SearchCompanyDataTypeOfficeType {
	if o == nil || IsNil(o.OfficeType) {
		var ret SearchCompanyDataTypeOfficeType
		return ret
	}
	return *o.OfficeType
}

// GetOfficeTypeOk returns a tuple with the OfficeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeCompanyDetails) GetOfficeTypeOk() (*SearchCompanyDataTypeOfficeType, bool) {
	if o == nil || IsNil(o.OfficeType) {
		return nil, false
	}
	return o.OfficeType, true
}

// HasOfficeType returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeCompanyDetails) HasOfficeType() bool {
	if o != nil && !IsNil(o.OfficeType) {
		return true
	}

	return false
}

// SetOfficeType gets a reference to the given SearchCompanyDataTypeOfficeType and assigns it to the OfficeType field.
func (o *SearchCompanyDataTypeCompanyDetails) SetOfficeType(v SearchCompanyDataTypeOfficeType) {
	o.OfficeType = &v
}

func (o SearchCompanyDataTypeCompanyDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeCompanyDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.VatCode) {
		toSerialize["vatCode"] = o.VatCode
	}
	if !IsNil(o.TaxCode) {
		toSerialize["taxCode"] = o.TaxCode
	}
	if !IsNil(o.LastUpdateDate) {
		toSerialize["lastUpdateDate"] = o.LastUpdateDate
	}
	if !IsNil(o.Cciaa) {
		toSerialize["cciaa"] = o.Cciaa
	}
	if !IsNil(o.ReaCode) {
		toSerialize["reaCode"] = o.ReaCode
	}
	if !IsNil(o.CrifNumber) {
		toSerialize["crifNumber"] = o.CrifNumber
	}
	if !IsNil(o.OfficeType) {
		toSerialize["officeType"] = o.OfficeType
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeCompanyDetails struct {
	value *SearchCompanyDataTypeCompanyDetails
	isSet bool
}

func (v NullableSearchCompanyDataTypeCompanyDetails) Get() *SearchCompanyDataTypeCompanyDetails {
	return v.value
}

func (v *NullableSearchCompanyDataTypeCompanyDetails) Set(val *SearchCompanyDataTypeCompanyDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeCompanyDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeCompanyDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeCompanyDetails(val *SearchCompanyDataTypeCompanyDetails) *NullableSearchCompanyDataTypeCompanyDetails {
	return &NullableSearchCompanyDataTypeCompanyDetails{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeCompanyDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeCompanyDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


