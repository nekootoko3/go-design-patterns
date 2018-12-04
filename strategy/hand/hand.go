package hand

type HandValue int

const (
	GUU HandValue = iota
	CHO
	PAA
)

var hands [3]*Hand

type Hand struct {
	handValue HandValue
}

func GetHand(handValue HandValue) *Hand {
	if hands[handValue] == nil {
		hands[handValue] = &Hand{handValue: handValue}
	}
	return hands[handValue]
}

func (h *Hand) IsStrongerThan(v *Hand) bool {
	return h.fight(v) == 1
}

func (h *Hand) fight(v *Hand) int {
	if h == v {
		return 0
	} else if ((h.handValue + 1) % 3) == v.handValue {
		return 1
	} else {
		return 2
	}
}
