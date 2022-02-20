package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter two cards following the pattern to compare the cards: CardValue CardSuit CardValue CardSuit - e.g. 2 Spade King Club. Result is -1, 0 and 1 respectively for a weaker, equal and stronger first card.")
	inputArgs := processInput(os.Args[1:])
	fmt.Printf("args %v", inputArgs)
}

func processInput(args []string) []Card {

}
