package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/abstract_factory/factory"
	"github.com/nekootoko3/go-design-patterns/abstract_factory/list_factory"
	"github.com/nekootoko3/go-design-patterns/abstract_factory/markdown_factory"
)

func main() {
	// html

	var f factory.Factory = list_factory.NewListFactory()

	yomi := f.NewLink("yomiuri", "https://yomiuri.jp")
	asahi := f.NewLink("asahi", "https://asahi.jp")

	yahoo := f.NewLink("yahoo", "https://yahoo.com")
	yahooJapan := f.NewLink("yahoo japan!", "https://yahoo.co.jp")

	google := f.NewLink("google", "https://google.com")
	excite := f.NewLink("excite", "https://excite.com")

	trayNews := f.NewTray("NEWS")
	trayNews.Add(yomi)
	trayNews.Add(asahi)

	trayYahoo := f.NewTray("yahoo")
	trayYahoo.Add(yahoo)
	trayYahoo.Add(yahooJapan)

	traySearchEngine := f.NewTray("Search Engines")
	traySearchEngine.Add(trayYahoo)
	traySearchEngine.Add(google)
	traySearchEngine.Add(excite)

	p := f.NewPage("My Favorites", "nekootoko3")
	p.Add(trayNews)
	p.Add(traySearchEngine)
	fmt.Println(p.Output())

	// markdown

	f = markdown_factory.NewMarkdownFactory()

	yomi = f.NewLink("yomiuri", "https://yomiuri.jp")
	asahi = f.NewLink("asahi", "https://asahi.jp")

	yahoo = f.NewLink("yahoo", "https://yahoo.com")
	yahooJapan = f.NewLink("yahoo japan!", "https://yahoo.co.jp")

	google = f.NewLink("google", "https://google.com")
	excite = f.NewLink("excite", "https://excite.com")

	trayNews = f.NewTray("NEWS")
	trayNews.Add(yomi)
	trayNews.Add(asahi)

	trayYahoo = f.NewTray("yahoo")
	trayYahoo.Add(yahoo)
	trayYahoo.Add(yahooJapan)

	traySearchEngine = f.NewTray("Search Engines")
	traySearchEngine.Add(trayYahoo)
	traySearchEngine.Add(google)
	traySearchEngine.Add(excite)

	p = f.NewPage("My Favorites", "nekootoko3")
	p.Add(trayNews)
	p.Add(traySearchEngine)
	fmt.Println(p.Output())
}
