package stock

type Warehouse interface {
	AddProduct(product Product) error
	UpdateQuantity(productID, quantity int) error
}
