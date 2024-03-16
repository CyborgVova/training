package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	tests := []struct {
		tname     string
		singleton *instance
		id        string
		name      string
	}{
		{
			tname:     "singleton is nil",
			singleton: GetInstance("2", "Seva"),
			id:        "2",
			name:      "Seva",
		}, {
			tname:     "have singleton instance",
			singleton: GetInstance("1", "Jhon"),
			id:        "2",
			name:      "Seva",
		},
	}

	for _, test := range tests {
		t.Run(test.tname, func(t *testing.T) {
			want := struct {
				id   string
				name string
			}{
				id:   test.id,
				name: test.name,
			}
			got := struct {
				id   string
				name string
			}{
				id:   test.singleton.Id,
				name: test.singleton.Name,
			}

			if want.id != got.id {
				t.Errorf("want: %s, got: %s\n", want.id, got.id)
			}

			if want.name != got.name {
				t.Errorf("want: %s, got: %s\n", want.name, got.name)
			}
		})
	}
}
