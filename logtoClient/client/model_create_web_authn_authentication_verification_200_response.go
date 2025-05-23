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

// checks if the CreateWebAuthnAuthenticationVerification200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebAuthnAuthenticationVerification200Response{}

// CreateWebAuthnAuthenticationVerification200Response struct for CreateWebAuthnAuthenticationVerification200Response
type CreateWebAuthnAuthenticationVerification200Response struct {
	// The unique ID for the WebAuthn authentication record, required to verify the WebAuthn authentication challenge.
	VerificationId string `json:"verificationId"`
	AuthenticationOptions CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions `json:"authenticationOptions"`
}

type _CreateWebAuthnAuthenticationVerification200Response CreateWebAuthnAuthenticationVerification200Response

// NewCreateWebAuthnAuthenticationVerification200Response instantiates a new CreateWebAuthnAuthenticationVerification200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebAuthnAuthenticationVerification200Response(verificationId string, authenticationOptions CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) *CreateWebAuthnAuthenticationVerification200Response {
	this := CreateWebAuthnAuthenticationVerification200Response{}
	this.VerificationId = verificationId
	this.AuthenticationOptions = authenticationOptions
	return &this
}

// NewCreateWebAuthnAuthenticationVerification200ResponseWithDefaults instantiates a new CreateWebAuthnAuthenticationVerification200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebAuthnAuthenticationVerification200ResponseWithDefaults() *CreateWebAuthnAuthenticationVerification200Response {
	this := CreateWebAuthnAuthenticationVerification200Response{}
	return &this
}

// GetVerificationId returns the VerificationId field value
func (o *CreateWebAuthnAuthenticationVerification200Response) GetVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationId
}

// GetVerificationIdOk returns a tuple with the VerificationId field value
// and a boolean to check if the value has been set.
func (o *CreateWebAuthnAuthenticationVerification200Response) GetVerificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationId, true
}

// SetVerificationId sets field value
func (o *CreateWebAuthnAuthenticationVerification200Response) SetVerificationId(v string) {
	o.VerificationId = v
}

// GetAuthenticationOptions returns the AuthenticationOptions field value
func (o *CreateWebAuthnAuthenticationVerification200Response) GetAuthenticationOptions() CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions {
	if o == nil {
		var ret CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions
		return ret
	}

	return o.AuthenticationOptions
}

// GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field value
// and a boolean to check if the value has been set.
func (o *CreateWebAuthnAuthenticationVerification200Response) GetAuthenticationOptionsOk() (*CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationOptions, true
}

// SetAuthenticationOptions sets field value
func (o *CreateWebAuthnAuthenticationVerification200Response) SetAuthenticationOptions(v CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) {
	o.AuthenticationOptions = v
}

func (o CreateWebAuthnAuthenticationVerification200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebAuthnAuthenticationVerification200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verificationId"] = o.VerificationId
	toSerialize["authenticationOptions"] = o.AuthenticationOptions
	return toSerialize, nil
}

func (o *CreateWebAuthnAuthenticationVerification200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"verificationId",
		"authenticationOptions",
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

	varCreateWebAuthnAuthenticationVerification200Response := _CreateWebAuthnAuthenticationVerification200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWebAuthnAuthenticationVerification200Response)

	if err != nil {
		return err
	}

	*o = CreateWebAuthnAuthenticationVerification200Response(varCreateWebAuthnAuthenticationVerification200Response)

	return err
}

type NullableCreateWebAuthnAuthenticationVerification200Response struct {
	value *CreateWebAuthnAuthenticationVerification200Response
	isSet bool
}

func (v NullableCreateWebAuthnAuthenticationVerification200Response) Get() *CreateWebAuthnAuthenticationVerification200Response {
	return v.value
}

func (v *NullableCreateWebAuthnAuthenticationVerification200Response) Set(val *CreateWebAuthnAuthenticationVerification200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebAuthnAuthenticationVerification200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebAuthnAuthenticationVerification200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebAuthnAuthenticationVerification200Response(val *CreateWebAuthnAuthenticationVerification200Response) *NullableCreateWebAuthnAuthenticationVerification200Response {
	return &NullableCreateWebAuthnAuthenticationVerification200Response{value: val, isSet: true}
}

func (v NullableCreateWebAuthnAuthenticationVerification200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebAuthnAuthenticationVerification200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


