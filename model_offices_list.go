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

// checks if the OfficesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfficesList{}

// OfficesList struct for OfficesList
type OfficesList struct {
	OfficeIds []string `json:"officeIds"`
}

type _OfficesList OfficesList

// NewOfficesList instantiates a new OfficesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfficesList(officeIds []string) *OfficesList {
	this := OfficesList{}
	this.OfficeIds = officeIds
	return &this
}

// NewOfficesListWithDefaults instantiates a new OfficesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfficesListWithDefaults() *OfficesList {
	this := OfficesList{}
	return &this
}

// GetOfficeIds returns the OfficeIds field value
func (o *OfficesList) GetOfficeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OfficeIds
}

// GetOfficeIdsOk returns a tuple with the OfficeIds field value
// and a boolean to check if the value has been set.
func (o *OfficesList) GetOfficeIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfficeIds, true
}

// SetOfficeIds sets field value
func (o *OfficesList) SetOfficeIds(v []string) {
	o.OfficeIds = v
}

func (o OfficesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfficesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["officeIds"] = o.OfficeIds
	return toSerialize, nil
}

func (o *OfficesList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varOfficesList := _OfficesList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOfficesList)

	if err != nil {
		return err
	}

	*o = OfficesList(varOfficesList)

	return err
}

type NullableOfficesList struct {
	value *OfficesList
	isSet bool
}

func (v NullableOfficesList) Get() *OfficesList {
	return v.value
}

func (v *NullableOfficesList) Set(val *OfficesList) {
	v.value = val
	v.isSet = true
}

func (v NullableOfficesList) IsSet() bool {
	return v.isSet
}

func (v *NullableOfficesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfficesList(val *OfficesList) *NullableOfficesList {
	return &NullableOfficesList{value: val, isSet: true}
}

func (v NullableOfficesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfficesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


