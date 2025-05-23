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

// checks if the UpdateApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApplicationRequest{}

// UpdateApplicationRequest struct for UpdateApplicationRequest
type UpdateApplicationRequest struct {
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	OidcClientMetadata *ListApplications200ResponseInnerOidcClientMetadata `json:"oidcClientMetadata,omitempty"`
	CustomClientMetadata *ListApplications200ResponseInnerCustomClientMetadata `json:"customClientMetadata,omitempty"`
	// arbitrary
	CustomData map[string]interface{} `json:"customData,omitempty"`
	ProtectedAppMetadata *UpdateApplicationRequestProtectedAppMetadata `json:"protectedAppMetadata,omitempty"`
	// Whether the application has admin access. User can enable the admin access for Machine-to-Machine apps.
	IsAdmin *bool `json:"isAdmin,omitempty"`
}

// NewUpdateApplicationRequest instantiates a new UpdateApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApplicationRequest() *UpdateApplicationRequest {
	this := UpdateApplicationRequest{}
	return &this
}

// NewUpdateApplicationRequestWithDefaults instantiates a new UpdateApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApplicationRequestWithDefaults() *UpdateApplicationRequest {
	this := UpdateApplicationRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateApplicationRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateApplicationRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateApplicationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateApplicationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateApplicationRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateApplicationRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateApplicationRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetOidcClientMetadata returns the OidcClientMetadata field value if set, zero value otherwise.
func (o *UpdateApplicationRequest) GetOidcClientMetadata() ListApplications200ResponseInnerOidcClientMetadata {
	if o == nil || IsNil(o.OidcClientMetadata) {
		var ret ListApplications200ResponseInnerOidcClientMetadata
		return ret
	}
	return *o.OidcClientMetadata
}

// GetOidcClientMetadataOk returns a tuple with the OidcClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequest) GetOidcClientMetadataOk() (*ListApplications200ResponseInnerOidcClientMetadata, bool) {
	if o == nil || IsNil(o.OidcClientMetadata) {
		return nil, false
	}
	return o.OidcClientMetadata, true
}

// HasOidcClientMetadata returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasOidcClientMetadata() bool {
	if o != nil && !IsNil(o.OidcClientMetadata) {
		return true
	}

	return false
}

// SetOidcClientMetadata gets a reference to the given ListApplications200ResponseInnerOidcClientMetadata and assigns it to the OidcClientMetadata field.
func (o *UpdateApplicationRequest) SetOidcClientMetadata(v ListApplications200ResponseInnerOidcClientMetadata) {
	o.OidcClientMetadata = &v
}

// GetCustomClientMetadata returns the CustomClientMetadata field value if set, zero value otherwise.
func (o *UpdateApplicationRequest) GetCustomClientMetadata() ListApplications200ResponseInnerCustomClientMetadata {
	if o == nil || IsNil(o.CustomClientMetadata) {
		var ret ListApplications200ResponseInnerCustomClientMetadata
		return ret
	}
	return *o.CustomClientMetadata
}

// GetCustomClientMetadataOk returns a tuple with the CustomClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequest) GetCustomClientMetadataOk() (*ListApplications200ResponseInnerCustomClientMetadata, bool) {
	if o == nil || IsNil(o.CustomClientMetadata) {
		return nil, false
	}
	return o.CustomClientMetadata, true
}

// HasCustomClientMetadata returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasCustomClientMetadata() bool {
	if o != nil && !IsNil(o.CustomClientMetadata) {
		return true
	}

	return false
}

// SetCustomClientMetadata gets a reference to the given ListApplications200ResponseInnerCustomClientMetadata and assigns it to the CustomClientMetadata field.
func (o *UpdateApplicationRequest) SetCustomClientMetadata(v ListApplications200ResponseInnerCustomClientMetadata) {
	o.CustomClientMetadata = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *UpdateApplicationRequest) GetCustomData() map[string]interface{} {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequest) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomData) {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]interface{} and assigns it to the CustomData field.
func (o *UpdateApplicationRequest) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetProtectedAppMetadata returns the ProtectedAppMetadata field value if set, zero value otherwise.
func (o *UpdateApplicationRequest) GetProtectedAppMetadata() UpdateApplicationRequestProtectedAppMetadata {
	if o == nil || IsNil(o.ProtectedAppMetadata) {
		var ret UpdateApplicationRequestProtectedAppMetadata
		return ret
	}
	return *o.ProtectedAppMetadata
}

// GetProtectedAppMetadataOk returns a tuple with the ProtectedAppMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequest) GetProtectedAppMetadataOk() (*UpdateApplicationRequestProtectedAppMetadata, bool) {
	if o == nil || IsNil(o.ProtectedAppMetadata) {
		return nil, false
	}
	return o.ProtectedAppMetadata, true
}

// HasProtectedAppMetadata returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasProtectedAppMetadata() bool {
	if o != nil && !IsNil(o.ProtectedAppMetadata) {
		return true
	}

	return false
}

// SetProtectedAppMetadata gets a reference to the given UpdateApplicationRequestProtectedAppMetadata and assigns it to the ProtectedAppMetadata field.
func (o *UpdateApplicationRequest) SetProtectedAppMetadata(v UpdateApplicationRequestProtectedAppMetadata) {
	o.ProtectedAppMetadata = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *UpdateApplicationRequest) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequest) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *UpdateApplicationRequest) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

func (o UpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.OidcClientMetadata) {
		toSerialize["oidcClientMetadata"] = o.OidcClientMetadata
	}
	if !IsNil(o.CustomClientMetadata) {
		toSerialize["customClientMetadata"] = o.CustomClientMetadata
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	if !IsNil(o.ProtectedAppMetadata) {
		toSerialize["protectedAppMetadata"] = o.ProtectedAppMetadata
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	return toSerialize, nil
}

type NullableUpdateApplicationRequest struct {
	value *UpdateApplicationRequest
	isSet bool
}

func (v NullableUpdateApplicationRequest) Get() *UpdateApplicationRequest {
	return v.value
}

func (v *NullableUpdateApplicationRequest) Set(val *UpdateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApplicationRequest(val *UpdateApplicationRequest) *NullableUpdateApplicationRequest {
	return &NullableUpdateApplicationRequest{value: val, isSet: true}
}

func (v NullableUpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


