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

// checks if the UpdateSignInExpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSignInExpRequest{}

// UpdateSignInExpRequest struct for UpdateSignInExpRequest
type UpdateSignInExpRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Color *UpdateSignInExpRequestColor `json:"color,omitempty"`
	Branding *ListApplicationOrganizations200ResponseInnerBranding `json:"branding,omitempty"`
	LanguageInfo *UpdateSignInExpRequestLanguageInfo `json:"languageInfo,omitempty"`
	AgreeToTermsPolicy *string `json:"agreeToTermsPolicy,omitempty"`
	SignIn *UpdateSignInExpRequestSignIn `json:"signIn,omitempty"`
	SignUp *UpdateSignInExpRequestSignUp `json:"signUp,omitempty"`
	SocialSignIn *GetSignInExp200ResponseSocialSignIn `json:"socialSignIn,omitempty"`
	// Specify the social sign-in connectors to display on the sign-in page.
	SocialSignInConnectorTargets []string `json:"socialSignInConnectorTargets,omitempty"`
	SignInMode *string `json:"signInMode,omitempty"`
	CustomCss NullableString `json:"customCss,omitempty"`
	// Custom content to display on experience flow pages. the page pathname will be the config key, the content will be the config value.
	CustomContent *map[string]string `json:"customContent,omitempty"`
	CustomUiAssets NullableGetSignInExp200ResponseCustomUiAssets `json:"customUiAssets,omitempty"`
	PasswordPolicy *GetSignInExp200ResponsePasswordPolicy `json:"passwordPolicy,omitempty"`
	Mfa *GetSignInExp200ResponseMfa `json:"mfa,omitempty"`
	SingleSignOnEnabled *bool `json:"singleSignOnEnabled,omitempty"`
	TermsOfUseUrl *UpdateSignInExpRequestTermsOfUseUrl `json:"termsOfUseUrl,omitempty"`
	PrivacyPolicyUrl *UpdateSignInExpRequestTermsOfUseUrl `json:"privacyPolicyUrl,omitempty"`
}

// NewUpdateSignInExpRequest instantiates a new UpdateSignInExpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSignInExpRequest() *UpdateSignInExpRequest {
	this := UpdateSignInExpRequest{}
	return &this
}

// NewUpdateSignInExpRequestWithDefaults instantiates a new UpdateSignInExpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSignInExpRequestWithDefaults() *UpdateSignInExpRequest {
	this := UpdateSignInExpRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *UpdateSignInExpRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetColor() UpdateSignInExpRequestColor {
	if o == nil || IsNil(o.Color) {
		var ret UpdateSignInExpRequestColor
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetColorOk() (*UpdateSignInExpRequestColor, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given UpdateSignInExpRequestColor and assigns it to the Color field.
func (o *UpdateSignInExpRequest) SetColor(v UpdateSignInExpRequestColor) {
	o.Color = &v
}

// GetBranding returns the Branding field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetBranding() ListApplicationOrganizations200ResponseInnerBranding {
	if o == nil || IsNil(o.Branding) {
		var ret ListApplicationOrganizations200ResponseInnerBranding
		return ret
	}
	return *o.Branding
}

// GetBrandingOk returns a tuple with the Branding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool) {
	if o == nil || IsNil(o.Branding) {
		return nil, false
	}
	return o.Branding, true
}

// HasBranding returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasBranding() bool {
	if o != nil && !IsNil(o.Branding) {
		return true
	}

	return false
}

// SetBranding gets a reference to the given ListApplicationOrganizations200ResponseInnerBranding and assigns it to the Branding field.
func (o *UpdateSignInExpRequest) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding) {
	o.Branding = &v
}

// GetLanguageInfo returns the LanguageInfo field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetLanguageInfo() UpdateSignInExpRequestLanguageInfo {
	if o == nil || IsNil(o.LanguageInfo) {
		var ret UpdateSignInExpRequestLanguageInfo
		return ret
	}
	return *o.LanguageInfo
}

// GetLanguageInfoOk returns a tuple with the LanguageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetLanguageInfoOk() (*UpdateSignInExpRequestLanguageInfo, bool) {
	if o == nil || IsNil(o.LanguageInfo) {
		return nil, false
	}
	return o.LanguageInfo, true
}

