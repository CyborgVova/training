package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	object := &Object{}
	proxy := NewProxy(object)

	want := "Start Run training\nEnd Run training\n"
	got := proxy.Run()

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}

	want = "Start Walking training\nEnd Walking training\n"
	got = proxy.Walking()

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
