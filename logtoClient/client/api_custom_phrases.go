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


// CustomPhrasesAPIService CustomPhrasesAPI service
type CustomPhrasesAPIService service

type ApiDeleteCustomPhraseRequest struct {
	ctx context.Context
	ApiService *CustomPhrasesAPIService
	languageTag string
}

func (r ApiDeleteCustomPhraseRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomPhraseExecute(r)
}

/*
DeleteCustomPhrase Delete custom phrase

Delete custom phrases for the specified language tag.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param languageTag
 @return ApiDeleteCustomPhraseRequest
*/
func (a *CustomPhrasesAPIService) DeleteCustomPhrase(ctx context.Context, languageTag string) ApiDeleteCustomPhraseRequest {
	return ApiDeleteCustomPhraseRequest{
		ApiService: a,
		ctx: ctx,
		languageTag: languageTag,
	}
}

// Execute executes the request
func (a *CustomPhrasesAPIService) DeleteCustomPhraseExecute(r ApiDeleteCustomPhraseRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPhrasesAPIService.DeleteCustomPhrase")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/custom-phrases/{languageTag}"
	localVarPath = strings.Replace(localVarPath, "{"+"languageTag"+"}", url.PathEscape(parameterValueToString(r.languageTag, "languageTag")), -1)

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

type ApiGetCustomPhraseRequest struct {
	ctx context.Context
	ApiService *CustomPhrasesAPIService
	languageTag string
}

func (r ApiGetCustomPhraseRequest) Execute() (*ListCustomPhrases200ResponseInner, *http.Response, error) {
	return r.ApiService.GetCustomPhraseExecute(r)
}

/*
GetCustomPhrase Get custom phrases

Get custom phrases for the specified language tag.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param languageTag
 @return ApiGetCustomPhraseRequest
*/
func (a *CustomPhrasesAPIService) GetCustomPhrase(ctx context.Context, languageTag string) ApiGetCustomPhraseRequest {
	return ApiGetCustomPhraseRequest{
		ApiService: a,
		ctx: ctx,
		languageTag: languageTag,
	}
}

// Execute executes the request
//  @return ListCustomPhrases200ResponseInner
func (a *CustomPhrasesAPIService) GetCustomPhraseExecute(r ApiGetCustomPhraseRequest) (*ListCustomPhrases200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCustomPhrases200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPhrasesAPIService.GetCustomPhrase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/custom-phrases/{languageTag}"
	localVarPath = strings.Replace(localVarPath, "{"+"languageTag"+"}", url.PathEscape(parameterValueToString(r.languageTag, "languageTag")), -1)

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

type ApiListCustomPhrasesRequest struct {
	ctx context.Context
	ApiService *CustomPhrasesAPIService
}

func (r ApiListCustomPhrasesRequest) Execute() ([]ListCustomPhrases200ResponseInner, *http.Response, error) {
	return r.ApiService.ListCustomPhrasesExecute(r)
}

/*
ListCustomPhrases Get all custom phrases

Get all custom phrases for all languages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCustomPhrasesRequest
*/
func (a *CustomPhrasesAPIService) ListCustomPhrases(ctx context.Context) ApiListCustomPhrasesRequest {
	return ApiListCustomPhrasesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ListCustomPhrases200ResponseInner
func (a *CustomPhrasesAPIService) ListCustomPhrasesExecute(r ApiListCustomPhrasesRequest) ([]ListCustomPhrases200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListCustomPhrases200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPhrasesAPIService.ListCustomPhrases")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/custom-phrases"

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

type ApiReplaceCustomPhraseRequest struct {
	ctx context.Context
	ApiService *CustomPhrasesAPIService
	languageTag string
	translationObject *TranslationObject
}

func (r ApiReplaceCustomPhraseRequest) TranslationObject(translationObject TranslationObject) ApiReplaceCustomPhraseRequest {
	r.translationObject = &translationObject
	return r
}

func (r ApiReplaceCustomPhraseRequest) Execute() (*ListCustomPhrases200ResponseInner, *http.Response, error) {
	return r.ApiService.ReplaceCustomPhraseExecute(r)
}

/*
ReplaceCustomPhrase Upsert custom phrases

Upsert custom phrases for the specified language tag. Upsert means that if the custom phrases already exist, they will be updated. Otherwise, they will be created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param languageTag
 @return ApiReplaceCustomPhraseRequest
*/
func (a *CustomPhrasesAPIService) ReplaceCustomPhrase(ctx context.Context, languageTag string) ApiReplaceCustomPhraseRequest {
	return ApiReplaceCustomPhraseRequest{
		ApiService: a,
		ctx: ctx,
		languageTag: languageTag,
	}
}

// Execute executes the request
//  @return ListCustomPhrases200ResponseInner
func (a *CustomPhrasesAPIService) ReplaceCustomPhraseExecute(r ApiReplaceCustomPhraseRequest) (*ListCustomPhrases200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCustomPhrases200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPhrasesAPIService.ReplaceCustomPhrase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/custom-phrases/{languageTag}"
	localVarPath = strings.Replace(localVarPath, "{"+"languageTag"+"}", url.PathEscape(parameterValueToString(r.languageTag, "languageTag")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.translationObject == nil {
		return localVarReturnValue, nil, reportError("translationObject is required and must be specified")
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
	localVarPostBody = r.translationObject
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
