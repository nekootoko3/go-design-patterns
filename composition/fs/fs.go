package fs

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/composition/entry"
)

type file struct {
	name string
	size int
}

func NewFile(name string) entry.Entry {
	return &file{
		name: name,
		size: len(name),
	}
}

func (f *file) GetName() string {
	return f.name
}

func (f *file) GetSize() int {
	return f.size
}

func (f *file) PrintList(prefix string) {
	fmt.Printf("%s/%s\n", prefix, f.name)
}

func (f *file) Add(entry entry.Entry) {
	fmt.Println("NO")
}

type dir struct {
	name    string
	entries []entry.Entry
}

func NewDir(name string) entry.Entry {
	return &dir{
		name:    name,
		entries: []entry.Entry{},
	}
}

func (d *dir) GetName() string {
	return d.name
}

func (d *dir) GetSize() int {
	size := 0
	for _, e := range d.entries {
		size += e.GetSize()
	}
	return size
}

func (d *dir) PrintList(prefix string) {
	fmt.Printf("%s/%s\n", prefix, d.name)
	for _, e := range d.entries {
		e.PrintList(fmt.Sprintf("%s/%s", prefix, d.name))
	}
}

func (d *dir) Add(entry entry.Entry) {
	d.entries = append(d.entries, entry)
}
