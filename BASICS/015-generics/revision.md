# Go Generics — Interview Revision

A quick revision note for Go generics.

## 1. What are Generics?

Generics let us write **one reusable function or type for multiple data types** while keeping **compile-time type safety**.

```go
func Add[T int | float64](a, b T) T {
    return a + b
}

fmt.Println(Add(10, 20))
fmt.Println(Add(10.5, 20.5))
```

`T` = **type parameter**

---

## 2. Basic Syntax

```go
func Print[T any](value T) {
    fmt.Println(value)
}
```

`[T any]` means `T` can be any type.

Go usually infers the type:

```go
Print(10)
Print("hello")
```

---

## 3. Constraints

Constraints define which types `T` can represent.

```go
type Number interface {
    int | int64 | float64
}

func Sum[T Number](a, b T) T {
    return a + b
}
```

This fails:

```go
func Add[T any](a, b T) T {
    return a + b // ❌
}
```

because `any` does not guarantee that `+` is supported.

---

## 4. `any` vs `comparable`

```go
any
```

= any Go type.

```go
comparable
```

= types that support `==` and `!=`.

```go
func Equal[T comparable](a, b T) bool {
    return a == b
}
```

---

## 5. `~` Operator

```go
type Number interface {
    ~int | ~float64
}
```

`~int` means `int` plus custom types whose underlying type is `int`.

```go
type Age int

func Double[T ~int](v T) T {
    return v * 2
}

Double(Age(20)) // ✅
```

Remember:

```text
int  → exact int
~int → int + custom types based on int
```

---

## 6. Generic Struct

```go
type Box[T any] struct {
    Value T
}

a := Box[int]{Value: 10}
b := Box[string]{Value: "Go"}
```

---

## 7. Useful Backend Example

```go
type Response[T any] struct {
    Data    T
    Message string
}
```

Usage:

```go
Response[User]
Response[[]User]
Response[Product]
```

---

## 8. Generic Slice Helper

```go
func First[T any](items []T) T {
    return items[0]
}
```

Works with:

```go
First([]int{1, 2, 3})
First([]string{"Go", "Java"})
```

---

## 9. Generic Map Example

```go
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))

    for k := range m {
        keys = append(keys, k)
    }

    return keys
}
```

`K` must be `comparable` because Go map keys must be comparable.

---

## 10. Generics vs Interfaces

### Interface

```go
type Writer interface {
    Write([]byte) error
}
```

Answers:

> **What behavior can this type perform?**

### Generics

```go
func First[T any](items []T) T
```

Answers:

> **Can the same algorithm work with several types?**

Remember:

```text
Interface → behavior abstraction
Generics  → type-safe code reuse
```

---

## 11. Important Interview Traps

### ❌ `any` does not guarantee `+`

```go
func Add[T any](a, b T) T {
    return a + b
}
```

### ❌ `any` does not guarantee `==`

```go
func Equal[T any](a, b T) bool {
    return a == b
}
```

Use:

```go
func Equal[T comparable](a, b T) bool {
    return a == b
}
```

### Generic methods cannot introduce their own new type parameters

Allowed:

```go
type Box[T any] struct {
    Value T
}

func (b Box[T]) Get() T {
    return b.Value
}
```

Not allowed:

```go
func (b Box[T]) Convert[U any]() U { // ❌
}
```

---

## 12. When Should I Use Generics?

Good use cases:

```text
✓ Slice/map helpers
✓ Stack / Queue
✓ Algorithms
✓ Generic containers
✓ Typed API responses
✓ Reusable libraries
```

Avoid generics when only one concrete type is actually needed.

---

# Interview Quick Recall

```text
Generics
→ reusable code across multiple types

T
→ type parameter

[T any]
→ any type

[T comparable]
→ supports == and !=

int | float64
→ type constraint / type set

~int
→ types whose underlying type is int

Generic Function
→ func Add[T Number](a, b T) T

Generic Type
→ type Box[T any] struct {}

Interface
→ common behavior

Generics
→ common algorithm/type logic
```

## One-Line Interview Answer

> **Go generics allow functions and types to work with multiple types using type parameters and constraints while preserving compile-time type safety.**
