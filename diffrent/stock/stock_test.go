package stock

import (
	"testing"
)

func TestAddProduct(t *testing.T) {
	storage := NewStorage()
	tests := []struct {
		name    string
		product *Product
		err     error
	}{
		{
			name: "new product",
			product: &Product{
				ID:       1,
				Name:     "Bread",
				Quantity: 20,
			},
			err: nil,
		}, {
			name: "exist product",
			product: &Product{
				ID:       1,
				Name:     "Bread",
				Quantity: 20,
			},
			err: SuchProductExist,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := storage.AddProduct(test.product)
			if err != test.err {
				t.Error()
			}
		})
	}
}

func TestUpdateQuantities(t *testing.T) {
	storage := NewStorage()

	tests := []struct {
		name      string
		productID int
		quantity  int
		err       error
		want      int
	}{
		{
			name:      "not enough product",
			productID: 1,
			quantity:  -4,
			err:       NotEnoughProduct,
			want:      3,
		}, {
			name:      "increase quantity",
			productID: 1,
			quantity:  3,
			err:       nil,
			want:      6,
		}, {
			name:      "decrease quantity",
			productID: 1,
			quantity:  -4,
			err:       nil,
			want:      2,
		}, {
			name:      "product not exist",
			productID: 2,
			quantity:  1,
			err:       SuchProductNotExist,
			want:      0,
		},
	}

	storage.AddProduct(&Product{
		ID:       1,
		Name:     "Bread",
		Quantity: 3,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := storage.UpdateQuantity(test.productID, test.quantity)
			if err != test.err {
				t.Error()
			}
			if err == nil && test.want != storage.Products[test.productID].Quantity {
				t.Error()
			}
		})
	}
}
