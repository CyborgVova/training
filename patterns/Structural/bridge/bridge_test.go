package bridge

import "testing"

func TestBridge(t *testing.T) {
	tests := []struct {
		name string
		want string
		car  Auto
	}{
		{
			name: "Audi Petrol",
			want: "It is Audi with Petrol engine\n",
			car:  &Audi{Engine: &PetrolEngine{}},
		}, {
			name: "Audi Petrol",
			want: "It is Audi with Diesel engine\n",
			car:  &Audi{Engine: &DieselEngine{}},
		}, {
			name: "BMW Petrol",
			want: "It is BMW with Petrol engine\n",
			car:  &BMW{Engine: &PetrolEngine{}},
		}, {
			name: "BMW Petrol",
			want: "It is BMW with Diesel engine\n",
			car:  &BMW{Engine: &DieselEngine{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := test.car.Complectation()

			if test.want != got {
				t.Errorf("want: %s, got: %s\n", test.want, got)
			}
		})
	}
}
