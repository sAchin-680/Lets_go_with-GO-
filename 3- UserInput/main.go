package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Example 1: Reading a string input
	fmt.Print("Enter your name: ")
	name, err := readString(reader)
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	fmt.Printf("Hello, %s!\n", name)

	// Example 2: Reading an integer input
	fmt.Print("Enter your age: ")
	age, err := readInt(reader)
	if err != nil {
		fmt.Println("Invalid input for age:", err)
		return
	}
	fmt.Printf("You are %d years old.\n", age)

	// Example 3: Reading a float input
	fmt.Print("Enter your height in meters: ")
	height, err := readFloat(reader)
	if err != nil {
		fmt.Println("Invalid input for height:", err)
		return
	}
	fmt.Printf("Your height is %.2f meters.\n", height)

	// Example 4: Reading multiple inputs in a single line
	fmt.Print("Enter your city and country (separated by a comma): ")
	city, country, err := readLocation(reader)
	if err != nil {
		fmt.Println("Invalid input for location:", err)
		return
	}
	fmt.Printf("You live in %s, %s.\n", city, country)
}

func readString(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readInt(reader *bufio.Reader) (int, error) {
	input, err := readString(reader)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(input)
}

func readFloat(reader *bufio.Reader) (float64, error) {
	input, err := readString(reader)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(input, 64)
}

func readLocation(reader *bufio.Reader) (string, string, error) {
	input, err := readString(reader)
	if err != nil {
		return "", "", err
	}
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("expected two parts separated by a comma")
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), nil
}
