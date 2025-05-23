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

// checks if the GetUserHasPassword200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserHasPassword200Response{}

// GetUserHasPassword200Response struct for GetUserHasPassword200Response
type GetUserHasPassword200Response struct {
	HasPassword bool `json:"hasPassword"`
}

type _GetUserHasPassword200Response GetUserHasPassword200Response

// NewGetUserHasPassword200Response instantiates a new GetUserHasPassword200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserHasPassword200Response(hasPassword bool) *GetUserHasPassword200Response {
	this := GetUserHasPassword200Response{}
	this.HasPassword = hasPassword
	return &this
}

// NewGetUserHasPassword200ResponseWithDefaults instantiates a new GetUserHasPassword200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserHasPassword200ResponseWithDefaults() *GetUserHasPassword200Response {
	this := GetUserHasPassword200Response{}
	return &this
}

// GetHasPassword returns the HasPassword field value
func (o *GetUserHasPassword200Response) GetHasPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value
// and a boolean to check if the value has been set.
func (o *GetUserHasPassword200Response) GetHasPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPassword, true
}

// SetHasPassword sets field value
func (o *GetUserHasPassword200Response) SetHasPassword(v bool) {
	o.HasPassword = v
}

func (o GetUserHasPassword200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserHasPassword200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasPassword"] = o.HasPassword
	return toSerialize, nil
}

func (o *GetUserHasPassword200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hasPassword",
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

	varGetUserHasPassword200Response := _GetUserHasPassword200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetUserHasPassword200Response)

	if err != nil {
		return err
	}

	*o = GetUserHasPassword200Response(varGetUserHasPassword200Response)

	return err
}

type NullableGetUserHasPassword200Response struct {
	value *GetUserHasPassword200Response
	isSet bool
}

func (v NullableGetUserHasPassword200Response) Get() *GetUserHasPassword200Response {
	return v.value
}

func (v *NullableGetUserHasPassword200Response) Set(val *GetUserHasPassword200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserHasPassword200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserHasPassword200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserHasPassword200Response(val *GetUserHasPassword200Response) *NullableGetUserHasPassword200Response {
	return &NullableGetUserHasPassword200Response{value: val, isSet: true}
}

func (v NullableGetUserHasPassword200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserHasPassword200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


