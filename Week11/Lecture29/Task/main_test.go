package main

import (
	"fmt"
	"testing"
)

func TestGroupBy(t *testing.T) {
	//Arrange
	expected := "map[John:[{John 1000} {John 1200}] Sara:[{Sara 2000} {Sara 1800}]]"

	t.Run("Test Group By Function", func(t *testing.T) {
		//Act
		actual := GroupBy([]Order{
			{Customer: "John", Amount: 1000},
			{Customer: "Sara", Amount: 2000},
			{Customer: "Sara", Amount: 1800},
			{Customer: "John", Amount: 1200},
		}, func(o Order) string { return o.Customer })

		//Assert
		if expected != fmt.Sprint(actual) {
			t.Errorf("Failed group by test expected %v, but got %v", expected, actual)
		}
	})
}
