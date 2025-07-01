package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	test := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{
			name:   "normal case",
			input:  []int{4, 8, 15, 18},
			target: 12,
			want:   []int{0, 1},
		},
		{
			name:   "empty input",
			input:  []int{},
			target: 0,
			want:   nil,
		},
		{
			name:   "single element",
			input:  []int{1},
			target: 2,
			want:   nil,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.input, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Two Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuplicateFinder(t *testing.T) {
	test := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "with duplicate",
			input: []int{1, 1, 2, 3, 3, 3, 4, 5, 5, 5, 5},
			want:  []int{1, 3, 3, 5, 5, 5},
		},
		{
			name:  "no duplicates",
			input: []int{1, 2, 3, 4, 5, 6},
			want:  []int{},
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []int{},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := DuplicateFinder(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DuplicateFinder() = %v, want = %v", got, tt.want)
			}
		})
	}
}
