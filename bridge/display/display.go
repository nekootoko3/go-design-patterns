package display

import (
	"math/rand"
	"time"

	"github.com/nekootoko3/go-design-patterns/bridge/display_impl"
)

type display struct {
	impl display_impl.DisplayImpl
}

func NewDisplay(impl display_impl.DisplayImpl) *display {
	return &display{impl: impl}
}

func (d *display) open() {
	d.impl.RawOpen()
}

func (d *display) print() {
	d.impl.RawPrint()
}

func (d *display) close() {
	d.impl.RawClose()
}

func (d *display) Display() {
	d.open()
	d.print()
	d.close()
}

type countDisplay struct {
	*display
}

func NewCountDisplay(d *display) *countDisplay {
	return &countDisplay{d}
}

func (cd *countDisplay) MultiDisplay(times int) {
	cd.open()
	for i := 0; i < times; i++ {
		cd.print()
	}
	cd.close()
}

func (cd *countDisplay) RandDisplay() {
	rand.Seed(time.Now().UnixNano())
	cd.MultiDisplay(rand.Intn(10))
}

type randDisplay struct {
	*display
}

func NewRandDisplay(d *display) *randDisplay {
	return &randDisplay{d}
}

func (rd *randDisplay) RandDisplay() {
	rd.open()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Intn(10); i++ {
		rd.print()
	}
	rd.close()
}
