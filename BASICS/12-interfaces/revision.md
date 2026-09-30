# Go Methods & Interfaces — Interview Revision Snippet

## 1. Method = Function attached to a type

```go
type User struct {
    Name string
}

func (u User) Print() {
    fmt.Println(u.Name)
}
```

**Remember:**  
`func (receiver Type) MethodName()` → method has a **receiver**.

---

## 2. Value Receiver vs Pointer Receiver

```go
func (u User) Rename(name string) {
    u.Name = name // changes copy only
}

func (u *User) Rename(name string) {
    u.Name = name // changes original
}
```

### Use pointer receiver `*T` when:
- You need to modify the value.
- Struct is large and copying is expensive.
- You want consistent method sets.

### Quick rule
```text
T receiver   → gets a copy
*T receiver  → works with original value
```

---

## 3. Interface = Behavior Contract

```go
type Speaker interface {
    Speak() string
}
```

Any type with:

```go
Speak() string
```

automatically satisfies `Speaker`.

```go
type Dog struct{}

func (Dog) Speak() string {
    return "Woof"
}
```

**No `implements` keyword in Go.**

---

## 4. Why Interfaces Matter

Interfaces provide:

- Loose coupling
- Dependency injection
- Easier testing / mocking
- Multiple implementations
- Behavior-based abstraction

```go
type Sender interface {
    Send() error
}

func Notify(s Sender) error {
    return s.Send()
}
```

Anything implementing `Send() error` can be passed to `Notify`.

---

## 5. Real Backend Pattern

```go
type UserRepository interface {
    FindByID(id int) (*User, error)
}

type UserService struct {
    repo UserRepository
}
```

Production implementation:

```go
type PostgresRepo struct{}

func (p *PostgresRepo) FindByID(id int) (*User, error) {
    return nil, nil
}
```

Testing implementation:

```go
type MockRepo struct{}

func (m *MockRepo) FindByID(id int) (*User, error) {
    return &User{Name: "Test"}, nil
}
```

**Key idea:** service depends on behavior, not PostgreSQL directly.

---

# INTERVIEW TRAPS

## 6. Method Set — VERY IMPORTANT

```go
type Animal interface {
    Speak()
}

type Dog struct{}

func (d *Dog) Speak() {}
```

This fails:

```go
var a Animal = Dog{} // ❌
```

This works:

```go
var a Animal = &Dog{} // ✅
```

### Memorize

```text
func (T) Method()
→ T and *T have the method

func (*T) Method()
→ only *T has the method for interface satisfaction
```

---

## 7. Automatic Addressing Can Trick You

```go
d := Dog{}
d.Speak()
```

May work even when:

```go
func (d *Dog) Speak()
```

because Go can treat it roughly like:

```go
(&d).Speak()
```

BUT:

```go
var a Animal = d
```

can still fail because interface assignment follows the **method set** rules.

---

## 8. Interface With Multiple Methods

```go
type Storage interface {
    Save(string) error
    Get(int) (string, error)
    Delete(int) error
}
```

A type must implement **every method** to satisfy the interface.

---

## 9. Prefer Small Interfaces

Better:

```go
type Reader interface {
    Read([]byte) (int, error)
}
```

instead of giant interfaces with many unrelated methods.

### Good Go principle

```text
Accept interfaces, return concrete types.
```

Useful default, not an absolute law.

---

## 10. Interface Embedding

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

`ReadWriter` requires both `Read()` and `Write()`.

---

## 11. `any` and `interface{}`

These are equivalent:

```go
any
interface{}
```

Example:

```go
var x any

x = 10
x = "hello"
x = User{}
```

Use `any` only when you genuinely accept arbitrary types.

---

## 12. Type Assertion

```go
var x any = "hello"

s, ok := x.(string)

if ok {
    fmt.Println(s)
}
```

Unsafe:

```go
s := x.(string)
```

Can panic if the value is not a string.

### Memorize

```go
value, ok := interfaceValue.(ConcreteType)
```

---

## 13. Type Switch

```go
func Print(v any) {
    switch x := v.(type) {
    case int:
        fmt.Println("int:", x)

    case string:
        fmt.Println("string:", x)

    default:
        fmt.Println("unknown")
    }
}
```

Special syntax:

```go
v.(type)
```

