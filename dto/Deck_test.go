package dto

import (
	"deck.com/dto/contract"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
	"testing"
)

func TestDeck_Init(t *testing.T) {
	d := new(Deck)
	d.Init()
	assert.Equal(t, true, d.initiated)
	assert.Equal(t, 52, len(d.cards))

	cmp.Panics(d.Init)
}

func TestDeck_Reset(t *testing.T) {
	d := new(Deck)
	cmp.Panics(d.Reset)
	d.Init()
	d.Pull(5)
	assert.Equal(t, 47, len(d.cards))
	d.Reset()
	assert.Equal(t, 52, len(d.cards))
}

func TestDeck_Pull(t *testing.T) {
	d := new(Deck)
	cmp.Panics(func() { d.Pull(5) })
	d.Init()
	cards := d.Pull(5)
	assert.Equal(t, 47, len(d.cards))
	assert.Equal(t, 5, len(cards))
}

func TestDeck_Search_Pull_Put(t *testing.T) {
	d := new(Deck)
	d.Init()

	assert.Equal(t, d.Search(contract.Spades, contract.Ace).CardPrintValue(), "TA (♠)")

	d.Pick(contract.Spades, contract.Ace)
	cmp.Nil(d.Search(contract.Spades, contract.Ace))

	d.Put(contract.Spades, contract.Ace)
	assert.Equal(t, d.Search(contract.Spades, contract.Ace).CardPrintValue(), "TA (♠)")
	cmp.Panics(func() { d.Put(contract.Spades, contract.Ace) })
}

func TestDeck_GetAllCards(t *testing.T) {
	d := new(Deck)
	d.Init()

	assert.Equal(t, d.GetAllCards(), "\n\ncard: 02 (♣)\ncard: 03 (♣)\ncard: 04 (♣)\ncard: 05 (♣)\ncard: 06 (♣)\ncard: 07 (♣)\ncard: 08 (♣)\ncard: 09 (♣)\ncard: 10 (♣)\ncard: TJ (♣)\ncard: TQ (♣)\ncard: TK (♣)\ncard: TA (♣)\ncard: 02 (♦)\ncard: 03 (♦)\ncard: 04 (♦)\ncard: 05 (♦)\ncard: 06 (♦)\ncard: 07 (♦)\ncard: 08 (♦)\ncard: 09 (♦)\ncard: 10 (♦)\ncard: TJ (♦)\ncard: TQ (♦)\ncard: TK (♦)\ncard: TA (♦)\ncard: 02 (♥)\ncard: 03 (♥)\ncard: 04 (♥)\ncard: 05 (♥)\ncard: 06 (♥)\ncard: 07 (♥)\ncard: 08 (♥)\ncard: 09 (♥)\ncard: 10 (♥)\ncard: TJ (♥)\ncard: TQ (♥)\ncard: TK (♥)\ncard: TA (♥)\ncard: 02 (♠)\ncard: 03 (♠)\ncard: 04 (♠)\ncard: 05 (♠)\ncard: 06 (♠)\ncard: 07 (♠)\ncard: 08 (♠)\ncard: 09 (♠)\ncard: 10 (♠)\ncard: TJ (♠)\ncard: TQ (♠)\ncard: TK (♠)\ncard: TA (♠)\n")
}

func TestDeck_Shuffle(t *testing.T) {
	d := new(Deck)
	d.Init()

	beforeShuffle := d.GetAllCards()
	d.Shuffle()
	afterShuffle := d.GetAllCards()

	if beforeShuffle == afterShuffle {
		t.Error("cards should change order")
	}
}
