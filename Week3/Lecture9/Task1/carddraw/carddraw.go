package carddraw

import (
	"Lecture9Task1/cardgame"
	"fmt"
)

type dealer interface {
	Deal() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	var cards []cardgame.Card
	for {
		dealtCard, err := dealer.Deal()
		if err != nil {
			if dealer.Done() {
				return cards, nil
			}

			wrappedErr := fmt.Errorf("got error dealing the cards: %w", err)

			return nil, wrappedErr
		}

		cards = append(cards, *dealtCard)
	}
}
