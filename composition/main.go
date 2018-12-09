package main

import (
	"github.com/nekootoko3/go-design-patterns/composition/fs"
)

func main() {
	f1 := fs.NewFile("yoshinaga")
	f2 := fs.NewFile("neko")
	f3 := fs.NewFile("inu")
	d1 := fs.NewDir("dir-1")
	d2 := fs.NewDir("dir-2")
	d3 := fs.NewDir("dir-3")

	d1.Add(f1)
	d2.Add(f2)
	d1.Add(d2)
	d3.Add(d1)
	d3.Add(f3)
	d3.Add(f1)
	d3.Add(f2)
	d3.PrintList("root")
}
