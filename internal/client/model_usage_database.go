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

// UsageDatabase struct for UsageDatabase
type UsageDatabase struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	UsedStorage *float32       `json:"usedStorage,omitempty"`
	Clusters    []UsageCluster `json:"clusters"`
}

// NewUsageDatabase instantiates a new UsageDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageDatabase(id string, name string, clusters []UsageCluster) *UsageDatabase {
	this := UsageDatabase{}
	this.Id = id
	this.Name = name
	this.Clusters = clusters
	return &this
}

// NewUsageDatabaseWithDefaults instantiates a new UsageDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageDatabaseWithDefaults() *UsageDatabase {
	this := UsageDatabase{}
	return &this
}

// GetId returns the Id field value
func (o *UsageDatabase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UsageDatabase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UsageDatabase) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *UsageDatabase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UsageDatabase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UsageDatabase) SetName(v string) {
	o.Name = v
}

// GetUsedStorage returns the UsedStorage field value if set, zero value otherwise.
func (o *UsageDatabase) GetUsedStorage() float32 {
	if o == nil || o.UsedStorage == nil {
		var ret float32
		return ret
	}
	return *o.UsedStorage
}

// GetUsedStorageOk returns a tuple with the UsedStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageDatabase) GetUsedStorageOk() (*float32, bool) {
	if o == nil || o.UsedStorage == nil {
		return nil, false
	}
	return o.UsedStorage, true
}

// HasUsedStorage returns a boolean if a field has been set.
func (o *UsageDatabase) HasUsedStorage() bool {
	if o != nil && o.UsedStorage != nil {
		return true
	}

	return false
}

// SetUsedStorage gets a reference to the given float32 and assigns it to the UsedStorage field.
func (o *UsageDatabase) SetUsedStorage(v float32) {
	o.UsedStorage = &v
}

// GetClusters returns the Clusters field value
func (o *UsageDatabase) GetClusters() []UsageCluster {
	if o == nil {
		var ret []UsageCluster
		return ret
	}

	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *UsageDatabase) GetClustersOk() ([]UsageCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clusters, true
}

// SetClusters sets field value
func (o *UsageDatabase) SetClusters(v []UsageCluster) {
	o.Clusters = v
}

func (o UsageDatabase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UsedStorage != nil {
		toSerialize["usedStorage"] = o.UsedStorage
	}
	if true {
		toSerialize["clusters"] = o.Clusters
	}
	return json.Marshal(toSerialize)
}

type NullableUsageDatabase struct {
	value *UsageDatabase
	isSet bool
}

func (v NullableUsageDatabase) Get() *UsageDatabase {
	return v.value
}

func (v *NullableUsageDatabase) Set(val *UsageDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageDatabase(val *UsageDatabase) *NullableUsageDatabase {
	return &NullableUsageDatabase{value: val, isSet: true}
}

func (v NullableUsageDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}