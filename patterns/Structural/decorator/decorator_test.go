package decorator

import "testing"

func TestDecorator(t *testing.T) {

	tests := []struct {
		name     string
		text     string
		apostrof string
		want     string
	}{
		{
			name:     "one",
			text:     "First row",
			apostrof: "\"",
			want:     "\"First row\"",
		}, {
			name:     "two",
			text:     "Second row",
			apostrof: "`",
			want:     "`Second row`",
		}, {
			name:     "three",
			text:     "Third row",
			apostrof: "'",
			want:     "'Third row'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			object := &Object{}
			decorator := &Decorator{Apostrof: test.apostrof, Object: object}

			want := test.text
			got := object.Print(test.text)

			if want != got {
				t.Errorf("want: %s, got: %s\n", want, got)
			}

			want = test.want
			got = decorator.Print(test.text)

			if want != got {
				t.Errorf("want: %s, got: %s\n", want, got)
			}
		})
	}
}
