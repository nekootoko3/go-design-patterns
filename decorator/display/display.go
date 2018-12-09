package display

import "fmt"

type Printer interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) string
}

type display struct {
	Printer
}

func NewDisplay(printer Printer) *display {
	return &display{printer}
}

func (d *display) Show() {
	for i := 0; i < d.GetRows(); i++ {
		fmt.Println(d.GetRowText(i))
	}
}

func Show(p Printer) {
	for i := 0; i < p.GetRows(); i++ {
		fmt.Println(p.GetRowText(i))
	}
}
