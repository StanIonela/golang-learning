package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	test := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty slice", []int{}, []int{}},
		{"one element", []int{5}, []int{5}},
		{"already sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"reverse order", []int{3, 2, 1}, []int{1, 2, 3}},
		{"duplicates", []int{5, 1, 5, 3}, []int{1, 3, 5, 5}},
		{"all same", []int{1, 1, 1}, []int{1, 1, 1}},
		{"negatives", []int{-1, -3, 0}, []int{-3, -1, 0}},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.in, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	test := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{5}, []int{5}},
		{"already sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"reverse order", []int{3, 2, 1}, []int{1, 2, 3}},
		{"with duplicates", []int{5, 1, 5, 3}, []int{1, 3, 5, 5}},
		{"negatives", []int{-1, -3, 0}, []int{-3, -1, 0}},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateRandomSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}
	return s
}

func BenchmarkBubbleSort(b *testing.B) {
	sizes := []int{100, 1000, 5000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("BubbleSort_%d", size), func(b *testing.B) {
			data := generateRandomSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				BubbleSort(append([]int{}, data...))
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	sizes := []int{100, 1000, 5000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("MergeSort_%d", size), func(b *testing.B) {
			data := generateRandomSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				MergeSort(append([]int{}, data...))
			}

		})
	}
}
