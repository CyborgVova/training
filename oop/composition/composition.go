package main

import "fmt"

type Car struct {
	Model string
	Color string
	Engine
}

type Engine struct {
	Volume float64
	Fuel   string
}

func (c Car) String() string {
	return fmt.Sprintf("Car:\n  Model: %s\n  Color: %s\n  Engine:\n    Volume: %.1f\n    Fuel: %s",
		c.Model, c.Color, c.Volume, c.Fuel)
}

func main() {
	car := Car{
		Model: "BMW",
		Color: "red",
		Engine: Engine{
			Volume: 4.2,
			Fuel:   "petrol",
		},
	}

	fmt.Println(car)
}
