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

// checks if the ApiInteractionMfaPutRequestOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionMfaPutRequestOneOf{}

// ApiInteractionMfaPutRequestOneOf struct for ApiInteractionMfaPutRequestOneOf
type ApiInteractionMfaPutRequestOneOf struct {
	Type string `json:"type"`
	Id string `json:"id"`
	RawId string `json:"rawId"`
	AuthenticatorAttachment *string `json:"authenticatorAttachment,omitempty"`
	ClientExtensionResults VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults `json:"clientExtensionResults"`
	Response VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse `json:"response"`
}

type _ApiInteractionMfaPutRequestOneOf ApiInteractionMfaPutRequestOneOf

// NewApiInteractionMfaPutRequestOneOf instantiates a new ApiInteractionMfaPutRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionMfaPutRequestOneOf(type_ string, id string, rawId string, clientExtensionResults VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, response VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) *ApiInteractionMfaPutRequestOneOf {
	this := ApiInteractionMfaPutRequestOneOf{}
	this.Type = type_
	this.Id = id
	this.RawId = rawId
	this.ClientExtensionResults = clientExtensionResults
	this.Response = response
	return &this
}

// NewApiInteractionMfaPutRequestOneOfWithDefaults instantiates a new ApiInteractionMfaPutRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionMfaPutRequestOneOfWithDefaults() *ApiInteractionMfaPutRequestOneOf {
	this := ApiInteractionMfaPutRequestOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *ApiInteractionMfaPutRequestOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionMfaPutRequestOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiInteractionMfaPutRequestOneOf) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ApiInteractionMfaPutRequestOneOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionMfaPutRequestOneOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiInteractionMfaPutRequestOneOf) SetId(v string) {
	o.Id = v
}

// GetRawId returns the RawId field value
func (o *ApiInteractionMfaPutRequestOneOf) GetRawId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawId
}

// GetRawIdOk returns a tuple with the RawId field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionMfaPutRequestOneOf) GetRawIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawId, true
}

// SetRawId sets field value
func (o *ApiInteractionMfaPutRequestOneOf) SetRawId(v string) {
	o.RawId = v
}

// GetAuthenticatorAttachment returns the AuthenticatorAttachment field value if set, zero value otherwise.
func (o *ApiInteractionMfaPutRequestOneOf) GetAuthenticatorAttachment() string {
	if o == nil || IsNil(o.AuthenticatorAttachment) {
		var ret string
		return ret
	}
	return *o.AuthenticatorAttachment
}

// GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionMfaPutRequestOneOf) GetAuthenticatorAttachmentOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticatorAttachment) {
		return nil, false
	}
	return o.AuthenticatorAttachment, true
}

// HasAuthenticatorAttachment returns a boolean if a field has been set.
func (o *ApiInteractionMfaPutRequestOneOf) HasAuthenticatorAttachment() bool {
	if o != nil && !IsNil(o.AuthenticatorAttachment) {
		return true
	}

	return false
}

// SetAuthenticatorAttachment gets a reference to the given string and assigns it to the AuthenticatorAttachment field.
func (o *ApiInteractionMfaPutRequestOneOf) SetAuthenticatorAttachment(v string) {
	o.AuthenticatorAttachment = &v
}

// GetClientExtensionResults returns the ClientExtensionResults field value
func (o *ApiInteractionMfaPutRequestOneOf) GetClientExtensionResults() VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults {
	if o == nil {
		var ret VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults
		return ret
	}

	return o.ClientExtensionResults
}

// GetClientExtensionResultsOk returns a tuple with the ClientExtensionResults field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionMfaPutRequestOneOf) GetClientExtensionResultsOk() (*VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientExtensionResults, true
}

// SetClientExtensionResults sets field value
func (o *ApiInteractionMfaPutRequestOneOf) SetClientExtensionResults(v VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults) {
	o.ClientExtensionResults = v
}

// GetResponse returns the Response field value
func (o *ApiInteractionMfaPutRequestOneOf) GetResponse() VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse {
	if o == nil {
		var ret VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionMfaPutRequestOneOf) GetResponseOk() (*VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *ApiInteractionMfaPutRequestOneOf) SetResponse(v VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) {
	o.Response = v
}

func (o ApiInteractionMfaPutRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionMfaPutRequestOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["rawId"] = o.RawId
	if !IsNil(o.AuthenticatorAttachment) {
		toSerialize["authenticatorAttachment"] = o.AuthenticatorAttachment
	}
	toSerialize["clientExtensionResults"] = o.ClientExtensionResults
	toSerialize["response"] = o.Response
	return toSerialize, nil
}

func (o *ApiInteractionMfaPutRequestOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"rawId",
		"clientExtensionResults",
		"response",
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

	varApiInteractionMfaPutRequestOneOf := _ApiInteractionMfaPutRequestOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionMfaPutRequestOneOf)

	if err != nil {
		return err
	}

	*o = ApiInteractionMfaPutRequestOneOf(varApiInteractionMfaPutRequestOneOf)

	return err
}

type NullableApiInteractionMfaPutRequestOneOf struct {
	value *ApiInteractionMfaPutRequestOneOf
	isSet bool
}

func (v NullableApiInteractionMfaPutRequestOneOf) Get() *ApiInteractionMfaPutRequestOneOf {
	return v.value
}

func (v *NullableApiInteractionMfaPutRequestOneOf) Set(val *ApiInteractionMfaPutRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionMfaPutRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionMfaPutRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionMfaPutRequestOneOf(val *ApiInteractionMfaPutRequestOneOf) *NullableApiInteractionMfaPutRequestOneOf {
	return &NullableApiInteractionMfaPutRequestOneOf{value: val, isSet: true}
}

func (v NullableApiInteractionMfaPutRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionMfaPutRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


