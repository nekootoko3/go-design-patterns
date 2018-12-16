package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/visitor/visitor"
)

func main() {
	rootDir := visitor.NewDir("root")
	usrDir := visitor.NewDir("usr")
	fileA := visitor.NewFile("A.txt", 1)

	rootDir.Add(usrDir)
	rootDir.Add(fileA)

	fileB := visitor.NewFile("B.html", 2)
	usrDir.Add(fileB)

	result := rootDir.Accept(visitor.NewListVisitor())
	fmt.Println(result)

	resultFileFind := rootDir.Accept(visitor.NewFileFindVisitor("html"))
	fmt.Println(resultFileFind)
}
