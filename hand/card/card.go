package card

import (
	"poker_service/hand/suit"
)

type Card struct {
	Value     string
	FaceValue FaceValue
	Suit      suit.Suit
}

func (c *Card) Validate() bool {
	if len(c.Value) == 2 {
		hand := c.Value

		faceValue := string(hand[0])
		if val, exists := FaceValueMap[faceValue]; exists {
			c.FaceValue = val
		} else {
			return false
		}

		cSuit := string(hand[1])
		if val, exists := suit.SuitMap[cSuit]; exists {
			c.Suit = val
		} else {
			return false
		}

		return true
	}
	return false
}

type FaceValue uint8

const (
	Two FaceValue = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Joker
	Queen
	King
	Ace
)

var FaceValueMap = map[string]FaceValue{
	"2": Two,
	"3": Three,
	"4": Four,
	"5": Five,
	"6": Six,
	"7": Seven,
	"8": Eight,
	"9": Nine,
	"T": Ten,
	"J": Joker,
	"Q": Queen,
	"K": King,
	"A": Ace,
}

func (f FaceValue) IsRoyal() bool {
	if f < Ten || f > Ace {
		return false
	}
	return true
}
