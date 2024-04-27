package stock

type Storage struct {
	Products map[int]*Product
}

func NewStorage() *Storage {
	return &Storage{Products: make(map[int]*Product)}
}

func (s *Storage) AddProduct(product *Product) error {
	if _, ok := s.Products[product.ID]; ok {
		return SuchProductExist
	}
	s.Products[product.ID] = product
	return nil
}

func (s *Storage) UpdateQuantity(productID, quantity int) error {
	if value, ok := s.Products[productID]; !ok {
		return SuchProductNotExist
	} else if value.Quantity+quantity < 0 {
		return NotEnoughProduct
	}

	s.Products[productID].Quantity += quantity
	return nil
}
