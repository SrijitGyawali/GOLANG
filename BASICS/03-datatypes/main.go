package main

import "fmt"

func main() {
	// --- Basic types in Go ---

	// Integers (whole numbers). int sizes to your machine (usually 64-bit).
	var i int = 42
	var big int64 = 9000000000

	// Floating point (decimals).
	var price float64 = 19.99

	// Boolean (true/false).
	var loggedIn bool = true

	// String (text). Always double quotes in Go.
	var greeting string = "hi"

	// Rune = a single Unicode character (an int32 under the hood).
	var letter rune = 'A'

	// %v prints the value in a default format; %T prints its TYPE.
	fmt.Printf("%v (%T)\n", i, i)
	fmt.Printf("%v (%T)\n", big, big)
	fmt.Printf("%v (%T)\n", price, price)
	fmt.Printf("%v (%T)\n", loggedIn, loggedIn)
	fmt.Printf("%v (%T)\n", greeting, greeting)
	fmt.Printf("%v (%T) -> as char: %c\n", letter, letter, letter)

	// --- Type conversion (Go does NOT auto-convert between types) ---
	var a int = 7
	var b float64 = 2.0
	// This would FAIL: a / b  (int and float64 can't mix)
	result := float64(a) / b // convert a to float64 first
	fmt.Println("7 / 2.0 =", result)
}
