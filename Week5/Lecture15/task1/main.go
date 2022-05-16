package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	dates := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022"}
	format := "Jan-02-2006"
	result, err := sortDates(format, dates...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func sortDates(format string, dates ...string) ([]string, error) {
	timeSlice := make([]time.Time, 0, len(dates))
	result := make([]string, 0, len(dates))

	for _, v := range dates {
		v, err := time.Parse(format, v)
		if err != nil {
			return nil, err
		}
		timeSlice = append(timeSlice, v)
	}

	sort.Slice(timeSlice, func(i, j int) bool {
		return timeSlice[i].Before(timeSlice[j])
	})

	for _, k := range timeSlice {
		d := k.String()
		result = append(result, d)
	}
	return result, nil
}
