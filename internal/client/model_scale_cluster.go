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

// ScaleCluster struct for ScaleCluster
type ScaleCluster struct {
	Size string `json:"size"`
}

// NewScaleCluster instantiates a new ScaleCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScaleCluster(size string) *ScaleCluster {
	this := ScaleCluster{}
	this.Size = size
	return &this
}

// NewScaleClusterWithDefaults instantiates a new ScaleCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScaleClusterWithDefaults() *ScaleCluster {
	this := ScaleCluster{}
	return &this
}

// GetSize returns the Size field value
func (o *ScaleCluster) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ScaleCluster) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ScaleCluster) SetSize(v string) {
	o.Size = v
}

func (o ScaleCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableScaleCluster struct {
	value *ScaleCluster
	isSet bool
}

func (v NullableScaleCluster) Get() *ScaleCluster {
	return v.value
}

func (v *NullableScaleCluster) Set(val *ScaleCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableScaleCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableScaleCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScaleCluster(val *ScaleCluster) *NullableScaleCluster {
	return &NullableScaleCluster{value: val, isSet: true}
}

func (v NullableScaleCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScaleCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
