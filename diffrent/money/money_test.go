package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMoney(t *testing.T) {
	tests := []struct {
		name      string
		haveMoney *Money
		addMoney  *Money
		want      *Money
	}{
		{
			name: "add zero",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 57,
			},
			addMoney: &Money{
				Rubles: 0,
				Kopeck: 0,
			},
			want: &Money{
				Rubles: 100,
				Kopeck: 57,
			},
		}, {
			name: "add rubles",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 57,
			},
			addMoney: &Money{
				Rubles: 23,
				Kopeck: 0,
			},
			want: &Money{
				Rubles: 123,
				Kopeck: 57,
			},
		}, {
			name: "add kopecks",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 57,
			},
			addMoney: &Money{
				Rubles: 0,
				Kopeck: 99,
			},
			want: &Money{
				Rubles: 101,
				Kopeck: 56,
			},
		}, {
			name: "add rubles and kopecks",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 57,
			},
			addMoney: &Money{
				Rubles: 11,
				Kopeck: 11,
			},
			want: &Money{
				Rubles: 111,
				Kopeck: 68,
			},
		}, {
			name: "add kopeck with increase rubles",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 57,
			},
			addMoney: &Money{
				Rubles: 0,
				Kopeck: 57,
			},
			want: &Money{
				Rubles: 101,
				Kopeck: 14,
			},
		},
		{
			name: "add kopeck without increase rubles",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 57,
			},
			addMoney: &Money{
				Rubles: 0,
				Kopeck: 3,
			},
			want: &Money{
				Rubles: 100,
				Kopeck: 60,
			},
		},
	}

	assert := assert.New(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.haveMoney.Add(test.addMoney)
			assert.Equal(test.haveMoney, test.want, "money should be the same")
		})
	}
}

func TestSubMoney(t *testing.T) {
	tests := []struct {
		name      string
		haveMoney *Money
		subMoney  *Money
		want      *Money
		err       error
	}{
		{
			name: "zero rubles and not enough kopeck",
			haveMoney: &Money{
				Rubles: 0,
				Kopeck: 45,
			},
			subMoney: &Money{
				Rubles: 0,
				Kopeck: 46,
			},
			want: &Money{
				Rubles: 0,
				Kopeck: 45,
			},
			err: NotEnoughMoney,
		}, {
			name: "equal rubles and not enough kopeck",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 45,
			},
			subMoney: &Money{
				Rubles: 100,
				Kopeck: 46,
			},
			want: &Money{
				Rubles: 100,
				Kopeck: 45,
			},
			err: NotEnoughMoney,
		}, {
			name: "equal rubles and equal kopeck",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 45,
			},
			subMoney: &Money{
				Rubles: 100,
				Kopeck: 45,
			},
			want: &Money{
				Rubles: 0,
				Kopeck: 0,
			},
			err: nil,
		}, {
			name: "less rubles and more kopeck",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 45,
			},
			subMoney: &Money{
				Rubles: 101,
				Kopeck: 44,
			},
			want: &Money{
				Rubles: 100,
				Kopeck: 45,
			},
			err: NotEnoughMoney,
		}, {
			name: "more rubles and less kopeck",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 45,
			},
			subMoney: &Money{
				Rubles: 99,
				Kopeck: 46,
			},
			want: &Money{
				Rubles: 0,
				Kopeck: 99,
			},
			err: nil,
		}, {
			name: "equal rubles and more kopeck",
			haveMoney: &Money{
				Rubles: 0,
				Kopeck: 46,
			},
			subMoney: &Money{
				Rubles: 0,
				Kopeck: 45,
			},
			want: &Money{
				Rubles: 0,
				Kopeck: 1,
			},
			err: nil,
		}, {
			name: "more rubles and equal kopeck",
			haveMoney: &Money{
				Rubles: 100,
				Kopeck: 45,
			},
			subMoney: &Money{
				Rubles: 99,
				Kopeck: 45,
			},
			want: &Money{
				Rubles: 1,
				Kopeck: 0,
			},
			err: nil,
		},
	}

	assert := assert.New(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.haveMoney.Sub(test.subMoney)
			assert.Equal(test.want, test.haveMoney)
			assert.Equal(test.err, err)
		})
	}
}
