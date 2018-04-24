package main

import (
	"fmt"
	"math/rand"

	"github.com/niftynei/algos/sort"
)

type sortFunc func([]int) []int

func generateRandom(length int) []int {
	result := make([]int, length)

	for index, _ := range result {
		result[index] = rand.Intn(length)
	}

	return result
}

func main() {
	/* ~/go/src/github.com/niftynei/algos $ go run proj.go
	   Sorting 500000 numbers

	   Heap took 156.757316ms
	   Insertion took 1m44.827435316s
	   Merge took 447.51028ms
	   Quick took 1m37.8403218s
	   Selection took 5m27.989212623s
	*/

	numToSort := 500000
	unsorted := generateRandom(numToSort)
	fmt.Printf("Sorting %d numbers\n\n", numToSort)

	sortTypes := []sortFunc{
		sort.Heap,
		sort.Insertion,
		sort.Merge,
		sort.Quick,
		sort.Selection,
	}

	for _, sortType := range sortTypes {
		sortType(unsorted)
	}
}
