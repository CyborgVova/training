package mediator

import "testing"

func TestMediator(t *testing.T) {
	product := map[string]float64{"carrot": 22.17, "potato": 9.34, "tomato": 37.67}
	buyer := &Buyer{money: 60.00}
	store := &Store{product: product}

	bank := &Bank{}
	bank.Union(store, buyer)

	buyer.SetMediator(bank)

	tests := []struct {
		name    string
		product string
		answer  string
	}{
		{
			name:    "first",
			product: "carrot",
			answer:  "purchase completed",
		}, {
			name:    "second",
			product: "tomato",
			answer:  "purchase completed",
		}, {
			name:    "third",
			product: "potato",
			answer:  "not enough money",
		}, {
			name:    "fourth",
			product: "brokkoli",
			answer:  "no such product",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := buyer.mediator.Paying(test.product)
			if err != nil {
				if test.answer != err.Error() {
					t.Errorf("want: %s, got: %s\n", test.answer, err.Error())
				}
				if got != "" {
					t.Errorf("want: %s, got: %s\n", "", got)
				}
			} else {
				if test.answer != got {
					t.Errorf("want: %s, got: %s\n", test.answer, got)
				}
			}
		})
	}
}
