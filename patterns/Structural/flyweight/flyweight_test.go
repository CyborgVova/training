package flyweight

import (
	"math"
	"testing"
)

func TestFlyWeight(t *testing.T) {
	circle := &Circle{Radius: 10.0}
	rectangle := &Rectangle{SideA: 10.0, SideB: 10.0}
	treeangle := &Treeangle{Side: 10.0, Height: 10.0}

	flyWeight := NewFlyWeight()
	flyWeight.Add("circle", circle)
	flyWeight.Add("rectangle", rectangle)
	flyWeight.Add("treeangle", treeangle)

	tests := []struct {
		name string
		want float64
	}{
		{
			name: "circle",
			want: 314.0,
		}, {
			name: "rectangle",
			want: 100.0,
		}, {
			name: "treeangle",
			want: 50.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			element, err := flyWeight.GetElement(test.name)
			if err != nil {
				t.Error("unexpected error: ", err.Error())
			}

			got := element.Area()

			if test.want != math.Round(got) {
				t.Errorf("want: %.6f, got: %.6f\n", test.want, got)
			}
		})
	}
}
