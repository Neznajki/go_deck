package dto

import (
	"deck.com/dto/contract"
)

type Hand Deck

func (h Hand) Init() {
	panic("implement me")
}

func (h Hand) Reset() {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Shuffle() {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Pull(size int) DeckInterface {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Pick(cardType contract.CTT, cardValue contract.CVT) CardInterface {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Search(cardType contract.CTT, cardValue contract.CVT) CardInterface {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Put(cardType contract.CTT, cardValue contract.CVT) {
	//TODO implement me
	panic("implement me")
}

func (h Hand) GetAllCards() string {
	return getCardsString(h.cards)
}

func DealCards(d *Deck, hand *Hand, amount int) DeckInterface {
	hand.initiated = true
	cards := d.Pull(amount)

	hand.cards = append(hand.cards, cards...)

	return hand
}
