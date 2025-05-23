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

// checks if the ListOrganizationApplications200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationApplications200ResponseInner{}

// ListOrganizationApplications200ResponseInner struct for ListOrganizationApplications200ResponseInner
type ListOrganizationApplications200ResponseInner struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Name string `json:"name"`
	// The internal client secret. Note it is only used for internal validation, and the actual secrets should be retrieved from `/api/applications/{id}/secrets` endpoints.
	// Deprecated
	Secret string `json:"secret"`
	Description NullableString `json:"description"`
	Type string `json:"type"`
	OidcClientMetadata ListApplications200ResponseInnerOidcClientMetadata `json:"oidcClientMetadata"`
	CustomClientMetadata ListApplications200ResponseInnerCustomClientMetadata `json:"customClientMetadata"`
	ProtectedAppMetadata NullableListApplications200ResponseInnerProtectedAppMetadata `json:"protectedAppMetadata"`
	// arbitrary
	CustomData map[string]interface{} `json:"customData"`
	IsThirdParty bool `json:"isThirdParty"`
	CreatedAt float32 `json:"createdAt"`
	OrganizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner `json:"organizationRoles"`
}

type _ListOrganizationApplications200ResponseInner ListOrganizationApplications200ResponseInner

// NewListOrganizationApplications200ResponseInner instantiates a new ListOrganizationApplications200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationApplications200ResponseInner(tenantId string, id string, name string, secret string, description NullableString, type_ string, oidcClientMetadata ListApplications200ResponseInnerOidcClientMetadata, customClientMetadata ListApplications200ResponseInnerCustomClientMetadata, protectedAppMetadata NullableListApplications200ResponseInnerProtectedAppMetadata, customData map[string]interface{}, isThirdParty bool, createdAt float32, organizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) *ListOrganizationApplications200ResponseInner {
	this := ListOrganizationApplications200ResponseInner{}
	this.TenantId = tenantId
	this.Id = id
	this.Name = name
	this.Secret = secret
	this.Description = description
	this.Type = type_
	this.OidcClientMetadata = oidcClientMetadata
	this.CustomClientMetadata = customClientMetadata
	this.ProtectedAppMetadata = protectedAppMetadata
	this.CustomData = customData
	this.IsThirdParty = isThirdParty
	this.CreatedAt = createdAt
	this.OrganizationRoles = organizationRoles
	return &this
}

// NewListOrganizationApplications200ResponseInnerWithDefaults instantiates a new ListOrganizationApplications200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationApplications200ResponseInnerWithDefaults() *ListOrganizationApplications200ResponseInner {
	this := ListOrganizationApplications200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListOrganizationApplications200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListOrganizationApplications200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *ListOrganizationApplications200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListOrganizationApplications200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListOrganizationApplications200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListOrganizationApplications200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetSecret returns the Secret field value
// Deprecated
func (o *ListOrganizationApplications200ResponseInner) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ListOrganizationApplications200ResponseInner) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
// Deprecated
func (o *ListOrganizationApplications200ResponseInner) SetSecret(v string) {
	o.Secret = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListOrganizationApplications200ResponseInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListOrganizationApplications200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ListOrganizationApplications200ResponseInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetType returns the Type field value
func (o *ListOrganizationApplications200ResponseInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListOrganizationApplications200ResponseInner) SetType(v string) {
	o.Type = v
}

// GetOidcClientMetadata returns the OidcClientMetadata field value
func (o *ListOrganizationApplications200ResponseInner) GetOidcClientMetadata() ListApplications200ResponseInnerOidcClientMetadata {
	if o == nil {
		var ret ListApplications200ResponseInnerOidcClientMetadata
		return ret
	}

	return o.OidcClientMetadata
}

// GetOidcClientMetadataOk returns a tuple with the OidcClientMetadata field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetOidcClientMetadataOk() (*ListApplications200ResponseInnerOidcClientMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcClientMetadata, true
}

// SetOidcClientMetadata sets field value
func (o *ListOrganizationApplications200ResponseInner) SetOidcClientMetadata(v ListApplications200ResponseInnerOidcClientMetadata) {
	o.OidcClientMetadata = v
}

// GetCustomClientMetadata returns the CustomClientMetadata field value
func (o *ListOrganizationApplications200ResponseInner) GetCustomClientMetadata() ListApplications200ResponseInnerCustomClientMetadata {
	if o == nil {
		var ret ListApplications200ResponseInnerCustomClientMetadata
		return ret
	}

	return o.CustomClientMetadata
}

