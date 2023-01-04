# \ChartApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Chart1Get**](ChartApi.md#Chart1Get) | **Get** /{apikey}/chart/{lat}/{lon}/{dec1}/{az1}/{kwp1} | Build chart token for one solar plane
[**Chart2Get**](ChartApi.md#Chart2Get) | **Get** /{apikey}/chart/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | Build chart token combined for two solar planes
[**Chart3Get**](ChartApi.md#Chart3Get) | **Get** /{apikey}/chart/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | Build chart token combined for three solar planes
[**ChartClearskyDayTokenJsGet**](ChartApi.md#ChartClearskyDayTokenJsGet) | **Get** /chart/clearsky/{day}/{token}.js | Script to render theoretic clear sky chart for a day in advance
[**ChartClearskyTokenJsGet**](ChartApi.md#ChartClearskyTokenJsGet) | **Get** /chart/clearsky/{token}.js | Script to render today theoretic clear sky chart
[**ChartDayTokenJsGet**](ChartApi.md#ChartDayTokenJsGet) | **Get** /chart/{day}/{token}.js | Script to render estimate chart for a day in advance
[**ChartHistoryDayTokenJsGet**](ChartApi.md#ChartHistoryDayTokenJsGet) | **Get** /chart/history/{day}/{token}.js | Script to render historic average chart for a day in advance
[**ChartHistoryTokenJsGet**](ChartApi.md#ChartHistoryTokenJsGet) | **Get** /chart/history/{token}.js | Script to render today historic average chart
[**ChartInitJsGet**](ChartApi.md#ChartInitJsGet) | **Get** /chart/init.js | Initialization script for charts
[**ChartTokenJsGet**](ChartApi.md#ChartTokenJsGet) | **Get** /chart/{token}.js | Script to render today estimate chart
[**ChartsGet**](ChartApi.md#ChartsGet) | **Get** /{apikey}/charts | List all charts for given API key


# **Chart1Get**
> ResponseBuildChartToken Chart1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)
Build chart token for one solar plane

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| Your personal API key | 
  **lat** | **float32**| Latitude of location | 
  **lon** | **float32**| Longitude of location | 
  **dec1** | **float32**| 1st solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az1** | **float32**| 1st solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp1** | **float32**| 1st solar plane max. peak power in kilo watt | 
 **optional** | ***ChartApiChart1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartApiChart1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseBuildChartToken**](ResponseBuildChartToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Chart2Get**
> ResponseBuildChartToken Chart2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)
Build chart token combined for two solar planes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| Your personal API key | 
  **lat** | **float32**| Latitude of location | 
  **lon** | **float32**| Longitude of location | 
  **dec1** | **float32**| 1st solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az1** | **float32**| 1st solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp1** | **float32**| 1st solar plane max. peak power in kilo watt | 
  **dec2** | **float32**| 2nd solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az2** | **float32**| 2nd solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp2** | **float32**| 2nd solar pülane max. peak power in kilo watt | 
 **optional** | ***ChartApiChart2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartApiChart2GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------









 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseBuildChartToken**](ResponseBuildChartToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Chart3Get**
> ResponseBuildChartToken Chart3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)
Build chart token combined for three solar planes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| Your personal API key | 
  **lat** | **float32**| Latitude of location | 
  **lon** | **float32**| Longitude of location | 
  **dec1** | **float32**| 1st solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az1** | **float32**| 1st solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp1** | **float32**| 1st solar plane max. peak power in kilo watt | 
  **dec2** | **float32**| 2nd solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az2** | **float32**| 2nd solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp2** | **float32**| 2nd solar pülane max. peak power in kilo watt | 
  **dec3** | **float32**| 3rd solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az3** | **float32**| 3rd solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp3** | **float32**| 3rd solar plane max. peak power in kilo watt peak | 
 **optional** | ***ChartApiChart3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartApiChart3GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------












 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseBuildChartToken**](ResponseBuildChartToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartClearskyDayTokenJsGet**
> ChartClearskyDayTokenJsGet(ctx, day, token)
Script to render theoretic clear sky chart for a day in advance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **day** | **float32**| Show days in advance: 0 &#x3D; today (default), max. 6 | 
  **token** | **string**| Token generated by the /{apikey}/chart/{lat}/{lon}/... functions | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartClearskyTokenJsGet**
> ChartClearskyTokenJsGet(ctx, token)
Script to render today theoretic clear sky chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Token generated by the /{apikey}/chart/{lat}/{lon}/... functions | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartDayTokenJsGet**
> ChartDayTokenJsGet(ctx, day, token)
Script to render estimate chart for a day in advance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **day** | **float32**| Show days in advance: 0 &#x3D; today (default), max. 6 | 
  **token** | **string**| Token generated by the /{apikey}/chart/{lat}/{lon}/... functions | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartHistoryDayTokenJsGet**
> ChartHistoryDayTokenJsGet(ctx, day, token)
Script to render historic average chart for a day in advance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **day** | **float32**| Show days in advance: 0 &#x3D; today (default), max. 6 | 
  **token** | **string**| Token generated by the /{apikey}/chart/{lat}/{lon}/... functions | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartHistoryTokenJsGet**
> ChartHistoryTokenJsGet(ctx, token)
Script to render today historic average chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Token generated by the /{apikey}/chart/{lat}/{lon}/... functions | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartInitJsGet**
> ChartInitJsGet(ctx, )
Initialization script for charts

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartTokenJsGet**
> ChartTokenJsGet(ctx, token)
Script to render today estimate chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Token generated by the /{apikey}/chart/{lat}/{lon}/... functions | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartsGet**
> ResponseCharts ChartsGet(ctx, apikey)
List all charts for given API key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| Your personal API key | 

### Return type

[**ResponseCharts**](ResponseCharts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

