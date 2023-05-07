package gametable

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"poker_service/hand"
	"poker_service/hand/card"
	"poker_service/hand/suit"
)

func TestGameTable_Distribute(t *testing.T) {
	cases := map[string]struct {
		hands       []hand.Hand
		expectedErr error
	}{
		"Should validate and distribute hands to players": {
			hands: []hand.Hand{
				{
					Cards: []card.Card{
						{Value: "9C", FaceValue: card.Nine, Suit: suit.C},
						{Value: "9D", FaceValue: card.Nine, Suit: suit.D},
						{Value: "8D", FaceValue: card.Eight, Suit: suit.D},
						{Value: "7C", FaceValue: card.Seven, Suit: suit.C},
						{Value: "3C", FaceValue: card.Three, Suit: suit.C},
					},
				},
				{
					Cards: []card.Card{
						{Value: "2S", FaceValue: card.Two, Suit: suit.S},
						{Value: "KD", FaceValue: card.King, Suit: suit.D},
						{Value: "TH", FaceValue: card.Ten, Suit: suit.H},
						{Value: "9H", FaceValue: card.Nine, Suit: suit.H},
						{Value: "9C", FaceValue: card.Nine, Suit: suit.C},
					},
				},
			},
		},
	}

	gameTable := NewGameTable(2)
	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			err := gameTable.Distribute(tc.hands)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestGameTable_FindAndSetWinner(t *testing.T) {
	cases := map[string]struct {
		hands               []hand.Hand
		expectedWinningHand int
	}{
		"Should find winner and increase the winning hands": {
			hands: []hand.Hand{
				{
					Cards: []card.Card{
						{Value: "9C", FaceValue: card.Nine, Suit: suit.C},
						{Value: "9D", FaceValue: card.Nine, Suit: suit.D},
						{Value: "8D", FaceValue: card.Eight, Suit: suit.D},
						{Value: "7C", FaceValue: card.Seven, Suit: suit.C},
						{Value: "3C", FaceValue: card.Three, Suit: suit.C},
					},
				},
				{
					Cards: []card.Card{
						{Value: "2S", FaceValue: card.Two, Suit: suit.S},
						{Value: "KD", FaceValue: card.King, Suit: suit.D},
						{Value: "TH", FaceValue: card.Ten, Suit: suit.H},
						{Value: "9H", FaceValue: card.Nine, Suit: suit.H},
						{Value: "9C", FaceValue: card.Nine, Suit: suit.C},
					},
				},
			},
		},
	}

	gameTable := NewGameTable(2)
	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			gameTable.FindWinner()

			assert.Equal(t, tc.expectedWinningHand, gameTable.Players[1].WinningHands)
		})
	}
}
