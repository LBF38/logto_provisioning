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

// checks if the UpdateSignInExp200ResponseColor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSignInExp200ResponseColor{}

// UpdateSignInExp200ResponseColor struct for UpdateSignInExp200ResponseColor
type UpdateSignInExp200ResponseColor struct {
	PrimaryColor string `json:"primaryColor" validate:"regexp=^#[\\\\da-f]{3}([\\\\da-f]{3})?$/i"`
	IsDarkModeEnabled bool `json:"isDarkModeEnabled"`
	DarkPrimaryColor string `json:"darkPrimaryColor" validate:"regexp=^#[\\\\da-f]{3}([\\\\da-f]{3})?$/i"`
}

type _UpdateSignInExp200ResponseColor UpdateSignInExp200ResponseColor

// NewUpdateSignInExp200ResponseColor instantiates a new UpdateSignInExp200ResponseColor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSignInExp200ResponseColor(primaryColor string, isDarkModeEnabled bool, darkPrimaryColor string) *UpdateSignInExp200ResponseColor {
	this := UpdateSignInExp200ResponseColor{}
	this.PrimaryColor = primaryColor
	this.IsDarkModeEnabled = isDarkModeEnabled
	this.DarkPrimaryColor = darkPrimaryColor
	return &this
}

// NewUpdateSignInExp200ResponseColorWithDefaults instantiates a new UpdateSignInExp200ResponseColor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSignInExp200ResponseColorWithDefaults() *UpdateSignInExp200ResponseColor {
	this := UpdateSignInExp200ResponseColor{}
	return &this
}

// GetPrimaryColor returns the PrimaryColor field value
func (o *UpdateSignInExp200ResponseColor) GetPrimaryColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExp200ResponseColor) GetPrimaryColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColor, true
}

// SetPrimaryColor sets field value
func (o *UpdateSignInExp200ResponseColor) SetPrimaryColor(v string) {
	o.PrimaryColor = v
}

// GetIsDarkModeEnabled returns the IsDarkModeEnabled field value
func (o *UpdateSignInExp200ResponseColor) GetIsDarkModeEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDarkModeEnabled
}

// GetIsDarkModeEnabledOk returns a tuple with the IsDarkModeEnabled field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExp200ResponseColor) GetIsDarkModeEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDarkModeEnabled, true
}

// SetIsDarkModeEnabled sets field value
func (o *UpdateSignInExp200ResponseColor) SetIsDarkModeEnabled(v bool) {
	o.IsDarkModeEnabled = v
}

// GetDarkPrimaryColor returns the DarkPrimaryColor field value
func (o *UpdateSignInExp200ResponseColor) GetDarkPrimaryColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DarkPrimaryColor
}

// GetDarkPrimaryColorOk returns a tuple with the DarkPrimaryColor field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExp200ResponseColor) GetDarkPrimaryColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DarkPrimaryColor, true
}

// SetDarkPrimaryColor sets field value
func (o *UpdateSignInExp200ResponseColor) SetDarkPrimaryColor(v string) {
	o.DarkPrimaryColor = v
}

func (o UpdateSignInExp200ResponseColor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSignInExp200ResponseColor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["primaryColor"] = o.PrimaryColor
	toSerialize["isDarkModeEnabled"] = o.IsDarkModeEnabled
	toSerialize["darkPrimaryColor"] = o.DarkPrimaryColor
	return toSerialize, nil
}

func (o *UpdateSignInExp200ResponseColor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"primaryColor",
		"isDarkModeEnabled",
		"darkPrimaryColor",
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

	varUpdateSignInExp200ResponseColor := _UpdateSignInExp200ResponseColor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSignInExp200ResponseColor)

	if err != nil {
		return err
	}

	*o = UpdateSignInExp200ResponseColor(varUpdateSignInExp200ResponseColor)

	return err
}

type NullableUpdateSignInExp200ResponseColor struct {
	value *UpdateSignInExp200ResponseColor
	isSet bool
}

func (v NullableUpdateSignInExp200ResponseColor) Get() *UpdateSignInExp200ResponseColor {
	return v.value
}

func (v *NullableUpdateSignInExp200ResponseColor) Set(val *UpdateSignInExp200ResponseColor) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSignInExp200ResponseColor) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSignInExp200ResponseColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSignInExp200ResponseColor(val *UpdateSignInExp200ResponseColor) *NullableUpdateSignInExp200ResponseColor {
	return &NullableUpdateSignInExp200ResponseColor{value: val, isSet: true}
}

func (v NullableUpdateSignInExp200ResponseColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSignInExp200ResponseColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


