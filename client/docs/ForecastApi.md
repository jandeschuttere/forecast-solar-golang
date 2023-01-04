# \ForecastApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**](ForecastApi.md#ApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get) | **Get** /{apikey}/estimate/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | 
[**ApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**](ForecastApi.md#ApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Get) | **Get** /{apikey}/estimate/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | 
[**ApikeyEstimateLatLonDec1Az1Kwp1Get**](ForecastApi.md#ApikeyEstimateLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/estimate/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**](ForecastApi.md#ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get) | **Get** /{apikey}/estimate/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | 
[**ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**](ForecastApi.md#ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Get) | **Get** /{apikey}/estimate/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | 
[**ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Get**](ForecastApi.md#ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/estimate/watthours/day/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**](ForecastApi.md#ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get) | **Get** /{apikey}/estimate/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | 
[**ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**](ForecastApi.md#ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Get) | **Get** /{apikey}/estimate/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | 
[**ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Get**](ForecastApi.md#ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/estimate/watthours/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**ApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**](ForecastApi.md#ApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get) | **Get** /{apikey}/estimate/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2}/{dec3}/{az3}/{kwp3} | 
[**ApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**](ForecastApi.md#ApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Get) | **Get** /{apikey}/estimate/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1}/{dec2}/{az2}/{kwp2} | 
[**ApikeyEstimateWattsLatLonDec1Az1Kwp1Get**](ForecastApi.md#ApikeyEstimateWattsLatLonDec1Az1Kwp1Get) | **Get** /{apikey}/estimate/watts/{lat}/{lon}/{dec1}/{az1}/{kwp1} | 
[**EstimateLatLonDecAzKwpGet**](ForecastApi.md#EstimateLatLonDecAzKwpGet) | **Get** /estimate/{lat}/{lon}/{dec}/{az}/{kwp} | 
[**EstimateWatthoursDayLatLonDecAzKwpGet**](ForecastApi.md#EstimateWatthoursDayLatLonDecAzKwpGet) | **Get** /estimate/watthours/day/{lat}/{lon}/{dec}/{az}/{kwp} | 
[**EstimateWatthoursLatLonDecAzKwpGet**](ForecastApi.md#EstimateWatthoursLatLonDecAzKwpGet) | **Get** /estimate/watthours/{lat}/{lon}/{dec}/{az}/{kwp} | 
[**EstimateWattsLatLonDecAzKwpGet**](ForecastApi.md#EstimateWattsLatLonDecAzKwpGet) | **Get** /estimate/watts/{lat}/{lon}/{dec}/{az}/{kwp} | 


# **ApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**
> ResponseFull ApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)


Get estimated watts, watt hours and watt hours per day for three planes

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
 **optional** | ***ForecastApiApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts struct

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

# **ApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**
> ResponseFull ApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)


Get estimated watts, watt hours and watt hours per day for two planes

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
 **optional** | ***ForecastApiApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts struct

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
 **optional** | ***ForecastApiApikeyEstimateLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateLatLonDec1Az1Kwp1GetOpts struct

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

# **ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**
> ResponseWattHoursDay ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)


Get estimated watt hours per day for three planes

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
 **optional** | ***ForecastApiApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts struct

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

# **ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**
> ResponseWattHoursDay ApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)


Get estimated hours per day for two planes

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
 **optional** | ***ForecastApiApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts struct

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
 **optional** | ***ForecastApiApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWatthoursDayLatLonDec1Az1Kwp1GetOpts struct

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

# **ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**
> ResponseWattHours ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)


Get estimated watt hours for three planes

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
 **optional** | ***ForecastApiApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts struct

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

# **ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**
> ResponseWattHours ApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)


Get estimated watt hours for two planes

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
 **optional** | ***ForecastApiApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWatthoursLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts struct

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
 **optional** | ***ForecastApiApikeyEstimateWatthoursLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWatthoursLatLonDec1Az1Kwp1GetOpts struct

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

# **ApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get**
> ResponseWatts ApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, dec3, az3, kwp3, optional)


Get estimated watts for three planes

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
 **optional** | ***ForecastApiApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Dec3Az3Kwp3GetOpts struct

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

# **ApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Get**
> ResponseWatts ApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2Get(ctx, apikey, lat, lon, dec1, az1, kwp1, dec2, az2, kwp2, optional)


Get estimated watts for two planes

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
 **optional** | ***ForecastApiApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWattsLatLonDec1Az1Kwp1Dec2Az2Kwp2GetOpts struct

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
 **optional** | ***ForecastApiApikeyEstimateWattsLatLonDec1Az1Kwp1GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiApikeyEstimateWattsLatLonDec1Az1Kwp1GetOpts struct

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

# **EstimateLatLonDecAzKwpGet**
> ResponseFull EstimateLatLonDecAzKwpGet(ctx, lat, lon, dec, az, kwp, optional)


Get estimated watts, watt hours and watt hours per day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lat** | **float32**| Latitude of location | 
  **lon** | **float32**| Longitude of location | 
  **dec** | **float32**| Solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az** | **float32**| Solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp** | **float32**| Solar plane max. peak power in kilo watt | 
 **optional** | ***ForecastApiEstimateLatLonDecAzKwpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiEstimateLatLonDecAzKwpGetOpts struct

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

# **EstimateWatthoursDayLatLonDecAzKwpGet**
> ResponseWattHoursDay EstimateWatthoursDayLatLonDecAzKwpGet(ctx, lat, lon, dec, az, kwp, optional)


Get estimated watt hours sum per day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lat** | **float32**| Latitude of location | 
  **lon** | **float32**| Longitude of location | 
  **dec** | **float32**| Solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az** | **float32**| Solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp** | **float32**| Solar plane max. peak power in kilo watt | 
 **optional** | ***ForecastApiEstimateWatthoursDayLatLonDecAzKwpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiEstimateWatthoursDayLatLonDecAzKwpGetOpts struct

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

# **EstimateWatthoursLatLonDecAzKwpGet**
> ResponseWattHours EstimateWatthoursLatLonDecAzKwpGet(ctx, lat, lon, dec, az, kwp, optional)


Get estimated watt hours data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lat** | **float32**| Latitude of location | 
  **lon** | **float32**| Longitude of location | 
  **dec** | **float32**| Solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az** | **float32**| Solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp** | **float32**| Solar plane max. peak power in kilo watt | 
 **optional** | ***ForecastApiEstimateWatthoursLatLonDecAzKwpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiEstimateWatthoursLatLonDecAzKwpGetOpts struct

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

# **EstimateWattsLatLonDecAzKwpGet**
> ResponseWatts EstimateWattsLatLonDecAzKwpGet(ctx, lat, lon, dec, az, kwp, optional)


Get estimated watts per day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lat** | **float32**| Latitude of location | 
  **lon** | **float32**| Longitude of location | 
  **dec** | **float32**| Solar plane declination, 0 &#x3D; horizontal, 90 &#x3D; vertical | 
  **az** | **float32**| Solar plane azimuth, West &#x3D; 90, South &#x3D; 0, East &#x3D; -90 | 
  **kwp** | **float32**| Solar plane max. peak power in kilo watt | 
 **optional** | ***ForecastApiEstimateWattsLatLonDecAzKwpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastApiEstimateWattsLatLonDecAzKwpGetOpts struct

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

