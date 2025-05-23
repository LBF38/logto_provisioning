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

// checks if the ListApplicationUserConsentOrganizations200ResponseOrganizationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplicationUserConsentOrganizations200ResponseOrganizationsInner{}

// ListApplicationUserConsentOrganizations200ResponseOrganizationsInner struct for ListApplicationUserConsentOrganizations200ResponseOrganizationsInner
type ListApplicationUserConsentOrganizations200ResponseOrganizationsInner struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	// arbitrary
	CustomData map[string]interface{} `json:"customData"`
	IsMfaRequired bool `json:"isMfaRequired"`
	Branding ListApplicationOrganizations200ResponseInnerBranding `json:"branding"`
	CreatedAt float32 `json:"createdAt"`
}

type _ListApplicationUserConsentOrganizations200ResponseOrganizationsInner ListApplicationUserConsentOrganizations200ResponseOrganizationsInner

// NewListApplicationUserConsentOrganizations200ResponseOrganizationsInner instantiates a new ListApplicationUserConsentOrganizations200ResponseOrganizationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplicationUserConsentOrganizations200ResponseOrganizationsInner(tenantId string, id string, name string, description NullableString, customData map[string]interface{}, isMfaRequired bool, branding ListApplicationOrganizations200ResponseInnerBranding, createdAt float32) *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner {
	this := ListApplicationUserConsentOrganizations200ResponseOrganizationsInner{}
	this.TenantId = tenantId
	this.Id = id
	this.Name = name
	this.Description = description
	this.CustomData = customData
	this.IsMfaRequired = isMfaRequired
	this.Branding = branding
	this.CreatedAt = createdAt
	return &this
}

// NewListApplicationUserConsentOrganizations200ResponseOrganizationsInnerWithDefaults instantiates a new ListApplicationUserConsentOrganizations200ResponseOrganizationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplicationUserConsentOrganizations200ResponseOrganizationsInnerWithDefaults() *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner {
	this := ListApplicationUserConsentOrganizations200ResponseOrganizationsInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetCustomData returns the CustomData field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetCustomData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// SetCustomData sets field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetIsMfaRequired returns the IsMfaRequired field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetIsMfaRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMfaRequired
}

// GetIsMfaRequiredOk returns a tuple with the IsMfaRequired field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetIsMfaRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMfaRequired, true
}

// SetIsMfaRequired sets field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetIsMfaRequired(v bool) {
	o.IsMfaRequired = v
}

// GetBranding returns the Branding field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetBranding() ListApplicationOrganizations200ResponseInnerBranding {
	if o == nil {
		var ret ListApplicationOrganizations200ResponseInnerBranding
		return ret
	}

	return o.Branding
}

// GetBrandingOk returns a tuple with the Branding field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branding, true
}

// SetBranding sets field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding) {
	o.Branding = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

func (o ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["customData"] = o.CustomData
	toSerialize["isMfaRequired"] = o.IsMfaRequired
	toSerialize["branding"] = o.Branding
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"name",
		"description",
		"customData",
		"isMfaRequired",
		"branding",
		"createdAt",
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

	varListApplicationUserConsentOrganizations200ResponseOrganizationsInner := _ListApplicationUserConsentOrganizations200ResponseOrganizationsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListApplicationUserConsentOrganizations200ResponseOrganizationsInner)

	if err != nil {
		return err
	}

	*o = ListApplicationUserConsentOrganizations200ResponseOrganizationsInner(varListApplicationUserConsentOrganizations200ResponseOrganizationsInner)

	return err
}

type NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner struct {
	value *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner
	isSet bool
}

func (v NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner) Get() *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner {
	return v.value
}

func (v *NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner) Set(val *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner(val *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) *NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner {
	return &NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner{value: val, isSet: true}
}

func (v NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplicationUserConsentOrganizations200ResponseOrganizationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


