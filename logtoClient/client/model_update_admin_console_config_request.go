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

// checks if the UpdateAdminConsoleConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAdminConsoleConfigRequest{}

// UpdateAdminConsoleConfigRequest struct for UpdateAdminConsoleConfigRequest
type UpdateAdminConsoleConfigRequest struct {
	SignInExperienceCustomized *bool `json:"signInExperienceCustomized,omitempty"`
	OrganizationCreated *bool `json:"organizationCreated,omitempty"`
	DevelopmentTenantMigrationNotification *GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification `json:"developmentTenantMigrationNotification,omitempty"`
	CheckedChargeNotification *GetAdminConsoleConfig200ResponseCheckedChargeNotification `json:"checkedChargeNotification,omitempty"`
}

// NewUpdateAdminConsoleConfigRequest instantiates a new UpdateAdminConsoleConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAdminConsoleConfigRequest() *UpdateAdminConsoleConfigRequest {
	this := UpdateAdminConsoleConfigRequest{}
	return &this
}

// NewUpdateAdminConsoleConfigRequestWithDefaults instantiates a new UpdateAdminConsoleConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAdminConsoleConfigRequestWithDefaults() *UpdateAdminConsoleConfigRequest {
	this := UpdateAdminConsoleConfigRequest{}
	return &this
}

// GetSignInExperienceCustomized returns the SignInExperienceCustomized field value if set, zero value otherwise.
func (o *UpdateAdminConsoleConfigRequest) GetSignInExperienceCustomized() bool {
	if o == nil || IsNil(o.SignInExperienceCustomized) {
		var ret bool
		return ret
	}
	return *o.SignInExperienceCustomized
}

// GetSignInExperienceCustomizedOk returns a tuple with the SignInExperienceCustomized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdminConsoleConfigRequest) GetSignInExperienceCustomizedOk() (*bool, bool) {
	if o == nil || IsNil(o.SignInExperienceCustomized) {
		return nil, false
	}
	return o.SignInExperienceCustomized, true
}

// HasSignInExperienceCustomized returns a boolean if a field has been set.
func (o *UpdateAdminConsoleConfigRequest) HasSignInExperienceCustomized() bool {
	if o != nil && !IsNil(o.SignInExperienceCustomized) {
		return true
	}

	return false
}

// SetSignInExperienceCustomized gets a reference to the given bool and assigns it to the SignInExperienceCustomized field.
func (o *UpdateAdminConsoleConfigRequest) SetSignInExperienceCustomized(v bool) {
	o.SignInExperienceCustomized = &v
}

// GetOrganizationCreated returns the OrganizationCreated field value if set, zero value otherwise.
func (o *UpdateAdminConsoleConfigRequest) GetOrganizationCreated() bool {
	if o == nil || IsNil(o.OrganizationCreated) {
		var ret bool
		return ret
	}
	return *o.OrganizationCreated
}

// GetOrganizationCreatedOk returns a tuple with the OrganizationCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdminConsoleConfigRequest) GetOrganizationCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.OrganizationCreated) {
		return nil, false
	}
	return o.OrganizationCreated, true
}

// HasOrganizationCreated returns a boolean if a field has been set.
func (o *UpdateAdminConsoleConfigRequest) HasOrganizationCreated() bool {
	if o != nil && !IsNil(o.OrganizationCreated) {
		return true
	}

	return false
}

// SetOrganizationCreated gets a reference to the given bool and assigns it to the OrganizationCreated field.
func (o *UpdateAdminConsoleConfigRequest) SetOrganizationCreated(v bool) {
	o.OrganizationCreated = &v
}

// GetDevelopmentTenantMigrationNotification returns the DevelopmentTenantMigrationNotification field value if set, zero value otherwise.
func (o *UpdateAdminConsoleConfigRequest) GetDevelopmentTenantMigrationNotification() GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification {
	if o == nil || IsNil(o.DevelopmentTenantMigrationNotification) {
		var ret GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification
		return ret
	}
	return *o.DevelopmentTenantMigrationNotification
}

