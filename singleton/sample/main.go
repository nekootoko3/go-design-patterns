package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/singleton/sample/singleton"
)

func main() {
	fmt.Println("Start")

	fmt.Println("Instance - 1")
	ins1 := singleton.GetInstance()
	fmt.Println("Instance - 2")
	ins2 := singleton.GetInstance()

	fmt.Printf("ins1 address: %p\n", ins1)
	fmt.Printf("ins2 address: %p\n", ins2)
	if ins1 == ins2 {
		fmt.Println("ins1 == ins2")
	}
}
