package command

import "testing"

func TestCommand(t *testing.T) {
	invoke := &Invoke{}

	invoke.AddCommandToInvoke(&ToggleOnCommand{&Reciever{}})
	invoke.AddCommandToInvoke(&ToggleOffCommand{&Reciever{}})

	got := invoke.Executor()

	want := "Toggle On\nToggle Off\n"

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
