package services

import (
	"encoding/json"
	"final/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func (ws *WeatherService) GetWeather(lat float64, long float64) (models.WeatherResponse, error) {
	weatherResponse := models.WeatherResponse{}
	apiKey := os.Getenv("apiKey")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, long, apiKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return weatherResponse, fmt.Errorf("error building the request to weather api: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return weatherResponse, fmt.Errorf("error getting the response from the weather api: %w", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return weatherResponse, fmt.Errorf("error reading the body of the response from the weather api: %w", err)
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return weatherResponse, fmt.Errorf("error deserializing the response body from the weather api: %w", err)
	}

	weatherResponse.FormattedTemp = fmt.Sprintf("%.2f", response.Main.Temp)
	weatherResponse.Description = response.Weather[0].Description
	weatherResponse.City = response.Name

	return weatherResponse, nil
}
