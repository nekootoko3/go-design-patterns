package iterator

// Abstract ------------

// Aggregate
type Aggregate interface {
	Iterator() Iterator
}

// Iterator
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Concrete -----------

// Book ------------

// Book represents Book struct
type Book struct {
	name string
}

// NewBook is a constructor of Book
func NewBook(name string) *Book {
	return &Book{name: name}
}

// GetName return a name of Book
func (b *Book) GetName() string {
	return b.name
}

// BookShelf ------------

// BookShelf implements Aggregate
type BookShelf struct {
	books []*Book
	last  int
}

// NewBookShelf is constructor of BookShelf
func NewBookShelf() *BookShelf {
	return &BookShelf{}
}

// GetBookAt returns Book
func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

// AppendBook append a Book to BookShelf
func (bs *BookShelf) AppendBook(book *Book) {
	bs.books = append(bs.books, book)
	bs.last++
}

func (bs *BookShelf) getLength() int {
	return bs.last
}

// Iterator returns Iterator of BookShelf
func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}

// BookShelf ------------

// bookShelfIterator implements Iterator
type bookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

// NewBookShelfIterator is a constructor of bookShelfIterator
func NewBookShelfIterator(bf *BookShelf) *bookShelfIterator {
	return &bookShelfIterator{
		bookShelf: bf,
		index:     0,
	}
}

// HasNext returns whether bookShelf has next Book or not
func (bsi *bookShelfIterator) HasNext() bool {
	if bsi.bookShelf.last <= bsi.index {
		return false
	}
	return true
}

// Next returns next Book
func (bsi *bookShelfIterator) Next() interface{} {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}
