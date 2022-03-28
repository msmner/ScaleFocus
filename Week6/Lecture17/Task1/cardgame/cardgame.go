package cardgame

type Card struct {
	Value CardValue
	Suit  CardSuit
}

type CardValue int

const (
	Two CardValue = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type CardSuit int

const (
	Club CardSuit = iota
	Diamond
	Heart
	Spade
)

func CompareCards(firstCard Card, secondCard Card) int {
	var result int
	if firstCard.Value < secondCard.Value {
		result = -1
	} else if firstCard.Value > secondCard.Value {
		result = 1
	} else {
		if firstCard.Suit < secondCard.Suit {
			result = -1
		} else if firstCard.Suit > secondCard.Suit {
			result = 1
		} else {
			result = 0
		}
	}
	return result
}

func MaxCard(cards []Card) Card {
	maxCard := cards[0]
	for _, c := range cards {
		result := CompareCards(maxCard, c)
		switch result {
		case 0, 1:
			continue
		case -1:
			maxCard = c
		}
	}

	return maxCard
}