// GetDevelopmentTenantMigrationNotificationOk returns a tuple with the DevelopmentTenantMigrationNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdminConsoleConfigRequest) GetDevelopmentTenantMigrationNotificationOk() (*GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification, bool) {
	if o == nil || IsNil(o.DevelopmentTenantMigrationNotification) {
		return nil, false
	}
	return o.DevelopmentTenantMigrationNotification, true
}

// HasDevelopmentTenantMigrationNotification returns a boolean if a field has been set.
func (o *UpdateAdminConsoleConfigRequest) HasDevelopmentTenantMigrationNotification() bool {
	if o != nil && !IsNil(o.DevelopmentTenantMigrationNotification) {
		return true
	}

	return false
}

// SetDevelopmentTenantMigrationNotification gets a reference to the given GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification and assigns it to the DevelopmentTenantMigrationNotification field.
func (o *UpdateAdminConsoleConfigRequest) SetDevelopmentTenantMigrationNotification(v GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification) {
	o.DevelopmentTenantMigrationNotification = &v
}

// GetCheckedChargeNotification returns the CheckedChargeNotification field value if set, zero value otherwise.
func (o *UpdateAdminConsoleConfigRequest) GetCheckedChargeNotification() GetAdminConsoleConfig200ResponseCheckedChargeNotification {
	if o == nil || IsNil(o.CheckedChargeNotification) {
		var ret GetAdminConsoleConfig200ResponseCheckedChargeNotification
		return ret
	}
	return *o.CheckedChargeNotification
}

// GetCheckedChargeNotificationOk returns a tuple with the CheckedChargeNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdminConsoleConfigRequest) GetCheckedChargeNotificationOk() (*GetAdminConsoleConfig200ResponseCheckedChargeNotification, bool) {
	if o == nil || IsNil(o.CheckedChargeNotification) {
		return nil, false
	}
	return o.CheckedChargeNotification, true
}

// HasCheckedChargeNotification returns a boolean if a field has been set.
func (o *UpdateAdminConsoleConfigRequest) HasCheckedChargeNotification() bool {
	if o != nil && !IsNil(o.CheckedChargeNotification) {
		return true
	}

	return false
}

// SetCheckedChargeNotification gets a reference to the given GetAdminConsoleConfig200ResponseCheckedChargeNotification and assigns it to the CheckedChargeNotification field.
func (o *UpdateAdminConsoleConfigRequest) SetCheckedChargeNotification(v GetAdminConsoleConfig200ResponseCheckedChargeNotification) {
	o.CheckedChargeNotification = &v
}

func (o UpdateAdminConsoleConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAdminConsoleConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignInExperienceCustomized) {
		toSerialize["signInExperienceCustomized"] = o.SignInExperienceCustomized
	}
	if !IsNil(o.OrganizationCreated) {
		toSerialize["organizationCreated"] = o.OrganizationCreated
	}
	if !IsNil(o.DevelopmentTenantMigrationNotification) {
		toSerialize["developmentTenantMigrationNotification"] = o.DevelopmentTenantMigrationNotification
	}
	if !IsNil(o.CheckedChargeNotification) {
		toSerialize["checkedChargeNotification"] = o.CheckedChargeNotification
	}
	return toSerialize, nil
}

type NullableUpdateAdminConsoleConfigRequest struct {
	value *UpdateAdminConsoleConfigRequest
	isSet bool
}

func (v NullableUpdateAdminConsoleConfigRequest) Get() *UpdateAdminConsoleConfigRequest {
	return v.value
}

func (v *NullableUpdateAdminConsoleConfigRequest) Set(val *UpdateAdminConsoleConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAdminConsoleConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAdminConsoleConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAdminConsoleConfigRequest(val *UpdateAdminConsoleConfigRequest) *NullableUpdateAdminConsoleConfigRequest {
	return &NullableUpdateAdminConsoleConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateAdminConsoleConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAdminConsoleConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


