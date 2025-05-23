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

// checks if the AddUserProfileRequestOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUserProfileRequestOneOf{}

// AddUserProfileRequestOneOf struct for AddUserProfileRequestOneOf
type AddUserProfileRequestOneOf struct {
	Type string `json:"type"`
	Value string `json:"value" validate:"regexp=^[A-Z_a-z]\\\\w*$"`
}

type _AddUserProfileRequestOneOf AddUserProfileRequestOneOf

// NewAddUserProfileRequestOneOf instantiates a new AddUserProfileRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserProfileRequestOneOf(type_ string, value string) *AddUserProfileRequestOneOf {
	this := AddUserProfileRequestOneOf{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewAddUserProfileRequestOneOfWithDefaults instantiates a new AddUserProfileRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserProfileRequestOneOfWithDefaults() *AddUserProfileRequestOneOf {
	this := AddUserProfileRequestOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *AddUserProfileRequestOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddUserProfileRequestOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddUserProfileRequestOneOf) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AddUserProfileRequestOneOf) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AddUserProfileRequestOneOf) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AddUserProfileRequestOneOf) SetValue(v string) {
	o.Value = v
}

func (o AddUserProfileRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUserProfileRequestOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *AddUserProfileRequestOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
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

	varAddUserProfileRequestOneOf := _AddUserProfileRequestOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddUserProfileRequestOneOf)

	if err != nil {
		return err
	}

	*o = AddUserProfileRequestOneOf(varAddUserProfileRequestOneOf)

	return err
}

type NullableAddUserProfileRequestOneOf struct {
	value *AddUserProfileRequestOneOf
	isSet bool
}

func (v NullableAddUserProfileRequestOneOf) Get() *AddUserProfileRequestOneOf {
	return v.value
}

func (v *NullableAddUserProfileRequestOneOf) Set(val *AddUserProfileRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUserProfileRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUserProfileRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUserProfileRequestOneOf(val *AddUserProfileRequestOneOf) *NullableAddUserProfileRequestOneOf {
	return &NullableAddUserProfileRequestOneOf{value: val, isSet: true}
}

func (v NullableAddUserProfileRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUserProfileRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


