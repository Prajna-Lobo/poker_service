package suit

import "errors"

type Suit uint8

const (
	// D for Diamonds
	D Suit = iota + 1
	//H for hearts
	H
	// S for spade
	S
	// C for Clubs
	C
)

var SuitMap = map[string]Suit{
	"D": D,
	"H": H,
	"S": S,
	"C": C,
}

func (s Suit) Validate() (bool, error) {
	if s > C || s < D {
		return false, errors.New("invalid suit")
	}
	return true, nil
}
