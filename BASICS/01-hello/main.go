// Package main is the entry point of an executable Go program.
// Every runnable Go program MUST have a package named "main"
// and a function named "main" inside it.
package main

// "fmt" is Go's standard formatting/printing library.
// import brings other packages into this file.
import "fmt"

func main() {
	// Println prints its arguments and adds a newline at the end.
	fmt.Println("Hello, Go!")

	// Print does NOT add a newline.
	fmt.Print("No newline here. ")
	fmt.Print("Still same line.\n") // \n is a manual newline

	// Printf lets you format values. %s = string, %d = integer, \n = newline.
	name := "PRa"
	age := 1 // "age" of your Go journey, in weeks :)
	fmt.Printf("Name: %s, learning Go for %d week(s)\n", name, age)
}
