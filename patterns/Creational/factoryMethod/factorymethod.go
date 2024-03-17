package factorymethod

const (
	A = "objectA"
	B = "objectB"
	C = "objectC"
)

type Factory interface {
	CreateProduct(string) Product
}

type ConcreteFactory struct{}

func NewFactory() Factory {
	return &ConcreteFactory{}
}

type Product interface {
	ProductId() string
}

func (cf *ConcreteFactory) CreateProduct(id string) Product {
	var product Product
	switch id {
	case A:
		product = &ProductA{Id: id}
	case B:
		product = &ProductB{Id: id}
	case C:
		product = &ProductC{Id: id}
	}
	return product
}

type ProductA struct{ Id string }

func (a *ProductA) ProductId() string {
	return a.Id
}

type ProductB struct{ Id string }

func (b *ProductB) ProductId() string {
	return b.Id
}

type ProductC struct{ Id string }

func (c *ProductC) ProductId() string {
	return c.Id
}
