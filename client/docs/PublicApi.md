# \PublicApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimateLatLonDecAzKwpGet**](PublicApi.md#EstimateLatLonDecAzKwpGet) | **Get** /estimate/{lat}/{lon}/{dec}/{az}/{kwp} | 
[**EstimateWatthoursDayLatLonDecAzKwpGet**](PublicApi.md#EstimateWatthoursDayLatLonDecAzKwpGet) | **Get** /estimate/watthours/day/{lat}/{lon}/{dec}/{az}/{kwp} | 
[**EstimateWatthoursLatLonDecAzKwpGet**](PublicApi.md#EstimateWatthoursLatLonDecAzKwpGet) | **Get** /estimate/watthours/{lat}/{lon}/{dec}/{az}/{kwp} | 
[**EstimateWattsLatLonDecAzKwpGet**](PublicApi.md#EstimateWattsLatLonDecAzKwpGet) | **Get** /estimate/watts/{lat}/{lon}/{dec}/{az}/{kwp} | 
[**HelpGet**](PublicApi.md#HelpGet) | **Get** /help | 
[**SwaggerYamlGet**](PublicApi.md#SwaggerYamlGet) | **Get** /swagger.yaml | 


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
 **optional** | ***PublicApiEstimateLatLonDecAzKwpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiEstimateLatLonDecAzKwpGetOpts struct

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
 **optional** | ***PublicApiEstimateWatthoursDayLatLonDecAzKwpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiEstimateWatthoursDayLatLonDecAzKwpGetOpts struct

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
 **optional** | ***PublicApiEstimateWatthoursLatLonDecAzKwpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiEstimateWatthoursLatLonDecAzKwpGetOpts struct

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
 **optional** | ***PublicApiEstimateWattsLatLonDecAzKwpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiEstimateWattsLatLonDecAzKwpGetOpts struct

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

