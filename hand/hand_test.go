package hand

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"poker_service/hand/card"
	"poker_service/hand/suit"
)

func TestHand_GetRank(t *testing.T) {
	cases := map[string]struct {
		hand         Hand
		expectedRank Rank
	}{
		"Should fetch rank as 10 when card combination is Royal Flush": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "TC", FaceValue: card.Ten, Suit: suit.C},
					{Value: "KC", FaceValue: card.King, Suit: suit.C},
					{Value: "QC", FaceValue: card.Queen, Suit: suit.C},
					{Value: "JC", FaceValue: card.Joker, Suit: suit.C},
					{Value: "AC", FaceValue: card.Ace, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 10,
		},
		"Should fetch rank as 9 when card combination is Straight Flush": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "6C", FaceValue: card.Six, Suit: suit.C},
					{Value: "5C", FaceValue: card.Five, Suit: suit.C},
					{Value: "4C", FaceValue: card.Four, Suit: suit.C},
					{Value: "3C", FaceValue: card.Three, Suit: suit.C},
					{Value: "2C", FaceValue: card.Two, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 9,
		},
		"Should fetch rank as 8 when card combination is Four of a Kind": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "6C", FaceValue: card.Six, Suit: suit.C},
					{Value: "6H", FaceValue: card.Six, Suit: suit.H},
					{Value: "6D", FaceValue: card.Six, Suit: suit.D},
					{Value: "6S", FaceValue: card.Six, Suit: suit.S},
					{Value: "2C", FaceValue: card.Two, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 8,
		},
		"Should fetch rank as 7 when card combination is FullHouse": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "6C", FaceValue: card.Six, Suit: suit.C},
					{Value: "6H", FaceValue: card.Six, Suit: suit.H},
					{Value: "6D", FaceValue: card.Six, Suit: suit.D},
					{Value: "2S", FaceValue: card.Two, Suit: suit.S},
					{Value: "2C", FaceValue: card.Two, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 7,
		},
		"Should fetch rank as 6 when card combination is Flush": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "6C", FaceValue: card.Six, Suit: suit.C},
					{Value: "7C", FaceValue: card.Seven, Suit: suit.C},
					{Value: "5C", FaceValue: card.Five, Suit: suit.C},
					{Value: "3C", FaceValue: card.Three, Suit: suit.C},
					{Value: "2C", FaceValue: card.Two, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 6,
		},
		"Should fetch rank as 5 when card combination is Straight": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "6C", FaceValue: card.Six, Suit: suit.C},
					{Value: "7C", FaceValue: card.Seven, Suit: suit.C},
					{Value: "5C", FaceValue: card.Five, Suit: suit.C},
					{Value: "4S", FaceValue: card.Four, Suit: suit.S},
					{Value: "3C", FaceValue: card.Three, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 5,
		},
		"Should fetch rank as 4 when card combination is Three of a kind": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "7C", FaceValue: card.Seven, Suit: suit.C},
					{Value: "7D", FaceValue: card.Seven, Suit: suit.D},
					{Value: "7S", FaceValue: card.Seven, Suit: suit.S},
					{Value: "4S", FaceValue: card.Four, Suit: suit.S},
					{Value: "3C", FaceValue: card.Three, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 4,
		},
		"Should fetch rank as 3 when card combination is Two Pair": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "7C", FaceValue: card.Seven, Suit: suit.C},
					{Value: "7D", FaceValue: card.Seven, Suit: suit.D},
					{Value: "5S", FaceValue: card.Five, Suit: suit.S},
					{Value: "5C", FaceValue: card.Five, Suit: suit.C},
					{Value: "3C", FaceValue: card.Three, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 3,
		},
		"Should fetch rank as 2 when card combination is Pair": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "7C", FaceValue: card.Seven, Suit: suit.C},
					{Value: "7D", FaceValue: card.Seven, Suit: suit.D},
					{Value: "4S", FaceValue: card.Four, Suit: suit.S},
					{Value: "5C", FaceValue: card.Five, Suit: suit.C},
					{Value: "3C", FaceValue: card.Three, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 2,
		},
		"Should fetch rank as 1 when card combination is High card": {
			hand: Hand{
				Cards: []card.Card{
					{Value: "QC", FaceValue: card.Queen, Suit: suit.C},
					{Value: "7D", FaceValue: card.Seven, Suit: suit.D},
					{Value: "4S", FaceValue: card.Four, Suit: suit.S},
					{Value: "5C", FaceValue: card.Five, Suit: suit.C},
					{Value: "3C", FaceValue: card.Three, Suit: suit.C},
				},
				Rank: 0,
			},
			expectedRank: 1,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			rank := tc.hand.GetRank()

			assert.Equal(t, tc.expectedRank, rank)
		})
	}
}