Only works inside a type switch.

---

# CLASSIC NIL INTERFACE TRAP

## 14. Interface Can Hold a Nil Pointer and Still Be Non-Nil

```go
var p *User = nil
var x any = p

fmt.Println(x == nil)
```

Output:

```text
false
```

Why?

Conceptually, an interface stores:

```text
(dynamic type, dynamic value)
```

Here:

```text
(*User, nil)
```

A truly nil interface is:

```text
(nil, nil)
```

Example:

```go
var x any
fmt.Println(x == nil) // true
```

---

## 15. `error` Nil Trap

```go
type MyError struct{}

func (*MyError) Error() string {
    return "error"
}

func doSomething() error {
    var err *MyError = nil
    return err
}
```

Then:

```go
err := doSomething()
fmt.Println(err == nil)
```

Output:

```text
false
```

Because the interface contains:

```text
(*MyError, nil)
```

Correct when there is no error:

```go
return nil
```

---

## 16. Compile-Time Interface Check

```go
var _ io.Reader = (*MyReader)(nil)
```

Meaning:

> Compiler: verify that `*MyReader` implements `io.Reader`.

Useful production pattern:

```go
var _ UserRepository = (*PostgresRepository)(nil)
```

---

## 17. Define Interfaces Near the Consumer

Instead of making a huge repository interface in the implementation package:

```go
type UserFinder interface {
    FindByID(int) (*User, error)
}
```

Define the small interface where the behavior is actually needed.

### Remember

```text
Interfaces describe what the consumer needs.
```

---

## 18. Methods Are Not Only for Structs

```go
type Celsius float64

func (c Celsius) Fahrenheit() float64 {
    return float64(c)*9/5 + 32
}
```

Methods can be attached to named types defined in your package.

---

## 19. Cannot Add Methods to Types You Don't Own

Invalid:

```go
func (s string) Hello() {} // ❌
```

Valid:

```go
type MyString string

func (s MyString) Hello() {} // ✅
```

---

# FAST INTERVIEW Q&A

### What is a method?
A function with a receiver attached to a type.

### Function vs method?
A method has a receiver; a normal function does not.

### Value receiver vs pointer receiver?
Value receiver gets a copy. Pointer receiver can mutate the original.

### Does Go use `implements`?
No. Interface implementation is implicit.

### What is a method set?
The methods available on `T` and `*T`, especially important for interface satisfaction.

### Can `T` satisfy an interface if the required method is defined on `*T`?
No. `*T` can; `T` cannot.

### Why use interfaces?
Decoupling, mocking, testing, multiple implementations, dependency injection.

### What is `any`?
Alias for `interface{}`.

### What is a type assertion?
Extracting a concrete type from an interface.

```go
v, ok := x.(MyType)
```

### What is a type switch?
Checks which concrete type an interface currently contains.

### Why can a nil pointer inside an interface make the interface non-nil?
Because the interface still contains a dynamic type.

### Does Go have inheritance?
Not traditional class inheritance. Go favors **composition + interfaces**.

---

# FINAL MENTAL MODEL

```text
TYPE
  ↓
contains data

METHOD
  ↓
defines behavior for the type

INTERFACE
  ↓
describes required behavior

CONCRETE TYPE
  ↓
implicitly satisfies interface
when its method set matches
```

Example:

```go
type PaymentProcessor interface {
    Pay(amount float64) error
}
```

Your code should think:

```text
"I do not care whether you are Stripe,
PayPal, Bank, Crypto, or a Mock.

Can you Pay()?"
```

That is the core idea behind Go interfaces.

---

# 30-SECOND REVISION

```text
METHOD
→ function + receiver

VALUE RECEIVER
→ copy

POINTER RECEIVER
→ original / mutation

INTERFACE
→ behavior contract

IMPLEMENTATION
→ implicit

T receiver method
→ available to T and *T

*T receiver method
→ interface satisfied only by *T

any
→ interface{}

TYPE ASSERTION
→ v, ok := x.(Type)

TYPE SWITCH
→ switch v := x.(type)

NIL INTERFACE
→ nil only when BOTH dynamic type and value are nil

GOOD GO DESIGN
→ small interfaces
→ consumer-defined interfaces
→ composition over inheritance
→ depend on behavior, not concrete implementation
```
