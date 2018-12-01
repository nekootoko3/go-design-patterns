package markdown_factory

import (
	"fmt"
	"strings"

	"github.com/nekootoko3/go-design-patterns/abstract_factory/factory"
)

type mdLink struct {
	caption string
	url     string
}

func (l *mdLink) ToString() string {
	return fmt.Sprintf("[%s](%s)\n", l.caption, l.url)
}

type mdTray struct {
	caption string
	items   []factory.Item
}

func (t *mdTray) ToString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("## %s\n", t.caption))
	for _, item := range t.items {

		b.WriteString(item.ToString())
	}

	return b.String()
}

func (t *mdTray) Add(item factory.Item) {
	t.items = append(t.items, item)
}

type mdPage struct {
	title  string
	author string
	items  []factory.Item
}

func (p *mdPage) Add(item factory.Item) {
	p.items = append(p.items, item)
}

func (p *mdPage) Output() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("# %s\n", p.title))
	b.WriteString(fmt.Sprintf("%s\n", p.author))
	for _, item := range p.items {
		b.WriteString(item.ToString())
	}

	return b.String()
}

type markdownFactory struct {
}

func NewMarkdownFactory() factory.Factory {
	return &markdownFactory{}
}

func (f *markdownFactory) NewLink(caption, url string) factory.Link {
	return &mdLink{
		caption: caption,
		url:     url,
	}
}

func (f *markdownFactory) NewTray(caption string) factory.Tray {
	return &mdTray{
		caption: caption,
	}
}

func (f *markdownFactory) NewPage(title, author string) factory.Page {
	return &mdPage{
		title:  title,
		author: author,
	}
}
