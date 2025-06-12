package main

import (
	"deck.com/dto"
	"fmt"
)

func main() {
	var deck = new(dto.Deck)

	deck.Init()

	deck.Shuffle()

	handCards := dto.DealCards(deck, new(dto.Hand), 5)

	fmt.Println(fmt.Sprintf("left in old deck: %s", deck.GetAllCards()))

	fmt.Println(fmt.Sprintf("pulled to hand: %s", handCards.GetAllCards()))
}
