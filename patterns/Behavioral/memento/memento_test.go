package memento

import (
	"testing"
)

func TestMemento(t *testing.T) {
	object := Object{Field: 100}
	state := NewStates(object)
	object.Field = 200
	state.Save(object)
	object.Field = 300
	state.Save(object)

	tests := []struct {
		name  string
		field int
		state Object
	}{
		{
			name:  "First",
			field: 300,
			state: state.GetLast(),
		}, {
			name:  "Second",
			field: 200,
			state: state.GetLast(),
		}, {
			name:  "Third",
			field: 100,
			state: state.GetLast(),
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
