package display_impl

import (
	"fmt"
	"strings"
)

type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

type stringDisplay struct {
	str   string
	width int
}

func NewStringDisplay(str string) *stringDisplay {
	return &stringDisplay{
		str:   str,
		width: len(str),
	}
}

func (sd *stringDisplay) RawOpen() {
	sd.printLine()
}

func (sd *stringDisplay) RawClose() {
	sd.printLine()
}

func (sd *stringDisplay) RawPrint() {
	fmt.Printf("|%s|\n", sd.str)
}

func (sd *stringDisplay) printLine() {
	fmt.Print("+")
	fmt.Print(strings.Repeat("-", sd.width))
	fmt.Print("+\n")
}
