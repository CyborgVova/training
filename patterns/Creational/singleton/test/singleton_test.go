package test

import (
	"testing"

	"github.com/cyborgvova/training/patterns/Creational/singleton"
)

func TestSingleton(t *testing.T) {
	tests := []struct {
		tname string
		id    string
		name  string
	}{
		{
			tname: "singleton is nil",
			id:    "2",
			name:  "Seva",
		}, {
			tname: "have singleton instance",
			id:    "1",
			name:  "Jhon",
		},
	}

	for _, test := range tests {
		t.Run(test.tname, func(t *testing.T) {
			slt := singleton.GetInstance(test.id, test.name)
			want := struct {
				id   string
				name string
			}{
				id:   "2",
				name: "Seva",
			}
			got := struct {
				id   string
				name string
			}{
				id:   slt.Id,
				name: slt.Name,
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
