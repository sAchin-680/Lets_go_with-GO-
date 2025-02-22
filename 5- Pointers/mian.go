// Package pointers provides utilities for working with pointers in Go.

// Pointers in Go are variables that store the memory address of another variable.
// They are useful for passing references to values and data structures, allowing
// functions to modify the original value or structure.

// In Go, a pointer is declared using the `*` operator followed by the type of the value
// it points to. For example, `*int` is a pointer to an integer.

// To get the memory address of a variable, use the `&` operator. For example, `&x`
// returns a pointer to the variable `x`.

// To access the value stored at a pointer, use the `*` operator. This is known as
// dereferencing the pointer. For example, `*p` returns the value stored at the memory
// address `p` points to.

// Example usage:

    package main

    import "fmt"

    func main() {
        var x int = 10
        var p *int = &x // p is a pointer to x

        fmt.Println("Value of x:", x)   // Output: Value of x: 10
        fmt.Println("Address of x:", p) // Output: Address of x: 0xc0000140a0 (example address)

        *p = 20 // Modify the value of x through the pointer p
        fmt.Println("New value of x:", x) // Output: New value of x: 20
    }

// Pointers are particularly useful when working with large data structures or when
// you need to modify the value of a variable inside a function. By passing a pointer
// to the function, you can ensure that the function operates on the original value
// rather than a copy.

// Note that Go does not support pointer arithmetic, which is common in languages like C.
// This design choice helps prevent common programming errors and improves code safety.