package gamehdl

type Play string

const (
	PlayHit   Play = "hit"
	PlayStand Play = "stand"
)

type Card struct {
	Number int    `json:"number"`
	Suit   string `json:"suit"`
}
type Hand []Card

func PlayCard(hands Hand) Play {
	if hands.Sum() < 5 {
		return PlayHit
	}

	return PlayStand
}

func (hand Hand) Sum() int {
	sum := 0
	for _, c := range hand {
		if c.Number > 10 {
			sum += 0
			continue
		}
		sum += c.Number
	}
	return sum % 10
}

func (hand Hand) IsSameSuit() bool {
	return hand[0].Suit == hand[1].Suit
}

func (hand Hand) IsSameNumber() bool {
	return hand[0].Number == hand[1].Number
}

func (hand Hand) IsHigh() bool {
	for _, card := range hand {
		if card.Number < 11 {
			return false
		}
	}

	return true
}
