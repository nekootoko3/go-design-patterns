package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/iterator/iterator"
)

func main() {
	bs := iterator.NewBookShelf()
	bs.AppendBook(iterator.NewBook("Kingdom"))
	bs.AppendBook(iterator.NewBook("The Life of Neko"))
	bs.AppendBook(iterator.NewBook("Holy Land"))
	bs.AppendBook(iterator.NewBook("Ema"))
	bsi := bs.Iterator()
	for bsi.HasNext() {
		book := bsi.Next()
		fmt.Println(book.(*(iterator.Book)).GetName())
	}
}
