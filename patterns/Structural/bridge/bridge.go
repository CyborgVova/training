package bridge

type Engine interface {
	EngineType() string
}

type PetrolEngine struct{}

func (p *PetrolEngine) EngineType() string {
	return "Petrol engine\n"
}

type DieselEngine struct{}

func (d *DieselEngine) EngineType() string {
	return "Diesel engine\n"
}

type Auto interface {
	Complectation() string
}

type BMW struct {
	Engine Engine
}

func (b *BMW) Complectation() string {
	return "It is BMW with " + b.Engine.EngineType()
}

type Audi struct {
	Engine Engine
}

func (a *Audi) Complectation() string {
	return "It is Audi with " + a.Engine.EngineType()
}
