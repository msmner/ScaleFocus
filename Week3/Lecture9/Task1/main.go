package main

import (
	"Lecture9Task1/carddraw"
	"Lecture9Task1/cardgame"
	"fmt"
	"log"
)

func main() {
	emptyDeck := &cardgame.Deck{}
	initDeck := emptyDeck.New()
	shuffledDeck := initDeck.Shuffle()
	cards, err := carddraw.DrawAllCards(shuffledDeck)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cards drawn are: %v", cards)
}
