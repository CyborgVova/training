package main

import (
	"fmt"
	"reflect"
)

type St struct {
	Name string
	Age  int8
}

func (s *St) PrintData() {
	fmt.Printf("Name: %s, Age: %v\n", s.Name, s.Age)
}

func (s *St) ReturnIntX2(x int) int {
	return x * 2
}

func GetReflectStruct(st any) {
	p := reflect.ValueOf(st)
	elem := p.Elem()

	fmt.Println("type structure:", elem.Type())

	for i := 0; i < elem.NumField(); i++ {
		fmt.Printf("field #%d: %v %v\n", i, elem.Type().Field(i).Name, elem.Field(i).Interface())
	}

	if elem.Field(1).CanSet() && elem.Field(1).Type().Kind() == reflect.Int8 {
		elem.Field(1).SetInt(44)
	}
}

func GetReflectVar(v any) {
	e := reflect.ValueOf(v).Elem()

	if e.CanSet() {
		switch e.Type().Kind() {
		case reflect.Int:
			e.SetInt(777)
		case reflect.String:
			e.SetString("Bye-Bye")
		}
	}
}

func main() {
	val := int(111)
	str := "Hello"

	st := St{
		Name: "Vova",
		Age:  43,
	}

	GetReflectStruct(&st)
	fmt.Println(st)

	GetReflectVar(&val)
	fmt.Println(val)

	GetReflectVar(&str)
	fmt.Println(str)

	p := reflect.ValueOf(&st)
	p.MethodByName("PrintData").Call([]reflect.Value{})
	rVal := []reflect.Value{reflect.ValueOf(int(21))}
	ret := p.MethodByName("ReturnIntX2").Call(rVal)
	fmt.Println(ret[0].Interface())
}
