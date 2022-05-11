/*
Exasol SaaS REST-API

## Authentication  The REST API can be used with your Personal Access Token (PAT). You don't know what a PAT is, check our documentation  [here](https://docs.exasol.com/saas/administration/access_mngt/access_token.htm).  After you created a PAT click on Authorize and add your PAT under BearerAuth.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AutoStop struct for AutoStop
type AutoStop struct {
	Enabled  bool  `json:"enabled"`
	IdleTime int32 `json:"idleTime"`
}

// NewAutoStop instantiates a new AutoStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoStop(enabled bool, idleTime int32) *AutoStop {
	this := AutoStop{}
	this.Enabled = enabled
	this.IdleTime = idleTime
	return &this
}

// NewAutoStopWithDefaults instantiates a new AutoStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoStopWithDefaults() *AutoStop {
	this := AutoStop{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *AutoStop) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AutoStop) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AutoStop) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIdleTime returns the IdleTime field value
func (o *AutoStop) GetIdleTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IdleTime
}

// GetIdleTimeOk returns a tuple with the IdleTime field value
// and a boolean to check if the value has been set.
func (o *AutoStop) GetIdleTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdleTime, true
}

// SetIdleTime sets field value
func (o *AutoStop) SetIdleTime(v int32) {
	o.IdleTime = v
}

func (o AutoStop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["idleTime"] = o.IdleTime
	}
	return json.Marshal(toSerialize)
}

type NullableAutoStop struct {
	value *AutoStop
	isSet bool
}

func (v NullableAutoStop) Get() *AutoStop {
	return v.value
}

func (v *NullableAutoStop) Set(val *AutoStop) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoStop) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoStop(val *AutoStop) *NullableAutoStop {
	return &NullableAutoStop{value: val, isSet: true}
}

func (v NullableAutoStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}