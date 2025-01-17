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

// checks if the SearchCompanyDataTypeManagersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCompanyDataTypeManagersInner{}

// SearchCompanyDataTypeManagersInner struct for SearchCompanyDataTypeManagersInner
type SearchCompanyDataTypeManagersInner struct {
	Name *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	Roles []SearchCompanyDataTypeManagersInnerRolesInner `json:"roles,omitempty"`
	Gender *SearchCompanyDataTypeManagersInnerGender `json:"gender,omitempty"`
	TaxCode *string `json:"taxCode,omitempty"`
	BirthDate *time.Time `json:"birthDate,omitempty"`
	Age *float32 `json:"age,omitempty"`
	BirthTown *string `json:"birthTown,omitempty"`
	IsLegalRepresentative *bool `json:"isLegalRepresentative,omitempty"`
}

// NewSearchCompanyDataTypeManagersInner instantiates a new SearchCompanyDataTypeManagersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCompanyDataTypeManagersInner() *SearchCompanyDataTypeManagersInner {
	this := SearchCompanyDataTypeManagersInner{}
	return &this
}

// NewSearchCompanyDataTypeManagersInnerWithDefaults instantiates a new SearchCompanyDataTypeManagersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCompanyDataTypeManagersInnerWithDefaults() *SearchCompanyDataTypeManagersInner {
	this := SearchCompanyDataTypeManagersInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchCompanyDataTypeManagersInner) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *SearchCompanyDataTypeManagersInner) SetSurname(v string) {
	o.Surname = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *SearchCompanyDataTypeManagersInner) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetRoles() []SearchCompanyDataTypeManagersInnerRolesInner {
	if o == nil || IsNil(o.Roles) {
		var ret []SearchCompanyDataTypeManagersInnerRolesInner
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetRolesOk() ([]SearchCompanyDataTypeManagersInnerRolesInner, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []SearchCompanyDataTypeManagersInnerRolesInner and assigns it to the Roles field.
func (o *SearchCompanyDataTypeManagersInner) SetRoles(v []SearchCompanyDataTypeManagersInnerRolesInner) {
	o.Roles = v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetGender() SearchCompanyDataTypeManagersInnerGender {
	if o == nil || IsNil(o.Gender) {
		var ret SearchCompanyDataTypeManagersInnerGender
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetGenderOk() (*SearchCompanyDataTypeManagersInnerGender, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given SearchCompanyDataTypeManagersInnerGender and assigns it to the Gender field.
func (o *SearchCompanyDataTypeManagersInner) SetGender(v SearchCompanyDataTypeManagersInnerGender) {
	o.Gender = &v
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetTaxCode() string {
	if o == nil || IsNil(o.TaxCode) {
		var ret string
		return ret
	}
	return *o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetTaxCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxCode) {
		return nil, false
	}
	return o.TaxCode, true
}

// HasTaxCode returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasTaxCode() bool {
	if o != nil && !IsNil(o.TaxCode) {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given string and assigns it to the TaxCode field.
func (o *SearchCompanyDataTypeManagersInner) SetTaxCode(v string) {
	o.TaxCode = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetBirthDate() time.Time {
	if o == nil || IsNil(o.BirthDate) {
		var ret time.Time
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetBirthDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given time.Time and assigns it to the BirthDate field.
func (o *SearchCompanyDataTypeManagersInner) SetBirthDate(v time.Time) {
	o.BirthDate = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetAge() float32 {
	if o == nil || IsNil(o.Age) {
		var ret float32
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetAgeOk() (*float32, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given float32 and assigns it to the Age field.
func (o *SearchCompanyDataTypeManagersInner) SetAge(v float32) {
	o.Age = &v
}

// GetBirthTown returns the BirthTown field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetBirthTown() string {
	if o == nil || IsNil(o.BirthTown) {
		var ret string
		return ret
	}
	return *o.BirthTown
}

// GetBirthTownOk returns a tuple with the BirthTown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetBirthTownOk() (*string, bool) {
	if o == nil || IsNil(o.BirthTown) {
		return nil, false
	}
	return o.BirthTown, true
}

// HasBirthTown returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasBirthTown() bool {
	if o != nil && !IsNil(o.BirthTown) {
		return true
	}

	return false
}

// SetBirthTown gets a reference to the given string and assigns it to the BirthTown field.
func (o *SearchCompanyDataTypeManagersInner) SetBirthTown(v string) {
	o.BirthTown = &v
}

// GetIsLegalRepresentative returns the IsLegalRepresentative field value if set, zero value otherwise.
func (o *SearchCompanyDataTypeManagersInner) GetIsLegalRepresentative() bool {
	if o == nil || IsNil(o.IsLegalRepresentative) {
		var ret bool
		return ret
	}
	return *o.IsLegalRepresentative
}

// GetIsLegalRepresentativeOk returns a tuple with the IsLegalRepresentative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCompanyDataTypeManagersInner) GetIsLegalRepresentativeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLegalRepresentative) {
		return nil, false
	}
	return o.IsLegalRepresentative, true
}

// HasIsLegalRepresentative returns a boolean if a field has been set.
func (o *SearchCompanyDataTypeManagersInner) HasIsLegalRepresentative() bool {
	if o != nil && !IsNil(o.IsLegalRepresentative) {
		return true
	}

	return false
}

// SetIsLegalRepresentative gets a reference to the given bool and assigns it to the IsLegalRepresentative field.
func (o *SearchCompanyDataTypeManagersInner) SetIsLegalRepresentative(v bool) {
	o.IsLegalRepresentative = &v
}

func (o SearchCompanyDataTypeManagersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCompanyDataTypeManagersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.TaxCode) {
		toSerialize["taxCode"] = o.TaxCode
	}
	if !IsNil(o.BirthDate) {
		toSerialize["birthDate"] = o.BirthDate
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if !IsNil(o.BirthTown) {
		toSerialize["birthTown"] = o.BirthTown
	}
	if !IsNil(o.IsLegalRepresentative) {
		toSerialize["isLegalRepresentative"] = o.IsLegalRepresentative
	}
	return toSerialize, nil
}

type NullableSearchCompanyDataTypeManagersInner struct {
	value *SearchCompanyDataTypeManagersInner
	isSet bool
}

func (v NullableSearchCompanyDataTypeManagersInner) Get() *SearchCompanyDataTypeManagersInner {
	return v.value
}

func (v *NullableSearchCompanyDataTypeManagersInner) Set(val *SearchCompanyDataTypeManagersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCompanyDataTypeManagersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCompanyDataTypeManagersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCompanyDataTypeManagersInner(val *SearchCompanyDataTypeManagersInner) *NullableSearchCompanyDataTypeManagersInner {
	return &NullableSearchCompanyDataTypeManagersInner{value: val, isSet: true}
}

func (v NullableSearchCompanyDataTypeManagersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCompanyDataTypeManagersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


