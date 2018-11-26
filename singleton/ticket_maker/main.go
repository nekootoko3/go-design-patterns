package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/singleton/ticket_maker/ticket_maker"
)

func main() {
	fmt.Println("Create ticket_maker 1")
	tm1 := ticket_maker.GetInstance()
	fmt.Println(tm1.GetNextTicketNumber())
	fmt.Println(tm1.GetNextTicketNumber())

	fmt.Println("Create ticket_maker 2")
	tm2 := ticket_maker.GetInstance()
	fmt.Println(tm2.GetNextTicketNumber())
	fmt.Println(tm2.GetNextTicketNumber())
}
