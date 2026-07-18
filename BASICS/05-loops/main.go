package main

import "fmt"

func main() {
	// Go has ONLY ONE loop keyword: "for".
	// It covers everything other languages do with for/while/do-while.

	// --- 1) Classic for loop: init; condition; post ---
	for i := 1; i <= 5; i++ {
		fmt.Println("count:", i)
	}

	// --- 2) "while" style: only a condition ---
	n := 3
	for n > 0 {
		fmt.Println("n =", n)
		n-- // decrease n, otherwise this loops forever
	}

	// --- 3) Infinite loop with break ---
	total := 0
	for {
		total++
		if total == 3 {
			break // exit the loop
		}
	}
	fmt.Println("total:", total)

	// --- 4) continue: skip the rest of THIS iteration ---
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("odd:", i)
	}

	// --- 5) range: loop over a collection ---
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("%d -> %s\n", index, fruit)
	}
	// If you don't need the index, use _ (the blank identifier).
	for _, fruit := range fruits {
		fmt.Println("fruit:", fruit)
	}
}
