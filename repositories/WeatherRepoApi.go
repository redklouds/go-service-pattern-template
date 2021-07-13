package repositories

/*
This is where you out logic to retrieve data from DB/API/External
*/
import (
	"create-go-svc/interfaces"
	"create-go-svc/models"
	"time"
)

type WeatherRepoRepository struct {
	WeatherApi interfaces.IWeatherRepo
}

func (repo *WeatherRepoRepository) GetWeatherData(time.Time) models.WeatherForecastData {
	return models.WeatherForecastData{}
}
