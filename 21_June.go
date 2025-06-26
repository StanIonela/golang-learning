package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//func measurement(s Shape) {
//	fmt.Println(s)
//	fmt.Println(s.area())
//	if c, ok := s.(Circle); ok {
//		fmt.Println("Circle with radius", c.radius)
//	}

//if r, ok := s.(Rectangle); ok {
//fmt.Printf("Rectangle with width %d and height %d", r.width, r.height)
//}
//}

func Inspect(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Printf("Circle with area %.2f\n", v.area())
	case Rectangle:
		fmt.Printf("Rectangle with area %2f\n", v.area())

	}
}

func main() {
	/*	r := Rectangle{width: 3, height: 4}
		c := Circle{radius: 5}
		//	measurement(r)
		//measurement(c)

		fmt.Println("Circle")
		Inspect(c)

		fmt.Println("Rectangle")
		Inspect(r)*/

	shapes := []Shape{
		Circle{radius: 6},
		Rectangle{width: 3, height: 4},
		Rectangle{width: 5, height: 6},
		Circle{2},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d\n", i+1)
		Inspect(shape)
	}
}
