package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a month and a year following the pattern: MONTH YEAR - e.g. 2 2020")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		year, month, err := processInput(input)
		if err != nil {
			fmt.Printf("Got Error %s: ", err)
			continue
		}

		daysInMonth, isLeapYear, err := daysInMonth(year, month)
		if err != nil {
			fmt.Printf("Got Error %s: ", err)
			continue
		}

		fmt.Printf("Days in month are %d. Year is leap year: %v\n", daysInMonth, isLeapYear)
	}
}

func processInput(args string) (int, int, error) {
	inputArr := strings.Split(args, " ")
	for i := range inputArr {
		inputArr[i] = strings.TrimSpace(inputArr[i])
	}

	month, err := strconv.Atoi(inputArr[0])
	if err != nil {
		formattedError := fmt.Errorf("error parsing month %v: %s", month, err)
		return 0, 0, formattedError
	}

	year, err := strconv.Atoi(inputArr[1])
	if err != nil {
		formattedError := fmt.Errorf("error parsing year %v: %s", year, err)
		return 0, 0, formattedError
	}

	return month, year, nil
}

func daysInMonth(month int, year int) (int, bool, error) {
	if month < 1 || month > 12 {
		formattedError := fmt.Errorf("invalid month %d", month)
		return 0, false, formattedError
	}

	isLeapYear := checkIfLeapYear(year)

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31, isLeapYear, nil
	case 4, 6, 9, 11:
		return 30, isLeapYear, nil
	case 2:
		if isLeapYear {
			return 29, isLeapYear, nil
		} else {
			return 28, isLeapYear, nil
		}
	}

	return 0, false, nil
}

func checkIfLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}
