package decorator

type Class interface {
	Print(string) string
}

type Object struct{}

func (o *Object) Print(text string) string {
	return text
}

type Decorator struct {
	Apostrof string
	Object   *Object
}

func (d *Decorator) Print(text string) string {
	return d.Apostrof + d.Object.Print(text) + d.Apostrof
}
