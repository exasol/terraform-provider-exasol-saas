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

// UpdateUser struct for UpdateUser
type UpdateUser struct {
	RoleID     *string  `json:"roleID,omitempty"`
	Databases  []string `json:"databases,omitempty"`
	DbUsername string   `json:"dbUsername"`
}

// NewUpdateUser instantiates a new UpdateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUser(dbUsername string) *UpdateUser {
	this := UpdateUser{}
	this.DbUsername = dbUsername
	return &this
}

// NewUpdateUserWithDefaults instantiates a new UpdateUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserWithDefaults() *UpdateUser {
	this := UpdateUser{}
	return &this
}

// GetRoleID returns the RoleID field value if set, zero value otherwise.
func (o *UpdateUser) GetRoleID() string {
	if o == nil || o.RoleID == nil {
		var ret string
		return ret
	}
	return *o.RoleID
}

// GetRoleIDOk returns a tuple with the RoleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetRoleIDOk() (*string, bool) {
	if o == nil || o.RoleID == nil {
		return nil, false
	}
	return o.RoleID, true
}

// HasRoleID returns a boolean if a field has been set.
func (o *UpdateUser) HasRoleID() bool {
	if o != nil && o.RoleID != nil {
		return true
	}

	return false
}

// SetRoleID gets a reference to the given string and assigns it to the RoleID field.
func (o *UpdateUser) SetRoleID(v string) {
	o.RoleID = &v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *UpdateUser) GetDatabases() []string {
	if o == nil || o.Databases == nil {
		var ret []string
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetDatabasesOk() ([]string, bool) {
	if o == nil || o.Databases == nil {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *UpdateUser) HasDatabases() bool {
	if o != nil && o.Databases != nil {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []string and assigns it to the Databases field.
func (o *UpdateUser) SetDatabases(v []string) {
	o.Databases = v
}

// GetDbUsername returns the DbUsername field value
func (o *UpdateUser) GetDbUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbUsername
}

// GetDbUsernameOk returns a tuple with the DbUsername field value
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetDbUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbUsername, true
}

// SetDbUsername sets field value
func (o *UpdateUser) SetDbUsername(v string) {
	o.DbUsername = v
}

func (o UpdateUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleID != nil {
		toSerialize["roleID"] = o.RoleID
	}
	if o.Databases != nil {
		toSerialize["databases"] = o.Databases
	}
	if true {
		toSerialize["dbUsername"] = o.DbUsername
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateUser struct {
	value *UpdateUser
	isSet bool
}

func (v NullableUpdateUser) Get() *UpdateUser {
	return v.value
}

func (v *NullableUpdateUser) Set(val *UpdateUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUser(val *UpdateUser) *NullableUpdateUser {
	return &NullableUpdateUser{value: val, isSet: true}
}

func (v NullableUpdateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}