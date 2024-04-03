package facade

type Action interface {
	In() string
	Out() string
}

type Human struct {
	Name string
}

type Home struct {
	Human *Human
}

func (h *Home) In() string {
	return h.Human.Name + " came home\n"
}

func (h *Home) Out() string {
	return h.Human.Name + " left home\n"
}

type Auto struct {
	Human *Human
}

func (h *Auto) In() string {
	return h.Human.Name + " got into the car\n"
}

func (h *Auto) Out() string {
	return h.Human.Name + " got out of the car\n"
}

type Job struct {
	Human *Human
}

func (h *Job) In() string {
	return h.Human.Name + " came job\n"
}

func (h *Job) Out() string {
	return h.Human.Name + " left job\n"
}

type Facade struct {
	Human *Human
	Home  *Home
	Auto  *Auto
	Job   *Job
}

func NewFacade(name string) *Facade {
	human := &Human{Name: name}
	return &Facade{
		Human: human,
		Home:  &Home{human},
		Auto:  &Auto{human},
		Job:   &Job{human},
	}
}
