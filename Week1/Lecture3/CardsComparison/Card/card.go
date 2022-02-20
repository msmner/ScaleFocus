package card

import "fmt"

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

func (cv CardValue) String() string {
	cardValues := [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14"}
	if cv < 2 || cv > 14 {
		return fmt.Sprintf("Card Value (%d)", int(cv))
	}
	return cardValues[cv+2]
}

func (cv CardValue) IsValid() bool {
	switch cv {
	case Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace:
		return true
	}
	return false
}

type CardSuit int

const (
	Club CardSuit = iota + 1
	Diamond
	Heart
	Spade
)

func (cs CardSuit) String() string {
	cardSuites := [...]string{"Club", "Diamond", "Heart", "Spade"}
	if cs < 1 || cs > 4 {
		return fmt.Sprintf("Card Value (%d)", int(cs))
	}
	return cardSuites[cs-1]
}

func (cs CardSuit) IsValid() bool {
	switch cs {
	case Club, Diamond, Heart, Spade:
		return true
	}
	return false
}

type Card struct {
	Value CardValue
	Suit  CardSuit
}
