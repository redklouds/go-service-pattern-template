package interfaces

/*
This is where you put your interfaces the services will need to define
-> Keeping Testablity in mind
*/

import (
	"create-go-svc/models"
	"time"
)

type IWeatherRepo interface {
	GetWeatherData(time.Time) models.WeatherForecastData
}
