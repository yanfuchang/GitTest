package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 2, 3, 3, 0}, []int{0, 1, 2, 3, 3, 3}},
		{"single element", []int{42}, []int{42}},
		{"empty", []int{}, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int{}, tc.input...)
			QuickSort(nums)
			if !reflect.DeepEqual(nums, tc.want) {
				t.Fatalf("QuickSort(%v) = %v, want %v", tc.input, nums, tc.want)
			}
		})
	}
}
