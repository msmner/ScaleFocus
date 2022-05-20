package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeather(t *testing.T) {
	weatherService := NewWeatherService()

	_, err := weatherService.GetWeather("a", "43.23")
	assert.EqualError(t, err, "error parsing latitude: strconv.ParseFloat: parsing \"a\": invalid syntax")

	_, err = weatherService.GetWeather("32.42", "a")
	assert.EqualError(t, err, "error parsing longitude: strconv.ParseFloat: parsing \"a\": invalid syntax")

	//expectedWeatherResponse := models.WeatherResponse{FormattedTemp: "32", Description: "test", City: "Sofiq"}

}
