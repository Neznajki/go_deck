package main

import (
	"deck.com/dto"
	"fmt"
	"os"
)

func main() {
	var deck = new(dto.Deck)

	err := deck.Init()
	if err != nil {
	}

	err = deck.Shuffle()
	if err != nil {
		handleError(err)
		os.Exit(2)
	}

	handCards, err := dto.DealCards(deck, new(dto.Hand), 5)
	if err != nil {
		handleError(err)
		os.Exit(3)
	}
	fmt.Println(fmt.Sprintf("left in old deck: %s", deck.GetAllCards()))

	fmt.Println(fmt.Sprintf("pulled to hand: %s", handCards.GetAllCards()))

}

func handleError(err error) {
	fmt.Println(err.Error())
}
