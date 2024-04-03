package strategy

import "testing"

func TestStrategy(t *testing.T) {
	strategy := &HeapSort{}
	object := NewObject(strategy)

	tests := []struct {
		name     string
		strategy Strategy
		source   []int
		want     []int
	}{
		{
			name:     "BubbleSort",
			strategy: &BubbleSort{},
			source:   []int{45, 34, 85, 22, 100},
			want:     []int{22, 34, 45, 85, 100},
		}, {
			name:     "HeapSort",
			strategy: &HeapSort{},
			source:   []int{45, 34, 85, 22, 100},
			want:     []int{22, 34, 45, 85, 100},
		}, {
			name:     "InsertSort",
			strategy: &InsertSort{},
			source:   []int{45, 34, 85, 22, 100},
			want:     []int{22, 34, 45, 85, 100},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			object.Toggler(test.strategy)
			object.Sort(test.source)
			got := test.source
			flag := false

			for i := range test.want {
				if test.want[i] != got[i] {
					flag = true
				}
			}

			if flag {
				t.Errorf("want: %v, got: %v\n", test.want, got)
			}
		})
	}
}
