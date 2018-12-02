package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/bridge/display"
	"github.com/nekootoko3/go-design-patterns/bridge/display_impl"
)

func main() {
	d1 := display.NewDisplay(display_impl.NewStringDisplay("Hello, World!"))
	d2 := display.NewDisplay(display_impl.NewStringDisplay("Hello, Universe!"))
	d3 := display.NewDisplay(display_impl.NewStringDisplay("Hello, Yamaguchi!"))
	cd := display.NewCountDisplay(d3)
	rd := display.NewRandDisplay(d3)
	d1.Display()
	d2.Display()
	cd.Display()
	fmt.Println("## 5 times")
	cd.MultiDisplay(5)
	fmt.Println("## rand times ( 0 - 10 )")
	cd.RandDisplay()
	rd.RandDisplay()
}
