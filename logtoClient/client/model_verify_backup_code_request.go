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

// checks if the VerifyBackupCodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyBackupCodeRequest{}

// VerifyBackupCodeRequest struct for VerifyBackupCodeRequest
type VerifyBackupCodeRequest struct {
	// The backup code to verify.
	Code string `json:"code"`
}

type _VerifyBackupCodeRequest VerifyBackupCodeRequest

// NewVerifyBackupCodeRequest instantiates a new VerifyBackupCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyBackupCodeRequest(code string) *VerifyBackupCodeRequest {
	this := VerifyBackupCodeRequest{}
	this.Code = code
	return &this
}

// NewVerifyBackupCodeRequestWithDefaults instantiates a new VerifyBackupCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyBackupCodeRequestWithDefaults() *VerifyBackupCodeRequest {
	this := VerifyBackupCodeRequest{}
	return &this
}

// GetCode returns the Code field value
func (o *VerifyBackupCodeRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *VerifyBackupCodeRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *VerifyBackupCodeRequest) SetCode(v string) {
	o.Code = v
}

func (o VerifyBackupCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyBackupCodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *VerifyBackupCodeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varVerifyBackupCodeRequest := _VerifyBackupCodeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifyBackupCodeRequest)

	if err != nil {
		return err
	}

	*o = VerifyBackupCodeRequest(varVerifyBackupCodeRequest)

	return err
}

type NullableVerifyBackupCodeRequest struct {
	value *VerifyBackupCodeRequest
	isSet bool
}

func (v NullableVerifyBackupCodeRequest) Get() *VerifyBackupCodeRequest {
	return v.value
}

func (v *NullableVerifyBackupCodeRequest) Set(val *VerifyBackupCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyBackupCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyBackupCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyBackupCodeRequest(val *VerifyBackupCodeRequest) *NullableVerifyBackupCodeRequest {
	return &NullableVerifyBackupCodeRequest{value: val, isSet: true}
}

func (v NullableVerifyBackupCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyBackupCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


