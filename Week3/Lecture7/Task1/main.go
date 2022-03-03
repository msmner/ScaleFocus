package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	emptyDeck := &Deck{}
	cardsDeck := emptyDeck.New()
	fmt.Printf("Newly generated cards are: %v\n", cardsDeck)
	cardsDeck = cardsDeck.Shuffle()
	fmt.Printf("Shuffled cards are: %v\n", cardsDeck)

	for i := 0; i < 53; i++ {
		dealtCardLoop := cardsDeck.Deal()
		if dealtCardLoop == nil {
			fmt.Printf("Deck is empty - can't deal any more cards")
			return
		}
		fmt.Printf("Dealt card is: %v Length of deck is: %d\n", dealtCardLoop, len(cardsDeck.cards))
	}

}

type Card struct {
	Suit  string
	Value string
}

type Deck struct {
	cards []Card
}

func (d *Deck) Deal() *Card {
	if len(d.cards) == 0 {
		return nil
	}

	dealtCard := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return &dealtCard
}

func (d *Deck) Shuffle() *Deck {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	return d
}

func (d *Deck) New() *Deck {
	cards := make([]Card, 0, 52)
	suits := []string{"Club", "Diamond", "Heart", "Spade"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	for _, s := range suits {
		for _, v := range values {
			card := Card{s, v}
			cards = append(cards, card)
		}
	}

	return &Deck{cards}
}
