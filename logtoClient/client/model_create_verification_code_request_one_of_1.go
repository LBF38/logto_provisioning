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

// checks if the CreateVerificationCodeRequestOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVerificationCodeRequestOneOf1{}

// CreateVerificationCodeRequestOneOf1 struct for CreateVerificationCodeRequestOneOf1
type CreateVerificationCodeRequestOneOf1 struct {
	Phone string `json:"phone" validate:"regexp=^\\\\d+$"`
}

type _CreateVerificationCodeRequestOneOf1 CreateVerificationCodeRequestOneOf1

// NewCreateVerificationCodeRequestOneOf1 instantiates a new CreateVerificationCodeRequestOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVerificationCodeRequestOneOf1(phone string) *CreateVerificationCodeRequestOneOf1 {
	this := CreateVerificationCodeRequestOneOf1{}
	this.Phone = phone
	return &this
}

// NewCreateVerificationCodeRequestOneOf1WithDefaults instantiates a new CreateVerificationCodeRequestOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVerificationCodeRequestOneOf1WithDefaults() *CreateVerificationCodeRequestOneOf1 {
	this := CreateVerificationCodeRequestOneOf1{}
	return &this
}

// GetPhone returns the Phone field value
func (o *CreateVerificationCodeRequestOneOf1) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *CreateVerificationCodeRequestOneOf1) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *CreateVerificationCodeRequestOneOf1) SetPhone(v string) {
	o.Phone = v
}

func (o CreateVerificationCodeRequestOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVerificationCodeRequestOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone"] = o.Phone
	return toSerialize, nil
}

func (o *CreateVerificationCodeRequestOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone",
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

	varCreateVerificationCodeRequestOneOf1 := _CreateVerificationCodeRequestOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateVerificationCodeRequestOneOf1)

	if err != nil {
		return err
	}

	*o = CreateVerificationCodeRequestOneOf1(varCreateVerificationCodeRequestOneOf1)

	return err
}

type NullableCreateVerificationCodeRequestOneOf1 struct {
	value *CreateVerificationCodeRequestOneOf1
	isSet bool
}

func (v NullableCreateVerificationCodeRequestOneOf1) Get() *CreateVerificationCodeRequestOneOf1 {
	return v.value
}

func (v *NullableCreateVerificationCodeRequestOneOf1) Set(val *CreateVerificationCodeRequestOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVerificationCodeRequestOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVerificationCodeRequestOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVerificationCodeRequestOneOf1(val *CreateVerificationCodeRequestOneOf1) *NullableCreateVerificationCodeRequestOneOf1 {
	return &NullableCreateVerificationCodeRequestOneOf1{value: val, isSet: true}
}

func (v NullableCreateVerificationCodeRequestOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVerificationCodeRequestOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


