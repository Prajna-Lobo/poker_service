package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"poker_service/gametable"
	"poker_service/hand"
	"poker_service/hand/card"
)

var NumOfPlayers = 2
var NumOfCardsPerHand = 5

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter 2 set of 5 cards for 2 players")

	gt := gametable.NewGameTable(NumOfPlayers)

	for scanner.Scan() {
		pHand := scanner.Text()

		// break the scanner steam when enter is pressed
		if len(pHand) == 0 {
			break
		}

		hands := formHands(pHand)

		err := gt.Distribute(hands)
		if err != nil {
			log.Print("distribution error ", err.Error())
			return
		}

		gt.FindWinner()
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal("scanner error ", err.Error())
	}

	gt.Print()
}

func formHands(pHand string) []hand.Hand {
	var hands []hand.Hand

	cardValues := strings.Split(pHand, " ")

	for i := 0; i < NumOfPlayers; i++ {
		var h hand.Hand

		start := i * NumOfCardsPerHand
		end := (i + 1) * NumOfCardsPerHand

		if end > len(cardValues) {
			end = len(cardValues)
		}

		handValues := cardValues[start:end]

		var cards []card.Card
		for k := 0; k < len(handValues); k++ {
			var c card.Card
			c.Value = handValues[k]

			cards = append(cards, c)
		}

		h.Cards = cards
		hands = append(hands, h)
	}

	return hands
}