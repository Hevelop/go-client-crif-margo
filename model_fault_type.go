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

// checks if the FaultType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FaultType{}

// FaultType struct for FaultType
type FaultType struct {
	Code *int32 `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Description *string `json:"description,omitempty"`
	Details []ErrorType `json:"details,omitempty"`
}

// NewFaultType instantiates a new FaultType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaultType() *FaultType {
	this := FaultType{}
	return &this
}

// NewFaultTypeWithDefaults instantiates a new FaultType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaultTypeWithDefaults() *FaultType {
	this := FaultType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FaultType) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultType) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FaultType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *FaultType) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FaultType) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultType) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FaultType) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FaultType) SetMessage(v string) {
	o.Message = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FaultType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FaultType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FaultType) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *FaultType) GetDetails() []ErrorType {
	if o == nil || IsNil(o.Details) {
		var ret []ErrorType
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultType) GetDetailsOk() ([]ErrorType, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *FaultType) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ErrorType and assigns it to the Details field.
func (o *FaultType) SetDetails(v []ErrorType) {
	o.Details = v
}

func (o FaultType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FaultType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableFaultType struct {
	value *FaultType
	isSet bool
}

func (v NullableFaultType) Get() *FaultType {
	return v.value
}

func (v *NullableFaultType) Set(val *FaultType) {
	v.value = val
	v.isSet = true
}

func (v NullableFaultType) IsSet() bool {
	return v.isSet
}

func (v *NullableFaultType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaultType(val *FaultType) *NullableFaultType {
	return &NullableFaultType{value: val, isSet: true}
}

func (v NullableFaultType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaultType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

