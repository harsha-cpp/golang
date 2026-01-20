package main

import (
	"fmt"
	"math"
)

func main() {
	//variable declaration

	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subraction: ", result)

	result = a * b
	fmt.Println("Multiplied ", result)

	result = a / b
	fmt.Println("Division ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const p float64 = 22 / 7
	fmt.Println(p)

	//overflow with signed integer
	var maxInt int64 = 9223372036854775807 // max value of 64 bit int
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	//overflow with unsigned integer
	var uMaxInt uint64 = 18446744073709551615 // max value of 64 bit uint
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	//underflow happens when smallest float value is divided by largest float number

	var smallFloat float64 = 1.0e-323
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
