package main

import "fmt"

type Car struct {
	Brand string
	Model string
}

type Driver struct {
	Name string
	Car  Car
}

func (d Driver) String() string {
	return fmt.Sprintf("Driver: %s\nCar:\n  Brand: %s\n  Model: %s\n", d.Name, d.Car.Brand, d.Car.Model)
}

func main() {
	driver := Driver{
		Name: "Vova",
		Car: Car{
			Brand: "BMW",
			Model: "E42",
		},
	}

	fmt.Println(driver)
}
