package cardgame

import (
	"math/rand"
	"time"
)

type Card struct {
	Value string
	Suit  string
}

type Deck struct {
	cards []Card
}

func (d *Deck) New() *Deck {
	cards := make([]Card, 0, 52)
	suits := []string{"Club", "Diamond", "Heart", "Spade"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	for _, s := range suits {
		for _, v := range values {
			card := Card{v, s}
			cards = append(cards, card)
		}
	}

	return &Deck{cards}
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
