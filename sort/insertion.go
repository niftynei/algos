package sort

import (
	"time"

	"github.com/niftynei/algos/timing"
)

func Insertion(array []int) (result []int) {
	defer timing.Timer(time.Now(), "Insertion")
	if len(array) == 0 || len(array) == 1 {
		return array
	}

	result = make([]int, 1, len(array))
	result[0] = array[0]
	array = array[1:]

	for _, elem := range array {
		insertAt := 0
		for resultIndex, resultElem := range result {
			if elem < resultElem {
				insertAt = resultIndex
				break
			}
		}
		result = insert(result, insertAt, elem)
	}

	return result
}

func insert(slice []int, index, value int) []int {
	slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}
