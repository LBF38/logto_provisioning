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

// checks if the GenerateBackupCodes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateBackupCodes200Response{}

// GenerateBackupCodes200Response struct for GenerateBackupCodes200Response
type GenerateBackupCodes200Response struct {
	// The unique verification ID of the newly created BackupCode verification record. This ID is required when adding the backup codes to the user profile via the Profile API.
	VerificationId string `json:"verificationId"`
	// The generated backup codes.
	Codes []string `json:"codes"`
}

type _GenerateBackupCodes200Response GenerateBackupCodes200Response

// NewGenerateBackupCodes200Response instantiates a new GenerateBackupCodes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateBackupCodes200Response(verificationId string, codes []string) *GenerateBackupCodes200Response {
	this := GenerateBackupCodes200Response{}
	this.VerificationId = verificationId
	this.Codes = codes
	return &this
}

// NewGenerateBackupCodes200ResponseWithDefaults instantiates a new GenerateBackupCodes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateBackupCodes200ResponseWithDefaults() *GenerateBackupCodes200Response {
	this := GenerateBackupCodes200Response{}
	return &this
}

// GetVerificationId returns the VerificationId field value
func (o *GenerateBackupCodes200Response) GetVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationId
}

// GetVerificationIdOk returns a tuple with the VerificationId field value
// and a boolean to check if the value has been set.
func (o *GenerateBackupCodes200Response) GetVerificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationId, true
}

// SetVerificationId sets field value
func (o *GenerateBackupCodes200Response) SetVerificationId(v string) {
	o.VerificationId = v
}

// GetCodes returns the Codes field value
func (o *GenerateBackupCodes200Response) GetCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value
// and a boolean to check if the value has been set.
func (o *GenerateBackupCodes200Response) GetCodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Codes, true
}

// SetCodes sets field value
func (o *GenerateBackupCodes200Response) SetCodes(v []string) {
	o.Codes = v
}

func (o GenerateBackupCodes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateBackupCodes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verificationId"] = o.VerificationId
	toSerialize["codes"] = o.Codes
	return toSerialize, nil
}

func (o *GenerateBackupCodes200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"verificationId",
		"codes",
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

	varGenerateBackupCodes200Response := _GenerateBackupCodes200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateBackupCodes200Response)

	if err != nil {
		return err
	}

	*o = GenerateBackupCodes200Response(varGenerateBackupCodes200Response)

	return err
}

type NullableGenerateBackupCodes200Response struct {
	value *GenerateBackupCodes200Response
	isSet bool
}

func (v NullableGenerateBackupCodes200Response) Get() *GenerateBackupCodes200Response {
	return v.value
}

func (v *NullableGenerateBackupCodes200Response) Set(val *GenerateBackupCodes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateBackupCodes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateBackupCodes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateBackupCodes200Response(val *GenerateBackupCodes200Response) *NullableGenerateBackupCodes200Response {
	return &NullableGenerateBackupCodes200Response{value: val, isSet: true}
}

func (v NullableGenerateBackupCodes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateBackupCodes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


