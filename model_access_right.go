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
	"bytes"
	"fmt"
)

// checks if the AccessRight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRight{}

// AccessRight struct for AccessRight
type AccessRight struct {
	AccessRight string `json:"accessRight"`
	OfficeIds []string `json:"officeIds"`
}

type _AccessRight AccessRight

// NewAccessRight instantiates a new AccessRight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRight(accessRight string, officeIds []string) *AccessRight {
	this := AccessRight{}
	this.AccessRight = accessRight
	this.OfficeIds = officeIds
	return &this
}

// NewAccessRightWithDefaults instantiates a new AccessRight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRightWithDefaults() *AccessRight {
	this := AccessRight{}
	return &this
}

// GetAccessRight returns the AccessRight field value
func (o *AccessRight) GetAccessRight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessRight
}

// GetAccessRightOk returns a tuple with the AccessRight field value
// and a boolean to check if the value has been set.
func (o *AccessRight) GetAccessRightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessRight, true
}

// SetAccessRight sets field value
func (o *AccessRight) SetAccessRight(v string) {
	o.AccessRight = v
}

// GetOfficeIds returns the OfficeIds field value
func (o *AccessRight) GetOfficeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OfficeIds
}

// GetOfficeIdsOk returns a tuple with the OfficeIds field value
// and a boolean to check if the value has been set.
func (o *AccessRight) GetOfficeIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfficeIds, true
}

// SetOfficeIds sets field value
func (o *AccessRight) SetOfficeIds(v []string) {
	o.OfficeIds = v
}

func (o AccessRight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessRight"] = o.AccessRight
	toSerialize["officeIds"] = o.OfficeIds
	return toSerialize, nil
}

func (o *AccessRight) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessRight",
		"officeIds",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccessRight := _AccessRight{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessRight)

	if err != nil {
		return err
	}

	*o = AccessRight(varAccessRight)

	return err
}

type NullableAccessRight struct {
	value *AccessRight
	isSet bool
}

func (v NullableAccessRight) Get() *AccessRight {
	return v.value
}

func (v *NullableAccessRight) Set(val *AccessRight) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRight) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRight(val *AccessRight) *NullableAccessRight {
	return &NullableAccessRight{value: val, isSet: true}
}

func (v NullableAccessRight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


