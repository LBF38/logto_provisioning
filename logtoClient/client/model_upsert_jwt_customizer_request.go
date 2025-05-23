/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpsertJwtCustomizerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertJwtCustomizerRequest{}

// UpsertJwtCustomizerRequest struct for UpsertJwtCustomizerRequest
type UpsertJwtCustomizerRequest struct {
	// The script of the JWT customizer.
	Script interface{} `json:"script,omitempty"`
	// The environment variables for the JWT customizer.
	EnvironmentVariables interface{} `json:"environmentVariables,omitempty"`
	// The sample context for the JWT customizer script testing purpose.
	ContextSample interface{} `json:"contextSample,omitempty"`
	// The sample raw token payload for the JWT customizer script testing purpose.
	TokenSample interface{} `json:"tokenSample,omitempty"`
}

// NewUpsertJwtCustomizerRequest instantiates a new UpsertJwtCustomizerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertJwtCustomizerRequest() *UpsertJwtCustomizerRequest {
	this := UpsertJwtCustomizerRequest{}
	return &this
}

// NewUpsertJwtCustomizerRequestWithDefaults instantiates a new UpsertJwtCustomizerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertJwtCustomizerRequestWithDefaults() *UpsertJwtCustomizerRequest {
	this := UpsertJwtCustomizerRequest{}
	return &this
}

// GetScript returns the Script field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertJwtCustomizerRequest) GetScript() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertJwtCustomizerRequest) GetScriptOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return &o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *UpsertJwtCustomizerRequest) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given interface{} and assigns it to the Script field.
func (o *UpsertJwtCustomizerRequest) SetScript(v interface{}) {
	o.Script = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertJwtCustomizerRequest) GetEnvironmentVariables() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertJwtCustomizerRequest) GetEnvironmentVariablesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return &o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *UpsertJwtCustomizerRequest) HasEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given interface{} and assigns it to the EnvironmentVariables field.
func (o *UpsertJwtCustomizerRequest) SetEnvironmentVariables(v interface{}) {
	o.EnvironmentVariables = v
}

// GetContextSample returns the ContextSample field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertJwtCustomizerRequest) GetContextSample() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ContextSample
}

// GetContextSampleOk returns a tuple with the ContextSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertJwtCustomizerRequest) GetContextSampleOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ContextSample) {
		return nil, false
	}
	return &o.ContextSample, true
}

// HasContextSample returns a boolean if a field has been set.
func (o *UpsertJwtCustomizerRequest) HasContextSample() bool {
	if o != nil && !IsNil(o.ContextSample) {
		return true
	}

	return false
}

// SetContextSample gets a reference to the given interface{} and assigns it to the ContextSample field.
func (o *UpsertJwtCustomizerRequest) SetContextSample(v interface{}) {
	o.ContextSample = v
}

// GetTokenSample returns the TokenSample field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertJwtCustomizerRequest) GetTokenSample() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TokenSample
}

// GetTokenSampleOk returns a tuple with the TokenSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertJwtCustomizerRequest) GetTokenSampleOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TokenSample) {
		return nil, false
	}
	return &o.TokenSample, true
}

// HasTokenSample returns a boolean if a field has been set.
func (o *UpsertJwtCustomizerRequest) HasTokenSample() bool {
	if o != nil && !IsNil(o.TokenSample) {
		return true
	}

	return false
}

// SetTokenSample gets a reference to the given interface{} and assigns it to the TokenSample field.
func (o *UpsertJwtCustomizerRequest) SetTokenSample(v interface{}) {
	o.TokenSample = v
}

func (o UpsertJwtCustomizerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertJwtCustomizerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	if o.EnvironmentVariables != nil {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if o.ContextSample != nil {
		toSerialize["contextSample"] = o.ContextSample
	}
	if o.TokenSample != nil {
		toSerialize["tokenSample"] = o.TokenSample
	}
	return toSerialize, nil
}

type NullableUpsertJwtCustomizerRequest struct {
	value *UpsertJwtCustomizerRequest
	isSet bool
}

func (v NullableUpsertJwtCustomizerRequest) Get() *UpsertJwtCustomizerRequest {
	return v.value
}

func (v *NullableUpsertJwtCustomizerRequest) Set(val *UpsertJwtCustomizerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertJwtCustomizerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertJwtCustomizerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertJwtCustomizerRequest(val *UpsertJwtCustomizerRequest) *NullableUpsertJwtCustomizerRequest {
	return &NullableUpsertJwtCustomizerRequest{value: val, isSet: true}
}

func (v NullableUpsertJwtCustomizerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertJwtCustomizerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


