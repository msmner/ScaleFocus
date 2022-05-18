package services

import (
	"encoding/json"
	"final/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type WeatherService struct{}

type Response struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
	Name    string    `json:"name"`
}

type Weather struct {
	Description string `json:"description"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (ws *WeatherService) GetWeather(latHeader string, lonHeader string) (models.WeatherResponse, error) {
	weatherResponse := models.WeatherResponse{}
	latitude, err := strconv.ParseFloat(latHeader, 32)
	if err != nil {
		return weatherResponse, fmt.Errorf("error parsing latitude: %w", err)
	}

	longitude, err := strconv.ParseFloat(lonHeader, 32)
	if err != nil {
		return weatherResponse, fmt.Errorf("error parsing longitude: %w", err)
	}

	apiKey := os.Getenv("apiKey")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", latitude, longitude, apiKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return weatherResponse, fmt.Errorf("error building request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return weatherResponse, fmt.Errorf("error making the request: %w", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return weatherResponse, fmt.Errorf("error reading the response body: %w", err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return weatherResponse, fmt.Errorf("error unmarshaling the response body: %w", err)
	}

	weatherResponse.FormattedTemp = fmt.Sprintf("%.2f", response.Main.Temp)
	weatherResponse.Description = response.Weather[0].Description
	weatherResponse.City = response.Name

	return weatherResponse, nil
}
