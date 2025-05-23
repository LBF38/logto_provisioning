/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ReplaceApplicationRolesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceApplicationRolesRequest{}

// ReplaceApplicationRolesRequest struct for ReplaceApplicationRolesRequest
type ReplaceApplicationRolesRequest struct {
	// An array of API resource role IDs to update for the application.
	RoleIds []string `json:"roleIds"`
}

type _ReplaceApplicationRolesRequest ReplaceApplicationRolesRequest

// NewReplaceApplicationRolesRequest instantiates a new ReplaceApplicationRolesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceApplicationRolesRequest(roleIds []string) *ReplaceApplicationRolesRequest {
	this := ReplaceApplicationRolesRequest{}
	this.RoleIds = roleIds
	return &this
}

// NewReplaceApplicationRolesRequestWithDefaults instantiates a new ReplaceApplicationRolesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceApplicationRolesRequestWithDefaults() *ReplaceApplicationRolesRequest {
	this := ReplaceApplicationRolesRequest{}
	return &this
}

// GetRoleIds returns the RoleIds field value
func (o *ReplaceApplicationRolesRequest) GetRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value
// and a boolean to check if the value has been set.
func (o *ReplaceApplicationRolesRequest) GetRoleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// SetRoleIds sets field value
func (o *ReplaceApplicationRolesRequest) SetRoleIds(v []string) {
	o.RoleIds = v
}

func (o ReplaceApplicationRolesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceApplicationRolesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleIds"] = o.RoleIds
	return toSerialize, nil
}

func (o *ReplaceApplicationRolesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roleIds",
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

	varReplaceApplicationRolesRequest := _ReplaceApplicationRolesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplaceApplicationRolesRequest)

	if err != nil {
		return err
	}

	*o = ReplaceApplicationRolesRequest(varReplaceApplicationRolesRequest)

	return err
}

type NullableReplaceApplicationRolesRequest struct {
	value *ReplaceApplicationRolesRequest
	isSet bool
}

func (v NullableReplaceApplicationRolesRequest) Get() *ReplaceApplicationRolesRequest {
	return v.value
}

func (v *NullableReplaceApplicationRolesRequest) Set(val *ReplaceApplicationRolesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceApplicationRolesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceApplicationRolesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceApplicationRolesRequest(val *ReplaceApplicationRolesRequest) *NullableReplaceApplicationRolesRequest {
	return &NullableReplaceApplicationRolesRequest{value: val, isSet: true}
}

func (v NullableReplaceApplicationRolesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceApplicationRolesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


