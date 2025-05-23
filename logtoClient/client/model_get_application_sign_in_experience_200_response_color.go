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

// checks if the GetApplicationSignInExperience200ResponseColor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApplicationSignInExperience200ResponseColor{}

// GetApplicationSignInExperience200ResponseColor struct for GetApplicationSignInExperience200ResponseColor
type GetApplicationSignInExperience200ResponseColor struct {
	PrimaryColor *string `json:"primaryColor,omitempty" validate:"regexp=^#[\\\\da-f]{3}([\\\\da-f]{3})?$/i"`
	IsDarkModeEnabled *bool `json:"isDarkModeEnabled,omitempty"`
	DarkPrimaryColor *string `json:"darkPrimaryColor,omitempty" validate:"regexp=^#[\\\\da-f]{3}([\\\\da-f]{3})?$/i"`
}

// NewGetApplicationSignInExperience200ResponseColor instantiates a new GetApplicationSignInExperience200ResponseColor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApplicationSignInExperience200ResponseColor() *GetApplicationSignInExperience200ResponseColor {
	this := GetApplicationSignInExperience200ResponseColor{}
	return &this
}

// NewGetApplicationSignInExperience200ResponseColorWithDefaults instantiates a new GetApplicationSignInExperience200ResponseColor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApplicationSignInExperience200ResponseColorWithDefaults() *GetApplicationSignInExperience200ResponseColor {
	this := GetApplicationSignInExperience200ResponseColor{}
	return &this
}

// GetPrimaryColor returns the PrimaryColor field value if set, zero value otherwise.
func (o *GetApplicationSignInExperience200ResponseColor) GetPrimaryColor() string {
	if o == nil || IsNil(o.PrimaryColor) {
		var ret string
		return ret
	}
	return *o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplicationSignInExperience200ResponseColor) GetPrimaryColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryColor) {
		return nil, false
	}
	return o.PrimaryColor, true
}

// HasPrimaryColor returns a boolean if a field has been set.
func (o *GetApplicationSignInExperience200ResponseColor) HasPrimaryColor() bool {
	if o != nil && !IsNil(o.PrimaryColor) {
		return true
	}

	return false
}

// SetPrimaryColor gets a reference to the given string and assigns it to the PrimaryColor field.
func (o *GetApplicationSignInExperience200ResponseColor) SetPrimaryColor(v string) {
	o.PrimaryColor = &v
}

// GetIsDarkModeEnabled returns the IsDarkModeEnabled field value if set, zero value otherwise.
func (o *GetApplicationSignInExperience200ResponseColor) GetIsDarkModeEnabled() bool {
	if o == nil || IsNil(o.IsDarkModeEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDarkModeEnabled
}

// GetIsDarkModeEnabledOk returns a tuple with the IsDarkModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplicationSignInExperience200ResponseColor) GetIsDarkModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDarkModeEnabled) {
		return nil, false
	}
	return o.IsDarkModeEnabled, true
}

// HasIsDarkModeEnabled returns a boolean if a field has been set.
func (o *GetApplicationSignInExperience200ResponseColor) HasIsDarkModeEnabled() bool {
	if o != nil && !IsNil(o.IsDarkModeEnabled) {
		return true
	}

	return false
}

// SetIsDarkModeEnabled gets a reference to the given bool and assigns it to the IsDarkModeEnabled field.
func (o *GetApplicationSignInExperience200ResponseColor) SetIsDarkModeEnabled(v bool) {
	o.IsDarkModeEnabled = &v
}

// GetDarkPrimaryColor returns the DarkPrimaryColor field value if set, zero value otherwise.
func (o *GetApplicationSignInExperience200ResponseColor) GetDarkPrimaryColor() string {
	if o == nil || IsNil(o.DarkPrimaryColor) {
		var ret string
		return ret
	}
	return *o.DarkPrimaryColor
}

// GetDarkPrimaryColorOk returns a tuple with the DarkPrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplicationSignInExperience200ResponseColor) GetDarkPrimaryColorOk() (*string, bool) {
	if o == nil || IsNil(o.DarkPrimaryColor) {
		return nil, false
	}
	return o.DarkPrimaryColor, true
}

// HasDarkPrimaryColor returns a boolean if a field has been set.
func (o *GetApplicationSignInExperience200ResponseColor) HasDarkPrimaryColor() bool {
	if o != nil && !IsNil(o.DarkPrimaryColor) {
		return true
	}

	return false
}

// SetDarkPrimaryColor gets a reference to the given string and assigns it to the DarkPrimaryColor field.
func (o *GetApplicationSignInExperience200ResponseColor) SetDarkPrimaryColor(v string) {
	o.DarkPrimaryColor = &v
}

func (o GetApplicationSignInExperience200ResponseColor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApplicationSignInExperience200ResponseColor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrimaryColor) {
		toSerialize["primaryColor"] = o.PrimaryColor
	}
	if !IsNil(o.IsDarkModeEnabled) {
		toSerialize["isDarkModeEnabled"] = o.IsDarkModeEnabled
	}
	if !IsNil(o.DarkPrimaryColor) {
		toSerialize["darkPrimaryColor"] = o.DarkPrimaryColor
	}
	return toSerialize, nil
}

type NullableGetApplicationSignInExperience200ResponseColor struct {
	value *GetApplicationSignInExperience200ResponseColor
	isSet bool
}

func (v NullableGetApplicationSignInExperience200ResponseColor) Get() *GetApplicationSignInExperience200ResponseColor {
	return v.value
}

func (v *NullableGetApplicationSignInExperience200ResponseColor) Set(val *GetApplicationSignInExperience200ResponseColor) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApplicationSignInExperience200ResponseColor) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApplicationSignInExperience200ResponseColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApplicationSignInExperience200ResponseColor(val *GetApplicationSignInExperience200ResponseColor) *NullableGetApplicationSignInExperience200ResponseColor {
	return &NullableGetApplicationSignInExperience200ResponseColor{value: val, isSet: true}
}

func (v NullableGetApplicationSignInExperience200ResponseColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApplicationSignInExperience200ResponseColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


