package main

import (
	"fmt"
	"github.com/dimitrijjedich/go-algorithms/algorithms"
)

func main() {
	arr := []int{5, 3, 8, 6, 2}
	fmt.Println("Original:", arr)
	fmt.Println("BubbleSort:", algorithms.BubbleSort(arr))
}
