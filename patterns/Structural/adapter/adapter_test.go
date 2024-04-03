package adapter

import "testing"

func TestAdapter(t *testing.T) {

	tests := []struct {
		name       string
		memoryCard MemoryCard
		want       string
	}{
		{
			name:       "sdCard",
			memoryCard: &SdCard{},
			want: "Connection ...\n" +
				"SdCard is connected\n" +
				"Data was copied\n",
		}, {
			name:       "microSd",
			memoryCard: &MicroSd{},
			want: "Connection ...\n" +
				"MicroSd is connected\n" +
				"Data was copied\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cardReader := &CardReader{MemoryCard: test.memoryCard}
			got := cardReader.ConnectToComputer()

			if test.want != got {
				t.Errorf("want: %s, got: %s\n", test.want, got)
			}
		})
	}
}
