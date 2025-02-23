// main.go
package main

import (
	"fmt"
)

func main() {
	// Creating a slice using a slice literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial slice:", numbers)

	// Appending elements to a slice
	numbers = append(numbers, 6, 7)
	fmt.Println("After appending elements:", numbers)

	// Slicing a slice
	subSlice := numbers[1:4]
	fmt.Println("Sub-slice:", subSlice)

	// Creating a slice with make
	madeSlice := make([]int, 5, 10)
	fmt.Println("Slice created with make:", madeSlice)
	fmt.Println("Length:", len(madeSlice), "Capacity:", cap(madeSlice))

	// Copying a slice
	copySlice := make([]int, len(numbers))
	copy(copySlice, numbers)
	fmt.Println("Copied slice:", copySlice)

	// Edge case: Slicing beyond capacity
	// Uncommenting the following line will cause a runtime panic
	// fmt.Println(numbers[:10])

	// Edge case: Appending to a nil slice
	var nilSlice []int
	nilSlice = append(nilSlice, 1, 2, 3)
	fmt.Println("Appending to a nil slice:", nilSlice)

	// Edge case: Slicing an empty slice
	emptySlice := []int{}
	emptySubSlice := emptySlice[:]
	fmt.Println("Slicing an empty slice:", emptySubSlice)

	// Use case: Removing an element from a slice
	indexToRemove := 2
	numbers = append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)
	fmt.Println("After removing element at index 2:", numbers)

	// Use case: Inserting an element into a slice
	indexToInsert := 2
	elementToInsert := 10
	numbers = append(numbers[:indexToInsert], append([]int{elementToInsert}, numbers[indexToInsert:]...)...)
	fmt.Println("After inserting element at index 2:", numbers)

	// Use case: Iterating over a slice
	fmt.Println("Iterating over the slice:")
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Edge case: Slicing with negative indices
	// Uncommenting the following line will cause a compile-time error
	// fmt.Println(numbers[-1:])

	// Use case: Clearing a slice
	numbers = numbers[:0]
	fmt.Println("After clearing the slice:", numbers)

	// Use case: Extending a slice
	numbers = append(numbers, 1, 2, 3, 4, 5)
	extendedSlice := numbers[:4]
	fmt.Println("Extended slice:", extendedSlice)

	// Use case: Checking if a slice is nil
	var anotherNilSlice []int
	if len(anotherNilSlice) == 0 {
		fmt.Println("The slice is nil")
	}

	// Use case: Creating a multi-dimensional slice
	multiDimSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Multi-dimensional slice:", multiDimSlice)
}