package contract

type CTT int

type CardTypeInterface interface {
	CardType() string
}

type CardType struct {
	cT CTT
}

const (
	Clubs CTT = iota
	Diamonds
	Hearts
	Spades
	FirstType = Clubs
	LastType  = Spades
)

var existingTypes = map[CTT]string{
	Clubs:    "♣",
	Diamonds: "♦",
	Hearts:   "♥",
	Spades:   "♠",
}

func (ct CardType) CardType() string {
	return existingTypes[ct.cT]
}

func NewCardType(cTT CTT) CardTypeInterface {
	return &CardType{cT: cTT}
}
