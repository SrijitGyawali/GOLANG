package main

import "fmt"

func main() {
	// --- Declaring variables ---

	// 1) Full form: var <name> <type> = <value>
	var country string = "Nepal"

	// 2) Type inference: Go figures out the type from the value.
	var city = "Kathmandu"

	// 3) Short declaration ":=" (only usable INSIDE functions).
	//    This is the most common style you'll see in Go code.
	language := "Go"

	// 4) Declare without a value -> gets the "zero value".
	//    Numbers -> 0, strings -> "", bool -> false.
	var count int
	var isReady bool

	fmt.Println(country, city, language, count, isReady)

	// --- Reassigning (no ":=" the second time, just "=") ---
	count = 5
	isReady = true
	fmt.Println("count:", count, "isReady:", isReady)

	// --- Multiple variables at once ---
	x, y := 10, 20
	fmt.Println("x + y =", x+y)

	// --- Constants: values that never change ---
	const pi = 3.14159
	fmt.Println("pi:", pi)
}
