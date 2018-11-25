package main

import "github.com/nekootoko3/go-design-patterns/factory_method/id_card"

func main() {
	factory := id_card.NewIDCardFactory()
	ic1 := factory.Create(factory, "neko")
	ic2 := factory.Create(factory, "inu")
	ic3 := factory.Create(factory, "okojo")
	ic1.Use()
	ic2.Use()
	ic3.Use()
	factory.ListOwners()
}
