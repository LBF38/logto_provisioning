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

// checks if the UpdateSignInExpRequestLanguageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSignInExpRequestLanguageInfo{}

// UpdateSignInExpRequestLanguageInfo Control the language detection policy for the sign-in page.
type UpdateSignInExpRequestLanguageInfo struct {
	AutoDetect bool `json:"autoDetect"`
	FallbackLanguage string `json:"fallbackLanguage"`
}

type _UpdateSignInExpRequestLanguageInfo UpdateSignInExpRequestLanguageInfo

// NewUpdateSignInExpRequestLanguageInfo instantiates a new UpdateSignInExpRequestLanguageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSignInExpRequestLanguageInfo(autoDetect bool, fallbackLanguage string) *UpdateSignInExpRequestLanguageInfo {
	this := UpdateSignInExpRequestLanguageInfo{}
	this.AutoDetect = autoDetect
	this.FallbackLanguage = fallbackLanguage
	return &this
}

// NewUpdateSignInExpRequestLanguageInfoWithDefaults instantiates a new UpdateSignInExpRequestLanguageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSignInExpRequestLanguageInfoWithDefaults() *UpdateSignInExpRequestLanguageInfo {
	this := UpdateSignInExpRequestLanguageInfo{}
	return &this
}

// GetAutoDetect returns the AutoDetect field value
func (o *UpdateSignInExpRequestLanguageInfo) GetAutoDetect() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoDetect
}

// GetAutoDetectOk returns a tuple with the AutoDetect field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequestLanguageInfo) GetAutoDetectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoDetect, true
}

// SetAutoDetect sets field value
func (o *UpdateSignInExpRequestLanguageInfo) SetAutoDetect(v bool) {
	o.AutoDetect = v
}

// GetFallbackLanguage returns the FallbackLanguage field value
func (o *UpdateSignInExpRequestLanguageInfo) GetFallbackLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FallbackLanguage
}

// GetFallbackLanguageOk returns a tuple with the FallbackLanguage field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequestLanguageInfo) GetFallbackLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FallbackLanguage, true
}

// SetFallbackLanguage sets field value
func (o *UpdateSignInExpRequestLanguageInfo) SetFallbackLanguage(v string) {
	o.FallbackLanguage = v
}

func (o UpdateSignInExpRequestLanguageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSignInExpRequestLanguageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["autoDetect"] = o.AutoDetect
	toSerialize["fallbackLanguage"] = o.FallbackLanguage
	return toSerialize, nil
}

func (o *UpdateSignInExpRequestLanguageInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"autoDetect",
		"fallbackLanguage",
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

	varUpdateSignInExpRequestLanguageInfo := _UpdateSignInExpRequestLanguageInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSignInExpRequestLanguageInfo)

	if err != nil {
		return err
	}

	*o = UpdateSignInExpRequestLanguageInfo(varUpdateSignInExpRequestLanguageInfo)

	return err
}

type NullableUpdateSignInExpRequestLanguageInfo struct {
	value *UpdateSignInExpRequestLanguageInfo
	isSet bool
}

func (v NullableUpdateSignInExpRequestLanguageInfo) Get() *UpdateSignInExpRequestLanguageInfo {
	return v.value
}

func (v *NullableUpdateSignInExpRequestLanguageInfo) Set(val *UpdateSignInExpRequestLanguageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSignInExpRequestLanguageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSignInExpRequestLanguageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSignInExpRequestLanguageInfo(val *UpdateSignInExpRequestLanguageInfo) *NullableUpdateSignInExpRequestLanguageInfo {
	return &NullableUpdateSignInExpRequestLanguageInfo{value: val, isSet: true}
}

func (v NullableUpdateSignInExpRequestLanguageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSignInExpRequestLanguageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


