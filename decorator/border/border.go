package border

import (
	"fmt"
	"strings"

	"github.com/nekootoko3/go-design-patterns/decorator/display"
)

type stringDisplay struct {
	str string
}

func NewStringDisplay(str string) display.Printer {
	return &stringDisplay{str: str}
}

func (sd *stringDisplay) GetColumns() int {
	return len(sd.str)
}

func (sd *stringDisplay) GetRows() int {
	return 1
}

func (sd *stringDisplay) GetRowText(row int) string {
	if row == 0 {
		return sd.str
	}

	return ""
}

type multiStringDisplay struct {
	strs      []string
	maxLength int
}

func NewMultiStringDisplay() display.Printer {
	return &multiStringDisplay{
		strs:      []string{},
		maxLength: 0,
	}
}

func MultiAdd(p display.Printer, str string) {
	md, ok := p.(*multiStringDisplay)
	if !ok {
		fmt.Println("not multi")
	}
	md.strs = append(md.strs, str)

	if len(str) > md.maxLength {
		md.maxLength = len(str)
	}
}

func (d *multiStringDisplay) GetColumns() int {
	return d.maxLength
}

func (d *multiStringDisplay) GetRows() int {
	return len(d.strs)
}

func (d *multiStringDisplay) GetRowText(row int) string {
	if row < len(d.strs) {
		str := d.strs[row]
		return fmt.Sprintf("%s%s", str, strings.Repeat(" ", d.maxLength-len(str)))
	}

	return ""
}

type sideBorder struct {
	printer    display.Printer
	borderChar string
}

func NewSideBorder(printer display.Printer, char string) *sideBorder {
	return &sideBorder{
		printer:    printer,
		borderChar: char,
	}
}

func (sb *sideBorder) GetColumns() int {
	return 1 + sb.printer.GetColumns() + 1
}

func (sb *sideBorder) GetRows() int {
	return sb.printer.GetRows()
}

func (sb *sideBorder) GetRowText(row int) string {
	return fmt.Sprintf("%s%s%s", sb.borderChar, sb.printer.GetRowText(row), sb.borderChar)
}

type fullBorder struct {
	printer display.Printer
}

func NewFullBorder(printer display.Printer) *fullBorder {
	return &fullBorder{printer: printer}
}

func (fb *fullBorder) GetColumns() int {
	return 1 + fb.printer.GetColumns() + 1
}

func (fb *fullBorder) GetRows() int {
	return 1 + fb.printer.GetRows() + 1
}

func (fb *fullBorder) GetRowText(row int) string {
	if row == 0 || row == (fb.printer.GetRows()+1) {
		return fmt.Sprintf("+%s+", fb.makeLine("-", fb.printer.GetColumns()))
	}

	return fmt.Sprintf("|%s|", fb.printer.GetRowText(row-1))
}

func (fb *fullBorder) makeLine(char string, cnt int) string {
	return strings.Repeat(char, cnt)
}

type upDownBoder struct {
	printer    display.Printer
	borderChar string
}

func NewUpDownBorder(printer display.Printer, char string) *upDownBoder {
	return &upDownBoder{
		printer:    printer,
		borderChar: char,
	}
}

func (udb *upDownBoder) GetColumns() int {
	return udb.printer.GetColumns()
}

func (udb *upDownBoder) GetRows() int {
	return 1 + udb.printer.GetRows() + 1
}

func (udb *upDownBoder) GetRowText(row int) string {
	if row == 0 || row == udb.GetRows()-1 {
		return strings.Repeat(udb.borderChar, udb.GetColumns())
	}

	return udb.printer.GetRowText(row - 1)
}
