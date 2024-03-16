package prototype

import "testing"

func TestPrototype(t *testing.T) {
	tests := []struct {
		city       string
		university string
	}{
		{
			city:       "Moscow",
			university: "MGU",
		},
		{
			city:       "Vladimir",
			university: "MFTI",
		},
		{
			city:       "Saint-Petersburg",
			university: "SPFA",
		},
	}

	for _, test := range tests {
		t.Run(test.city, func(t *testing.T) {
			prototype := NewPrototype(test.city, test.university)
			got := (prototype.Clone()).(*Prototype)
			if test.city != got.GetCity() {
				t.Errorf("City: want: %s, got: %s\n", test.city, got)
			}

			if test.university != got.GetUniversity() {
				t.Errorf("University: want: %s, got: %s\n", test.university, got)
			}

		})

	}
}
