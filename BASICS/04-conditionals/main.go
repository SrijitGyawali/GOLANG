package main

import "fmt"

func main() {
	score := 72

	// --- if / else if / else ---
	// Note: no parentheses around the condition, but braces {} are required.
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 60 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: F")
	}

	// --- if with a short statement ---
	// You can declare a variable that only exists inside the if/else.
	if remainder := score % 2; remainder == 0 {
		fmt.Println(score, "is even")
	} else {
		fmt.Println(score, "is odd")
	}

	// --- switch: cleaner than a long if/else chain ---
	day := "Sat"
	switch day {
	case "Sat", "Sun": // multiple values in one case
		fmt.Println("Weekend")
	case "Mon":
		fmt.Println("Start of the week")
	default:
		fmt.Println("A weekday")
	}

	// --- switch with no value acts like if/else ---
	switch {
	case score > 50:
		fmt.Println("Passed")
	default:
		fmt.Println("Failed")
	}
}
