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

func ConvertSuitToEnum(cardSuit string) int {
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

func ConvertEnumToSuit(cardSuitInt int) string {
	var result string
	switch cardSuitInt {
	case 0:
		result = "Club"
	case 1:
		result = "Diamond"
	case 2:
		result = "Heart"
	case 3:
		result = "Spade"
	}

	return result
}

func NewCard(value CardValue, suit CardSuit) Card {
	return Card{value, suit}
}

func MaxCard(cards []Card) Card {
	var card Card
	for i := range cards {
		if i == len(cards)-1 {
			return card
		}
		result := CompareCards(cards[i], cards[i+1])
		switch result {
		case 0, 1:
			card = cards[i]
		case -1:
			card = cards[i+1]
		}
	}

	return card
}
