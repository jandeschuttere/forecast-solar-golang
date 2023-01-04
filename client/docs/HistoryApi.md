# \HistoryApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**](HistoryApi.md#ApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get) | **Get** /{apikey}/history/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | 
[**ApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**](HistoryApi.md#ApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Get) | **Get** /{apikey}/history/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | 
[**ApikeyHistoryLatLonDec1Az1Kwp1Get**](HistoryApi.md#ApikeyHistoryLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/history/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**](HistoryApi.md#ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get) | **Get** /{apikey}/history/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | 
[**ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**](HistoryApi.md#ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Get) | **Get** /{apikey}/history/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | 
[**ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Get**](HistoryApi.md#ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/history/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**](HistoryApi.md#ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get) | **Get** /{apikey}/history/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | 
[**ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**](HistoryApi.md#ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Get) | **Get** /{apikey}/history/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | 
[**ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Get**](HistoryApi.md#ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/history/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**](HistoryApi.md#ApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get) | **Get** /{apikey}/history/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | 
[**ApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**](HistoryApi.md#ApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Get) | **Get** /{apikey}/history/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | 
[**ApikeyHistoryWattsLatLonDec1Az1Kwp1Get**](HistoryApi.md#ApikeyHistoryWattsLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/history/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 


# **ApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**
> ResponseFull ApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)


Get real average watts, watt hours and watt hours per day for three planes

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
 **optional** | ***HistoryApiApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts struct

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

# **ApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**
> ResponseFull ApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)


Get real average watts, watt hours and watt hours per day for two planes

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
 **optional** | ***HistoryApiApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts struct

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
 **optional** | ***HistoryApiApikeyHistoryLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryLatLonDec1Az1Kwp1GetOpts struct

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

# **ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**
> ResponseFull ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)


Get real average watt hours per day for three planes

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
 **optional** | ***HistoryApiApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts struct

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

# **ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**
> ResponseFull ApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)


Get real average watt hours per day for two planes

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
 **optional** | ***HistoryApiApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts struct

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
 **optional** | ***HistoryApiApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWatthoursDayLatLonDec1Az1Kwp1GetOpts struct

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

# **ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**
> ResponseFull ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)


Get real average watt hours for three planes

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
 **optional** | ***HistoryApiApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts struct

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

# **ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**
> ResponseFull ApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)


Get real average watt hours for two planes

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
 **optional** | ***HistoryApiApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts struct

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
 **optional** | ***HistoryApiApikeyHistoryWatthoursLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWatthoursLatLonDec1Az1Kwp1GetOpts struct

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

# **ApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**
> ResponseFull ApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)


Get real average watts for three planes

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
 **optional** | ***HistoryApiApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts struct

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

# **ApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**
> ResponseFull ApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)


Get real average watts for two planes

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
 **optional** | ***HistoryApiApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts struct

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
 **optional** | ***HistoryApiApikeyHistoryWattsLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiApikeyHistoryWattsLatLonDec1Az1Kwp1GetOpts struct

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

