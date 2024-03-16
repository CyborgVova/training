package chainofresponsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	handle :=
		&HandlerA{
			next: &HandlerB{
				next: &HandlerC{
					next: nil,
				},
			},
		}

	tests := []struct {
		name    string
		request string
		want    string
	}{
		{
			name:    "A",
			request: "A",
			want:    "response from A handler",
		}, {
			name:    "B",
			request: "B",
			want:    "response from B handler",
		}, {
			name:    "C",
			request: "C",
			want:    "response from C handler",
		}, {
			name:    "D",
			request: "D",
			want:    "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := handle.HandleFunc(test.request)
			if test.want != got {
				t.Errorf("want: %s, got: %s\n", test.want, got)
			}
		})
	}
}
