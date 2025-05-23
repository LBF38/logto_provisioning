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

// checks if the CreateHookRequestConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHookRequestConfig{}

// CreateHookRequestConfig struct for CreateHookRequestConfig
type CreateHookRequestConfig struct {
	Url string `json:"url"`
	Headers *map[string]string `json:"headers,omitempty"`
	// Now the retry times is fixed to 3. Keep for backward compatibility.
	// Deprecated
	Retries *float32 `json:"retries,omitempty"`
}

type _CreateHookRequestConfig CreateHookRequestConfig

// NewCreateHookRequestConfig instantiates a new CreateHookRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHookRequestConfig(url string) *CreateHookRequestConfig {
	this := CreateHookRequestConfig{}
	this.Url = url
	return &this
}

// NewCreateHookRequestConfigWithDefaults instantiates a new CreateHookRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHookRequestConfigWithDefaults() *CreateHookRequestConfig {
	this := CreateHookRequestConfig{}
	return &this
}

// GetUrl returns the Url field value
func (o *CreateHookRequestConfig) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateHookRequestConfig) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateHookRequestConfig) SetUrl(v string) {
	o.Url = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *CreateHookRequestConfig) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHookRequestConfig) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *CreateHookRequestConfig) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *CreateHookRequestConfig) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
// Deprecated
func (o *CreateHookRequestConfig) GetRetries() float32 {
	if o == nil || IsNil(o.Retries) {
		var ret float32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateHookRequestConfig) GetRetriesOk() (*float32, bool) {
	if o == nil || IsNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *CreateHookRequestConfig) HasRetries() bool {
	if o != nil && !IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given float32 and assigns it to the Retries field.
// Deprecated
func (o *CreateHookRequestConfig) SetRetries(v float32) {
	o.Retries = &v
}

func (o CreateHookRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHookRequestConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	return toSerialize, nil
}

func (o *CreateHookRequestConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varCreateHookRequestConfig := _CreateHookRequestConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateHookRequestConfig)

	if err != nil {
		return err
	}

	*o = CreateHookRequestConfig(varCreateHookRequestConfig)

	return err
}

type NullableCreateHookRequestConfig struct {
	value *CreateHookRequestConfig
	isSet bool
}

func (v NullableCreateHookRequestConfig) Get() *CreateHookRequestConfig {
	return v.value
}

func (v *NullableCreateHookRequestConfig) Set(val *CreateHookRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHookRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHookRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHookRequestConfig(val *CreateHookRequestConfig) *NullableCreateHookRequestConfig {
	return &NullableCreateHookRequestConfig{value: val, isSet: true}
}

func (v NullableCreateHookRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHookRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


