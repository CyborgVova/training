package decorator

import "testing"

func TestDecorator(t *testing.T) {
	object := &Object{}
	decorator := &Decorator{Object: object}

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
		}, {
			name:     "default",
			text:     "Default behavior",
			apostrof: "",
			want:     "Default behavior",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			want := test.text
			got := object.Print(test.text)

			if want != got {
				t.Errorf("want: %s, got: %s\n", want, got)
			}

			decorator.Apostrof = test.apostrof
			want = test.want
			got = decorator.Print(test.text)

			if want != got {
				t.Errorf("want: %s, got: %s\n", want, got)
			}
		})
	}
}