// HasLanguageInfo returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasLanguageInfo() bool {
	if o != nil && !IsNil(o.LanguageInfo) {
		return true
	}

	return false
}

// SetLanguageInfo gets a reference to the given UpdateSignInExpRequestLanguageInfo and assigns it to the LanguageInfo field.
func (o *UpdateSignInExpRequest) SetLanguageInfo(v UpdateSignInExpRequestLanguageInfo) {
	o.LanguageInfo = &v
}

// GetAgreeToTermsPolicy returns the AgreeToTermsPolicy field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetAgreeToTermsPolicy() string {
	if o == nil || IsNil(o.AgreeToTermsPolicy) {
		var ret string
		return ret
	}
	return *o.AgreeToTermsPolicy
}

// GetAgreeToTermsPolicyOk returns a tuple with the AgreeToTermsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetAgreeToTermsPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.AgreeToTermsPolicy) {
		return nil, false
	}
	return o.AgreeToTermsPolicy, true
}

// HasAgreeToTermsPolicy returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasAgreeToTermsPolicy() bool {
	if o != nil && !IsNil(o.AgreeToTermsPolicy) {
		return true
	}

	return false
}

// SetAgreeToTermsPolicy gets a reference to the given string and assigns it to the AgreeToTermsPolicy field.
func (o *UpdateSignInExpRequest) SetAgreeToTermsPolicy(v string) {
	o.AgreeToTermsPolicy = &v
}

// GetSignIn returns the SignIn field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetSignIn() UpdateSignInExpRequestSignIn {
	if o == nil || IsNil(o.SignIn) {
		var ret UpdateSignInExpRequestSignIn
		return ret
	}
	return *o.SignIn
}

// GetSignInOk returns a tuple with the SignIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetSignInOk() (*UpdateSignInExpRequestSignIn, bool) {
	if o == nil || IsNil(o.SignIn) {
		return nil, false
	}
	return o.SignIn, true
}

// HasSignIn returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasSignIn() bool {
	if o != nil && !IsNil(o.SignIn) {
		return true
	}

	return false
}

// SetSignIn gets a reference to the given UpdateSignInExpRequestSignIn and assigns it to the SignIn field.
func (o *UpdateSignInExpRequest) SetSignIn(v UpdateSignInExpRequestSignIn) {
	o.SignIn = &v
}

// GetSignUp returns the SignUp field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetSignUp() UpdateSignInExpRequestSignUp {
	if o == nil || IsNil(o.SignUp) {
		var ret UpdateSignInExpRequestSignUp
		return ret
	}
	return *o.SignUp
}

// GetSignUpOk returns a tuple with the SignUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetSignUpOk() (*UpdateSignInExpRequestSignUp, bool) {
	if o == nil || IsNil(o.SignUp) {
		return nil, false
	}
	return o.SignUp, true
}

// HasSignUp returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasSignUp() bool {
	if o != nil && !IsNil(o.SignUp) {
		return true
	}

	return false
}

// SetSignUp gets a reference to the given UpdateSignInExpRequestSignUp and assigns it to the SignUp field.
func (o *UpdateSignInExpRequest) SetSignUp(v UpdateSignInExpRequestSignUp) {
	o.SignUp = &v
}

// GetSocialSignIn returns the SocialSignIn field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetSocialSignIn() GetSignInExp200ResponseSocialSignIn {
	if o == nil || IsNil(o.SocialSignIn) {
		var ret GetSignInExp200ResponseSocialSignIn
		return ret
	}
	return *o.SocialSignIn
}

// GetSocialSignInOk returns a tuple with the SocialSignIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetSocialSignInOk() (*GetSignInExp200ResponseSocialSignIn, bool) {
	if o == nil || IsNil(o.SocialSignIn) {
		return nil, false
	}
	return o.SocialSignIn, true
}

// HasSocialSignIn returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasSocialSignIn() bool {
	if o != nil && !IsNil(o.SocialSignIn) {
		return true
	}

	return false
}

// SetSocialSignIn gets a reference to the given GetSignInExp200ResponseSocialSignIn and assigns it to the SocialSignIn field.
func (o *UpdateSignInExpRequest) SetSocialSignIn(v GetSignInExp200ResponseSocialSignIn) {
	o.SocialSignIn = &v
}

