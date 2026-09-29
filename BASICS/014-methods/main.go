package methods

import (
	"fmt"
)

type User struct {
	Name string
}

// Value Type
func (u User) Print() {
	fmt.Println(u.Name)
}

// refrence type
func (u *User) Rename(name string) {
	u.Name = name
}

func main() {
	u := User{Name: "Alice"}
	u.Print()
	u.Rename("Srijit")
	u.Print()

}
