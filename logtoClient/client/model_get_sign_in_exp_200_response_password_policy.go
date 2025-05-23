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

// checks if the GetSignInExp200ResponsePasswordPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignInExp200ResponsePasswordPolicy{}

// GetSignInExp200ResponsePasswordPolicy Password policies to adjust the password strength requirements.
type GetSignInExp200ResponsePasswordPolicy struct {
	Length *GetSignInExp200ResponsePasswordPolicyLength `json:"length,omitempty"`
	CharacterTypes *GetSignInExp200ResponsePasswordPolicyCharacterTypes `json:"characterTypes,omitempty"`
	Rejects *GetSignInExp200ResponsePasswordPolicyRejects `json:"rejects,omitempty"`
}

// NewGetSignInExp200ResponsePasswordPolicy instantiates a new GetSignInExp200ResponsePasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignInExp200ResponsePasswordPolicy() *GetSignInExp200ResponsePasswordPolicy {
	this := GetSignInExp200ResponsePasswordPolicy{}
	var length GetSignInExp200ResponsePasswordPolicyLength
	this.Length = &length
	var characterTypes GetSignInExp200ResponsePasswordPolicyCharacterTypes 
	this.CharacterTypes = &characterTypes
	var rejects GetSignInExp200ResponsePasswordPolicyRejects 
	this.Rejects = &rejects
	return &this
}

// NewGetSignInExp200ResponsePasswordPolicyWithDefaults instantiates a new GetSignInExp200ResponsePasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignInExp200ResponsePasswordPolicyWithDefaults() *GetSignInExp200ResponsePasswordPolicy {
	this := GetSignInExp200ResponsePasswordPolicy{}
	var length GetSignInExp200ResponsePasswordPolicyLength 
	this.Length = &length
	var characterTypes GetSignInExp200ResponsePasswordPolicyCharacterTypes 
	this.CharacterTypes = &characterTypes
	var rejects GetSignInExp200ResponsePasswordPolicyRejects 
	this.Rejects = &rejects
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *GetSignInExp200ResponsePasswordPolicy) GetLength() GetSignInExp200ResponsePasswordPolicyLength {
	if o == nil || IsNil(o.Length) {
		var ret GetSignInExp200ResponsePasswordPolicyLength
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponsePasswordPolicy) GetLengthOk() (*GetSignInExp200ResponsePasswordPolicyLength, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *GetSignInExp200ResponsePasswordPolicy) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given GetSignInExp200ResponsePasswordPolicyLength and assigns it to the Length field.
func (o *GetSignInExp200ResponsePasswordPolicy) SetLength(v GetSignInExp200ResponsePasswordPolicyLength) {
	o.Length = &v
}

// GetCharacterTypes returns the CharacterTypes field value if set, zero value otherwise.
func (o *GetSignInExp200ResponsePasswordPolicy) GetCharacterTypes() GetSignInExp200ResponsePasswordPolicyCharacterTypes {
	if o == nil || IsNil(o.CharacterTypes) {
		var ret GetSignInExp200ResponsePasswordPolicyCharacterTypes
		return ret
	}
	return *o.CharacterTypes
}

// GetCharacterTypesOk returns a tuple with the CharacterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponsePasswordPolicy) GetCharacterTypesOk() (*GetSignInExp200ResponsePasswordPolicyCharacterTypes, bool) {
	if o == nil || IsNil(o.CharacterTypes) {
		return nil, false
	}
	return o.CharacterTypes, true
}

// HasCharacterTypes returns a boolean if a field has been set.
func (o *GetSignInExp200ResponsePasswordPolicy) HasCharacterTypes() bool {
	if o != nil && !IsNil(o.CharacterTypes) {
		return true
	}

	return false
}

// SetCharacterTypes gets a reference to the given GetSignInExp200ResponsePasswordPolicyCharacterTypes and assigns it to the CharacterTypes field.
func (o *GetSignInExp200ResponsePasswordPolicy) SetCharacterTypes(v GetSignInExp200ResponsePasswordPolicyCharacterTypes) {
	o.CharacterTypes = &v
}

// GetRejects returns the Rejects field value if set, zero value otherwise.
func (o *GetSignInExp200ResponsePasswordPolicy) GetRejects() GetSignInExp200ResponsePasswordPolicyRejects {
	if o == nil || IsNil(o.Rejects) {
		var ret GetSignInExp200ResponsePasswordPolicyRejects
		return ret
	}
	return *o.Rejects
}

// GetRejectsOk returns a tuple with the Rejects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponsePasswordPolicy) GetRejectsOk() (*GetSignInExp200ResponsePasswordPolicyRejects, bool) {
	if o == nil || IsNil(o.Rejects) {
		return nil, false
	}
	return o.Rejects, true
}

// HasRejects returns a boolean if a field has been set.
func (o *GetSignInExp200ResponsePasswordPolicy) HasRejects() bool {
	if o != nil && !IsNil(o.Rejects) {
		return true
	}

	return false
}

// SetRejects gets a reference to the given GetSignInExp200ResponsePasswordPolicyRejects and assigns it to the Rejects field.
func (o *GetSignInExp200ResponsePasswordPolicy) SetRejects(v GetSignInExp200ResponsePasswordPolicyRejects) {
	o.Rejects = &v
}

func (o GetSignInExp200ResponsePasswordPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignInExp200ResponsePasswordPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.CharacterTypes) {
		toSerialize["characterTypes"] = o.CharacterTypes
	}
	if !IsNil(o.Rejects) {
		toSerialize["rejects"] = o.Rejects
	}
	return toSerialize, nil
}

type NullableGetSignInExp200ResponsePasswordPolicy struct {
	value *GetSignInExp200ResponsePasswordPolicy
	isSet bool
}

func (v NullableGetSignInExp200ResponsePasswordPolicy) Get() *GetSignInExp200ResponsePasswordPolicy {
	return v.value
}

func (v *NullableGetSignInExp200ResponsePasswordPolicy) Set(val *GetSignInExp200ResponsePasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignInExp200ResponsePasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignInExp200ResponsePasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignInExp200ResponsePasswordPolicy(val *GetSignInExp200ResponsePasswordPolicy) *NullableGetSignInExp200ResponsePasswordPolicy {
	return &NullableGetSignInExp200ResponsePasswordPolicy{value: val, isSet: true}
}

func (v NullableGetSignInExp200ResponsePasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignInExp200ResponsePasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


