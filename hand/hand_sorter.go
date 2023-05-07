package hand

import (
	"sort"

	"poker_service/hand/card"
	"poker_service/hand/suit"
)

type Rank float32

const (
	HighCard Rank = iota + 1
	Pair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type SortHands struct {
	hand             Hand
	FirstCardSuite   suit.Suit
	HighestFaceValue card.FaceValue
	FaceValues       map[card.FaceValue]int
	Suits            map[suit.Suit]int
}

func NewSortHands(h Hand) *SortHands {
	sortHands := SortHands{
		hand:       h,
		FaceValues: make(map[card.FaceValue]int, len(h.Cards)),
		Suits:      make(map[suit.Suit]int, len(h.Cards)),
	}
	sortHands.setEvaluationData()

	return &sortHands
}

func (s *SortHands) setEvaluationData() {
	sort.Slice(s.hand.Cards, func(i, j int) bool {
		return s.hand.Cards[i].FaceValue < s.hand.Cards[j].FaceValue
	})

	sortedCards := s.hand.Cards

	for i, c := range sortedCards {
		if i == 0 {
			s.FirstCardSuite = c.Suit
		}
		if i == len(sortedCards)-1 {
			s.HighestFaceValue = c.FaceValue
		}

		s.FaceValues[c.FaceValue] += 1
		s.Suits[c.Suit] += 1
	}
}

func (s *SortHands) IsSameSuite() bool {
	if len(s.Suits) == 1 {
		return true
	}
	return false
}

func (s *SortHands) IsStraight() bool {
	if len(s.FaceValues) != 5 {
		return false
	}

	for i := 0; i < len(s.hand.Cards)-1; i++ {
		if s.hand.Cards[i+1].FaceValue != s.hand.Cards[i].FaceValue+1 {
			return false
		}
	}
	return true
}

func (s *SortHands) HasPair() bool {
	if len(s.FaceValues) == 5 {
		return false
	}
	return true
}

func (s *SortHands) HasTwoPairs() bool {
	if len(s.FaceValues) != 3 {
		return false
	}

	return true
}

func (s *SortHands) HasAnyOfAKind(count int) bool {
	for _, counter := range s.FaceValues {
		if counter == count {
			return true
		}
	}
	return false
}

func (s *SortHands) HasFullHouse() bool {
	if len(s.FaceValues) > 2 {
		return false
	}
	return true
}

func (s *SortHands) AreAllRoyal() bool {
	if len(s.FaceValues) != 5 {
		return false
	}

	for faceValue := range s.FaceValues {
		if !faceValue.IsRoyal() {
			return false
		}
	}
	return true
}
