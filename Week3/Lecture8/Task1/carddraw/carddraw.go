package carddraw

import "Lecture8Task1/cardgame"

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	var cards []cardgame.Card
	for {
		dealtCard := dealer.Deal()
		if dealtCard == nil {
			return cards
		}

		cards = append(cards, *dealtCard)
	}
}
