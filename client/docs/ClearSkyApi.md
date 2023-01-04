# \ClearSkyApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Clearsky1Get**](ClearSkyApi.md#Clearsky1Get) | **Get** /{apikey}/clearsky/{lat}/{lon}/{dec1}/{az1}/{kwp1} | Get values for one plane
[**Clearsky2Get**](ClearSkyApi.md#Clearsky2Get) | **Get** /{apikey}/clearsky/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | Get values for two planes
[**Clearsky3Get**](ClearSkyApi.md#Clearsky3Get) | **Get** /{apikey}/clearsky/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | Get values for three planes
[**ClearskyWatthours1Get**](ClearSkyApi.md#ClearskyWatthours1Get) | **Get** /{apikey}/clearsky/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1} | Get theoretic clear sky watt hours for one plane
[**ClearskyWatthours2Get**](ClearSkyApi.md#ClearskyWatthours2Get) | **Get** /{apikey}/clearsky/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | Get theoretic clear sky watt hours for two planes
[**ClearskyWatthours3Get**](ClearSkyApi.md#ClearskyWatthours3Get) | **Get** /{apikey}/clearsky/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | Get theoretic clear sky watt hours three planes
[**ClearskyWatthoursDay1Get**](ClearSkyApi.md#ClearskyWatthoursDay1Get) | **Get** /{apikey}/clearsky/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1} | Get theoretic clear sky watt hours per day for one plane
[**ClearskyWatthoursDay2Get**](ClearSkyApi.md#ClearskyWatthoursDay2Get) | **Get** /{apikey}/clearsky/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | Get theoretic clear sky watt hours per day for two planes
[**ClearskyWatthoursDay3Get**](ClearSkyApi.md#ClearskyWatthoursDay3Get) | **Get** /{apikey}/clearsky/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | Get theoretic clear sky watt hours per day for three planes
[**ClearskyWatts1Get**](ClearSkyApi.md#ClearskyWatts1Get) | **Get** /{apikey}/clearsky/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1} | Get watts values for one plane
[**ClearskyWatts2Get**](ClearSkyApi.md#ClearskyWatts2Get) | **Get** /{apikey}/clearsky/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | Get theoretic clear sky watts for two planes
[**ClearskyWatts3Get**](ClearSkyApi.md#ClearskyWatts3Get) | **Get** /{apikey}/clearsky/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | Get theoretic clear sky watt for three planes


# **Clearsky1Get**
> ResponseFull Clearsky1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)
Get values for one plane

Get theoretic clear sky watts, watt hours and watt hours per day for one plane

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
 **optional** | ***ClearSkyApiClearsky1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearsky1GetOpts struct

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

# **Clearsky2Get**
> ResponseFull Clearsky2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)
Get values for two planes

Get theoretic clear sky watts, watt hours and watt hours per day for two planes

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
 **optional** | ***ClearSkyApiClearsky2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearsky2GetOpts struct

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

# **Clearsky3Get**
> ResponseFull Clearsky3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)
Get values for three planes

Get theoretic clear sky watts, watt hours and watt hours per day for three planes

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
 **optional** | ***ClearSkyApiClearsky3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearsky3GetOpts struct

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

# **ClearskyWatthours1Get**
> ResponseFull ClearskyWatthours1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)
Get theoretic clear sky watt hours for one plane

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
 **optional** | ***ClearSkyApiClearskyWatthours1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatthours1GetOpts struct

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

# **ClearskyWatthours2Get**
> ResponseFull ClearskyWatthours2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)
Get theoretic clear sky watt hours for two planes

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
 **optional** | ***ClearSkyApiClearskyWatthours2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatthours2GetOpts struct

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

# **ClearskyWatthours3Get**
> ResponseFull ClearskyWatthours3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)
Get theoretic clear sky watt hours three planes

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
 **optional** | ***ClearSkyApiClearskyWatthours3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatthours3GetOpts struct

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

# **ClearskyWatthoursDay1Get**
> ResponseFull ClearskyWatthoursDay1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)
Get theoretic clear sky watt hours per day for one plane

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
 **optional** | ***ClearSkyApiClearskyWatthoursDay1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatthoursDay1GetOpts struct

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

# **ClearskyWatthoursDay2Get**
> ResponseFull ClearskyWatthoursDay2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)
Get theoretic clear sky watt hours per day for two planes

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
 **optional** | ***ClearSkyApiClearskyWatthoursDay2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatthoursDay2GetOpts struct

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

# **ClearskyWatthoursDay3Get**
> ResponseFull ClearskyWatthoursDay3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)
Get theoretic clear sky watt hours per day for three planes

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
 **optional** | ***ClearSkyApiClearskyWatthoursDay3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatthoursDay3GetOpts struct

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

# **ClearskyWatts1Get**
> ResponseFull ClearskyWatts1Get(ctx, apikey, lat, lon, dec1, az1, kwp1, optional)
Get watts values for one plane

Get theoretic clear sky watts for one plane

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
 **optional** | ***ClearSkyApiClearskyWatts1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatts1GetOpts struct

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

# **ClearskyWatts2Get**
> ResponseFull ClearskyWatts2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)
Get theoretic clear sky watts for two planes

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
 **optional** | ***ClearSkyApiClearskyWatts2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatts2GetOpts struct

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

# **ClearskyWatts3Get**
> ResponseFull ClearskyWatts3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)
Get theoretic clear sky watt for three planes

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
 **optional** | ***ClearSkyApiClearskyWatts3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClearSkyApiClearskyWatts3GetOpts struct

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

