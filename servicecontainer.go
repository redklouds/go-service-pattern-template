package main

import (
	"create-go-svc/controllers"
	"sync"
)

type kernel struct{}

var (
	ENV           string
	k             *kernel
	containerOnce sync.Once
)

type IServiceContainer interface {
	InjectWeatherController() *controllers.WeatherController
}

//Base off Selected Deployed Environment inject correct Controllers
func (k *kernel) InjectWeatherController() *controllers.WeatherController {
	/*Make Configurable */
	//TODO Make Configurable
	ENV := "DEV"

	var weatherController controllers.WeatherController
	switch ENV {
	case "DEV":
		weatherController = controllers.WeatherController{}
		break
	case "Mock":
		break
	case "PROD":
		break
	default:
		break
	}
	return &weatherController
}

//Ensure Single instance of Container Exists
func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
