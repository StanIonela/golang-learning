package main

import "fmt"

func main() {
	var items []string
	var input string

	for {
		fmt.Println("Enter item or done: ")
		fmt.Scan(&input)

		if input == "done" {
			break
		}
		items = append(items, input)
	}
	fmt.Println("Final shopping list: ")
	for item, list := range items {
		fmt.Printf("%d. %s\n", item+1, list)
	}
}
