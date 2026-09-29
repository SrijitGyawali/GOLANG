# Go Methods

Methods are functions that belong to a type. They let you keep data and the behavior that works with that data together.

This folder practices:

- The difference between a normal function and a method
- Receivers
- Value receivers
- Pointer receivers
- Changing a struct through a pointer receiver
- When to use `*T`

## 1. A normal function

A normal function receives a value as an argument:

```go
package main

import "fmt"

type User struct {
    Name string
}

func PrintUser(u User) {
    fmt.Println(u.Name)
}

func main() {
    u := User{Name: "Ram"}
    PrintUser(u)
}
```

`PrintUser` works with a `User`, but it is not owned by the `User` type. You call it by passing the value as an argument:

```go
PrintUser(u)
```

## 2. A method

A method is a function with a receiver. The receiver connects the function to a type:

```go
func (u User) Print() {
    fmt.Println(u.Name)
}
```

Now `Print` belongs to `User`, so it is called with dot notation:

```go
u := User{Name: "Ram"}
u.Print()
```

Think of the difference this way:

```text
function -> does something with values passed to it
method   -> behavior belonging to a type
```

Methods keep related data and behavior together.

## 3. The receiver

In this method:

```go
func (u User) Print() {
    fmt.Println(u.Name)
}
```

`u User` is the receiver:

- `u` is the receiver variable name.
- `User` is the receiver type.
- The method can access fields through `u`, such as `u.Name`.

The receiver is similar to the object used before the dot:

```go
u.Print()
```

Inside `Print`, that same value is available as `u`.

## 4. Value receivers

A value receiver receives a copy of the struct:

```go
func (u User) Rename(name string) {
    u.Name = name
}
```

Example:

```go
u := User{Name: "Ram"}
u.Rename("Hari")
fmt.Println(u.Name)
```

Output:

```text
Ram
```

Why does it still print `Ram`? The method changed its local copy of `u`, not the original `u` in `main`.

A value receiver is useful when the method should not modify the original value.

## 5. Pointer receivers

A pointer receiver receives the address of the struct:

```go
func (u *User) Rename(name string) {
    u.Name = name
}
```

Example:

```go
u := User{Name: "Ram"}
u.Rename("Hari")
fmt.Println(u.Name)
```

Output:

```text
Hari
```

This time, `Rename` changes the original `User` because `u *User` points to the original struct instead of receiving a copy.

Go automatically takes the address of an addressable variable when calling a pointer method, so this works:

```go
u.Rename("Hari")
```

It is conceptually similar to:

```go
(&u).Rename("Hari")
```

## 6. The current example: `Account.Deposit`

The runnable example in `main.go` is:

```go
type Account struct {
    Balance float64
}

func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}
```

`Deposit` uses a pointer receiver because depositing money must update the original account.

The program creates an account:

```go
a := Account{Balance: 4500.79}
```

Then it prints the starting balance:

```text
4500.79
```

Next, it deposits `50000`:

```go
a.Deposit(50000)
```

The pointer receiver updates `a.Balance` in the original account. The final balance is:

```text
54500.79
```

The calculation is:

```text
4500.79 + 50000 = 54500.79
```

## 7. When should you use a pointer receiver?

Usually use a pointer receiver, `*T`, when:

1. The method modifies the struct.
2. The struct is large and copying it would be expensive.
3. You want the type's methods to use a consistent receiver style.

For example, `Deposit` should use `*Account` because it changes `Balance`:

```go
func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}
```

Use a value receiver when the method only reads the value and copying it is acceptable:

```go
func (u User) Print() {
    fmt.Println(u.Name)
}
```

## 8. Value receiver vs pointer receiver

| Receiver | Receives | Can modify original? | Example |
|---|---|---:|---|
| `u User` | A copy | No | `Print` |
| `u *User` | Address of original | Yes | `Rename` |
| `a *Account` | Address of original | Yes | `Deposit` |

The most important question is:

> Should this method change the original value?

- No: a value receiver may be appropriate.
- Yes: use a pointer receiver.

## 9. How to run this example

From the repository root:

```powershell
go run "BASICS/014-methods/main.go"
```

Or from inside the `BASICS` folder:

```powershell
go run ./014-methods
```

Expected output:

```text
4500.79
54500.79
```

## 10. Practice exercises

### Exercise 1: Test a value receiver

Uncomment the `User` code in `main.go` and use a value receiver for `Rename`:

```go
func (u User) Rename(name string) {
    u.Name = name
}
```

Predict the output before running the program. The name should remain unchanged.

### Exercise 2: Change to a pointer receiver

Change the receiver to:

```go
func (u *User) Rename(name string) {
    u.Name = name
}
```

Run the program again. The original name should now change.

### Exercise 3: Add a withdrawal method

Add a pointer-receiver method that changes the account balance:

```go
func (a *Account) Withdraw(amount float64) {
    a.Balance -= amount
}
```

Call it from `main` and check the new balance.

### Exercise 4: Add a read-only method

Add a value-receiver method that prints the balance without changing it:

```go
func (a Account) PrintBalance() {
    fmt.Println(a.Balance)
}
```

Notice that this method only reads the account, while `Deposit` changes it.

## Key takeaway

A method is a function attached to a type. The receiver determines whether the method works with a copy or with the original value:

```go
func (u User) Print()       // value receiver: works with a copy
func (u *User) Rename(...)  // pointer receiver: can change the original
```

For methods that modify a struct, use a pointer receiver so the change remains after the method returns.
