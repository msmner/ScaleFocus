package main

import (
	"Lecture8Task1/carddraw"
	"Lecture8Task1/cardgame"
	"fmt"
)

func main() {
	emptyDeck := &cardgame.Deck{}
	initDeck := emptyDeck.New()
	shuffledDeck := initDeck.Shuffle()
	cards := carddraw.DrawAllCards(shuffledDeck)
	fmt.Printf("Cards drawn are: %v", cards)
}
