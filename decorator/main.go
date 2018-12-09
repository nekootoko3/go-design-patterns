package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/decorator/border"
	"github.com/nekootoko3/go-design-patterns/decorator/display"
)

func main() {
	d1 := border.NewStringDisplay("nekootoko3")
	display.Show(d1)
	d2 := border.NewSideBorder(d1, "/")
	display.Show(d2)
	d3 := border.NewFullBorder(d2)
	display.Show(d3)
	fmt.Println("")
	d4 := border.NewUpDownBorder(d3, "*")
	display.Show(d4)
	fmt.Println("")
	d5 := border.NewFullBorder(d4)
	display.Show(d5)
	fmt.Println("")

	d6 := border.NewUpDownBorder(d1, "*")
	display.Show(d6)

	d7 := border.NewMultiStringDisplay()
	border.MultiAdd(d7, "neko")
	border.MultiAdd(d7, "inu")
	border.MultiAdd(d7, "pandakopanda")
	display.Show(d7)
	d8 := border.NewUpDownBorder(d7, "*")
	display.Show(d8)
	d9 := border.NewFullBorder(d8)
	display.Show(d9)
}
