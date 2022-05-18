package controllers

import (
	"final/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	weatherResponseJSON = `{"formattedTemp":"33","description":"test","city":"sofia"}
`
)

type MockWeatherService struct{}

func (ws *MockWeatherService) GetWeather(latHeader string, lonHeader string) (models.WeatherResponse, error) {
	weatherResponse := models.WeatherResponse{FormattedTemp: "33", Description: "test", City: "sofia"}
	return weatherResponse, nil
}

func TestGetWeather(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	req.Header.Set("lat", "33")
	req.Header.Set("lon", "33")
	c := e.NewContext(req, rec)
	mockService := MockWeatherService{}
	weatherController := NewWeatherController(&mockService)
	if assert.NoError(t, weatherController.GetWeather(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, weatherResponseJSON, rec.Body.String())
	}
}
