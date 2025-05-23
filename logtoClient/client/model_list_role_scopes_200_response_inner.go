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

// checks if the ListRoleScopes200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRoleScopes200ResponseInner{}

// ListRoleScopes200ResponseInner struct for ListRoleScopes200ResponseInner
type ListRoleScopes200ResponseInner struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	ResourceId string `json:"resourceId"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	CreatedAt float32 `json:"createdAt"`
	Resource GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource `json:"resource"`
}

type _ListRoleScopes200ResponseInner ListRoleScopes200ResponseInner

// NewListRoleScopes200ResponseInner instantiates a new ListRoleScopes200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRoleScopes200ResponseInner(tenantId string, id string, resourceId string, name string, description NullableString, createdAt float32, resource GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource) *ListRoleScopes200ResponseInner {
	this := ListRoleScopes200ResponseInner{}
	this.TenantId = tenantId
	this.Id = id
	this.ResourceId = resourceId
	this.Name = name
	this.Description = description
	this.CreatedAt = createdAt
	this.Resource = resource
	return &this
}

// NewListRoleScopes200ResponseInnerWithDefaults instantiates a new ListRoleScopes200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRoleScopes200ResponseInnerWithDefaults() *ListRoleScopes200ResponseInner {
	this := ListRoleScopes200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListRoleScopes200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListRoleScopes200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListRoleScopes200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *ListRoleScopes200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListRoleScopes200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListRoleScopes200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetResourceId returns the ResourceId field value
func (o *ListRoleScopes200ResponseInner) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ListRoleScopes200ResponseInner) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ListRoleScopes200ResponseInner) SetResourceId(v string) {
	o.ResourceId = v
}

// GetName returns the Name field value
func (o *ListRoleScopes200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListRoleScopes200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListRoleScopes200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListRoleScopes200ResponseInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListRoleScopes200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ListRoleScopes200ResponseInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListRoleScopes200ResponseInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListRoleScopes200ResponseInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListRoleScopes200ResponseInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetResource returns the Resource field value
func (o *ListRoleScopes200ResponseInner) GetResource() GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource {
	if o == nil {
		var ret GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ListRoleScopes200ResponseInner) GetResourceOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ListRoleScopes200ResponseInner) SetResource(v GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource) {
	o.Resource = v
}

func (o ListRoleScopes200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRoleScopes200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

func (o *ListRoleScopes200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"resourceId",
		"name",
		"description",
		"createdAt",
		"resource",
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

	varListRoleScopes200ResponseInner := _ListRoleScopes200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListRoleScopes200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListRoleScopes200ResponseInner(varListRoleScopes200ResponseInner)

	return err
}

type NullableListRoleScopes200ResponseInner struct {
	value *ListRoleScopes200ResponseInner
	isSet bool
}

func (v NullableListRoleScopes200ResponseInner) Get() *ListRoleScopes200ResponseInner {
	return v.value
}

func (v *NullableListRoleScopes200ResponseInner) Set(val *ListRoleScopes200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListRoleScopes200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListRoleScopes200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRoleScopes200ResponseInner(val *ListRoleScopes200ResponseInner) *NullableListRoleScopes200ResponseInner {
	return &NullableListRoleScopes200ResponseInner{value: val, isSet: true}
}

func (v NullableListRoleScopes200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRoleScopes200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


