package gametable

import (
	"fmt"

	"poker_service/hand"
	"poker_service/player"
)

type GameTable struct {
	Players []player.Player
}

func NewGameTable(numOfPlayers int) *GameTable {
	return &GameTable{
		Players: initializePlayers(numOfPlayers),
	}
}

type IGameTable interface {
	Distribute(handsPerPlayer []hand.Hand)
}

func initializePlayers(numOfPlayers int) []player.Player {
	players := make([]player.Player, numOfPlayers, numOfPlayers)
	for i := 0; i < numOfPlayers; i++ {
		var p player.Player
		p.ID = i + 1

		players[i] = p
	}

	return players
}

func (g *GameTable) Distribute(handsPerPlayer []hand.Hand) error {
	for i := 0; i < len(g.Players); i++ {
		isValid, err := handsPerPlayer[i].Validate()
		if err != nil || !isValid {
			return err
		}
		g.Players[i].Hand = handsPerPlayer[i]
	}
	return nil
}

func (g *GameTable) FindWinner() {
	var maxRanks hand.Rank
	var index int

	for i := 0; i < len(g.Players); i++ {
		cards := g.Players[i].Hand.Cards
		rank := g.Players[i].Hand.GetRank()

		if maxRanks == rank && i >= 0 {
			index = g.breakATie(index, i, len(cards)-1)
		}

		if maxRanks < rank {
			maxRanks = rank
			index = i
		}
	}

	if len(g.Players) > 0 && len(g.Players) > index {
		g.Players[index].IncrementScore()
	}
}

func (g *GameTable) Print() {
	for i := 0; i < len(g.Players); i++ {
		fmt.Printf("Player %d : %d\n", g.Players[i].ID, g.Players[i].WinningHands)
	}
}

func (g *GameTable) breakATie(prevPlayer, nextPlayer int, cardIndex int) int {
	if g.Players != nil {
		player1 := g.Players[prevPlayer].Hand.Cards
		player2 := g.Players[nextPlayer].Hand.Cards

		if cardIndex > 0 {
			if player1[cardIndex].FaceValue > player2[cardIndex].FaceValue {
				return prevPlayer
			}

			if player1[cardIndex].FaceValue < player2[cardIndex].FaceValue {
				return nextPlayer
			}

			g.breakATie(prevPlayer, nextPlayer, cardIndex-1)
		}
	}
	return prevPlayer
}
