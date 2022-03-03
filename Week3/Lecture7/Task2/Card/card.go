package card

import (
	"fmt"
	"math/rand"
	"time"
)

type CardComparator func(cOne Card, cTwo Card) int

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

func IsValidSuit(arg string) bool {
	switch arg {
	case "Club", "Diamond", "Heart", "Spade":
		return true
	}
	return false
}

func IsValidValue(arg int) bool {
	switch arg {
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
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

func (cs CardSuit) String() string {
	var result string
	switch cs {
	case 0:
		result = "Club"
	case 1:
		result = "Diamond"
	case 2:
		result = "Heart"
	case 3:
		result = "Spade"
	default:
		return fmt.Sprintf("%d", int(cs))
	}
	return result
}

func (cv CardValue) String() string {
	var result string
	switch cv {
	case 2:
		result = "2"
	case 3:
		result = "3"
	case 4:
		result = "4"
	case 5:
		result = "5"
	case 6:
		result = "6"
	case 7:
		result = "7"
	case 8:
		result = "8"
	case 9:
		result = "9"
	case 10:
		result = "10"
	case 11:
		result = "Jack"
	case 12:
		result = "Queen"
	case 13:
		result = "King"
	case 14:
		result = "Ace"
	default:
		return fmt.Sprintf("%d", int(cv))
	}
	return result
}

func NewCard(value CardValue, suit CardSuit) Card {
	return Card{value, suit}
}

func MaxCard(cards []Card, comparatorFunc CardComparator) Card {
	var maxCard Card
	for i := range cards {
		isMax := true
		for k := range cards {
			result := comparatorFunc(cards[i], cards[k])
			switch result {
			case 0, 1:
				continue
			case -1:
				isMax = false
			}
			if !isMax {
				break
			}
		}
		if isMax {
			maxCard = cards[i]
			break
		}
	}

	return maxCard
}

func CreateDeckOfCards() []Card {
	rand.Seed(time.Now().UnixMilli())
	dataPointCount := 10
	cardSuits := []string{"Club", "Diamond", "Heart", "Spade"}
	cardValues := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	cardsDeck := make([]Card, dataPointCount)
	for i := range cardsDeck {
		cardsDeck[i] = NewCard(CardValue(cardValues[rand.Intn(len(cardValues))]), CardSuit(ConvertSuitToEnum((cardSuits[rand.Intn(len(cardSuits))]))))
	}

	return cardsDeck
}
