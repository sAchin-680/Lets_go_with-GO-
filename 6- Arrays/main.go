package main

import "fmt"

func main() {
	// Integer array
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Integer array:", numbers)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])
	numbers[2] = 35
	fmt.Println("Modified third element:", numbers[2])

	fmt.Println("Iterating over the integer array:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element at index %d: %d\n", i, numbers[i])
	}

	fmt.Println("Iterating over the integer array using range:")
	for index, value := range numbers {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}

	// Multi-dimensional array
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multi-dimensional array:", matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, matrix[i][j])
		}
	}

	// String array
	var fruits [3]string = [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("String array:", fruits)
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Fruit at index %d: %s\n", i, fruits[i])
	}

	// Boolean array
	var flags [4]bool = [4]bool{true, false, true, false}
	fmt.Println("Boolean array:", flags)
	for i := 0; i < len(flags); i++ {
		fmt.Printf("Flag at index %d: %t\n", i, flags[i])
	}

	// Float array
	var decimals [3]float64 = [3]float64{1.1, 2.2, 3.3}
	fmt.Println("Float array:", decimals)
	for i := 0; i < len(decimals); i++ {
		fmt.Printf("Decimal at index %d: %.1f\n", i, decimals[i])
	}
}
