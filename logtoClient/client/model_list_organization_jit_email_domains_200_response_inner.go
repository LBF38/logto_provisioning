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

// checks if the ListOrganizationJitEmailDomains200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationJitEmailDomains200ResponseInner{}

// ListOrganizationJitEmailDomains200ResponseInner struct for ListOrganizationJitEmailDomains200ResponseInner
type ListOrganizationJitEmailDomains200ResponseInner struct {
	TenantId string `json:"tenantId"`
	OrganizationId string `json:"organizationId"`
	EmailDomain string `json:"emailDomain"`
}

type _ListOrganizationJitEmailDomains200ResponseInner ListOrganizationJitEmailDomains200ResponseInner

// NewListOrganizationJitEmailDomains200ResponseInner instantiates a new ListOrganizationJitEmailDomains200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationJitEmailDomains200ResponseInner(tenantId string, organizationId string, emailDomain string) *ListOrganizationJitEmailDomains200ResponseInner {
	this := ListOrganizationJitEmailDomains200ResponseInner{}
	this.TenantId = tenantId
	this.OrganizationId = organizationId
	this.EmailDomain = emailDomain
	return &this
}

// NewListOrganizationJitEmailDomains200ResponseInnerWithDefaults instantiates a new ListOrganizationJitEmailDomains200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationJitEmailDomains200ResponseInnerWithDefaults() *ListOrganizationJitEmailDomains200ResponseInner {
	this := ListOrganizationJitEmailDomains200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListOrganizationJitEmailDomains200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationJitEmailDomains200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListOrganizationJitEmailDomains200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *ListOrganizationJitEmailDomains200ResponseInner) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationJitEmailDomains200ResponseInner) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *ListOrganizationJitEmailDomains200ResponseInner) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetEmailDomain returns the EmailDomain field value
func (o *ListOrganizationJitEmailDomains200ResponseInner) GetEmailDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailDomain
}

// GetEmailDomainOk returns a tuple with the EmailDomain field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationJitEmailDomains200ResponseInner) GetEmailDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailDomain, true
}

// SetEmailDomain sets field value
func (o *ListOrganizationJitEmailDomains200ResponseInner) SetEmailDomain(v string) {
	o.EmailDomain = v
}

func (o ListOrganizationJitEmailDomains200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationJitEmailDomains200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["emailDomain"] = o.EmailDomain
	return toSerialize, nil
}

func (o *ListOrganizationJitEmailDomains200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"organizationId",
		"emailDomain",
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

	varListOrganizationJitEmailDomains200ResponseInner := _ListOrganizationJitEmailDomains200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListOrganizationJitEmailDomains200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListOrganizationJitEmailDomains200ResponseInner(varListOrganizationJitEmailDomains200ResponseInner)

	return err
}

type NullableListOrganizationJitEmailDomains200ResponseInner struct {
	value *ListOrganizationJitEmailDomains200ResponseInner
	isSet bool
}

func (v NullableListOrganizationJitEmailDomains200ResponseInner) Get() *ListOrganizationJitEmailDomains200ResponseInner {
	return v.value
}

func (v *NullableListOrganizationJitEmailDomains200ResponseInner) Set(val *ListOrganizationJitEmailDomains200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationJitEmailDomains200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationJitEmailDomains200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationJitEmailDomains200ResponseInner(val *ListOrganizationJitEmailDomains200ResponseInner) *NullableListOrganizationJitEmailDomains200ResponseInner {
	return &NullableListOrganizationJitEmailDomains200ResponseInner{value: val, isSet: true}
}

func (v NullableListOrganizationJitEmailDomains200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationJitEmailDomains200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


