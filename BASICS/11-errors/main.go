package main

import (
	"errors"
	"fmt"
)

// Go has NO exceptions. Instead, functions return an "error" as their
// last value. If err is nil, everything went fine.

// A sentinel error: a named error value you can compare against.
var ErrDivideByZero = errors.New("cannot divide by zero")

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil // nil means "no error"
}

// You can also build errors with context using fmt.Errorf.
func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, fmt.Errorf("insufficient funds: have %d, need %d", balance, amount)
	}
	return balance - amount, nil
}

func main() {
	// The standard Go pattern you'll write thousands of times:
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	_, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	}

	// errors.Is checks whether an error matches a known sentinel.
	if errors.Is(err, ErrDivideByZero) {
		fmt.Println("(and it was specifically a divide-by-zero)")
	}

	newBalance, err := withdraw(100, 150)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("balance:", newBalance)
}
