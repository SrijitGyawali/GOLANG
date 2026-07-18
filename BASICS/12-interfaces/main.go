package main

import "fmt"

// An interface defines BEHAVIOR (a set of methods), not data.
// Any type that has these methods automatically satisfies the interface —
// you never write "implements". This is called structural typing.

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// Rectangle has an Area() method -> it IS a Shape.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle has an Area() method -> it IS a Shape too.
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// This function accepts ANY Shape, without knowing the concrete type.
func printArea(s Shape) {
	fmt.Printf("%T area = %.2f\n", s, s.Area())
}

func main() {
	printArea(Rectangle{Width: 3, Height: 4})
	printArea(Circle{Radius: 5})

	// A slice of the interface type can hold different concrete types.
	shapes := []Shape{
		Rectangle{Width: 2, Height: 2},
		Circle{Radius: 1},
	}
	for _, s := range shapes {
		printArea(s)
	}
}
