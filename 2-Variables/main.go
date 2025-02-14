package main
import "fmt"

// main is the entry point of the program. It demonstrates variable declaration,
// type inference, short variable declaration, and constants in Go. The function
// prints the values and types of various variables to the console.
func main() {
		var username string = "John Doe"
		fmt.Println(username)
		fmt.Printf("Variable username is of type: %T \n", username)

		// Variable declaration
		var age int = 30
		fmt.Println("Age:", age)

		// Type inference
		var isEmployed = true
		fmt.Println("Is Employed:", isEmployed)

		// Short variable declaration
		country := "USA"
		fmt.Println("Country:", country)

		// Constants
		const pi = 3.14
		fmt.Println("Pi:", pi)

		const (
			companyName = "Tech Corp"
			yearFounded = 1998
		)
		fmt.Println("Company Name:", companyName)
		fmt.Println("Year Founded:", yearFounded)
	}
