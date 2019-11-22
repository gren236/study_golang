package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

type Cubic interface {
	Diagonal() float64
}

// Nested interface combines multiple interfaces
type Shape interface {
	Geometry
	Cubic
}

type Rect struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perim() float64 {
	return 2 * r.Width + 2 * r.Height
}

func (r *Rect) Diagonal() float64 {
	return math.Sqrt(r.Width * r.Width + r.Height * r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}

type AnyType interface{}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func sum(a, b AnyType) AnyType {
	return a.(int) + b.(int)
}

func main() {
	// Geometry is static type, but it's a storage for "Type => value"
	// Now it's nil => nil
	var s Geometry
	fmt.Println(s)

	// Here it makes r Rect => struct{...}
	r := Rect{12.1, 4}
	c := Circle{12.3}

	measure(r)
	measure(c)

	fmt.Printf("Type of r is %T\n", r)

	s = &r
	r.Height = 5
	fmt.Println(s)

	// Dynamic function
	fmt.Println(sum(4, 5))

	// Type assertion - extracting struct from interface
	// fmt.Println(s.Diagonal) - wont work because it holds a static value of interface Geometry
	fmt.Println(s.(*Rect).Diagonal())

	// Use type switch to make compoarisons of types!

	// Nested interface
	var shape Shape
	shape = &Rect{15, 4}
	// shape = Circle{12} Won't work as Circle not inplementing Diagonal() method
	fmt.Println(shape.Diagonal())

	// If one method has pointer receiver - than you should always pass concrete types as pointers!
	// shape = Rect{15, 4} - won't work!
}
