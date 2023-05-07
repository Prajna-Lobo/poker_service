package hand

import (
	"errors"

	"poker_service/hand/card"
)

type Hand struct {
	Cards []card.Card
	Rank  float32
}

func (h *Hand) Validate() (bool, error) {
	for i, c := range h.Cards {
		if !c.Validate() {
			return false, errors.New("not a valid hand")
		}
		h.Cards[i] = c
	}

	return true, nil
}

func (h Hand) GetRank() Rank {
	hands := NewSortHands(h)

	straight := hands.IsStraight()
	flush := hands.IsSameSuite()
	royal := hands.AreAllRoyal()

	if flush && straight && royal {
		return RoyalFlush
	}
	if flush && straight {
		return StraightFlush
	}
	if hands.HasAnyOfAKind(4) {
		return FourOfAKind
	}
	if hands.HasFullHouse() {
		return FullHouse
	}
	if flush {
		return Flush
	}
	if straight {
		return Straight
	}
	if hands.HasAnyOfAKind(3) {
		return ThreeOfAKind
	}
	if hands.HasTwoPairs() {
		return TwoPairs
	}
	if hands.HasPair() {
		return Pair
	}

	return HighCard
}
