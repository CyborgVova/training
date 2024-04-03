package flyweight

import (
	"errors"
	"math"
)

type Shape interface {
	Area() float64
}

type FlyWeight struct {
	Elements map[string]Shape
}

func NewFlyWeight() *FlyWeight {
	return &FlyWeight{Elements: map[string]Shape{}}
}

func (f *FlyWeight) Add(name string, shape Shape) error {
	if _, ok := f.Elements[name]; ok {
		return errors.New("such element exist yet")
	}
	f.Elements[name] = shape
	return nil
}

func (f *FlyWeight) GetElement(name string) (Shape, error) {
	if _, ok := f.Elements[name]; !ok {
		return nil, errors.New("such element not exist")
	}
	return f.Elements[name], nil
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Rectangle struct {
	SideA float64
	SideB float64
}

func (r *Rectangle) Area() float64 {
	return r.SideA * r.SideB
}

type Treeangle struct {
	Side   float64
	Height float64
}

func (t *Treeangle) Area() float64 {
	return t.Side * t.Height * 0.5
}
