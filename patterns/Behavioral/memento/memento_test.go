package memento

import (
	"testing"
)

func TestMemento(t *testing.T) {
	object := Object{Field: 100}
	states := NewStates(object)
	object.Field = 200
	states.Save(object)
	object.Field = 300
	states.Save(object)

	tests := []struct {
		name  string
		field int
		state Object
	}{
		{
			name:  "First",
			field: 300,
			state: states.GetLast(),
		}, {
			name:  "Second",
			field: 200,
			state: states.GetLast(),
		}, {
			name:  "Third",
			field: 100,
			state: states.GetLast(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.field
			got := test.state.Field

			if want != got {
				t.Errorf("want: %d, got: %d\n", want, got)
			}
		})
	}

}
