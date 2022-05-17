package interfaces

import "final/models"

type IWeatherService interface {
	GetWeather(latHeader string, lonHeader string) (models.WeatherResponse, error)
}
