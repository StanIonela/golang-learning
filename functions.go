package main

import "fmt"

func add() int {
	var a, b int
	fmt.Println("Enter two numbers")
	fmt.Scan(&a, &b)
	return a + b
}

func subtract() int {
	var a, b int
	fmt.Println("Enter two numbers")
	fmt.Scan(&a, &b)
	return a - b
}

func multiply() int {
	var a, b int
	fmt.Println("Enter two numbers")
	fmt.Scan(&a, &b)
	return a * b
}

func divide() int {
	var a, b int
	fmt.Println("Enter two numbers")
	fmt.Scan(&a, &b)
	return a / b
}

func main() {
	res := add()
	fmt.Printf("Sum = %d\n", res)

	res1 := multiply()
	fmt.Printf("Multiply = %d\n", res1)

	res2 := subtract()
	fmt.Printf("Subtract = %d\n", res2)

	res3 := divide()
	fmt.Printf("Divide = %d\n", res3)
}
