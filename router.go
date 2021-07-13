package main

import (
	"sync"

	"github.com/go-chi/chi"
)

var (
	m          *Router
	routerOnce sync.Once
)

type ChiRouter interface {
	InitRouter() *chi.Mux
}
type Router struct{}

//Define Routes Here
func (router *Router) InitRouter() *chi.Mux {
	r := chi.NewRouter()

	//Get the controller to from container to inject
	WeatherController := ServiceContainer().InjectWeatherController()

	//Route Group
	r.Route("/api", func(r chi.Router) {
		r.Get("/WeatherForecast", WeatherController.GetWeatherForecast)
	})
	return r
}

//Ensure Single Instance exists - Singleton
func NewChiRouter() ChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &Router{}
		})
	}
	return m
}