// GetCustomClientMetadataOk returns a tuple with the CustomClientMetadata field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetCustomClientMetadataOk() (*ListApplications200ResponseInnerCustomClientMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomClientMetadata, true
}

// SetCustomClientMetadata sets field value
func (o *ListOrganizationApplications200ResponseInner) SetCustomClientMetadata(v ListApplications200ResponseInnerCustomClientMetadata) {
	o.CustomClientMetadata = v
}

// GetProtectedAppMetadata returns the ProtectedAppMetadata field value
// If the value is explicit nil, the zero value for ListApplications200ResponseInnerProtectedAppMetadata will be returned
func (o *ListOrganizationApplications200ResponseInner) GetProtectedAppMetadata() ListApplications200ResponseInnerProtectedAppMetadata {
	if o == nil || o.ProtectedAppMetadata.Get() == nil {
		var ret ListApplications200ResponseInnerProtectedAppMetadata
		return ret
	}

	return *o.ProtectedAppMetadata.Get()
}

// GetProtectedAppMetadataOk returns a tuple with the ProtectedAppMetadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListOrganizationApplications200ResponseInner) GetProtectedAppMetadataOk() (*ListApplications200ResponseInnerProtectedAppMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectedAppMetadata.Get(), o.ProtectedAppMetadata.IsSet()
}

// SetProtectedAppMetadata sets field value
func (o *ListOrganizationApplications200ResponseInner) SetProtectedAppMetadata(v ListApplications200ResponseInnerProtectedAppMetadata) {
	o.ProtectedAppMetadata.Set(&v)
}

// GetCustomData returns the CustomData field value
func (o *ListOrganizationApplications200ResponseInner) GetCustomData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// SetCustomData sets field value
func (o *ListOrganizationApplications200ResponseInner) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetIsThirdParty returns the IsThirdParty field value
func (o *ListOrganizationApplications200ResponseInner) GetIsThirdParty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsThirdParty
}

// GetIsThirdPartyOk returns a tuple with the IsThirdParty field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetIsThirdPartyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsThirdParty, true
}

// SetIsThirdParty sets field value
func (o *ListOrganizationApplications200ResponseInner) SetIsThirdParty(v bool) {
	o.IsThirdParty = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListOrganizationApplications200ResponseInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListOrganizationApplications200ResponseInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetOrganizationRoles returns the OrganizationRoles field value
func (o *ListOrganizationApplications200ResponseInner) GetOrganizationRoles() []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner {
	if o == nil {
		var ret []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner
		return ret
	}

	return o.OrganizationRoles
}

// GetOrganizationRolesOk returns a tuple with the OrganizationRoles field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationApplications200ResponseInner) GetOrganizationRolesOk() ([]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationRoles, true
}

// SetOrganizationRoles sets field value
func (o *ListOrganizationApplications200ResponseInner) SetOrganizationRoles(v []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) {
	o.OrganizationRoles = v
}

func (o ListOrganizationApplications200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationApplications200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["secret"] = o.Secret
	toSerialize["description"] = o.Description.Get()
	toSerialize["type"] = o.Type
	toSerialize["oidcClientMetadata"] = o.OidcClientMetadata
	toSerialize["customClientMetadata"] = o.CustomClientMetadata
	toSerialize["protectedAppMetadata"] = o.ProtectedAppMetadata.Get()
	toSerialize["customData"] = o.CustomData
	toSerialize["isThirdParty"] = o.IsThirdParty
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["organizationRoles"] = o.OrganizationRoles
	return toSerialize, nil
}

func (o *ListOrganizationApplications200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"name",
		"secret",
		"description",
		"type",
		"oidcClientMetadata",
		"customClientMetadata",
		"protectedAppMetadata",
		"customData",
		"isThirdParty",
		"createdAt",
		"organizationRoles",
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

	varListOrganizationApplications200ResponseInner := _ListOrganizationApplications200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListOrganizationApplications200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListOrganizationApplications200ResponseInner(varListOrganizationApplications200ResponseInner)

	return err
}

type NullableListOrganizationApplications200ResponseInner struct {
	value *ListOrganizationApplications200ResponseInner
	isSet bool
}

func (v NullableListOrganizationApplications200ResponseInner) Get() *ListOrganizationApplications200ResponseInner {
	return v.value
}

func (v *NullableListOrganizationApplications200ResponseInner) Set(val *ListOrganizationApplications200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationApplications200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationApplications200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationApplications200ResponseInner(val *ListOrganizationApplications200ResponseInner) *NullableListOrganizationApplications200ResponseInner {
	return &NullableListOrganizationApplications200ResponseInner{value: val, isSet: true}
}

func (v NullableListOrganizationApplications200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationApplications200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


