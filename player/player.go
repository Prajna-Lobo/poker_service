package player

import (
	"poker_service/hand"
)

type Player struct {
	ID           int
	Hand         hand.Hand
	WinningHands int
}

type IPlayer interface {
	Validate() (bool, error)
	IncrementScore()
	GetWinningHands() int
}

func (p *Player) Validate() (bool, error) {
	isValidHand, err := p.Hand.Validate()
	if err != nil || !isValidHand {
		return false, err
	}
	return true, nil
}

func (p *Player) IncrementScore() {
	p.WinningHands += 1
}
