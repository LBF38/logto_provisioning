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

// checks if the GetJwtCustomizer200ResponseOneOf1TokenSample type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetJwtCustomizer200ResponseOneOf1TokenSample{}

// GetJwtCustomizer200ResponseOneOf1TokenSample struct for GetJwtCustomizer200ResponseOneOf1TokenSample
type GetJwtCustomizer200ResponseOneOf1TokenSample struct {
	Jti *string `json:"jti,omitempty"`
	Aud *GetJwtCustomizer200ResponseOneOfTokenSampleAud `json:"aud,omitempty"`
	Scope *string `json:"scope,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	Kind *string `json:"kind,omitempty"`
}

// NewGetJwtCustomizer200ResponseOneOf1TokenSample instantiates a new GetJwtCustomizer200ResponseOneOf1TokenSample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJwtCustomizer200ResponseOneOf1TokenSample() *GetJwtCustomizer200ResponseOneOf1TokenSample {
	this := GetJwtCustomizer200ResponseOneOf1TokenSample{}
	return &this
}

// NewGetJwtCustomizer200ResponseOneOf1TokenSampleWithDefaults instantiates a new GetJwtCustomizer200ResponseOneOf1TokenSample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJwtCustomizer200ResponseOneOf1TokenSampleWithDefaults() *GetJwtCustomizer200ResponseOneOf1TokenSample {
	this := GetJwtCustomizer200ResponseOneOf1TokenSample{}
	return &this
}

// GetJti returns the Jti field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetJti() string {
	if o == nil || IsNil(o.Jti) {
		var ret string
		return ret
	}
	return *o.Jti
}

// GetJtiOk returns a tuple with the Jti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetJtiOk() (*string, bool) {
	if o == nil || IsNil(o.Jti) {
		return nil, false
	}
	return o.Jti, true
}

// HasJti returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) HasJti() bool {
	if o != nil && !IsNil(o.Jti) {
		return true
	}

	return false
}

// SetJti gets a reference to the given string and assigns it to the Jti field.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) SetJti(v string) {
	o.Jti = &v
}

// GetAud returns the Aud field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetAud() GetJwtCustomizer200ResponseOneOfTokenSampleAud {
	if o == nil || IsNil(o.Aud) {
		var ret GetJwtCustomizer200ResponseOneOfTokenSampleAud
		return ret
	}
	return *o.Aud
}

// GetAudOk returns a tuple with the Aud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetAudOk() (*GetJwtCustomizer200ResponseOneOfTokenSampleAud, bool) {
	if o == nil || IsNil(o.Aud) {
		return nil, false
	}
	return o.Aud, true
}

// HasAud returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) HasAud() bool {
	if o != nil && !IsNil(o.Aud) {
		return true
	}

	return false
}

// SetAud gets a reference to the given GetJwtCustomizer200ResponseOneOfTokenSampleAud and assigns it to the Aud field.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) SetAud(v GetJwtCustomizer200ResponseOneOfTokenSampleAud) {
	o.Aud = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) SetScope(v string) {
	o.Scope = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) SetClientId(v string) {
	o.ClientId = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GetJwtCustomizer200ResponseOneOf1TokenSample) SetKind(v string) {
	o.Kind = &v
}

func (o GetJwtCustomizer200ResponseOneOf1TokenSample) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetJwtCustomizer200ResponseOneOf1TokenSample) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jti) {
		toSerialize["jti"] = o.Jti
	}
	if !IsNil(o.Aud) {
		toSerialize["aud"] = o.Aud
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return toSerialize, nil
}

type NullableGetJwtCustomizer200ResponseOneOf1TokenSample struct {
	value *GetJwtCustomizer200ResponseOneOf1TokenSample
	isSet bool
}

func (v NullableGetJwtCustomizer200ResponseOneOf1TokenSample) Get() *GetJwtCustomizer200ResponseOneOf1TokenSample {
	return v.value
}

func (v *NullableGetJwtCustomizer200ResponseOneOf1TokenSample) Set(val *GetJwtCustomizer200ResponseOneOf1TokenSample) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJwtCustomizer200ResponseOneOf1TokenSample) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJwtCustomizer200ResponseOneOf1TokenSample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJwtCustomizer200ResponseOneOf1TokenSample(val *GetJwtCustomizer200ResponseOneOf1TokenSample) *NullableGetJwtCustomizer200ResponseOneOf1TokenSample {
	return &NullableGetJwtCustomizer200ResponseOneOf1TokenSample{value: val, isSet: true}
}

func (v NullableGetJwtCustomizer200ResponseOneOf1TokenSample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJwtCustomizer200ResponseOneOf1TokenSample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


