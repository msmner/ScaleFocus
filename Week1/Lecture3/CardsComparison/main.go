package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	card "github.com/msmner/ScaleFocus/tree/main/Week1/Lecture3/Card"
)

func main() {
	fmt.Println("Enter two cards following the pattern to compare the cards: CardValue CardSuit CardValue CardSuit - e.g. 2 Spade 11 Club. Result is -1, 0 and 1 respectively for a weaker, equal and stronger first card.")
	for {
		reader := bufio.NewReader(os.Stdin)
		cardsInputStr, _ := reader.ReadString('\n')
		cards, err := processInput(cardsInputStr)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		result := card.CompareCards(cards[0], cards[1])
		fmt.Printf("Result is: %d\n", result)
	}
}

func processInput(args string) ([]card.Card, error) {
	cardsInput := strings.Split(args, " ")
	for i := range cardsInput {
		cardsInput[i] = strings.TrimSpace(cardsInput[i])
	}

	if len(cardsInput) > 4 || len(cardsInput) < 1 {
		return nil, fmt.Errorf("input must be consistent of 4 arguments! You have got %d", len(args))
	}

	firstCardValue, _ := strconv.Atoi(cardsInput[0])
	firstCardSuit := cardsInput[1]
	firstCardSuitEnum := convertSuitToEnum(firstCardSuit)

	secondCardValue, _ := strconv.Atoi(cardsInput[2])
	secondCardSuit := cardsInput[3]
	secondCardSuitEnum := convertSuitToEnum(secondCardSuit)

	if !card.IsValidValue(firstCardValue) || !card.IsValidValue(secondCardValue) {
		return nil, errors.New("card value out of bounds - must be between 2 and 14")
	}

	if !card.IsValidSuit(firstCardSuit) || !card.IsValidSuit(secondCardSuit) {
		return nil, errors.New("card suit is invalid - must be either Club or Diamond or Heart or Spade")
	}

	firstCard := card.Card{Value: card.CardValue(firstCardValue), Suit: card.CardSuit(firstCardSuitEnum)}
	secondCard := card.Card{Value: card.CardValue(secondCardValue), Suit: card.CardSuit(secondCardSuitEnum)}

	var cards []card.Card
	cards = append(cards, firstCard, secondCard)

	return cards, nil
}

func convertSuitToEnum(cardSuit string) int {
	var result int
	switch cardSuit {
	case "Club":
		result = 0
	case "Diamond":
		result = 1
	case "Heart":
		result = 2
	case "Spade":
		result = 3
	}

	return result
}
