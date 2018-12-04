package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nekootoko3/go-design-patterns/strategy/player"
	"github.com/nekootoko3/go-design-patterns/strategy/strategy"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i1 := int64(rand.Int())
	i2 := int64(rand.Int())
	neko := player.NewPlayer("nekootoko3", strategy.NewProbStrategy(i1))
	inu := player.NewPlayer("inuotoko3", strategy.NewWinningStrategy(i2))
	fmt.Println(i1, i2)
	fmt.Println(neko, inu)
}
