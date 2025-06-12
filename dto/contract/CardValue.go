package contract

type CVT int

type CardValueInterface interface {
	CardValue() string
}

type CardValue struct {
	cV CVT
}

const (
	Two CVT = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	FirstValue = Two
	LastValue  = Ace
)

var existingValues = map[CVT]string{
	Two:   "02",
	Three: "03",
	Four:  "04",
	Five:  "05",
	Six:   "06",
	Seven: "07",
	Eight: "08",
	Nine:  "09",
	Ten:   "10",
	Jack:  "TJ",
	Queen: "TQ",
	King:  "TK",
	Ace:   "TA",
}

func (cv CardValue) CardValue() string {
	return existingValues[cv.cV]
}

func NewCardValue(cVT CVT) CardValueInterface {
	return &CardValue{cVT}
}
