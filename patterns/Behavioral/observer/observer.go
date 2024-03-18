package observer

type Publisher interface {
	AddConsumer(...Subscriber)
	SetState(string)
	Notify()
}

type Producer struct {
	Subscribers []Subscriber
	State       string
}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) AddConsumer(subscriber ...Subscriber) {
	for _, sub := range subscriber {
		p.Subscribers = append(p.Subscribers, sub)
	}
}

func (p *Producer) SetState(state string) {
	p.State = state
}

func (p *Producer) Notify() {
	for _, observer := range p.Subscribers {
		observer.Update(p.State)
	}
}

type Subscriber interface {
	Update(string)
}

type Consumer struct {
	State string
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Update(state string) {
	c.State = state
}
