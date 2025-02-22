// This Go program demonstrates the use of variables in Go.
// Variables in Go are explicitly declared and used to store data of various types.
// The `var` keyword is used to declare a variable, followed by the variable name and type.
// Go also supports type inference, where the type of the variable is determined by the value assigned to it.
// Additionally, Go has shorthand syntax for variable declaration and initialization using the `:=` operator.
// This program will cover different ways to declare and initialize variables, as well as the scope and lifetime of variables.
package main

import "fmt"

func main() {
	// Declaring a variable with explicit type
	var name string
	name = "John Doe"
	fmt.Println("Name:", name)

	// Declaring and initializing a variable with explicit type
	var age int = 30
	fmt.Println("Age:", age)

	// Type inference during variable declaration and initialization
	var height = 5.9
	fmt.Println("Height:", height)

	// Shorthand syntax for variable declaration and initialization
	weight := 70.5
	fmt.Println("Weight:", weight)

	// Multiple variable declaration and initialization
	var city, country = "New York", "USA"
	fmt.Println("City:", city, "Country:", country)

	// Declaring multiple variables of different types
	var (
		isEmployed bool    = true
		salary     float64 = 50000.0
	)
	fmt.Println("Employed:", isEmployed, "Salary:", salary)

	// Declaring a constant
	const pi = 3.14
	fmt.Println("Pi:", pi)

	// Using iota for enumerated constants
	const (
		Red = iota
		Green
		Blue
	)
	fmt.Println("Red:", Red, "Green:", Green, "Blue:", Blue)

	// Using variables in a function
	fmt.Println("Sum of 5 and 3:", add(5, 3))

	// Shadowing variables
	x := 10
	fmt.Println("Outer x:", x)
	{
		x := 20
		fmt.Println("Inner x:", x)
	}
	fmt.Println("Outer x after inner scope:", x)
}

// Function to add two integers
func add(a int, b int) int {
	return a + b
}