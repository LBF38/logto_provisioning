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

// checks if the CreateUserMfaVerificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserMfaVerificationRequest{}

// CreateUserMfaVerificationRequest struct for CreateUserMfaVerificationRequest
type CreateUserMfaVerificationRequest struct {
	Type CreateUserMfaVerificationRequestType `json:"type"`
}

type _CreateUserMfaVerificationRequest CreateUserMfaVerificationRequest

// NewCreateUserMfaVerificationRequest instantiates a new CreateUserMfaVerificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserMfaVerificationRequest(type_ CreateUserMfaVerificationRequestType) *CreateUserMfaVerificationRequest {
	this := CreateUserMfaVerificationRequest{}
	this.Type = type_
	return &this
}

// NewCreateUserMfaVerificationRequestWithDefaults instantiates a new CreateUserMfaVerificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserMfaVerificationRequestWithDefaults() *CreateUserMfaVerificationRequest {
	this := CreateUserMfaVerificationRequest{}
	return &this
}

// GetType returns the Type field value
func (o *CreateUserMfaVerificationRequest) GetType() CreateUserMfaVerificationRequestType {
	if o == nil {
		var ret CreateUserMfaVerificationRequestType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateUserMfaVerificationRequest) GetTypeOk() (*CreateUserMfaVerificationRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateUserMfaVerificationRequest) SetType(v CreateUserMfaVerificationRequestType) {
	o.Type = v
}

func (o CreateUserMfaVerificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserMfaVerificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CreateUserMfaVerificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCreateUserMfaVerificationRequest := _CreateUserMfaVerificationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUserMfaVerificationRequest)

	if err != nil {
		return err
	}

	*o = CreateUserMfaVerificationRequest(varCreateUserMfaVerificationRequest)

	return err
}

type NullableCreateUserMfaVerificationRequest struct {
	value *CreateUserMfaVerificationRequest
	isSet bool
}

func (v NullableCreateUserMfaVerificationRequest) Get() *CreateUserMfaVerificationRequest {
	return v.value
}

func (v *NullableCreateUserMfaVerificationRequest) Set(val *CreateUserMfaVerificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserMfaVerificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserMfaVerificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserMfaVerificationRequest(val *CreateUserMfaVerificationRequest) *NullableCreateUserMfaVerificationRequest {
	return &NullableCreateUserMfaVerificationRequest{value: val, isSet: true}
}

func (v NullableCreateUserMfaVerificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserMfaVerificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


