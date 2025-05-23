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

// checks if the AddUserProfileRequestOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUserProfileRequestOneOf3{}

// AddUserProfileRequestOneOf3 struct for AddUserProfileRequestOneOf3
type AddUserProfileRequestOneOf3 struct {
	Type string `json:"type"`
	VerificationId string `json:"verificationId"`
}

type _AddUserProfileRequestOneOf3 AddUserProfileRequestOneOf3

// NewAddUserProfileRequestOneOf3 instantiates a new AddUserProfileRequestOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserProfileRequestOneOf3(type_ string, verificationId string) *AddUserProfileRequestOneOf3 {
	this := AddUserProfileRequestOneOf3{}
	this.Type = type_
	this.VerificationId = verificationId
	return &this
}

// NewAddUserProfileRequestOneOf3WithDefaults instantiates a new AddUserProfileRequestOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserProfileRequestOneOf3WithDefaults() *AddUserProfileRequestOneOf3 {
	this := AddUserProfileRequestOneOf3{}
	return &this
}

// GetType returns the Type field value
func (o *AddUserProfileRequestOneOf3) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddUserProfileRequestOneOf3) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddUserProfileRequestOneOf3) SetType(v string) {
	o.Type = v
}

// GetVerificationId returns the VerificationId field value
func (o *AddUserProfileRequestOneOf3) GetVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationId
}

// GetVerificationIdOk returns a tuple with the VerificationId field value
// and a boolean to check if the value has been set.
func (o *AddUserProfileRequestOneOf3) GetVerificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationId, true
}

// SetVerificationId sets field value
func (o *AddUserProfileRequestOneOf3) SetVerificationId(v string) {
	o.VerificationId = v
}

func (o AddUserProfileRequestOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUserProfileRequestOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["verificationId"] = o.VerificationId
	return toSerialize, nil
}

func (o *AddUserProfileRequestOneOf3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"verificationId",
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

	varAddUserProfileRequestOneOf3 := _AddUserProfileRequestOneOf3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddUserProfileRequestOneOf3)

	if err != nil {
		return err
	}

	*o = AddUserProfileRequestOneOf3(varAddUserProfileRequestOneOf3)

	return err
}

type NullableAddUserProfileRequestOneOf3 struct {
	value *AddUserProfileRequestOneOf3
	isSet bool
}

func (v NullableAddUserProfileRequestOneOf3) Get() *AddUserProfileRequestOneOf3 {
	return v.value
}

func (v *NullableAddUserProfileRequestOneOf3) Set(val *AddUserProfileRequestOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUserProfileRequestOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUserProfileRequestOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUserProfileRequestOneOf3(val *AddUserProfileRequestOneOf3) *NullableAddUserProfileRequestOneOf3 {
	return &NullableAddUserProfileRequestOneOf3{value: val, isSet: true}
}

func (v NullableAddUserProfileRequestOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUserProfileRequestOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


