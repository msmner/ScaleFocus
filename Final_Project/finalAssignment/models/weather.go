package models

type WeatherResponse struct {
	FormattedTemp string `json:"formattedTemp"`
	Description   string `json:"description"`
	City          string `json:"city"`
}
