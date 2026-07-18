package main

import "fmt"

// A struct groups related fields together — Go's version of a "class"
// (but without inheritance). This is how you model real things.
type User struct {
	Name  string
	Email string
	Age   int
}

// A METHOD is a function attached to a type.
// (u User) is the "receiver" — it makes Greet belong to User.
func (u User) Greet() string {
	return "Hi, I'm " + u.Name
}

// A pointer receiver (*User) lets the method MODIFY the struct.
func (u *User) HaveBirthday() {
	u.Age++
}

func main() {
	// Create a struct value using field names (clearest way).
	alice := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   30,
	}

	// Access fields with a dot.
	fmt.Println(alice.Name, "-", alice.Email)

	// Call a method.
	fmt.Println(alice.Greet())

	// Modify via a pointer method.
	alice.HaveBirthday()
	fmt.Println("After birthday, age:", alice.Age)

	// Zero-value struct: every field starts at its zero value.
	var empty User
	fmt.Printf("empty user: %+v\n", empty) // %+v shows field names

	// Structs are copied by value. Changing a copy doesn't touch the original.
	copyOfAlice := alice
	copyOfAlice.Name = "Not Alice"
	fmt.Println("original still:", alice.Name)
}
