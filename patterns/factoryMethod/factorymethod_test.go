package factorymethod

import "testing"

func TestFactoryMethod(t *testing.T) {
	factory := NewFactory()

	products := []struct {
		name    string
		product Product
		want    string
	}{
		{
			name:    "A",
			product: factory.CreateProduct(A),
			want:    A,
		}, {
			name:    "B",
			product: factory.CreateProduct(B),
			want:    B,
		}, {
			name:    "C",
			product: factory.CreateProduct(C),
			want:    C,
		},
	}

	for _, p := range products {
		t.Run(p.name, func(t *testing.T) {
			got := p.product.ProductId()

			if p.want != got {
				t.Errorf("want: %s, got: %s\n", p.want, got)
			}
		})
	}
}
