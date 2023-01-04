/*
 * Forecast.Solar
 *
 * ## Restful API for Solar production forecast data and Weather forecast data. ### Retrieve data based on a location, the declination and orientation of solar panels.  #### <a href='https://forecast.solar'>Homepage</a> &bull; <a href='https://doc.forecast.solar'>Documentation</a>  <small>made by <a href='https://knutkohl.consulting'>Knut Kohl . Consulting</a></small>
 *
 * API version: v5.69.0.1072
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import "time"

// Production forecast per day
type ResponseWattHoursDay struct {
	Result  *map[time.Time]float64 `json:"result,omitempty"`
	Message *MessageOk             `json:"message,omitempty"`
}
