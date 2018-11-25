package template_method

import "fmt"

type printer interface {
	open()
	print()
	close()
}

type abstractDisplay struct {
	printer
}

func NewAbstractDisplay(printer printer) *abstractDisplay {
	return &abstractDisplay{printer}
}

func (ad *abstractDisplay) Display() {
	ad.open()
	for i := 0; i < 5; i++ {
		ad.print()
	}
	ad.close()
}

type charDisplay struct {
	str string
}

func NewCharDisplay(str string) *charDisplay {
	return &charDisplay{str: str}
}

func (cd *charDisplay) open() {
	fmt.Print("<<")
}
func (cd *charDisplay) print() {
	fmt.Print(cd.str)
}
func (cd *charDisplay) close() {
	fmt.Print(">>\n")
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

func (sd *stringDisplay) open() {
	sd.printLine()
}
func (sd *stringDisplay) print() {
	fmt.Println("|", sd.str, "|")
}
func (sd *stringDisplay) close() {
	sd.printLine()
}
func (sd *stringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n")
}
