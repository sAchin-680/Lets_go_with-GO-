package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "Welcome to Go!"
	fmt.Println(Welcome)
	fmt.Print("Enter your name: ")

	// Create a new reader to read input from the standard input
	reader := bufio.NewReader(os.Stdin)
	
	// Read the input until a newline character is encountered
	name, err := reader.ReadString('\n')
	
	// Check if there was an error while reading the input
	if err != nil {
		// Print the error message and exit
		fmt.Println("An error occurred:", err)
		return
	}

	// Print a personalized greeting
	fmt.Printf("Hello, %s", name)
}