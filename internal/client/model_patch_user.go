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

// PatchUser struct for PatchUser
type PatchUser struct {
	RoleID     *string         `json:"roleID,omitempty"`
	Databases  *PatchDatabases `json:"databases,omitempty"`
	DbUsername *string         `json:"dbUsername,omitempty"`
}

// NewPatchUser instantiates a new PatchUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchUser() *PatchUser {
	this := PatchUser{}
	return &this
}

// NewPatchUserWithDefaults instantiates a new PatchUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchUserWithDefaults() *PatchUser {
	this := PatchUser{}
	return &this
}

// GetRoleID returns the RoleID field value if set, zero value otherwise.
func (o *PatchUser) GetRoleID() string {
	if o == nil || o.RoleID == nil {
		var ret string
		return ret
	}
	return *o.RoleID
}

// GetRoleIDOk returns a tuple with the RoleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUser) GetRoleIDOk() (*string, bool) {
	if o == nil || o.RoleID == nil {
		return nil, false
	}
	return o.RoleID, true
}

// HasRoleID returns a boolean if a field has been set.
func (o *PatchUser) HasRoleID() bool {
	if o != nil && o.RoleID != nil {
		return true
	}

	return false
}

// SetRoleID gets a reference to the given string and assigns it to the RoleID field.
func (o *PatchUser) SetRoleID(v string) {
	o.RoleID = &v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *PatchUser) GetDatabases() PatchDatabases {
	if o == nil || o.Databases == nil {
		var ret PatchDatabases
		return ret
	}
	return *o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUser) GetDatabasesOk() (*PatchDatabases, bool) {
	if o == nil || o.Databases == nil {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *PatchUser) HasDatabases() bool {
	if o != nil && o.Databases != nil {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given PatchDatabases and assigns it to the Databases field.
func (o *PatchUser) SetDatabases(v PatchDatabases) {
	o.Databases = &v
}

// GetDbUsername returns the DbUsername field value if set, zero value otherwise.
func (o *PatchUser) GetDbUsername() string {
	if o == nil || o.DbUsername == nil {
		var ret string
		return ret
	}
	return *o.DbUsername
}

// GetDbUsernameOk returns a tuple with the DbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUser) GetDbUsernameOk() (*string, bool) {
	if o == nil || o.DbUsername == nil {
		return nil, false
	}
	return o.DbUsername, true
}

// HasDbUsername returns a boolean if a field has been set.
func (o *PatchUser) HasDbUsername() bool {
	if o != nil && o.DbUsername != nil {
		return true
	}

	return false
}

// SetDbUsername gets a reference to the given string and assigns it to the DbUsername field.
func (o *PatchUser) SetDbUsername(v string) {
	o.DbUsername = &v
}

func (o PatchUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleID != nil {
		toSerialize["roleID"] = o.RoleID
	}
	if o.Databases != nil {
		toSerialize["databases"] = o.Databases
	}
	if o.DbUsername != nil {
		toSerialize["dbUsername"] = o.DbUsername
	}
	return json.Marshal(toSerialize)
}

type NullablePatchUser struct {
	value *PatchUser
	isSet bool
}

func (v NullablePatchUser) Get() *PatchUser {
	return v.value
}

func (v *NullablePatchUser) Set(val *PatchUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchUser(val *PatchUser) *NullablePatchUser {
	return &NullablePatchUser{value: val, isSet: true}
}

func (v NullablePatchUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
