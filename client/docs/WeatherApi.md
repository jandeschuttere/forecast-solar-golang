# \WeatherApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WeatherGet**](WeatherApi.md#WeatherGet) | **Get** /{apikey}/weather/{lat}/{lon} | Get weather forecast for a location


# **WeatherGet**
> ResponseWeather WeatherGet(ctx, apikey, lat, lon, optional)
Get weather forecast for a location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| Your personal API key | 
  **lat** | **float32**| Latitude of location | 
  **lon** | **float32**| Longitude of location | 
 **optional** | ***WeatherApiWeatherGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WeatherApiWeatherGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xDelimiter** | **optional.String**| Dataset delimiter in CSV response, default &#39;|&#39; | [default to |]
 **xSeparator** | **optional.String**| Value separator in CSV response, default &#39;;&#39;, for a tabulator use string &#39;TAB&#39; | [default to ;]

### Return type

[**ResponseWeather**](ResponseWeather.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

