/*
Exasol SaaS REST-API

## Authentication  The REST API can be used with your Personal Access Token (PAT). You don't know what a PAT is, check our documentation  [here](https://docs.exasol.com/saas/administration/access_mngt/access_token.htm).  After you created a PAT click on Authorize and add your PAT under BearerAuth.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// ClustersApiService ClustersApi service
type ClustersApiService service

type ApiCreateClusterRequest struct {
	ctx           context.Context
	ApiService    *ClustersApiService
	accountID     string
	databaseID    string
	createCluster *CreateCluster
}

func (r ApiCreateClusterRequest) CreateCluster(createCluster CreateCluster) ApiCreateClusterRequest {
	r.createCluster = &createCluster
	return r
}

func (r ApiCreateClusterRequest) Execute() (*Cluster, *http.Response, error) {
	return r.ApiService.CreateClusterExecute(r)
}

/*
CreateCluster Create cluster

Create a new cluster

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @return ApiCreateClusterRequest
*/
func (a *ClustersApiService) CreateCluster(ctx context.Context, accountID string, databaseID string) ApiCreateClusterRequest {
	return ApiCreateClusterRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
	}
}

// Execute executes the request
//  @return Cluster
func (a *ClustersApiService) CreateClusterExecute(r ApiCreateClusterRequest) (*Cluster, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Cluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.CreateCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createCluster == nil {
		return localVarReturnValue, nil, reportError("createCluster is required and must be specified")
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
	localVarPostBody = r.createCluster
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiDeleteClusterRequest struct {
	ctx        context.Context
	ApiService *ClustersApiService
	accountID  string
	databaseID string
	clusterID  string
}

func (r ApiDeleteClusterRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClusterExecute(r)
}

/*
DeleteCluster Delete cluster

Delete the cluster, if main cluster is deleted, the whole database will be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @param clusterID Cluster ID
 @return ApiDeleteClusterRequest
*/
func (a *ClustersApiService) DeleteCluster(ctx context.Context, accountID string, databaseID string, clusterID string) ApiDeleteClusterRequest {
	return ApiDeleteClusterRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
		clusterID:  clusterID,
	}
}

// Execute executes the request
func (a *ClustersApiService) DeleteClusterExecute(r ApiDeleteClusterRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.DeleteCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters/{clusterID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", url.PathEscape(parameterToString(r.clusterID, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetClusterRequest struct {
	ctx        context.Context
	ApiService *ClustersApiService
	accountID  string
	databaseID string
	clusterID  string
}

func (r ApiGetClusterRequest) Execute() (*Cluster, *http.Response, error) {
	return r.ApiService.GetClusterExecute(r)
}

/*
GetCluster Get cluster

Get cluster

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @param clusterID Cluster ID
 @return ApiGetClusterRequest
*/
func (a *ClustersApiService) GetCluster(ctx context.Context, accountID string, databaseID string, clusterID string) ApiGetClusterRequest {
	return ApiGetClusterRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
		clusterID:  clusterID,
	}
}

// Execute executes the request
//  @return Cluster
func (a *ClustersApiService) GetClusterExecute(r ApiGetClusterRequest) (*Cluster, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Cluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters/{clusterID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", url.PathEscape(parameterToString(r.clusterID, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiGetClusterConnectionRequest struct {
	ctx        context.Context
	ApiService *ClustersApiService
	accountID  string
	databaseID string
	clusterID  string
}

func (r ApiGetClusterConnectionRequest) Execute() (*Connections, *http.Response, error) {
	return r.ApiService.GetClusterConnectionExecute(r)
}

/*
GetClusterConnection Get connection information

Get connection information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @param clusterID Cluster ID
 @return ApiGetClusterConnectionRequest
*/
func (a *ClustersApiService) GetClusterConnection(ctx context.Context, accountID string, databaseID string, clusterID string) ApiGetClusterConnectionRequest {
	return ApiGetClusterConnectionRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
		clusterID:  clusterID,
	}
}

// Execute executes the request
//  @return Connections
func (a *ClustersApiService) GetClusterConnectionExecute(r ApiGetClusterConnectionRequest) (*Connections, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Connections
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetClusterConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters/{clusterID}/connect"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", url.PathEscape(parameterToString(r.clusterID, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiListClustersRequest struct {
	ctx        context.Context
	ApiService *ClustersApiService
	accountID  string
	databaseID string
}

func (r ApiListClustersRequest) Execute() ([]Cluster, *http.Response, error) {
	return r.ApiService.ListClustersExecute(r)
}

/*
ListClusters List clusters

List clusters from a database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @return ApiListClustersRequest
*/
func (a *ClustersApiService) ListClusters(ctx context.Context, accountID string, databaseID string) ApiListClustersRequest {
	return ApiListClustersRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
	}
}

// Execute executes the request
//  @return []Cluster
func (a *ClustersApiService) ListClustersExecute(r ApiListClustersRequest) ([]Cluster, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Cluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.ListClusters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiScaleClusterRequest struct {
	ctx          context.Context
	ApiService   *ClustersApiService
	accountID    string
	databaseID   string
	clusterID    string
	scaleCluster *ScaleCluster
}

func (r ApiScaleClusterRequest) ScaleCluster(scaleCluster ScaleCluster) ApiScaleClusterRequest {
	r.scaleCluster = &scaleCluster
	return r
}

func (r ApiScaleClusterRequest) Execute() (*http.Response, error) {
	return r.ApiService.ScaleClusterExecute(r)
}

/*
ScaleCluster Scale cluster

Scale cluster

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @param clusterID Cluster ID
 @return ApiScaleClusterRequest
*/
func (a *ClustersApiService) ScaleCluster(ctx context.Context, accountID string, databaseID string, clusterID string) ApiScaleClusterRequest {
	return ApiScaleClusterRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
		clusterID:  clusterID,
	}
}

// Execute executes the request
func (a *ClustersApiService) ScaleClusterExecute(r ApiScaleClusterRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.ScaleCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters/{clusterID}/scale"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", url.PathEscape(parameterToString(r.clusterID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.scaleCluster == nil {
		return nil, reportError("scaleCluster is required and must be specified")
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
	localVarPostBody = r.scaleCluster
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiStartClusterRequest struct {
	ctx        context.Context
	ApiService *ClustersApiService
	accountID  string
	databaseID string
	clusterID  string
}

func (r ApiStartClusterRequest) Execute() (*http.Response, error) {
	return r.ApiService.StartClusterExecute(r)
}

/*
StartCluster Start cluster

Start cluster

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @param clusterID Cluster ID
 @return ApiStartClusterRequest
*/
func (a *ClustersApiService) StartCluster(ctx context.Context, accountID string, databaseID string, clusterID string) ApiStartClusterRequest {
	return ApiStartClusterRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
		clusterID:  clusterID,
	}
}

// Execute executes the request
func (a *ClustersApiService) StartClusterExecute(r ApiStartClusterRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.StartCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters/{clusterID}/start"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", url.PathEscape(parameterToString(r.clusterID, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiStopClusterRequest struct {
	ctx        context.Context
	ApiService *ClustersApiService
	accountID  string
	databaseID string
	clusterID  string
}

func (r ApiStopClusterRequest) Execute() (*http.Response, error) {
	return r.ApiService.StopClusterExecute(r)
}

/*
StopCluster Stop cluster

Stop cluster

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @param clusterID Cluster ID
 @return ApiStopClusterRequest
*/
func (a *ClustersApiService) StopCluster(ctx context.Context, accountID string, databaseID string, clusterID string) ApiStopClusterRequest {
	return ApiStopClusterRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
		clusterID:  clusterID,
	}
}

// Execute executes the request
func (a *ClustersApiService) StopClusterExecute(r ApiStopClusterRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.StopCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters/{clusterID}/stop"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", url.PathEscape(parameterToString(r.clusterID, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateClusterRequest struct {
	ctx           context.Context
	ApiService    *ClustersApiService
	accountID     string
	databaseID    string
	clusterID     string
	updateCluster *UpdateCluster
}

func (r ApiUpdateClusterRequest) UpdateCluster(updateCluster UpdateCluster) ApiUpdateClusterRequest {
	r.updateCluster = &updateCluster
	return r
}

func (r ApiUpdateClusterRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateClusterExecute(r)
}

/*
UpdateCluster Update cluster

Update the cluster with the specified ID.
 Only metadata can be changed. To scale the cluster size, use the scale cluster endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID
 @param databaseID Database ID
 @param clusterID Cluster ID
 @return ApiUpdateClusterRequest
*/
func (a *ClustersApiService) UpdateCluster(ctx context.Context, accountID string, databaseID string, clusterID string) ApiUpdateClusterRequest {
	return ApiUpdateClusterRequest{
		ApiService: a,
		ctx:        ctx,
		accountID:  accountID,
		databaseID: databaseID,
		clusterID:  clusterID,
	}
}

// Execute executes the request
func (a *ClustersApiService) UpdateClusterExecute(r ApiUpdateClusterRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpdateCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{accountID}/databases/{databaseID}/clusters/{clusterID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterToString(r.accountID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseID"+"}", url.PathEscape(parameterToString(r.databaseID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", url.PathEscape(parameterToString(r.clusterID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateCluster == nil {
		return nil, reportError("updateCluster is required and must be specified")
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
	localVarPostBody = r.updateCluster
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v APIError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
