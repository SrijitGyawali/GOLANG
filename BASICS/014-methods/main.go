package main

import (
	"fmt"
)

type User struct {
	Name string
}

// // Value Type
// func (u User) Print() {
// 	fmt.Println(u.Name)
// }

// // refrence type
// func (u *User) Rename(name string) {
// 	u.Name = name
// }

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func main() {
	// u := User{Name: "Alice"}
	// u.Print()
	// u.Rename("Srijit")
	// u.Print()

	a := Account{Balance: 4500.79}
	fmt.Println(a.Balance)
	a.Deposit(50000)
	fmt.Println(a.Balance)

}
