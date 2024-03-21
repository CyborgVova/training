package flyweight

import "testing"

func TestFlyWeight(t *testing.T) {
	want := ""
	got := ""

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
