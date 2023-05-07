package player_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"poker_service/hand"
	"poker_service/hand/card"
	"poker_service/hand/suit"
	"poker_service/player"
)

func TestPlayer_Validate(t *testing.T) {
	cases := map[string]struct {
		hand           hand.Hand
		expectedOutput bool
		expectedErr    error
	}{
		"Successfully validate player hand": {
			hand: hand.Hand{
				Cards: []card.Card{
					{Value: "9C", FaceValue: card.Nine, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedOutput: true,
		},
		"Should validate hand and return false for invalid card": {
			hand: hand.Hand{
				Cards: []card.Card{
					{Value: "XS", FaceValue: card.Nine, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedOutput: false,
			expectedErr:    errors.New("not a valid hand"),
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			player1 := player.Player{
				ID:           1,
				Hand:         tc.hand,
				WinningHands: 0,
			}

			validate, err := player1.Validate()

			assert.Equal(t, tc.expectedOutput, validate)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestPlayer_IncrementScore(t *testing.T) {
	cases := map[string]struct {
		hand           hand.Hand
		expectedOutput int
	}{
		"Successfully increase the score of a player": {
			hand: hand.Hand{
				Cards: []card.Card{
					{Value: "9C", FaceValue: card.Nine, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedOutput: 1,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			player1 := player.Player{
				ID:           1,
				Hand:         tc.hand,
				WinningHands: 0,
			}

			player1.IncrementScore()

			assert.Equal(t, tc.expectedOutput, player1.WinningHands)
		})
	}
}
