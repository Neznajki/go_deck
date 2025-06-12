package dto

import (
	"deck.com/dto/contract"
	"fmt"
)

type CardInterface interface {
	CardPrintValue() string
	Equals(nci CardInterface) bool
}

type Card struct {
	cT contract.CardTypeInterface
	cV contract.CardValueInterface
}

func (c Card) CardPrintValue() string {
	return fmt.Sprintf("%s (%s)", c.cV.CardValue(), c.cT.CardType())
}

func NewCard(cardType contract.CTT, cardValue contract.CVT) CardInterface {
	return &Card{cT: contract.NewCardType(cardType), cV: contract.NewCardValue(cardValue)}
}

func (c Card) Equals(nci CardInterface) bool {
	var nc = nci.(*Card)
	return c.cV.CardValue() == nc.cV.CardValue() && c.cT.CardType() == nc.cT.CardType()
}
