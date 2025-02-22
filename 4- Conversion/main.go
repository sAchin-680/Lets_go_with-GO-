// main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String to Integer
	strToInt("123")

	// Integer to String
	intToStr(456)

	// String to Float
	strToFloat("123.45")

	// Float to String
	floatToStr(678.90)

	// String to Boolean
	strToBool("true")

	// Boolean to String
	boolToStr(false)

	// Integer to Float
	intToFloat(789)

	// Float to Integer
	floatToInt(123.99)
}

func strToInt(str string) {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("String to Integer:", intValue)
	}
}

func intToStr(intValue int) {
	str := strconv.Itoa(intValue)
	fmt.Println("Integer to String:", str)
}

func strToFloat(str string) {
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println("String to Float:", floatValue)
	}
}

func floatToStr(floatValue float64) {
	str := strconv.FormatFloat(floatValue, 'f', 2, 64)
	fmt.Println("Float to String:", str)
}

func strToBool(str string) {
	boolValue, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println("Error converting string to bool:", err)
	} else {
		fmt.Println("String to Boolean:", boolValue)
	}
}

func boolToStr(boolValue bool) {
	str := strconv.FormatBool(boolValue)
	fmt.Println("Boolean to String:", str)
}

func intToFloat(intValue int) {
	floatValue := float64(intValue)
	fmt.Println("Integer to Float:", floatValue)
}

func floatToInt(floatValue float64) {
	intValue := int(floatValue)
	fmt.Println("Float to Integer:", intValue)
}