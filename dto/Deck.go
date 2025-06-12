package dto

import (
	"deck.com/dto/contract"
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type DeckInterface interface {
	Init()
	Reset()
	Shuffle()
	Pull(size int) DeckInterface
	Pick(cardType contract.CTT, cardValue contract.CVT) CardInterface
	Search(cardType contract.CTT, cardValue contract.CVT) CardInterface
	Put(cardType contract.CTT, cardValue contract.CVT)
	GetAllCards() string
}

type Deck struct {
	initiated bool
	cards     []CardInterface
}

func (d *Deck) Init() {
	if d.initiated {
		fmt.Println(errors.New("already initiated"))
		os.Exit(1)
	}
	d.initiated = true

	recreateDeck(d)
}

func (d *Deck) Reset() {
	if !d.initiated {
		fmt.Println(errors.New("already initiated"))
		os.Exit(1)
	}
	d.initiated = true

	recreateDeck(d)
}

func (d *Deck) Shuffle() {
	if !d.initiated {
		fmt.Println(errors.New("already initiated"))
		os.Exit(1)
	}

	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Pull(size int) []CardInterface {
	if len(d.cards) < size {
		fmt.Println(errors.New("already initiated"))
		os.Exit(1)
	}
	result := d.cards[:size]

	d.cards = d.cards[size:]

	return result
}

func remove(d *Deck, card CardInterface) {
	for i, c := range d.cards {
		if c == card {
			l := len(d.cards)
			d.cards[i] = d.cards[l-1]
			d.cards = d.cards[:l-1]
		}
	}
}

func (d *Deck) Pick(cardType contract.CTT, cardValue contract.CVT) CardInterface {
	var result = d.Search(cardType, cardValue)

	remove(d, result)

	return result
}

func (d *Deck) Search(cardType contract.CTT, cardValue contract.CVT) (CardInterface, error) {
	newCard := NewCard(cardType, cardValue)
	for _, card := range d.cards {
		if card.Equals(newCard) {
			return card, nil
		}
	}

	fmt.Println(errors.New("card already withdrawn"))
	os.Exit(1)
}

func (d *Deck) Put(cardType contract.CTT, cardValue contract.CVT) error {
	var _, err = d.Search(cardType, cardValue)
	var card = NewCard(cardType, cardValue)
	if err == nil {
		return errors.New(fmt.Sprintf("Deck already have %s card", card.CardPrintValue()))
	}

	d.cards = append(d.cards, card)

	return nil
}

func recreateDeck(d *Deck) {
	d.cards = []CardInterface{}

	for i := contract.FirstType; i <= contract.LastType; i++ {
		for k := contract.FirstValue; k <= contract.LastValue; k++ {
			d.cards = append(d.cards, NewCard(i, k))
		}
	}

	//d.cards = append(d.cards, createAllTypes(contract.Three)...)
	//d.cards = append(d.cards, createAllTypes(contract.Four)...)
	//d.cards = append(d.cards, createAllTypes(contract.Five)...)
	//d.cards = append(d.cards, createAllTypes(contract.Six)...)
	//d.cards = append(d.cards, createAllTypes(contract.Seven)...)
	//d.cards = append(d.cards, createAllTypes(contract.Eight)...)
	//d.cards = append(d.cards, createAllTypes(contract.Nine)...)
	//d.cards = append(d.cards, createAllTypes(contract.Ten)...)
	//d.cards = append(d.cards, createAllTypes(contract.Jack)...)
	//d.cards = append(d.cards, createAllTypes(contract.Queen)...)
	//d.cards = append(d.cards, createAllTypes(contract.King)...)
	//d.cards = append(d.cards, createAllTypes(contract.Ace)...)
}

//func createAllTypes(cvt contract.CVT) []CardInterface {
//	return []CardInterface{
//		NewCard(contract.Clubs, cvt),
//		NewCard(contract.Diamonds, cvt),
//		NewCard(contract.Hearts, cvt),
//		NewCard(contract.Spades, cvt),
//	}
//}

func (d *Deck) GetAllCards() string {
	return getCardsString(d.cards)
}

func getCardsString(cards []CardInterface) string {
	var result = fmt.Sprintln(fmt.Sprintln(""))

	for _, card := range cards {
		result += fmt.Sprintln(fmt.Sprintf("card: %s", card.CardPrintValue()))
	}

	return result
}
