package dto

import (
	"deck.com/dto/contract"
	"gotest.tools/v3/assert"
	"testing"
)

func TestNewCard(t *testing.T) {
	c := NewCard(contract.Spades, contract.Ace)

	if c == nil {
		t.Error("NewCard() returned nil")
	}
}

func TestCard_CardPrintValue(t *testing.T) {
	assert.Equal(t, NewCard(contract.Spades, contract.Ace).CardPrintValue(), "TA (♠)")
}

func TestCard_Equals(t *testing.T) {
	c := NewCard(contract.Spades, contract.Ace)

	assert.Equal(t, true, c.Equals(NewCard(contract.Spades, contract.Ace)))
	assert.Equal(t, false, c.Equals(NewCard(contract.Spades, contract.King)))
	assert.Equal(t, false, c.Equals(NewCard(contract.Hearts, contract.Ace)))
}
