package strategy

import (
	"math/rand"

	"github.com/nekootoko3/go-design-patterns/strategy/hand"
)

// Strategy  represents represents strategy that decide next hand
type Strategy interface {
	NextHand() *hand.Hand
	Study(win bool)
}

// winning strategy

type winningStrategy struct {
	random   *rand.Rand
	won      bool
	prevHand *hand.Hand
}

func NewWinningStrategy(seed int64) *winningStrategy {
	return &winningStrategy{
		random: rand.New(rand.NewSource(seed)),
		won:    false,
	}
}

func (s *winningStrategy) NextHand() *hand.Hand {
	if !s.won {
		handValue := hand.HandValue(s.random.Intn(3))
		s.prevHand = hand.GetHand(handValue)
	}
	return s.prevHand
}

func (s *winningStrategy) Study(win bool) {
	s.won = win
}

// prob strategy

type probStrategy struct {
	random           *rand.Rand
	prevHandValue    hand.HandValue
	currentHandValue hand.HandValue
	history          [][]int
}

func NewProbStrategy(seed int64) *probStrategy {
	hist := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	return &probStrategy{
		random:  rand.New(rand.NewSource(seed)),
		history: hist,
	}
}

func (s *probStrategy) NextHand() *hand.Hand {
	bet := s.random.Intn(s.getSum(s.currentHandValue))
	hv := 0
	if bet < hv {
		hv = 0
	} else if bet < s.history[s.currentHandValue][0]+s.history[s.currentHandValue][1] {
		hv = 1
	} else {
		hv = 2
	}
	s.prevHandValue = s.currentHandValue
	s.currentHandValue = hand.HandValue(hv)
	return hand.GetHand(s.currentHandValue)
}

func (s *probStrategy) Study(win bool) {
	// 今度やる
}

func (s *probStrategy) getSum(hv hand.HandValue) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += s.history[hv][i]
	}
	return sum
}
