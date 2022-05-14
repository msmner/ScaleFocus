package main

import "fmt"

type Order struct {
	Customer string
	Amount   int
}

func main() {
	results := GroupBy([]Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}, func(o Order) string { return o.Customer })

	fmt.Println(results)
}

func GroupBy[T any, U comparable](col []T, keyFn func(T) U) map[U][]T {
	result := make(map[U][]T)
	for i, v := range col {
		if _, ok := result[keyFn(v)]; !ok {
			result[keyFn(v)] = make([]T, 0, len(col))
		}
		result[keyFn(v)] = append(result[keyFn(v)], col[i])
	}

	return result
}
