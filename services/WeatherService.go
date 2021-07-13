package services

import (
	"create-go-svc/interfaces"
	"create-go-svc/models"
	"time"
)

/*
Place your Service layer logic here
*/

type WeatherService struct {
	WeatherRepo interfaces.IWeatherService
}

func (service *WeatherService) GetWeatherForecasts(reqTime time.Time) models.WeatherForecast {
	forcastData := service.WeatherRepo.GetWeatherForecasts(reqTime)
	result := models.WeatherForecast{
		Date: forcastData.Date,
	}
	if forcastData.Temperature < 0 {
		result.TemperatureC = forcastData.Temperature
	} else {
		result.TemperatureF = forcastData.Temperature
	}
	return result
}
