package cardgame

import "testing"

func TestCompareCards(t *testing.T) {
	//Arrange
	firstCard1 := Card{CardValue(3), CardSuit(2)}
	secondCard1 := Card{CardValue(2), CardSuit(2)}
	firstCardMinus1 := Card{CardValue(2), CardSuit(2)}
	secondCardMinus1 := Card{CardValue(3), CardSuit(2)}
	firstCardEqualValueFirstSuit := Card{CardValue(2), CardSuit(3)}
	secondCardEqualValueFirstSuit := Card{CardValue(2), CardSuit(2)}
	firstCardEqualValueSecondSuit := Card{CardValue(2), CardSuit(2)}
	secondCardEqualValueSecondSuit := Card{CardValue(2), CardSuit(3)}
	firstCardEqual := Card{CardValue(2), CardSuit(2)}
	secondCardEqual := Card{CardValue(2), CardSuit(2)}

	cards := make([]Card, 0, 2)
	cards = append(cards, firstCard1, secondCard1)

	t.Run("Test Max Card Function", func(t *testing.T) {
		expectedCard := firstCard1

		//Act
		actualCard := MaxCard(cards)

		//Assert
		if expectedCard != actualCard {
			t.Errorf("Failed max card test expected %v, but got %v", expectedCard, actualCard)
		}
	})

	t.Run("Test Compare Cards Function", func(t *testing.T) {
		expectedResult1 := 1
		expectedResultMinus1 := -1
		expectedResultEqualValue1 := 1
		expectedResultEqualValueMinus1 := -1
		expectedResultEqual := 0

		//Act
		actualResult1 := CompareCards(firstCard1, secondCard1)
		actualResultMinus1 := CompareCards(firstCardMinus1, secondCardMinus1)
		actualResultEqualValue1 := CompareCards(firstCardEqualValueFirstSuit, secondCardEqualValueFirstSuit)
		actualResultEqualValueMinus1 := CompareCards(firstCardEqualValueSecondSuit, secondCardEqualValueSecondSuit)
		actualResultEqual := CompareCards(firstCardEqual, secondCardEqual)

		//Assert
		if expectedResult1 != actualResult1 {
			t.Errorf("Failed positive 1 test expected %d, but got %d", expectedResultMinus1, actualResultMinus1)
		}

		if expectedResultMinus1 != actualResultMinus1 {
			t.Errorf("Failed negative 1 test expected %d, but got %d", expectedResultMinus1, actualResult1)
		}

		if expectedResultEqualValue1 != actualResultEqualValue1 {
			t.Errorf("Failed positive 1 test expected %d, but got %d", expectedResultEqualValue1, actualResultEqualValue1)
		}

		if expectedResultEqualValueMinus1 != actualResultEqualValueMinus1 {
			t.Errorf("Failed positive 1 test expected %d, but got %d", expectedResultEqualValueMinus1, actualResultEqualValueMinus1)
		}

		if expectedResultEqual != actualResultEqual {
			t.Errorf("Failed positive 1 test expected %d, but got %d", expectedResultEqual, actualResultEqual)
		}
	})
}
