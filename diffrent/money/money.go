package money

import (
	"fmt"
)

type Money struct {
	Rubles uint64
	Kopeck uint8
}

type ErrorMoney string

var (
	NotEnoughMoney ErrorMoney = "not enough money"
)

func (e ErrorMoney) Error() string {
	return string(e)
}

func (m *Money) PrintMoney() {
	fmt.Printf("%d.%d", m.Rubles, m.Kopeck)
}

func (m *Money) Equal(other *Money) bool {
	return m.Rubles == other.Rubles && m.Kopeck == other.Kopeck
}

func (m *Money) Add(inc *Money) {
	m.Kopeck += inc.Kopeck
	if m.Kopeck > 99 {
		m.Kopeck %= 100
		m.Rubles++
	}
	m.Rubles += inc.Rubles
}

func (m *Money) Sub(dec *Money) error {
	if m.Rubles < dec.Rubles || (m.Rubles == dec.Rubles && m.Kopeck < dec.Kopeck) {
		return NotEnoughMoney
	}

	if m.Kopeck < dec.Kopeck {
		m.Kopeck += 100 - dec.Kopeck
		m.Rubles--
	} else {
		m.Kopeck -= dec.Kopeck
	}
	m.Rubles -= dec.Rubles
	return nil
}
