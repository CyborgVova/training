package prototype

type Clonable interface {
	Clone() Clonable
}

type Prototype struct {
	city       string
	university string
}

func NewPrototype(city string, university string) Clonable {
	return &Prototype{city: city, university: university}
}

func (p *Prototype) Clone() Clonable {
	return &Prototype{city: p.GetCity(), university: p.GetUniversity()}
}

func (p *Prototype) GetCity() string {
	return p.city
}

func (p *Prototype) SetCity() string {
	return p.city
}

func (p *Prototype) GetUniversity() string {
	return p.university
}

func (p *Prototype) SetUniversity(university string) {
	p.university = university
}
