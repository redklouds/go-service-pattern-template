package models

import "time"

/*
Define your service Models here
*/

type Summaries struct {
	Summary []string
}

type WeatherForecast struct {
	Date         time.Time
	TemperatureC int
	TemperatureF int
	Summary      string
}

type WeatherForecastData struct {
	Date        time.Time `json:"date"`
	Temperature int       `json:"temperature"`
}
