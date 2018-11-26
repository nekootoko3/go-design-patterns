package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/singleton/triple/triple"
)

func main() {
	t1 := triple.GetInstance(1)
	t2 := triple.GetInstance(2)
	t3 := triple.GetInstance(3)
	dup_t1 := triple.GetInstance(1)

	t1.String()
	t2.String()
	t3.String()
	dup_t1.String()

	if t1 == dup_t1 {
		fmt.Printf("t1: %p\n", t1)
		fmt.Printf("dup_t1: %p\n", t1)
	}
}
