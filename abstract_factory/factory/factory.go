package factory

type Item interface {
	ToString() string
}

type Link interface {
	Item
}

type Tray interface {
	Item
	Add(item Item)
}

type Page interface {
	Add(item Item)
	Output() string
}

type Factory interface {
	NewLink(caption, url string) Link
	NewTray(caption string) Tray
	NewPage(title, author string) Page
}
