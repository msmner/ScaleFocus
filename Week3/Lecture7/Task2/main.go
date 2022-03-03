package main

import (
	card "Lecture7Task2/Card"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter two cards following the pattern to compare the cards: CardValue CardSuit CardValue CardSuit - e.g. 2 Spade 11 Club. Result is -1, 0 and 1 respectively for a weaker, equal and stronger first card.")
	for {
		reader := bufio.NewReader(os.Stdin)
		cardsInputStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		cards, err := processInput(cardsInputStr)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		// compare the two cards from input
		result := card.CompareCards(cards[0], cards[1])

		fmt.Printf("Result of comparing the two cards is: %d\n", result)

		//Task 2 max card
		// Make a random deck
		cardsDeck := card.CreateDeckOfCards()

		// find max card in the deck
		maxDeckCard := card.MaxCard(cardsDeck, card.CompareCards)

		//print result
		fmt.Printf("Cards deck is: %v\n", cardsDeck)
		fmt.Printf("Result of comparing the deck of cards is: %v\n", maxDeckCard)
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
	firstCardSuitEnum := card.ConvertSuitToEnum(firstCardSuit)

	secondCardValue, _ := strconv.Atoi(cardsInput[2])
	secondCardSuit := cardsInput[3]
	secondCardSuitEnum := card.ConvertSuitToEnum(secondCardSuit)

	if !card.IsValidValue(firstCardValue) || !card.IsValidValue(secondCardValue) {
		return nil, errors.New("card value out of bounds - must be between 2 and 14")
	}

	if !card.IsValidSuit(firstCardSuit) || !card.IsValidSuit(secondCardSuit) {
		return nil, errors.New("card suit is invalid - must be either Club or Diamond or Heart or Spade")
	}

	firstCard := card.NewCard(card.CardValue(firstCardValue), card.CardSuit(firstCardSuitEnum))
	secondCard := card.NewCard(card.CardValue(secondCardValue), card.CardSuit(secondCardSuitEnum))

	var cards []card.Card
	cards = append(cards, firstCard, secondCard)

	return cards, nil
}
