package main

import (
	"fmt"

	"github.com/niftynei/algos/sort"
)

func main() {
	fmt.Println("hello")

	unsorted := []int{9, 5, 2, 4, 7, 8, 1, 6, -2}
	fmt.Println(sort.Selection(unsorted))
	fmt.Println(sort.Insertion(unsorted))
	fmt.Println(sort.Heap(unsorted))
}
