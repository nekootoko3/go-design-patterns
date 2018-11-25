package framework

type Creater interface {
	CreateProduct(owner string) Product
	RegisterProduct(Product)
}

type Factory struct {
}

func (f *Factory) Create(factory Creater, owner string) Product {
	p := factory.CreateProduct(owner)
	factory.RegisterProduct(p)
	return p
}
