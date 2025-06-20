package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	list := []int{4, 5}
	list = append(list, 6)
	mapRange := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println("arr", arr)
	fmt.Println("list", list)

	for i, item := range list {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	for m, a := range mapRange {
		fmt.Printf("%s -> %s\n", m, a)
	}

	shopping_list := make(map[string]int)
	shopping_list["flour"]++
	shopping_list["banana"] += 2
	fmt.Println("shopping_list = ", shopping_list)

	delete(shopping_list, "banana")

	fmt.Println("shopping_list = ", shopping_list)
	for shop, list := range shopping_list {
		fmt.Printf("%s -> %s\n", shop, list)
	}
}
