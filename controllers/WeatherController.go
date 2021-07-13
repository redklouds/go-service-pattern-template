package controllers

/*
Define your service controllers here
*/

import (
	"create-go-svc/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type WeatherController struct{}

func (controller *WeatherController) GetWeatherForecast(w http.ResponseWriter, req *http.Request) {
	Summaries := [...]string{"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}

	s1 := rand.NewSource(time.Now().Unix())
	rng := rand.New(s1)

	var result []models.WeatherForecast
	for i := 0; i < 5; i++ {
		_weather := models.WeatherForecast{
			Date:         time.Now().AddDate(0, 0, 5),
			TemperatureC: rng.Intn(55-(-20)) + -20,
			Summary:      Summaries[rng.Intn(len(Summaries))],
		}
		result = append(result, _weather)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
	return

}
