# \PersonalApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApikeyEstimateLatLonDec1Az1Kwp1Get**](PersonalApi.md#ApikeyEstimateLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/estimate/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Get**](PersonalApi.md#ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/estimate/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Get**](PersonalApi.md#ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/estimate/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyEstimateWattsLatLonDec1Az1Kwp1Get**](PersonalApi.md#ApikeyEstimateWattsLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/estimate/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyHistoryLatLonDec1Az1Kwp1Get**](PersonalApi.md#ApikeyHistoryLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/history/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Get**](PersonalApi.md#ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/history/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Get**](PersonalApi.md#ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/history/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyHistoryWattsLatLonDec1Az1Kwp1Get**](PersonalApi.md#ApikeyHistoryWattsLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/history/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**HelpGet**](PersonalApi.md#HelpGet) | **Get** /help | 
[**SwaggerYamlGet**](PersonalApi.md#SwaggerYamlGet) | **Get** /swagger.yaml | 


# **ApikeyEstimateLatLonDec1Az1Kwp1Get**
> ResponseFull ApikeyEstimateLatLonDec1Az1Kwp1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)


Get estimated watts, watt hours and watt hours per day for one plane

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
 **optional** | ***PersonalApiApikeyEstimateLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalApiApikeyEstimateLatLonDec1Az1Kwp1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **time** | **optional.String**| Set to &#39;utc&#39; to return timestamps in ISO 8601, e.g. 2004-02-12T15:19:21+00:00 | 
 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseFull**](ResponseFull.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Get**
> ResponseWattHoursDay ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)


Get estimated watt hours per day for one plane

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
 **optional** | ***PersonalApiApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalApiApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **time** | **optional.String**| Set to &#39;utc&#39; to return timestamps in ISO 8601, e.g. 2004-02-12T15:19:21+00:00 | 
 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseWattHoursDay**](ResponseWattHoursDay.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Get**
> ResponseWattHours ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)


Get estimated watt hours for one plane

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
 **optional** | ***PersonalApiApikeyEstimateWatthoursLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalApiApikeyEstimateWatthoursLatLonDec1Az1Kwp1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **time** | **optional.String**| Set to &#39;utc&#39; to return timestamps in ISO 8601, e.g. 2004-02-12T15:19:21+00:00 | 
 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseWattHours**](ResponseWattHours.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApikeyEstimateWattsLatLonDec1Az1Kwp1Get**
> ResponseWatts ApikeyEstimateWattsLatLonDec1Az1Kwp1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)


Get estimated watts for one plane

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
 **optional** | ***PersonalApiApikeyEstimateWattsLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalApiApikeyEstimateWattsLatLonDec1Az1Kwp1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **time** | **optional.String**| Set to &#39;utc&#39; to return timestamps in ISO 8601, e.g. 2004-02-12T15:19:21+00:00 | 
 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseWatts**](ResponseWatts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApikeyHistoryLatLonDec1Az1Kwp1Get**
> ResponseFull ApikeyHistoryLatLonDec1Az1Kwp1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)


Get real average watts, watt hours and watt hours per day for one plane

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
 **optional** | ***PersonalApiApikeyHistoryLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalApiApikeyHistoryLatLonDec1Az1Kwp1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseFull**](ResponseFull.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Get**
> ResponseFull ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)


Get real average watt hours per day for one plane

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
 **optional** | ***PersonalApiApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalApiApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseFull**](ResponseFull.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Get**
> ResponseFull ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)


Get real average watt hours for one plane

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
 **optional** | ***PersonalApiApikeyHistoryWatthoursLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalApiApikeyHistoryWatthoursLatLonDec1Az1Kwp1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseFull**](ResponseFull.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApikeyHistoryWattsLatLonDec1Az1Kwp1Get**
> ResponseFull ApikeyHistoryWattsLatLonDec1Az1Kwp1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)


Get real average watts for one plane

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
 **optional** | ***PersonalApiApikeyHistoryWattsLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalApiApikeyHistoryWattsLatLonDec1Az1Kwp1GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseFull**](ResponseFull.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HelpGet**
> string HelpGet(ctx, )


Refer to URL to explore the API in Swagger UI

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwaggerYamlGet**
> string SwaggerYamlGet(ctx, )


API documentation in YAML for Swagger UI

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

