package algorithms_test

import (
	"github.com/dimitrijjedich/go-algorithms/algorithms"
	"reflect"
	"testing"
)

func TestRedBlackTreeSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"unsorted array", []int{5, 3, 8, 6, 2}, []int{2, 3, 5, 6, 8}},
		{"non consecutive array", []int{9, 3, 5, 1, 7}, []int{1, 3, 5, 7, 9}},
		{"empty array", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"duplicates", []int{3, 1, 2, 3, 1}, []int{1, 1, 2, 3, 3}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := algorithms.RedBlackTreeSort(testCase.input)

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Failed %s: expected %v, got %v", testCase.name, testCase.expected, result)
			}
		})
	}
}
