package main

import "fmt"

func print_list(m map[string]int) {
	for i, j := range m {
		fmt.Printf("%s -> %d\n", i, j)
	}
}
func main() {
	shopping_list := make(map[string]int)
	for {
		var item string
		var input string
		fmt.Print("Add item or write done or -item: ")
		n, err := fmt.Scan(&item)
		if item == "done" {
			break
		}

		if item == "-item" {
			fmt.Println("What item you want to delete? ")
			print_list(shopping_list)
			fmt.Scan(&input)
			delete(shopping_list, input)

			continue
		}
		if err != nil || n == 0 || item == "" {
			fmt.Println("Invalid or empty input, please try again.")
			continue
		}
		shopping_list[item]++
	}
	fmt.Println("Shopping List: ")
	print_list(shopping_list)
}
