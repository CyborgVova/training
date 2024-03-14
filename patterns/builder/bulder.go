package builder

type Builder interface {
	MakeFoundation(string)
	MakeBuild(string)
	MakeFinishing(string)
}

type Director struct {
	Builder Builder
}

func (d *Director) BuildAHouse() {
	d.Builder.MakeFoundation("Foundation: ")
	d.Builder.MakeBuild("Build: ")
	d.Builder.MakeFinishing("Finishing: ")
}

type ConcreteBuilder struct {
	Stage *Stage
}

func (cb *ConcreteBuilder) MakeFoundation(s string) {
	cb.Stage.Readiness += s + "ready\n"
}

func (cb *ConcreteBuilder) MakeBuild(s string) {
	cb.Stage.Readiness += s + "ready\n"
}

func (cb *ConcreteBuilder) MakeFinishing(s string) {
	cb.Stage.Readiness += s + "ready\n"
}

type Stage struct {
	Readiness string
}

func (s *Stage) Result() string {
	return s.Readiness
}
