package state

import "testing"

func TestState(t *testing.T) {
	object := NewObject()

	tests := []struct {
		name string
		want string
		got  string
	}{
		{
			name: "First",
			want: "GREEN",
		}, {
			name: "Second",
			want: "RED",
		}, {
			name: "Third",
			want: "GREEN",
		}, {
			name: "Fourth",
			want: "RED",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			object.Trigger()
			got := object.GetState()
			if test.want != got {
				t.Errorf("want: %s, got: %s\n", test.want, got)
			}
		})
	}
}
