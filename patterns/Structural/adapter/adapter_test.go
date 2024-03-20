package adapter

import "testing"

func TestAdapter(t *testing.T) {
	sdCard := &SdCard{}
	cardReader := &CardReader{SdCard: sdCard}

	want := "Connection ...\n" +
		"Connection is done\n" +
		"Data was copied\n"
	got := cardReader.ConnectToComputer()

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
