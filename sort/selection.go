package sort

import (
	"time"

	"github.com/niftynei/algos/timing"
)

func Selection(array []int) (result []int) {
	defer timing.Timer(time.Now(), "Selection")
	result = make([]int, len(array))

	next := 0
	for len(array) > 0 {
		min := array[0]
		at := 0
		for index, elem := range array {
			if elem < min {
				min = elem
				at = index
			}
		}
		result[next] = min
		array = remove(array, at)
		next += 1
	}

	return result
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
