package main

import (
	"fmt"
)

func BubbleSort(values []int) []int {
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	return values
}

func MergeSort(values []int) []int {
	if len(values) < 2 {
		return values
	}

	first := MergeSort(values[:len(values)/2])
	second := MergeSort(values[len(values)/2:])
	return Merge(first, second)
}

func Merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

func main() {
	array := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(BubbleSort(array))

	array2 := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	fmt.Println(MergeSort(array2))
}
