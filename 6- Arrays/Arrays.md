# Arrays in Go

Arrays are a fundamental data structure in Go that allow you to store a fixed-size sequence of elements of the same type. This guide will help you understand how to work with arrays in Go.

## Table of Contents
- [Introduction](#introduction)
- [Declaring Arrays](#declaring-arrays)
- [Initializing Arrays](#initializing-arrays)
- [Accessing Array Elements](#accessing-array-elements)
- [Iterating Over Arrays](#iterating-over-arrays)
- [Array Length](#array-length)
- [Multidimensional Arrays](#multidimensional-arrays)
- [Conclusion](#conclusion)

## Introduction
An array is a collection of elements that are identified by an index or key. In Go, arrays have a fixed length, which means that once an array is declared, its size cannot be changed.

## Declaring Arrays
To declare an array in Go, you specify the type of its elements and the number of elements it can hold.

```go
var arrayName [size]elementType
```

Example:
```go
var numbers [5]int
```

## Initializing Arrays
You can initialize an array at the time of declaration using an array literal.

Example:
```go
var numbers = [5]int{1, 2, 3, 4, 5}
```

You can also let Go infer the length of the array:
```go
var numbers = [...]int{1, 2, 3, 4, 5}
```

## Accessing Array Elements
Array elements are accessed using the index, which starts from 0.

Example:
```go
numbers[0] = 10
fmt.Println(numbers[0]) // Output: 10
```

## Iterating Over Arrays
You can iterate over an array using a `for` loop or a `for range` loop.

Example using `for` loop:
```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

Example using `for range` loop:
```go
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

## Array Length
The length of an array can be determined using the `len` function.

Example:
```go
fmt.Println(len(numbers)) // Output: 5
```

## Multidimensional Arrays
Go supports multidimensional arrays, which are arrays of arrays.

Example:
```go
var matrix [3][3]int
matrix[0][0] = 1
matrix[1][1] = 2
matrix[2][2] = 3
```

You can also initialize a multidimensional array at the time of declaration:
```go
var matrix = [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

## Conclusion
Arrays in Go are a powerful tool for storing and managing collections of data. Understanding how to declare, initialize, and manipulate arrays is essential for effective Go programming.