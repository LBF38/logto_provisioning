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

// ListJwtCustomizers200ResponseInner - struct for ListJwtCustomizers200ResponseInner
type ListJwtCustomizers200ResponseInner struct {
	ListJwtCustomizers200ResponseInnerOneOf *ListJwtCustomizers200ResponseInnerOneOf
	ListJwtCustomizers200ResponseInnerOneOf1 *ListJwtCustomizers200ResponseInnerOneOf1
}

// ListJwtCustomizers200ResponseInnerOneOfAsListJwtCustomizers200ResponseInner is a convenience function that returns ListJwtCustomizers200ResponseInnerOneOf wrapped in ListJwtCustomizers200ResponseInner
func ListJwtCustomizers200ResponseInnerOneOfAsListJwtCustomizers200ResponseInner(v *ListJwtCustomizers200ResponseInnerOneOf) ListJwtCustomizers200ResponseInner {
	return ListJwtCustomizers200ResponseInner{
		ListJwtCustomizers200ResponseInnerOneOf: v,
	}
}

// ListJwtCustomizers200ResponseInnerOneOf1AsListJwtCustomizers200ResponseInner is a convenience function that returns ListJwtCustomizers200ResponseInnerOneOf1 wrapped in ListJwtCustomizers200ResponseInner
func ListJwtCustomizers200ResponseInnerOneOf1AsListJwtCustomizers200ResponseInner(v *ListJwtCustomizers200ResponseInnerOneOf1) ListJwtCustomizers200ResponseInner {
	return ListJwtCustomizers200ResponseInner{
		ListJwtCustomizers200ResponseInnerOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListJwtCustomizers200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListJwtCustomizers200ResponseInnerOneOf
	err = newStrictDecoder(data).Decode(&dst.ListJwtCustomizers200ResponseInnerOneOf)
	if err == nil {
		jsonListJwtCustomizers200ResponseInnerOneOf, _ := json.Marshal(dst.ListJwtCustomizers200ResponseInnerOneOf)
		if string(jsonListJwtCustomizers200ResponseInnerOneOf) == "{}" { // empty struct
			dst.ListJwtCustomizers200ResponseInnerOneOf = nil
		} else {
			if err = validator.Validate(dst.ListJwtCustomizers200ResponseInnerOneOf); err != nil {
				dst.ListJwtCustomizers200ResponseInnerOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.ListJwtCustomizers200ResponseInnerOneOf = nil
	}

	// try to unmarshal data into ListJwtCustomizers200ResponseInnerOneOf1
	err = newStrictDecoder(data).Decode(&dst.ListJwtCustomizers200ResponseInnerOneOf1)
	if err == nil {
		jsonListJwtCustomizers200ResponseInnerOneOf1, _ := json.Marshal(dst.ListJwtCustomizers200ResponseInnerOneOf1)
		if string(jsonListJwtCustomizers200ResponseInnerOneOf1) == "{}" { // empty struct
			dst.ListJwtCustomizers200ResponseInnerOneOf1 = nil
		} else {
			if err = validator.Validate(dst.ListJwtCustomizers200ResponseInnerOneOf1); err != nil {
				dst.ListJwtCustomizers200ResponseInnerOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.ListJwtCustomizers200ResponseInnerOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListJwtCustomizers200ResponseInnerOneOf = nil
		dst.ListJwtCustomizers200ResponseInnerOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListJwtCustomizers200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListJwtCustomizers200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListJwtCustomizers200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.ListJwtCustomizers200ResponseInnerOneOf != nil {
		return json.Marshal(&src.ListJwtCustomizers200ResponseInnerOneOf)
	}

	if src.ListJwtCustomizers200ResponseInnerOneOf1 != nil {
		return json.Marshal(&src.ListJwtCustomizers200ResponseInnerOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListJwtCustomizers200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ListJwtCustomizers200ResponseInnerOneOf != nil {
		return obj.ListJwtCustomizers200ResponseInnerOneOf
	}

	if obj.ListJwtCustomizers200ResponseInnerOneOf1 != nil {
		return obj.ListJwtCustomizers200ResponseInnerOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableListJwtCustomizers200ResponseInner struct {
	value *ListJwtCustomizers200ResponseInner
	isSet bool
}

func (v NullableListJwtCustomizers200ResponseInner) Get() *ListJwtCustomizers200ResponseInner {
	return v.value
}

func (v *NullableListJwtCustomizers200ResponseInner) Set(val *ListJwtCustomizers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListJwtCustomizers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListJwtCustomizers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListJwtCustomizers200ResponseInner(val *ListJwtCustomizers200ResponseInner) *NullableListJwtCustomizers200ResponseInner {
	return &NullableListJwtCustomizers200ResponseInner{value: val, isSet: true}
}

func (v NullableListJwtCustomizers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListJwtCustomizers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


