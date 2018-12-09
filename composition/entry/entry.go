package entry

type Entry interface {
	GetName() string
	GetSize() int
	PrintList(prefix string)
	Add(entry Entry)
}
