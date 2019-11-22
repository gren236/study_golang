package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Foo struct {
	Bar *Foo
}

type Data struct {
	int
	bool
}

type Bar struct {
	Data
}

type Dog struct {
	Action func()
}

func NewPerson(name string, age int) *Person {
	p := Person{
		Name: name,
		Age: age,
	}

	return &p
}

func main() {
	// Omitted fields will be zero-valued
	fmt.Println(Person{Name: "Fred"})

	// Creating structs using positions
	fmt.Println(Person{"George", 17})

	sp := NewPerson("Jon", 15)
	fmt.Println(sp)
	fmt.Println(sp.Name)

	var foo1 Foo
	foo2 := Foo{Bar: &foo1}
	foo1.Bar = &foo2
	fmt.Println(foo1.Bar.Bar.Bar) // ... and so on :)
	fmt.Printf("type1: %T\n", foo1)

	// Structs are first-class too
	someStruct := struct {
		Foo string
		Bar float32
	}{
		Foo: "hello-world",
		Bar: 42,
	}

	fmt.Println(someStruct)
	fmt.Printf("type2: %T\n", someStruct)

	// Structs with anonymous fields
	saf := Data{}
	saf.int = 44
	saf.bool = true
	fmt.Println(saf)

	// Nested promoted fields - if nested struct is anonymous, it's fields are available at parent level
	spf := Bar{}
	spf.bool = true
	spf.int = 1997
	fmt.Println(spf)

	// Structs can also have functions in fields
	dog := Dog{}
	dog.Action = func() {
		fmt.Println("Bark!")
	}

	dog.Action()
}