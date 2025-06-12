package dto

import (
	"deck.com/dto/contract"
)

type Hand Deck

func (h Hand) Init() error {
	panic("implement me")
}

func (h Hand) Reset() error {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Shuffle() error {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Pull(size int) (DeckInterface, error) {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Pick(cardType contract.CTT, cardValue contract.CVT) (CardInterface, error) {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Search(cardType contract.CTT, cardValue contract.CVT) (CardInterface, error) {
	//TODO implement me
	panic("implement me")
}

func (h Hand) Put(cardType contract.CTT, cardValue contract.CVT) error {
	//TODO implement me
	panic("implement me")
}

func (h Hand) GetAllCards() string {
	return getCardsString(h.cards)
}

func DealCards(d *Deck, hand *Hand, amount int) (DeckInterface, error) {
	hand.initiated = true
	cards, err := d.Pull(amount)

	if err != nil {
		return nil, err
	}

	hand.cards = append(hand.cards, cards...)

	return hand, nil
}
