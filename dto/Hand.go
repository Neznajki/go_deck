package dto

import (
	"deck.com/dto/contract"
)

type Hand Deck

func (Hand) Init() {
	panic("implement me")
}

func (Hand) Reset() {
	//TODO implement me
	panic("implement me")
}

func (Hand) Shuffle() {
	//TODO implement me
	panic("implement me")
}

func (Hand) Pull(int) DeckInterface {
	//TODO implement me
	panic("implement me")
}

func (Hand) Pick(contract.CTT, contract.CVT) CardInterface {
	//TODO implement me
	panic("implement me")
}

func (Hand) Search(contract.CTT, contract.CVT) CardInterface {
	//TODO implement me
	panic("implement me")
}

func (Hand) Put(contract.CTT, contract.CVT) {
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
