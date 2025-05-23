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

// CheckPasswordWithDefaultSignInExperience200Response - struct for CheckPasswordWithDefaultSignInExperience200Response
type CheckPasswordWithDefaultSignInExperience200Response struct {
	CheckPasswordWithDefaultSignInExperience200ResponseOneOf *CheckPasswordWithDefaultSignInExperience200ResponseOneOf
	CheckPasswordWithDefaultSignInExperience200ResponseOneOf1 *CheckPasswordWithDefaultSignInExperience200ResponseOneOf1
}

// CheckPasswordWithDefaultSignInExperience200ResponseOneOfAsCheckPasswordWithDefaultSignInExperience200Response is a convenience function that returns CheckPasswordWithDefaultSignInExperience200ResponseOneOf wrapped in CheckPasswordWithDefaultSignInExperience200Response
func CheckPasswordWithDefaultSignInExperience200ResponseOneOfAsCheckPasswordWithDefaultSignInExperience200Response(v *CheckPasswordWithDefaultSignInExperience200ResponseOneOf) CheckPasswordWithDefaultSignInExperience200Response {
	return CheckPasswordWithDefaultSignInExperience200Response{
		CheckPasswordWithDefaultSignInExperience200ResponseOneOf: v,
	}
}

// CheckPasswordWithDefaultSignInExperience200ResponseOneOf1AsCheckPasswordWithDefaultSignInExperience200Response is a convenience function that returns CheckPasswordWithDefaultSignInExperience200ResponseOneOf1 wrapped in CheckPasswordWithDefaultSignInExperience200Response
func CheckPasswordWithDefaultSignInExperience200ResponseOneOf1AsCheckPasswordWithDefaultSignInExperience200Response(v *CheckPasswordWithDefaultSignInExperience200ResponseOneOf1) CheckPasswordWithDefaultSignInExperience200Response {
	return CheckPasswordWithDefaultSignInExperience200Response{
		CheckPasswordWithDefaultSignInExperience200ResponseOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CheckPasswordWithDefaultSignInExperience200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CheckPasswordWithDefaultSignInExperience200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf)
	if err == nil {
		jsonCheckPasswordWithDefaultSignInExperience200ResponseOneOf, _ := json.Marshal(dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf)
		if string(jsonCheckPasswordWithDefaultSignInExperience200ResponseOneOf) == "{}" { // empty struct
			dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf = nil
		} else {
			if err = validator.Validate(dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf); err != nil {
				dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf = nil
	}

	// try to unmarshal data into CheckPasswordWithDefaultSignInExperience200ResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1)
	if err == nil {
		jsonCheckPasswordWithDefaultSignInExperience200ResponseOneOf1, _ := json.Marshal(dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1)
		if string(jsonCheckPasswordWithDefaultSignInExperience200ResponseOneOf1) == "{}" { // empty struct
			dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1 = nil
		} else {
			if err = validator.Validate(dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1); err != nil {
				dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf = nil
		dst.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CheckPasswordWithDefaultSignInExperience200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CheckPasswordWithDefaultSignInExperience200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CheckPasswordWithDefaultSignInExperience200Response) MarshalJSON() ([]byte, error) {
	if src.CheckPasswordWithDefaultSignInExperience200ResponseOneOf != nil {
		return json.Marshal(&src.CheckPasswordWithDefaultSignInExperience200ResponseOneOf)
	}

	if src.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1 != nil {
		return json.Marshal(&src.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CheckPasswordWithDefaultSignInExperience200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CheckPasswordWithDefaultSignInExperience200ResponseOneOf != nil {
		return obj.CheckPasswordWithDefaultSignInExperience200ResponseOneOf
	}

	if obj.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1 != nil {
		return obj.CheckPasswordWithDefaultSignInExperience200ResponseOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableCheckPasswordWithDefaultSignInExperience200Response struct {
	value *CheckPasswordWithDefaultSignInExperience200Response
	isSet bool
}

func (v NullableCheckPasswordWithDefaultSignInExperience200Response) Get() *CheckPasswordWithDefaultSignInExperience200Response {
	return v.value
}

func (v *NullableCheckPasswordWithDefaultSignInExperience200Response) Set(val *CheckPasswordWithDefaultSignInExperience200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckPasswordWithDefaultSignInExperience200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckPasswordWithDefaultSignInExperience200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckPasswordWithDefaultSignInExperience200Response(val *CheckPasswordWithDefaultSignInExperience200Response) *NullableCheckPasswordWithDefaultSignInExperience200Response {
	return &NullableCheckPasswordWithDefaultSignInExperience200Response{value: val, isSet: true}
}

func (v NullableCheckPasswordWithDefaultSignInExperience200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckPasswordWithDefaultSignInExperience200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


