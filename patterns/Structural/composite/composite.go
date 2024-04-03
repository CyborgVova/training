package composite

type Composite interface {
	Add(...Composite)
	Remove(...Composite)
	Print() string
}

type Word struct {
	Composites []Composite
}

func (w *Word) Add(composite ...Composite) {
	w.Composites = append(w.Composites, composite...)
}

func (w *Word) Remove(composite ...Composite) {
	for _, c := range composite {
		for i, word := range w.Composites {
			if word == c {
				w.Composites[i] = nil
				break
			}
		}
	}
}

func (w *Word) Print() (result string) {
	for _, text := range w.Composites {
		switch v := text.(type) {
		case *Character:
			result += v.Print()
		case *Word:
			for _, ch := range v.Composites {
				result += ch.Print()
			}
		}
	}
	return
}

type Character struct {
	Character rune
}

func (c *Character) Add(composite ...Composite) {}

func (c *Character) Remove(composite ...Composite) {}

func (c *Character) Print() string {
	return (string)(c.Character)
}
