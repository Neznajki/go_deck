package contract

import (
	"fmt"
	"testing"
)
import "gotest.tools/v3/assert"

func TestNewCardType(t *testing.T) {
	cardType := NewCardType(Clubs)
	if cardType == nil {
		t.Error("NewCardType(Clubs) returned nil")
	}

	assert.Equal(t, fmt.Sprintf("%+v", cardType), "&{CTT:0}")
}

func TestCardType_CardType(t *testing.T) {
	assert.Equal(t, NewCardType(Clubs).CardType(), "♣")
	assert.Equal(t, NewCardType(Diamonds).CardType(), "♦")
	assert.Equal(t, NewCardType(Hearts).CardType(), "♥")
	assert.Equal(t, NewCardType(Spades).CardType(), "♠")
}
