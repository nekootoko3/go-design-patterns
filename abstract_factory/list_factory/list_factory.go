package list_factory

import (
	"fmt"
	"strings"

	"github.com/nekootoko3/go-design-patterns/abstract_factory/factory"
)

// Link

type listLink struct {
	caption string
	url     string
}

func (l *listLink) ToString() string {
	return fmt.Sprintf("<li><a href=\"%s\">%s</a><li>\n", l.url, l.caption)
}

// Tray

type listTray struct {
	caption string
	items   []factory.Item
}

func (t *listTray) ToString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("<li>%s\n", t.caption))
	b.WriteString("<ul>\n")
	for _, item := range t.items {
		b.WriteString(item.ToString())
	}
	b.WriteString("</ul>\n")
	b.WriteString("</li>\n")
	return b.String()
}

func (t *listTray) Add(item factory.Item) {
	t.items = append(t.items, item)
}

// Page

type listPage struct {
	title    string
	author   string
	contents []factory.Item
}

func (p *listPage) Add(item factory.Item) {
	p.contents = append(p.contents, item)
}

func (p *listPage) Output() string {
	var b strings.Builder
	b.WriteString("<body>\n")
	b.WriteString(fmt.Sprintf("<h1>%s</h1>\n", p.title))
	b.WriteString(fmt.Sprintf("<h2>%s</h2>\n", p.author))
	for _, c := range p.contents {
		b.WriteString(c.ToString())
	}
	b.WriteString("</body>\n")
	return b.String()
}

type listFactory struct {
}

func NewListFactory() factory.Factory {
	return &listFactory{}
}

func (f *listFactory) NewLink(caption, url string) factory.Link {
	return &listLink{
		caption: caption,
		url:     url,
	}
}
func (f *listFactory) NewTray(caption string) factory.Tray {
	return &listTray{
		caption: caption,
		items:   []factory.Item{},
	}
}

func (f *listFactory) NewPage(title, author string) factory.Page {
	return &listPage{
		title:    title,
		author:   author,
		contents: []factory.Item{},
	}
}
