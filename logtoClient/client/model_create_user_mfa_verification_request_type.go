/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// CreateUserMfaVerificationRequestType - The type of MFA verification to create.
type CreateUserMfaVerificationRequestType struct {
	String *string
}

// stringAsCreateUserMfaVerificationRequestType is a convenience function that returns string wrapped in CreateUserMfaVerificationRequestType
func StringAsCreateUserMfaVerificationRequestType(v *string) CreateUserMfaVerificationRequestType {
	return CreateUserMfaVerificationRequestType{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateUserMfaVerificationRequestType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateUserMfaVerificationRequestType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateUserMfaVerificationRequestType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateUserMfaVerificationRequestType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateUserMfaVerificationRequestType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateUserMfaVerificationRequestType struct {
	value *CreateUserMfaVerificationRequestType
	isSet bool
}

func (v NullableCreateUserMfaVerificationRequestType) Get() *CreateUserMfaVerificationRequestType {
	return v.value
}

func (v *NullableCreateUserMfaVerificationRequestType) Set(val *CreateUserMfaVerificationRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserMfaVerificationRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserMfaVerificationRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserMfaVerificationRequestType(val *CreateUserMfaVerificationRequestType) *NullableCreateUserMfaVerificationRequestType {
	return &NullableCreateUserMfaVerificationRequestType{value: val, isSet: true}
}

func (v NullableCreateUserMfaVerificationRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserMfaVerificationRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


