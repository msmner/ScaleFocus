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

		year, month := processInput(input)

		daysInMonth, isLeapYear := daysInMonth(year, month)

		fmt.Printf("Days in month are %d. Year is leap year: %v\n", daysInMonth, isLeapYear)
	}
}

func processInput(args string) (int, int) {
	inputArr := strings.Split(args, " ")
	for i := range inputArr {
		inputArr[i] = strings.TrimSpace(inputArr[i])
	}

	month, err := strconv.Atoi(inputArr[0])
	if err != nil {
		fmt.Printf("Error parsing month %v: %s", month, err)
	}

	year, err := strconv.Atoi(inputArr[1])
	if err != nil {
		fmt.Printf("Error parsing year %v: %s", year, err)
	}

	return month, year
}

func daysInMonth(month int, year int) (int, bool) {
	if month < 1 || month > 12 {
		fmt.Printf("Invalid month %d", month)
		return 0, false
	}

	isLeapYear := checkIfLeapYear(year)

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31, isLeapYear
	case 4, 6, 9, 11:
		return 30, isLeapYear
	case 2:
		if isLeapYear {
			return 29, isLeapYear
		} else {
			return 28, isLeapYear
		}
	}

	return 0, false
}

func checkIfLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}