// GetSocialSignInConnectorTargets returns the SocialSignInConnectorTargets field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetSocialSignInConnectorTargets() []string {
	if o == nil || IsNil(o.SocialSignInConnectorTargets) {
		var ret []string
		return ret
	}
	return o.SocialSignInConnectorTargets
}

// GetSocialSignInConnectorTargetsOk returns a tuple with the SocialSignInConnectorTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetSocialSignInConnectorTargetsOk() ([]string, bool) {
	if o == nil || IsNil(o.SocialSignInConnectorTargets) {
		return nil, false
	}
	return o.SocialSignInConnectorTargets, true
}

// HasSocialSignInConnectorTargets returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasSocialSignInConnectorTargets() bool {
	if o != nil && !IsNil(o.SocialSignInConnectorTargets) {
		return true
	}

	return false
}

// SetSocialSignInConnectorTargets gets a reference to the given []string and assigns it to the SocialSignInConnectorTargets field.
func (o *UpdateSignInExpRequest) SetSocialSignInConnectorTargets(v []string) {
	o.SocialSignInConnectorTargets = v
}

// GetSignInMode returns the SignInMode field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetSignInMode() string {
	if o == nil || IsNil(o.SignInMode) {
		var ret string
		return ret
	}
	return *o.SignInMode
}

// GetSignInModeOk returns a tuple with the SignInMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetSignInModeOk() (*string, bool) {
	if o == nil || IsNil(o.SignInMode) {
		return nil, false
	}
	return o.SignInMode, true
}

// HasSignInMode returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasSignInMode() bool {
	if o != nil && !IsNil(o.SignInMode) {
		return true
	}

	return false
}

// SetSignInMode gets a reference to the given string and assigns it to the SignInMode field.
func (o *UpdateSignInExpRequest) SetSignInMode(v string) {
	o.SignInMode = &v
}

// GetCustomCss returns the CustomCss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSignInExpRequest) GetCustomCss() string {
	if o == nil || IsNil(o.CustomCss.Get()) {
		var ret string
		return ret
	}
	return *o.CustomCss.Get()
}

// GetCustomCssOk returns a tuple with the CustomCss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSignInExpRequest) GetCustomCssOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomCss.Get(), o.CustomCss.IsSet()
}

// HasCustomCss returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasCustomCss() bool {
	if o != nil && o.CustomCss.IsSet() {
		return true
	}

	return false
}

// SetCustomCss gets a reference to the given NullableString and assigns it to the CustomCss field.
func (o *UpdateSignInExpRequest) SetCustomCss(v string) {
	o.CustomCss.Set(&v)
}
// SetCustomCssNil sets the value for CustomCss to be an explicit nil
func (o *UpdateSignInExpRequest) SetCustomCssNil() {
	o.CustomCss.Set(nil)
}

// UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
func (o *UpdateSignInExpRequest) UnsetCustomCss() {
	o.CustomCss.Unset()
}

// GetCustomContent returns the CustomContent field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetCustomContent() map[string]string {
	if o == nil || IsNil(o.CustomContent) {
		var ret map[string]string
		return ret
	}
	return *o.CustomContent
}

// GetCustomContentOk returns a tuple with the CustomContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetCustomContentOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomContent) {
		return nil, false
	}
	return o.CustomContent, true
}

// HasCustomContent returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasCustomContent() bool {
	if o != nil && !IsNil(o.CustomContent) {
		return true
	}

	return false
}

// SetCustomContent gets a reference to the given map[string]string and assigns it to the CustomContent field.
func (o *UpdateSignInExpRequest) SetCustomContent(v map[string]string) {
	o.CustomContent = &v
}

// GetCustomUiAssets returns the CustomUiAssets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSignInExpRequest) GetCustomUiAssets() GetSignInExp200ResponseCustomUiAssets {
	if o == nil || IsNil(o.CustomUiAssets.Get()) {
		var ret GetSignInExp200ResponseCustomUiAssets
		return ret
	}
	return *o.CustomUiAssets.Get()
}

// GetCustomUiAssetsOk returns a tuple with the CustomUiAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSignInExpRequest) GetCustomUiAssetsOk() (*GetSignInExp200ResponseCustomUiAssets, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomUiAssets.Get(), o.CustomUiAssets.IsSet()
}

