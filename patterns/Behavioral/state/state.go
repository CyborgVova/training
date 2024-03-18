package state

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

type Object struct {
	State State
}

// Here default state set as StateA
func NewObject() *Object {
	return &Object{State: &StateA{}}
}

// The method name can be anything
func (o *Object) Trigger() {
	o.State = o.State.Trigger()
}

// The method name can be anything
func (o *Object) GetState() string {
	return o.State.GetState()
}
