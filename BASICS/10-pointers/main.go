package main

import "fmt"

// A pointer holds the ADDRESS of a value, not the value itself.
// Two key symbols:
//   &x  -> "address of x"  (get a pointer)
//   *p  -> "value at p"    (follow the pointer)

// Pass by value: gets a COPY, so the original is unchanged.
func tryToDouble(n int) {
	n = n * 2
}

// Pass by pointer: can change the original.
func double(n *int) {
	*n = *n * 2 // follow the pointer and update what it points to
}

func main() {
	x := 10

	tryToDouble(x)
	fmt.Println("after tryToDouble:", x) // still 10

	double(&x) // pass the ADDRESS of x
	fmt.Println("after double:", x) // now 20

	// Looking at a pointer directly:
	p := &x
	fmt.Println("address (p):", p)  // something like 0xc0000...
	fmt.Println("value (*p):", *p)  // 20

	*p = 99 // change x through the pointer
	fmt.Println("x is now:", x)

	// nil is the zero value of a pointer (points to nothing).
	var q *int
	fmt.Println("nil pointer:", q)
}
