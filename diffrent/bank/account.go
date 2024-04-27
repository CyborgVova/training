package account

type Account struct {
	name    string
	balance float64
}

type AccountError string

var (
	NotEnoughMoney     AccountError = "not enough money"
	InputNegativeValue AccountError = "input negative value"
)

func (a AccountError) Error() string {
	return string(a)
}

func NewAccount(owner string) *Account {
	return &Account{name: owner}
}

func (a *Account) SetBalance(quantity float64) error {
	if quantity < 0 {
		return InputNegativeValue
	}
	a.balance = quantity
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(quantity float64) error {
	if quantity < 0 {
		return InputNegativeValue
	}
	a.balance += quantity
	return nil
}

func (a *Account) Withdraw(quantity float64) error {
	if quantity < 0 {
		return InputNegativeValue
	}
	if a.balance-quantity < 0 {
		return NotEnoughMoney
	}
	a.balance -= quantity
	return nil
}
