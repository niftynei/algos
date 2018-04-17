package main

import (
	"fmt"

	"github.com/niftynei/algos/sort"
)

func main() {
	unsorted := []int{9, 5, 2, 4, 7, 8, 1, 6, -2}

	fmt.Printf("Unsorted array: %v\n\n", unsorted)

	fmt.Println(sort.Selection(unsorted))
	fmt.Println(sort.Insertion(unsorted))
	fmt.Println(sort.Heap(unsorted))
	fmt.Println(sort.Quick(unsorted))
}
