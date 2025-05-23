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

// checks if the AssignApplicationRolesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignApplicationRolesRequest{}

// AssignApplicationRolesRequest struct for AssignApplicationRolesRequest
type AssignApplicationRolesRequest struct {
	// An array of API resource role IDs to assign.
	RoleIds []string `json:"roleIds"`
}

type _AssignApplicationRolesRequest AssignApplicationRolesRequest

// NewAssignApplicationRolesRequest instantiates a new AssignApplicationRolesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignApplicationRolesRequest(roleIds []string) *AssignApplicationRolesRequest {
	this := AssignApplicationRolesRequest{}
	this.RoleIds = roleIds
	return &this
}

// NewAssignApplicationRolesRequestWithDefaults instantiates a new AssignApplicationRolesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignApplicationRolesRequestWithDefaults() *AssignApplicationRolesRequest {
	this := AssignApplicationRolesRequest{}
	return &this
}

// GetRoleIds returns the RoleIds field value
func (o *AssignApplicationRolesRequest) GetRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value
// and a boolean to check if the value has been set.
func (o *AssignApplicationRolesRequest) GetRoleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// SetRoleIds sets field value
func (o *AssignApplicationRolesRequest) SetRoleIds(v []string) {
	o.RoleIds = v
}

func (o AssignApplicationRolesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignApplicationRolesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleIds"] = o.RoleIds
	return toSerialize, nil
}

func (o *AssignApplicationRolesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAssignApplicationRolesRequest := _AssignApplicationRolesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssignApplicationRolesRequest)

	if err != nil {
		return err
	}

	*o = AssignApplicationRolesRequest(varAssignApplicationRolesRequest)

	return err
}

type NullableAssignApplicationRolesRequest struct {
	value *AssignApplicationRolesRequest
	isSet bool
}

func (v NullableAssignApplicationRolesRequest) Get() *AssignApplicationRolesRequest {
	return v.value
}

func (v *NullableAssignApplicationRolesRequest) Set(val *AssignApplicationRolesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignApplicationRolesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignApplicationRolesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignApplicationRolesRequest(val *AssignApplicationRolesRequest) *NullableAssignApplicationRolesRequest {
	return &NullableAssignApplicationRolesRequest{value: val, isSet: true}
}

func (v NullableAssignApplicationRolesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignApplicationRolesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


