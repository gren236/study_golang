package main

import (
	"fmt"
	"math"
)

// Figures stuff
type Rect struct {
	width, height int
}

func (r *Rect) Area() int {
	return r.width * r.height
}

func (r Rect) Perim() int {
	return 2*r.width + 2*r.height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Employee stuff
type Employee struct {
	FirstName, LastName string
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) UpdateFirstName(name string) {
	e.FirstName = name
}

// Misc
type SomeString string

// Methods can be declared for ANY type from the same package!
func (s *SomeString) Update(str string) {
	*s = SomeString(str)
}

func main() {
	r := Rect{width: 10, height: 5}

	fmt.Println("area: ", r.Area())
	fmt.Println("perim: ", r.Perim())

	// Go handles type conversions for pointers by itself
	rp := &r
	fmt.Println("area: ", rp.Area())
	fmt.Println("perim:", rp.Perim())

	// You can name methods the same for different structs
	c := Circle{32.5}
	fmt.Println(c.Area())

	emp := Employee{"George", "Smith"}
	fmt.Println(emp.FullName())
	emp.UpdateFirstName("John")
	fmt.Println(emp.FullName())

	// We can pass pointer to value receiver
	fmt.Println((&emp).FullName())

	// Usage of methods on non-structs
	var somes SomeString
	somes = "Hello!"
	somes.Update("Foo!")
	fmt.Println(somes)
}
