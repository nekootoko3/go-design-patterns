package visitor

import (
	"fmt"
	"regexp"
)

type visitor interface {
	visitFile(f *file) string
	visitDir(d *dir) string
}

type element interface {
	Accept(v visitor) string
}

type entry interface {
	element
	getName() string
	getSize() int
	Add(e entry)
}

type defaultEntry struct {
	entry
	name string
}

func (de *defaultEntry) getName() string {
	return de.name
}

func (de *defaultEntry) print(e entry) string {
	return fmt.Sprintf("%s (%d)\n", e.getName(), e.getSize())
}

type file struct {
	*defaultEntry
	size int
}

func (f *file) getSize() int {
	return f.size
}

func (f *file) Add(e entry) {}

func (f *file) Accept(v visitor) string {
	return v.visitFile(f)
}

type dir struct {
	*defaultEntry
	entries []entry
}

func (d *dir) getSize() int {
	size := 0
	for _, e := range d.entries {
		size += e.getSize()
	}
	return size
}

func (d *dir) Add(e entry) {
	d.entries = append(d.entries, e)
}

func (d *dir) Accept(v visitor) string {
	return v.visitDir(d)
}

type listVisitor struct {
	currentDir string
}

func NewListVisitor() *listVisitor {
	return &listVisitor{}
}

func (lv *listVisitor) visitFile(f *file) string {
	return fmt.Sprintf("%s/%s", lv.currentDir, f.print(f))
}

func (lv *listVisitor) visitDir(d *dir) string {
	saveDir := lv.currentDir
	list := fmt.Sprintf("%s/%s", lv.currentDir, d.print(d))
	lv.currentDir += "/" + d.getName()
	for _, e := range d.entries {
		list += e.Accept(lv)
	}
	lv.currentDir = saveDir
	return list
}

type fileFindVisitor struct {
	currentDir string
	regex      *regexp.Regexp
}

func NewFileFindVisitor(regexStr string) *fileFindVisitor {
	regex := regexp.MustCompile(regexStr)
	return &fileFindVisitor{regex: regex}
}

func (ffv *fileFindVisitor) visitFile(f *file) string {
	if ffv.regex.MatchString(f.getName()) {
		return fmt.Sprintf("%s/%s", ffv.currentDir, f.print(f))
	}
	return ""
}

func (ffv *fileFindVisitor) visitDir(d *dir) string {
	saveDir := ffv.currentDir
	list := ""
	ffv.currentDir += "/" + d.getName()
	for _, e := range d.entries {
		list += e.Accept(ffv)
	}
	ffv.currentDir = saveDir
	return list
}

func NewFile(name string, size int) *file {
	return &file{
		defaultEntry: &defaultEntry{name: name},
		size:         size,
	}
}

func NewDir(name string) *dir {
	return &dir{defaultEntry: &defaultEntry{name: name}}
}
