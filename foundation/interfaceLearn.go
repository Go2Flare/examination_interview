package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Object interface {
	Perimeter() float32
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * (c.radius * c.radius)
}

func main() {
	var s Shape = Circle{3}
	value1, ok1 := s.(Shape)
	value2, ok2 := s.(Object)

	fmt.Println(value1, ok1)
	fmt.Printf("%T\t", s.(Shape).(Shape).(Shape).(Shape))
	fmt.Println(value2, ok2)
}
