package main

import "fmt"

// A basic function: name, parameters (with types), and a return type.
func add(a int, b int) int {
	return a + b
}

// When parameters share a type, you can write it once.
func multiply(a, b int) int {
	return a * b
}

// Go can return MULTIPLE values. This is used everywhere (especially errors).
func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false // false = "not ok"
	}
	return a / b, true
}

// Named return values + a "variadic" parameter (zero or more ints).
func sum(numbers ...int) (result int) {
	for _, n := range numbers {
		result += n
	}
	return // returns "result" automatically
}

func main() {
	fmt.Println("add:", add(2, 3))
	fmt.Println("multiply:", multiply(4, 5))

	quotient, ok := divide(10, 2)
	fmt.Println("divide 10/2:", quotient, "ok:", ok)

	_, ok = divide(10, 0)
	fmt.Println("divide 10/0 ok:", ok)

	fmt.Println("sum:", sum(1, 2, 3, 4, 5))

	// Functions are values too — you can store one in a variable.
	square := func(x int) int { return x * x }
	fmt.Println("square(6):", square(6))
}
