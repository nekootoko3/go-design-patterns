package id_card

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/factory_method/framework"
)

type idCardFactory struct {
	*framework.Factory
	latest    int
	ownerList map[int]string
}

func NewIDCardFactory() *idCardFactory {
	return &idCardFactory{
		latest:    1,
		ownerList: map[int]string{},
	}
}

func (idf *idCardFactory) CreateProduct(owner string) framework.Product {
	fmt.Printf("Creating %s's Card\n", owner)
	return newIDCard(idf.latest, owner)
}

func (idf *idCardFactory) RegisterProduct(product framework.Product) {
	idf.ownerList[idf.latest] = product.(*idCard).GetOwner()
	idf.latest++
}

func (idf *idCardFactory) ListOwners() {
	for id := 1; id < idf.latest; id++ {
		if owner, ok := idf.ownerList[id]; ok {
			fmt.Printf("%d: %s\n", id, owner)
		}
	}
}
