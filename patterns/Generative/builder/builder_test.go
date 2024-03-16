package builder

import "testing"

func TestBuilder(t *testing.T) {
	stage := &Stage{}
	concreteBuilder := &ConcreteBuilder{Stage: stage}
	director := &Director{Builder: concreteBuilder}
	director.BuildAHouse()

	want := stage.Result()
	got := "Foundation: ready\n" +
		"Build: ready\n" +
		"Finishing: ready\n"

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
