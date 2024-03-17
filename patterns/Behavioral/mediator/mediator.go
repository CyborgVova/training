package mediator

import (
	"errors"
)

type Mediator interface {
	Paying(string) (string, error)
}

type Store struct {
	product map[string]float64
	money   float64
}

func (s *Store) Sell(product string) {
	s.money += s.product[product]
	delete(s.product, product)
}

type Buyer struct {
	mediator Mediator
	purchase []string
	money    float64
}

func (b *Buyer) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

func (b *Buyer) Buy(product string, price float64) {
	b.purchase = append(b.purchase, product)
	b.money -= price
}

type Bank struct {
	store *Store
	buyer *Buyer
}

func (b *Bank) Union(store *Store, buyer *Buyer) {
	b.store = store
	b.buyer = buyer
}

func (b *Bank) Paying(product string) (string, error) {
	if cost, ok := b.store.product[product]; ok {
		if cost <= b.buyer.money {
			b.buyer.Buy(product, cost)
			b.store.Sell(product)

			return "purchase completed", nil
		} else {
			return "", errors.New("not enough money")
		}
	}
	return "", errors.New("no such product")
}
