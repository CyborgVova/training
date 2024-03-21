package facade

import "testing"

func TestFacade(t *testing.T) {
	facade := NewFacade("Igor Pupkin")

	myDay := facade.Home.Out() +
		facade.Auto.In() +
		facade.Auto.Out() +
		facade.Job.In() +
		facade.Job.Out() +
		facade.Auto.In() +
		facade.Auto.Out() +
		facade.Home.In()

	want := "Igor Pupkin left home\n" +
		"Igor Pupkin got into the car\n" +
		"Igor Pupkin got out of the car\n" +
		"Igor Pupkin came job\n" +
		"Igor Pupkin left job\n" +
		"Igor Pupkin got into the car\n" +
		"Igor Pupkin got out of the car\n" +
		"Igor Pupkin came home\n"

	if want != myDay {
		t.Errorf("want: %s, got: %s\n", want, myDay)
	}
}