// HasCustomUiAssets returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasCustomUiAssets() bool {
	if o != nil && o.CustomUiAssets.IsSet() {
		return true
	}

	return false
}

// SetCustomUiAssets gets a reference to the given NullableGetSignInExp200ResponseCustomUiAssets and assigns it to the CustomUiAssets field.
func (o *UpdateSignInExpRequest) SetCustomUiAssets(v GetSignInExp200ResponseCustomUiAssets) {
	o.CustomUiAssets.Set(&v)
}
// SetCustomUiAssetsNil sets the value for CustomUiAssets to be an explicit nil
func (o *UpdateSignInExpRequest) SetCustomUiAssetsNil() {
	o.CustomUiAssets.Set(nil)
}

// UnsetCustomUiAssets ensures that no value is present for CustomUiAssets, not even an explicit nil
func (o *UpdateSignInExpRequest) UnsetCustomUiAssets() {
	o.CustomUiAssets.Unset()
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetPasswordPolicy() GetSignInExp200ResponsePasswordPolicy {
	if o == nil || IsNil(o.PasswordPolicy) {
		var ret GetSignInExp200ResponsePasswordPolicy
		return ret
	}
	return *o.PasswordPolicy
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetPasswordPolicyOk() (*GetSignInExp200ResponsePasswordPolicy, bool) {
	if o == nil || IsNil(o.PasswordPolicy) {
		return nil, false
	}
	return o.PasswordPolicy, true
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasPasswordPolicy() bool {
	if o != nil && !IsNil(o.PasswordPolicy) {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given GetSignInExp200ResponsePasswordPolicy and assigns it to the PasswordPolicy field.
func (o *UpdateSignInExpRequest) SetPasswordPolicy(v GetSignInExp200ResponsePasswordPolicy) {
	o.PasswordPolicy = &v
}

// GetMfa returns the Mfa field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetMfa() GetSignInExp200ResponseMfa {
	if o == nil || IsNil(o.Mfa) {
		var ret GetSignInExp200ResponseMfa
		return ret
	}
	return *o.Mfa
}

// GetMfaOk returns a tuple with the Mfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetMfaOk() (*GetSignInExp200ResponseMfa, bool) {
	if o == nil || IsNil(o.Mfa) {
		return nil, false
	}
	return o.Mfa, true
}

// HasMfa returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasMfa() bool {
	if o != nil && !IsNil(o.Mfa) {
		return true
	}

	return false
}

// SetMfa gets a reference to the given GetSignInExp200ResponseMfa and assigns it to the Mfa field.
func (o *UpdateSignInExpRequest) SetMfa(v GetSignInExp200ResponseMfa) {
	o.Mfa = &v
}

// GetSingleSignOnEnabled returns the SingleSignOnEnabled field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetSingleSignOnEnabled() bool {
	if o == nil || IsNil(o.SingleSignOnEnabled) {
		var ret bool
		return ret
	}
	return *o.SingleSignOnEnabled
}

// GetSingleSignOnEnabledOk returns a tuple with the SingleSignOnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetSingleSignOnEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SingleSignOnEnabled) {
		return nil, false
	}
	return o.SingleSignOnEnabled, true
}

// HasSingleSignOnEnabled returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasSingleSignOnEnabled() bool {
	if o != nil && !IsNil(o.SingleSignOnEnabled) {
		return true
	}

	return false
}

// SetSingleSignOnEnabled gets a reference to the given bool and assigns it to the SingleSignOnEnabled field.
func (o *UpdateSignInExpRequest) SetSingleSignOnEnabled(v bool) {
	o.SingleSignOnEnabled = &v
}

// GetTermsOfUseUrl returns the TermsOfUseUrl field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetTermsOfUseUrl() UpdateSignInExpRequestTermsOfUseUrl {
	if o == nil || IsNil(o.TermsOfUseUrl) {
		var ret UpdateSignInExpRequestTermsOfUseUrl
		return ret
	}
	return *o.TermsOfUseUrl
}

// GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetTermsOfUseUrlOk() (*UpdateSignInExpRequestTermsOfUseUrl, bool) {
	if o == nil || IsNil(o.TermsOfUseUrl) {
		return nil, false
	}
	return o.TermsOfUseUrl, true
}

