package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		fmt.Println("Cannot divide by zero")
		return 0
	}

	return a / b
}

func main() {
	var a, b, i int

	fmt.Println("Enter two numbers")
	fmt.Scan(&a, &b)

	fmt.Println("Choose the operation:\n 1. add\n 2. subtract\n 3. multiply\n 4. divide")
	fmt.Scan(&i)

	switch i {
	case 1:
		fmt.Printf("Sum = %d", add(a, b))
	case 2:
		fmt.Printf("Subtract = %d", subtract(a, b))
	case 3:
		fmt.Printf("Multiply = %d", multiply(a, b))
	case 4:
		fmt.Printf("Divide = %d", divide(a, b))
	default:
		fmt.Println("Invalid operation")
	}
}
