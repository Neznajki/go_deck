package contract

import (
	"testing"
)
import "gotest.tools/v3/assert"

func TestNewCardType(t *testing.T) {
	if NewCardType(Clubs) == nil {
		t.Error("NewCardType(Clubs) returned nil")
	}
}

func TestCardType_CardType(t *testing.T) {
	assert.Equal(t, NewCardType(Clubs).CardType(), "♣")
	assert.Equal(t, NewCardType(Diamonds).CardType(), "♦")
	assert.Equal(t, NewCardType(Hearts).CardType(), "♥")
	assert.Equal(t, NewCardType(Spades).CardType(), "♠")
}
