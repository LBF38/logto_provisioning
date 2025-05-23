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

// checks if the VerifyVerificationCodeVerificationRequestIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyVerificationCodeVerificationRequestIdentifier{}

// VerifyVerificationCodeVerificationRequestIdentifier The identifier (email address or phone number) to verify the code against. Must match the identifier used to send the verification code.
type VerifyVerificationCodeVerificationRequestIdentifier struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

type _VerifyVerificationCodeVerificationRequestIdentifier VerifyVerificationCodeVerificationRequestIdentifier

// NewVerifyVerificationCodeVerificationRequestIdentifier instantiates a new VerifyVerificationCodeVerificationRequestIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyVerificationCodeVerificationRequestIdentifier(type_ string, value string) *VerifyVerificationCodeVerificationRequestIdentifier {
	this := VerifyVerificationCodeVerificationRequestIdentifier{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewVerifyVerificationCodeVerificationRequestIdentifierWithDefaults instantiates a new VerifyVerificationCodeVerificationRequestIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyVerificationCodeVerificationRequestIdentifierWithDefaults() *VerifyVerificationCodeVerificationRequestIdentifier {
	this := VerifyVerificationCodeVerificationRequestIdentifier{}
	return &this
}

// GetType returns the Type field value
func (o *VerifyVerificationCodeVerificationRequestIdentifier) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VerifyVerificationCodeVerificationRequestIdentifier) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VerifyVerificationCodeVerificationRequestIdentifier) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *VerifyVerificationCodeVerificationRequestIdentifier) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VerifyVerificationCodeVerificationRequestIdentifier) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VerifyVerificationCodeVerificationRequestIdentifier) SetValue(v string) {
	o.Value = v
}

func (o VerifyVerificationCodeVerificationRequestIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyVerificationCodeVerificationRequestIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *VerifyVerificationCodeVerificationRequestIdentifier) UnmarshalJSON(data []byte) (err error) {
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

	varVerifyVerificationCodeVerificationRequestIdentifier := _VerifyVerificationCodeVerificationRequestIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifyVerificationCodeVerificationRequestIdentifier)

	if err != nil {
		return err
	}

	*o = VerifyVerificationCodeVerificationRequestIdentifier(varVerifyVerificationCodeVerificationRequestIdentifier)

	return err
}

type NullableVerifyVerificationCodeVerificationRequestIdentifier struct {
	value *VerifyVerificationCodeVerificationRequestIdentifier
	isSet bool
}

func (v NullableVerifyVerificationCodeVerificationRequestIdentifier) Get() *VerifyVerificationCodeVerificationRequestIdentifier {
	return v.value
}

func (v *NullableVerifyVerificationCodeVerificationRequestIdentifier) Set(val *VerifyVerificationCodeVerificationRequestIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyVerificationCodeVerificationRequestIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyVerificationCodeVerificationRequestIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyVerificationCodeVerificationRequestIdentifier(val *VerifyVerificationCodeVerificationRequestIdentifier) *NullableVerifyVerificationCodeVerificationRequestIdentifier {
	return &NullableVerifyVerificationCodeVerificationRequestIdentifier{value: val, isSet: true}
}

func (v NullableVerifyVerificationCodeVerificationRequestIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyVerificationCodeVerificationRequestIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


