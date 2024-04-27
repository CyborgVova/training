package account

import (
	"testing"
)

func TestSetBalance(t *testing.T) {
	acc := NewAccount("Vova")
	tests := []struct {
		name     string
		quantity float64
		want     float64
		err      error
	}{
		{
			name:     "negative quantity",
			quantity: -1,
			want:     0,
			err:      InputNegativeValue,
		}, {
			name:     "positive quantity",
			quantity: 1,
			want:     1,
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := acc.SetBalance(test.quantity)
			if err != test.err || acc.balance != test.want {
				t.Error()
			}

		})
	}
}

func TestGetBalance(t *testing.T) {
	acc := NewAccount("Vova")
	test := struct {
		name    string
		balance float64
		want    float64
	}{
		name:    "get balance",
		balance: 100,
		want:    100,
	}

	err := acc.SetBalance(test.balance)
	if err != nil {
		t.Error()
	}

	got := acc.GetBalance()
	if got != test.want {
		t.Error()
	}
}

func TestDeposit(t *testing.T) {
	acc := NewAccount("Vova")
	tests := []struct {
		name     string
		quantity float64
		err      error
		want     float64
	}{
		{
			name:     "negative quantity",
			quantity: -1,
			err:      InputNegativeValue,
			want:     0,
		}, {
			name:     "add balance",
			quantity: 100,
			err:      nil,
			want:     100,
		}, {
			name:     "add balance yet",
			quantity: 100,
			err:      nil,
			want:     200,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := acc.Deposit(test.quantity)
			if err != test.err {
				t.Error()
			}
			if test.want != acc.GetBalance() {
				t.Error()
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	acc := NewAccount("Vova")
	err := acc.SetBalance(100)
	if err != nil {
		t.Error()
	}

	tests := []struct {
		name     string
		quantity float64
		err      error
		want     float64
	}{
		{
			name:     "negative quantity",
			quantity: -1,
			err:      InputNegativeValue,
			want:     100,
		}, {
			name:     "not enough money",
			quantity: 200,
			err:      NotEnoughMoney,
			want:     100,
		}, {
			name:     "withdraw ok",
			quantity: 50,
			err:      nil,
			want:     50,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := acc.Withdraw(test.quantity)
			if err != test.err {
				t.Error()
			}
			if test.want != acc.GetBalance() {
				t.Error()
			}
		})
	}
}
