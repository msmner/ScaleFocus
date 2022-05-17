package controllers

import (
	"final/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WeatherController struct {
	weatherService interfaces.IWeatherService
}

func NewWeatherController(ws interfaces.IWeatherService) *WeatherController {
	return &WeatherController{weatherService: ws}
}

func (wc *WeatherController) GetWeather(c echo.Context) error {
	req := c.Request()
	headers := req.Header
	latHeader := headers.Get("lat")
	lonHeader := headers.Get("lon")
	weather, err := wc.weatherService.GetWeather(latHeader, lonHeader)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, weather)
}
