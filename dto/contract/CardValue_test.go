package contract

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestNewCardValue(t *testing.T) {
	cardValue := NewCardValue(Ace)
	if cardValue == nil {
		t.Error("NewCardValue(Ace) returned nil")
	}

	assert.Equal(t, fmt.Sprintf("%+v", cardValue), "&{cV:12}")
}

func TestCardValue_CardValue(t *testing.T) {
	assert.Equal(t, NewCardValue(Two).CardValue(), "02")
	assert.Equal(t, NewCardValue(Three).CardValue(), "03")
	assert.Equal(t, NewCardValue(Four).CardValue(), "04")
	assert.Equal(t, NewCardValue(Five).CardValue(), "05")
	assert.Equal(t, NewCardValue(Six).CardValue(), "06")
	assert.Equal(t, NewCardValue(Seven).CardValue(), "07")
	assert.Equal(t, NewCardValue(Eight).CardValue(), "08")
	assert.Equal(t, NewCardValue(Nine).CardValue(), "09")
	assert.Equal(t, NewCardValue(Ten).CardValue(), "10")
	assert.Equal(t, NewCardValue(Jack).CardValue(), "TJ")
	assert.Equal(t, NewCardValue(Queen).CardValue(), "TQ")
	assert.Equal(t, NewCardValue(King).CardValue(), "TK")
	assert.Equal(t, NewCardValue(Ace).CardValue(), "TA")
}
