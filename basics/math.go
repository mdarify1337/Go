package main

import (
	"fmt"
)

func calculater(a int, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		}
		fmt.Println("Error: Division by zero")
		return 0
	default:
		fmt.Println("Error: Unknown operation")
		return 0
	}
}
