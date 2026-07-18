package main

import "fmt"

func main() {
	// A map stores key -> value pairs (like a dictionary / hash table).
	// map[KeyType]ValueType

	// Create a map with some values.
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("map:", ages)

	// Read a value by key.
	fmt.Println("Alice is", ages["Alice"])

	// Add or update.
	ages["Charlie"] = 40
	ages["Bob"] = 26
	fmt.Println("after updates:", ages)

	// Delete a key.
	delete(ages, "Alice")
	fmt.Println("after delete:", ages)

	// --- The "comma ok" idiom: check if a key exists ---
	// Reading a missing key returns the zero value (0 here), which is
	// ambiguous. The second value tells you if the key was really there.
	age, exists := ages["Alice"]
	fmt.Println("Alice age:", age, "exists:", exists)

	// An empty map you can add to later:
	counts := make(map[string]int)
	counts["clicks"]++ // missing key starts at 0, then becomes 1
	counts["clicks"]++
	fmt.Println("clicks:", counts["clicks"])

	// Loop over a map (order is RANDOM in Go — never rely on it).
	for name, a := range ages {
		fmt.Printf("%s -> %d\n", name, a)
	}
}
