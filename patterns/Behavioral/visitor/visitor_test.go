package visitor

import "testing"

func TestVisitor(t *testing.T) {
	zoo := &Zoo{}
	cinema := &Cinema{}
	circus := &Circus{}

	city := &City{}
	city.AddingPlaces(zoo, cinema, circus)

	tests := []struct {
		name    string
		visitor Visitor
		want    string
	}{
		{
			name:    "Vasja",
			visitor: &Vasja{},
			want: "Vasja visited the Zoo\n" +
				"Vasja visited the Cinema\n" +
				"Vasja visited the Circus\n",
		}, {
			name:    "Fedja",
			visitor: &Fedja{},
			want: "Fedja visited the Zoo\n" +
				"Fedja visited the Cinema\n" +
				"Fedja visited the Circus\n",
		}, {
			name:    "Sasha",
			visitor: &Sasha{},
			want: "Sasha visited the Zoo\n" +
				"Sasha visited the Cinema\n" +
				"Sasha visited the Circus\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := city.Accept(test.visitor)

			if test.want != got {
				t.Errorf("want: %s, got: %s\n", test.want, got)
			}
		})
	}
}
