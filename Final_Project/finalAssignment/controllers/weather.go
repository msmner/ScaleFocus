package controllers

import (
	"final/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type WeatherController struct {
	weatherService *services.WeatherService
}

func NewWeatherController(ws *services.WeatherService) *WeatherController {
	return &WeatherController{weatherService: ws}
}

func (wc *WeatherController) GetWeather(c echo.Context) (err error) {
	req := c.Request()
	headers := req.Header
	latHeader := headers.Get("lat")
	lonHeader := headers.Get("lon")
	latitude, err := strconv.ParseFloat(latHeader, 32)
	if err != nil {
		return err
	}
	longitude, err := strconv.ParseFloat(lonHeader, 32)
	if err != nil {
		return err
	}

	weather, err := wc.weatherService.GetWeather(latitude, longitude)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, weather)
}
