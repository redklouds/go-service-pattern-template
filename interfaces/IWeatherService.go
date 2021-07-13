package interfaces

import (
	"create-go-svc/models"
	"time"
)

type IWeatherService interface {
	GetWeatherForecasts(time.Time) models.WeatherForecastData
}
