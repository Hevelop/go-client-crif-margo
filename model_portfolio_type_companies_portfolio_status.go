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

// checks if the PortfolioTypeCompaniesPortfolioStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioTypeCompaniesPortfolioStatus{}

// PortfolioTypeCompaniesPortfolioStatus struct for PortfolioTypeCompaniesPortfolioStatus
type PortfolioTypeCompaniesPortfolioStatus struct {
	Prospect *int32 `json:"prospect,omitempty"`
	Client *int32 `json:"client,omitempty"`
	Supplier *int32 `json:"supplier,omitempty"`
	Lead *int32 `json:"lead,omitempty"`
	Competitor *int32 `json:"competitor,omitempty"`
}

// NewPortfolioTypeCompaniesPortfolioStatus instantiates a new PortfolioTypeCompaniesPortfolioStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioTypeCompaniesPortfolioStatus() *PortfolioTypeCompaniesPortfolioStatus {
	this := PortfolioTypeCompaniesPortfolioStatus{}
	return &this
}

// NewPortfolioTypeCompaniesPortfolioStatusWithDefaults instantiates a new PortfolioTypeCompaniesPortfolioStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioTypeCompaniesPortfolioStatusWithDefaults() *PortfolioTypeCompaniesPortfolioStatus {
	this := PortfolioTypeCompaniesPortfolioStatus{}
	return &this
}

// GetProspect returns the Prospect field value if set, zero value otherwise.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetProspect() int32 {
	if o == nil || IsNil(o.Prospect) {
		var ret int32
		return ret
	}
	return *o.Prospect
}

// GetProspectOk returns a tuple with the Prospect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetProspectOk() (*int32, bool) {
	if o == nil || IsNil(o.Prospect) {
		return nil, false
	}
	return o.Prospect, true
}

// HasProspect returns a boolean if a field has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) HasProspect() bool {
	if o != nil && !IsNil(o.Prospect) {
		return true
	}

	return false
}

// SetProspect gets a reference to the given int32 and assigns it to the Prospect field.
func (o *PortfolioTypeCompaniesPortfolioStatus) SetProspect(v int32) {
	o.Prospect = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetClient() int32 {
	if o == nil || IsNil(o.Client) {
		var ret int32
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetClientOk() (*int32, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given int32 and assigns it to the Client field.
func (o *PortfolioTypeCompaniesPortfolioStatus) SetClient(v int32) {
	o.Client = &v
}

// GetSupplier returns the Supplier field value if set, zero value otherwise.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetSupplier() int32 {
	if o == nil || IsNil(o.Supplier) {
		var ret int32
		return ret
	}
	return *o.Supplier
}

// GetSupplierOk returns a tuple with the Supplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetSupplierOk() (*int32, bool) {
	if o == nil || IsNil(o.Supplier) {
		return nil, false
	}
	return o.Supplier, true
}

// HasSupplier returns a boolean if a field has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) HasSupplier() bool {
	if o != nil && !IsNil(o.Supplier) {
		return true
	}

	return false
}

// SetSupplier gets a reference to the given int32 and assigns it to the Supplier field.
func (o *PortfolioTypeCompaniesPortfolioStatus) SetSupplier(v int32) {
	o.Supplier = &v
}

// GetLead returns the Lead field value if set, zero value otherwise.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetLead() int32 {
	if o == nil || IsNil(o.Lead) {
		var ret int32
		return ret
	}
	return *o.Lead
}

// GetLeadOk returns a tuple with the Lead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetLeadOk() (*int32, bool) {
	if o == nil || IsNil(o.Lead) {
		return nil, false
	}
	return o.Lead, true
}

// HasLead returns a boolean if a field has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) HasLead() bool {
	if o != nil && !IsNil(o.Lead) {
		return true
	}

	return false
}

// SetLead gets a reference to the given int32 and assigns it to the Lead field.
func (o *PortfolioTypeCompaniesPortfolioStatus) SetLead(v int32) {
	o.Lead = &v
}

// GetCompetitor returns the Competitor field value if set, zero value otherwise.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetCompetitor() int32 {
	if o == nil || IsNil(o.Competitor) {
		var ret int32
		return ret
	}
	return *o.Competitor
}

// GetCompetitorOk returns a tuple with the Competitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) GetCompetitorOk() (*int32, bool) {
	if o == nil || IsNil(o.Competitor) {
		return nil, false
	}
	return o.Competitor, true
}

// HasCompetitor returns a boolean if a field has been set.
func (o *PortfolioTypeCompaniesPortfolioStatus) HasCompetitor() bool {
	if o != nil && !IsNil(o.Competitor) {
		return true
	}

	return false
}

// SetCompetitor gets a reference to the given int32 and assigns it to the Competitor field.
func (o *PortfolioTypeCompaniesPortfolioStatus) SetCompetitor(v int32) {
	o.Competitor = &v
}

func (o PortfolioTypeCompaniesPortfolioStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioTypeCompaniesPortfolioStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prospect) {
		toSerialize["prospect"] = o.Prospect
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Supplier) {
		toSerialize["supplier"] = o.Supplier
	}
	if !IsNil(o.Lead) {
		toSerialize["lead"] = o.Lead
	}
	if !IsNil(o.Competitor) {
		toSerialize["competitor"] = o.Competitor
	}
	return toSerialize, nil
}

type NullablePortfolioTypeCompaniesPortfolioStatus struct {
	value *PortfolioTypeCompaniesPortfolioStatus
	isSet bool
}

func (v NullablePortfolioTypeCompaniesPortfolioStatus) Get() *PortfolioTypeCompaniesPortfolioStatus {
	return v.value
}

func (v *NullablePortfolioTypeCompaniesPortfolioStatus) Set(val *PortfolioTypeCompaniesPortfolioStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioTypeCompaniesPortfolioStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioTypeCompaniesPortfolioStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioTypeCompaniesPortfolioStatus(val *PortfolioTypeCompaniesPortfolioStatus) *NullablePortfolioTypeCompaniesPortfolioStatus {
	return &NullablePortfolioTypeCompaniesPortfolioStatus{value: val, isSet: true}
}

func (v NullablePortfolioTypeCompaniesPortfolioStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioTypeCompaniesPortfolioStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


