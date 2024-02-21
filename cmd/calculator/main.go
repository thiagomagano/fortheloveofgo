package main

import (
	"calculator"
	"fmt"
)

func main() {
	addResult := calculator.Add(2, 2)
	subResult := calculator.Subtract(2, 2)
	multiResult := calculator.Multiply(2, 2)
	DivideResult, _ := calculator.Divide(2, 2)
	sqrtResult, _ := calculator.Sqrt(2)

	fmt.Printf("2 + 2 = %v \n", addResult)
	fmt.Printf("2 - 2 = %v \n", subResult)
	fmt.Printf("2 * 2 = %v \n", multiResult)
	fmt.Printf("2 / 2 = %v \n", DivideResult)
	fmt.Printf("sqrt(2) = %0.5v \n", sqrtResult)
}
