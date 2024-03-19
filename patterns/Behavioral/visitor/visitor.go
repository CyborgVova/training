package visitor

type Visitor interface {
	VisitToZoo(*Zoo) string
	VisitToCinema(*Cinema) string
	VisitToCircus(*Circus) string
}

type Place interface {
	Accept(Visitor) string
}

type Zoo struct{}

func (z *Zoo) Accept(v Visitor) string {
	return v.VisitToZoo(z) + "visited the Zoo\n"
}

type Cinema struct{}

func (c *Cinema) Accept(v Visitor) string {
	return v.VisitToCinema(c) + "visited the Cinema\n"
}

type Circus struct{}

func (c *Circus) Accept(v Visitor) string {
	return v.VisitToCircus(c) + "visited the Circus\n"
}

type City struct {
	places []Place
}

func (c *City) AddingPlaces(places ...Place) {
	c.places = append(c.places, places...)
}

func (c *City) Accept(v Visitor) (result string) {
	for _, place := range c.places {
		result += place.Accept(v)
	}
	return
}

type Vasja struct{}

func (vasja *Vasja) VisitToZoo(z *Zoo) string {
	return "Vasja "
}

func (vasja *Vasja) VisitToCinema(c *Cinema) string {
	return "Vasja "
}

func (vasja *Vasja) VisitToCircus(c *Circus) string {
	return "Vasja "
}

type Fedja struct{}

func (fedja *Fedja) VisitToZoo(z *Zoo) string {
	return "Fedja "
}

func (fedja *Fedja) VisitToCinema(c *Cinema) string {
	return "Fedja "
}

func (fedja *Fedja) VisitToCircus(c *Circus) string {
	return "Fedja "
}

type Sasha struct{}

func (sasha *Sasha) VisitToZoo(z *Zoo) string {
	return "Sasha "
}

func (sasha *Sasha) VisitToCinema(c *Cinema) string {
	return "Sasha "
}

func (sasha *Sasha) VisitToCircus(c *Circus) string {
	return "Sasha "
}
