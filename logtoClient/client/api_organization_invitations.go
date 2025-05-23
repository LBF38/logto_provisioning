/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// OrganizationInvitationsAPIService OrganizationInvitationsAPI service
type OrganizationInvitationsAPIService service

type ApiCreateOrganizationInvitationRequest struct {
	ctx context.Context
	ApiService *OrganizationInvitationsAPIService
	createOrganizationInvitationRequest *CreateOrganizationInvitationRequest
}

// The organization invitation to create.
func (r ApiCreateOrganizationInvitationRequest) CreateOrganizationInvitationRequest(createOrganizationInvitationRequest CreateOrganizationInvitationRequest) ApiCreateOrganizationInvitationRequest {
	r.createOrganizationInvitationRequest = &createOrganizationInvitationRequest
	return r
}

func (r ApiCreateOrganizationInvitationRequest) Execute() (*GetOrganizationInvitation200Response, *http.Response, error) {
	return r.ApiService.CreateOrganizationInvitationExecute(r)
}

/*
CreateOrganizationInvitation Create organization invitation

Create an organization invitation and optionally send it via email. The tenant should have an email connector configured if you want to send the invitation via email at this point.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrganizationInvitationRequest
*/
func (a *OrganizationInvitationsAPIService) CreateOrganizationInvitation(ctx context.Context) ApiCreateOrganizationInvitationRequest {
	return ApiCreateOrganizationInvitationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetOrganizationInvitation200Response
func (a *OrganizationInvitationsAPIService) CreateOrganizationInvitationExecute(r ApiCreateOrganizationInvitationRequest) (*GetOrganizationInvitation200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationInvitation200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationInvitationsAPIService.CreateOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/organization-invitations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInvitationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createOrganizationInvitationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateOrganizationInvitationMessageRequest struct {
	ctx context.Context
	ApiService *OrganizationInvitationsAPIService
	id string
	createOrganizationInvitationRequestMessagePayloadOneOf *CreateOrganizationInvitationRequestMessagePayloadOneOf
}

// The message payload for the \&quot;OrganizationInvitation\&quot; template to use when sending the invitation via email.
func (r ApiCreateOrganizationInvitationMessageRequest) CreateOrganizationInvitationRequestMessagePayloadOneOf(createOrganizationInvitationRequestMessagePayloadOneOf CreateOrganizationInvitationRequestMessagePayloadOneOf) ApiCreateOrganizationInvitationMessageRequest {
	r.createOrganizationInvitationRequestMessagePayloadOneOf = &createOrganizationInvitationRequestMessagePayloadOneOf
	return r
}

func (r ApiCreateOrganizationInvitationMessageRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateOrganizationInvitationMessageExecute(r)
}

/*
CreateOrganizationInvitationMessage Resend invitation message

Resend the invitation message to the invitee.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the organization invitation.
 @return ApiCreateOrganizationInvitationMessageRequest
*/
func (a *OrganizationInvitationsAPIService) CreateOrganizationInvitationMessage(ctx context.Context, id string) ApiCreateOrganizationInvitationMessageRequest {
	return ApiCreateOrganizationInvitationMessageRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *OrganizationInvitationsAPIService) CreateOrganizationInvitationMessageExecute(r ApiCreateOrganizationInvitationMessageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationInvitationsAPIService.CreateOrganizationInvitationMessage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/organization-invitations/{id}/message"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInvitationRequestMessagePayloadOneOf == nil {
		return nil, reportError("createOrganizationInvitationRequestMessagePayloadOneOf is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createOrganizationInvitationRequestMessagePayloadOneOf
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteOrganizationInvitationRequest struct {
	ctx context.Context
	ApiService *OrganizationInvitationsAPIService
	id string
}

func (r ApiDeleteOrganizationInvitationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationInvitationExecute(r)
}

/*
DeleteOrganizationInvitation Delete organization invitation

Delete an organization invitation by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the organization invitation.
 @return ApiDeleteOrganizationInvitationRequest
*/
func (a *OrganizationInvitationsAPIService) DeleteOrganizationInvitation(ctx context.Context, id string) ApiDeleteOrganizationInvitationRequest {
	return ApiDeleteOrganizationInvitationRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *OrganizationInvitationsAPIService) DeleteOrganizationInvitationExecute(r ApiDeleteOrganizationInvitationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationInvitationsAPIService.DeleteOrganizationInvitation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/organization-invitations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetOrganizationInvitationRequest struct {
	ctx context.Context
	ApiService *OrganizationInvitationsAPIService
	id string
}

func (r ApiGetOrganizationInvitationRequest) Execute() (*GetOrganizationInvitation200Response, *http.Response, error) {
	return r.ApiService.GetOrganizationInvitationExecute(r)
}

/*
GetOrganizationInvitation Get organization invitation

Get an organization invitation by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the organization invitation.
 @return ApiGetOrganizationInvitationRequest
*/
func (a *OrganizationInvitationsAPIService) GetOrganizationInvitation(ctx context.Context, id string) ApiGetOrganizationInvitationRequest {
	return ApiGetOrganizationInvitationRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetOrganizationInvitation200Response
func (a *OrganizationInvitationsAPIService) GetOrganizationInvitationExecute(r ApiGetOrganizationInvitationRequest) (*GetOrganizationInvitation200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationInvitation200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationInvitationsAPIService.GetOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/organization-invitations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrganizationInvitationsRequest struct {
	ctx context.Context
	ApiService *OrganizationInvitationsAPIService
	organizationId *string
	inviterId *string
	invitee *string
}

func (r ApiListOrganizationInvitationsRequest) OrganizationId(organizationId string) ApiListOrganizationInvitationsRequest {
	r.organizationId = &organizationId
	return r
}

func (r ApiListOrganizationInvitationsRequest) InviterId(inviterId string) ApiListOrganizationInvitationsRequest {
	r.inviterId = &inviterId
	return r
}

func (r ApiListOrganizationInvitationsRequest) Invitee(invitee string) ApiListOrganizationInvitationsRequest {
	r.invitee = &invitee
	return r
}

func (r ApiListOrganizationInvitationsRequest) Execute() ([]GetOrganizationInvitation200Response, *http.Response, error) {
	return r.ApiService.ListOrganizationInvitationsExecute(r)
}

/*
ListOrganizationInvitations Get organization invitations

Get organization invitations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOrganizationInvitationsRequest
*/
func (a *OrganizationInvitationsAPIService) ListOrganizationInvitations(ctx context.Context) ApiListOrganizationInvitationsRequest {
	return ApiListOrganizationInvitationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GetOrganizationInvitation200Response
func (a *OrganizationInvitationsAPIService) ListOrganizationInvitationsExecute(r ApiListOrganizationInvitationsRequest) ([]GetOrganizationInvitation200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationInvitation200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationInvitationsAPIService.ListOrganizationInvitations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/organization-invitations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.organizationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "organizationId", r.organizationId, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.invitee != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitee", r.invitee, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReplaceOrganizationInvitationStatusRequest struct {
	ctx context.Context
	ApiService *OrganizationInvitationsAPIService
	id string
	replaceOrganizationInvitationStatusRequest *ReplaceOrganizationInvitationStatusRequest
}

// The organization invitation status to update.
func (r ApiReplaceOrganizationInvitationStatusRequest) ReplaceOrganizationInvitationStatusRequest(replaceOrganizationInvitationStatusRequest ReplaceOrganizationInvitationStatusRequest) ApiReplaceOrganizationInvitationStatusRequest {
	r.replaceOrganizationInvitationStatusRequest = &replaceOrganizationInvitationStatusRequest
	return r
}

func (r ApiReplaceOrganizationInvitationStatusRequest) Execute() (*GetOrganizationInvitation200Response, *http.Response, error) {
	return r.ApiService.ReplaceOrganizationInvitationStatusExecute(r)
}

/*
ReplaceOrganizationInvitationStatus Update organization invitation status

Update the status of an organization invitation by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the organization invitation.
 @return ApiReplaceOrganizationInvitationStatusRequest
*/
func (a *OrganizationInvitationsAPIService) ReplaceOrganizationInvitationStatus(ctx context.Context, id string) ApiReplaceOrganizationInvitationStatusRequest {
	return ApiReplaceOrganizationInvitationStatusRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetOrganizationInvitation200Response
func (a *OrganizationInvitationsAPIService) ReplaceOrganizationInvitationStatusExecute(r ApiReplaceOrganizationInvitationStatusRequest) (*GetOrganizationInvitation200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationInvitation200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationInvitationsAPIService.ReplaceOrganizationInvitationStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/organization-invitations/{id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.replaceOrganizationInvitationStatusRequest == nil {
		return localVarReturnValue, nil, reportError("replaceOrganizationInvitationStatusRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.replaceOrganizationInvitationStatusRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
