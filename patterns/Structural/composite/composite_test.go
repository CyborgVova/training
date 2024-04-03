package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {
	chH := &Character{Character: 'H'}
	che := &Character{Character: 'e'}
	chl := &Character{Character: 'l'}
	cho := &Character{Character: 'o'}
	chSpace := &Character{Character: ' '}
	chComma := &Character{Character: ','}
	chW := &Character{Character: 'W'}
	chr := &Character{Character: 'r'}
	chd := &Character{Character: 'd'}
	chExcl := &Character{Character: '!'}
	chV := &Character{Character: 'V'}
	chv := &Character{Character: 'v'}
	cha := &Character{Character: 'a'}

	hello := &Word{}
	hello.Add(chH, che, chl, chl, cho)

	world := &Word{}
	world.Add(chW, cho, chr, chl, chd)

	vova := &Word{}
	vova.Add(chV, cho, chv, cha)

	sentence := &Word{}
	sentence.Add(hello, chComma, chSpace, world, chSpace, chExcl, chExcl, chExcl)

	want := "Hello, World !!!"
	got := sentence.Print()
	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}

	sentence.Remove(chSpace, world, chSpace, chExcl, chExcl, chExcl)
	sentence.Add(chSpace, vova, chSpace, chExcl, chExcl, chExcl)

	want = "Hello, Vova !!!"
	got = sentence.Print()
	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
