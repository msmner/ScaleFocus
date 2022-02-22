package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cities, prices := citiesAndPrices()
	result := groupSlices(cities, prices)
	fmt.Printf("Map is: %v", result)
}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}

func groupSlices(cities []string, prices []int) map[string][]int {
	var mapResult = make(map[string][]int)
	for i, v := range cities {
		mapResult[v] = append(mapResult[v], prices[i])
	}

	return mapResult
}
