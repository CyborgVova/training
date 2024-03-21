package decorator

import "testing"

func TestDecorator(t *testing.T) {
	want := ""
	got := ""

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
