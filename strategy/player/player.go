package player

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/strategy/hand"
	"github.com/nekootoko3/go-design-patterns/strategy/strategy"
)

type player struct {
	name      string
	strategy  strategy.Strategy
	winCount  int
	loseCount int
	gameCount int
}

func NewPlayer(name string, strategy strategy.Strategy) *player {
	return &player{
		name:      name,
		strategy:  strategy,
		winCount:  0,
		loseCount: 0,
	}
}

func (p *player) NextHand() *hand.Hand {
	return p.strategy.NextHand()
}

func (p *player) Win() {
	p.strategy.Study(true)
	p.winCount++
	p.gameCount++
}

func (p *player) Lose() {
	p.strategy.Study(false)
	p.loseCount++
	p.gameCount++
}

func (p *player) ToString() string {
	return fmt.Sprintf("[%s] win: %d, lose: %d", p.name, p.winCount, p.loseCount)
}