// HasTermsOfUseUrl returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasTermsOfUseUrl() bool {
	if o != nil && !IsNil(o.TermsOfUseUrl) {
		return true
	}

	return false
}

// SetTermsOfUseUrl gets a reference to the given UpdateSignInExpRequestTermsOfUseUrl and assigns it to the TermsOfUseUrl field.
func (o *UpdateSignInExpRequest) SetTermsOfUseUrl(v UpdateSignInExpRequestTermsOfUseUrl) {
	o.TermsOfUseUrl = &v
}

// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value if set, zero value otherwise.
func (o *UpdateSignInExpRequest) GetPrivacyPolicyUrl() UpdateSignInExpRequestTermsOfUseUrl {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		var ret UpdateSignInExpRequestTermsOfUseUrl
		return ret
	}
	return *o.PrivacyPolicyUrl
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequest) GetPrivacyPolicyUrlOk() (*UpdateSignInExpRequestTermsOfUseUrl, bool) {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		return nil, false
	}
	return o.PrivacyPolicyUrl, true
}

// HasPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *UpdateSignInExpRequest) HasPrivacyPolicyUrl() bool {
	if o != nil && !IsNil(o.PrivacyPolicyUrl) {
		return true
	}

	return false
}

// SetPrivacyPolicyUrl gets a reference to the given UpdateSignInExpRequestTermsOfUseUrl and assigns it to the PrivacyPolicyUrl field.
func (o *UpdateSignInExpRequest) SetPrivacyPolicyUrl(v UpdateSignInExpRequestTermsOfUseUrl) {
	o.PrivacyPolicyUrl = &v
}

func (o UpdateSignInExpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSignInExpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Branding) {
		toSerialize["branding"] = o.Branding
	}
	if !IsNil(o.LanguageInfo) {
		toSerialize["languageInfo"] = o.LanguageInfo
	}
	if !IsNil(o.AgreeToTermsPolicy) {
		toSerialize["agreeToTermsPolicy"] = o.AgreeToTermsPolicy
	}
	if !IsNil(o.SignIn) {
		toSerialize["signIn"] = o.SignIn
	}
	if !IsNil(o.SignUp) {
		toSerialize["signUp"] = o.SignUp
	}
	if !IsNil(o.SocialSignIn) {
		toSerialize["socialSignIn"] = o.SocialSignIn
	}
	if !IsNil(o.SocialSignInConnectorTargets) {
		toSerialize["socialSignInConnectorTargets"] = o.SocialSignInConnectorTargets
	}
	if !IsNil(o.SignInMode) {
		toSerialize["signInMode"] = o.SignInMode
	}
	if o.CustomCss.IsSet() {
		toSerialize["customCss"] = o.CustomCss.Get()
	}
	if !IsNil(o.CustomContent) {
		toSerialize["customContent"] = o.CustomContent
	}
	if o.CustomUiAssets.IsSet() {
		toSerialize["customUiAssets"] = o.CustomUiAssets.Get()
	}
	if !IsNil(o.PasswordPolicy) {
		toSerialize["passwordPolicy"] = o.PasswordPolicy
	}
	if !IsNil(o.Mfa) {
		toSerialize["mfa"] = o.Mfa
	}
	if !IsNil(o.SingleSignOnEnabled) {
		toSerialize["singleSignOnEnabled"] = o.SingleSignOnEnabled
	}
	if !IsNil(o.TermsOfUseUrl) {
		toSerialize["termsOfUseUrl"] = o.TermsOfUseUrl
	}
	if !IsNil(o.PrivacyPolicyUrl) {
		toSerialize["privacyPolicyUrl"] = o.PrivacyPolicyUrl
	}
	return toSerialize, nil
}

type NullableUpdateSignInExpRequest struct {
	value *UpdateSignInExpRequest
	isSet bool
}

func (v NullableUpdateSignInExpRequest) Get() *UpdateSignInExpRequest {
	return v.value
}

func (v *NullableUpdateSignInExpRequest) Set(val *UpdateSignInExpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSignInExpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSignInExpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSignInExpRequest(val *UpdateSignInExpRequest) *NullableUpdateSignInExpRequest {
	return &NullableUpdateSignInExpRequest{value: val, isSet: true}
}

func (v NullableUpdateSignInExpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSignInExpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


