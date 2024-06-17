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

// checks if the UserType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserType{}

// UserType struct for UserType
type UserType struct {
	Id *string `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	// This item identifies the user's role in the hierarchy.
	Type *string `json:"type,omitempty"`
	Language *string `json:"language,omitempty"`
	// It contains the position of the user inside the subscription hierarchy
	DistinguishName *string `json:"distinguishName,omitempty"`
}

// NewUserType instantiates a new UserType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserType() *UserType {
	this := UserType{}
	return &this
}

// NewUserTypeWithDefaults instantiates a new UserType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypeWithDefaults() *UserType {
	this := UserType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserType) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UserType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UserType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UserType) SetCode(v string) {
	o.Code = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserType) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserType) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserType) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserType) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserType) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserType) SetLastName(v string) {
	o.LastName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserType) SetType(v string) {
	o.Type = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *UserType) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *UserType) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *UserType) SetLanguage(v string) {
	o.Language = &v
}

// GetDistinguishName returns the DistinguishName field value if set, zero value otherwise.
func (o *UserType) GetDistinguishName() string {
	if o == nil || IsNil(o.DistinguishName) {
		var ret string
		return ret
	}
	return *o.DistinguishName
}

// GetDistinguishNameOk returns a tuple with the DistinguishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetDistinguishNameOk() (*string, bool) {
	if o == nil || IsNil(o.DistinguishName) {
		return nil, false
	}
	return o.DistinguishName, true
}

// HasDistinguishName returns a boolean if a field has been set.
func (o *UserType) HasDistinguishName() bool {
	if o != nil && !IsNil(o.DistinguishName) {
		return true
	}

	return false
}

// SetDistinguishName gets a reference to the given string and assigns it to the DistinguishName field.
func (o *UserType) SetDistinguishName(v string) {
	o.DistinguishName = &v
}

func (o UserType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.DistinguishName) {
		toSerialize["distinguishName"] = o.DistinguishName
	}
	return toSerialize, nil
}

type NullableUserType struct {
	value *UserType
	isSet bool
}

func (v NullableUserType) Get() *UserType {
	return v.value
}

func (v *NullableUserType) Set(val *UserType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserType(val *UserType) *NullableUserType {
	return &NullableUserType{value: val, isSet: true}
}

func (v NullableUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


