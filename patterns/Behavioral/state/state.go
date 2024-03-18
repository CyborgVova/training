package state

type Object struct {
	State State
}

func NewObject() *Object {
	return &Object{State: &StateA{}}
}

func (o *Object) Trigger() {
	o.State = o.State.Trigger()
}

func (o *Object) GetState() string {
	return o.State.GetState()
}

type State interface {
	Trigger() State
	GetState() string
}

type StateA struct{}

func (s *StateA) GetState() string {
	return "RED"
}

func (s *StateA) Trigger() State {
	return &StateB{}
}

type StateB struct{}

func (s *StateB) GetState() string {
	return "GREEN"
}

func (s *StateB) Trigger() State {
	return &StateA{}
}
