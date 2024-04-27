package polimorfizm

import "fmt"

type Entity struct {
	Animal Animal
}

func NewEntity(animal Animal) *Entity {
	return &Entity{Animal: animal}
}

func (e *Entity) MakeVoice() {
	e.Animal.Voice()
}

type Cat struct {
	Name string
}

func NewCat(name string) *Cat {
	return &Cat{Name: name}
}

func (c *Cat) Voice() {
	fmt.Printf("cat '%s' say: Meoow\n", c.Name)
}

type Dog struct {
	Name string
}

func NewDog(name string) *Dog {
	return &Dog{Name: name}
}

func (d *Dog) Voice() {
	fmt.Printf("dog '%s' say: Woof-Woof\n", d.Name)
}
