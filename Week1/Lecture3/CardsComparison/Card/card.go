package card

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

func IsValidValue(arg int) bool {
	switch arg {
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		return true
	}
	return false
}

type CardSuit int

const (
	Club CardSuit = iota
	Diamond
	Heart
	Spade
)

func IsValidSuit(arg string) bool {
	switch arg {
	case "Club", "Diamond", "Heart", "Spade":
		return true
	}
	return false
}

func CompareCards(firstCard Card, secondCard Card) int {
	if firstCard.Value < secondCard.Value {
		return -1
	} else if firstCard.Value > secondCard.Value {
		return 1
	} else {
		if firstCard.Suit < secondCard.Suit {
			return -1
		} else if firstCard.Suit > secondCard.Suit {
			return 1
		} else {
			return 0
		}
	}
}
