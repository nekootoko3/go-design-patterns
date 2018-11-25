package id_card

import "fmt"

type idCard struct {
	id    int
	owner string
}

func newIDCard(id int, owner string) *idCard {
	return &idCard{
		id:    id,
		owner: owner,
	}
}

func (ic *idCard) Use() {
	fmt.Printf("Use %s's Card\n", ic.owner)
}

func (ic *idCard) GetOwner() string {
	return ic.owner
}
