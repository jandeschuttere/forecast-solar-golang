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

type MessageOk struct {
	Code      int32             `json:"code,omitempty"`
	Type_     string            `json:"type,omitempty"`
	Text      string            `json:"text,omitempty"`
	Info      *MessageInfo      `json:"info,omitempty"`
	RateLimit *MessageRateLimit `json:"ratelimit,omitempty"`
}

type MessageInfo struct {
	Latitude  float64   `json:"latitude,omitempty"`
	Longitude float64   `json:"longitude,omitempty"`
	Distance  int       `json:"distance,omitempty"`
	Place     string    `json:"place,omitempty"`
	Timezone  string    `json:"timezone,omitempty"`
	Time      time.Time `json:"time,omitempty"`
	TimeUtc   time.Time `json:"time_utc,omitempty"`
}

type MessageRateLimit struct {
	Period    int `json:"period,omitempty"`
	Limit     int `json:"limit,omitempty"`
	Remaining int `json:"remaining,omitempty"`
}