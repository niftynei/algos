package sort

import (
	"time"

	"github.com/niftynei/algos/timing"
)

func quicksort(array []int, firstElem int, lastElem int) {
	if (lastElem - firstElem) > 0 {
		pivot := partition(array, firstElem, lastElem)
		quicksort(array, firstElem, pivot-1)
		quicksort(array, pivot+1, lastElem)
	}
}

func partition(array []int, firstElem int, lastElem int) (firstHigh int) {
	pivot := lastElem

	firstHigh = firstElem
	for i := firstElem; i < lastElem; i++ {
		if array[i] < array[pivot] {
			swap(array, i, firstHigh)
			firstHigh++
		}
	}
	swap(array, pivot, firstHigh)

	return firstHigh
}

func swap(array []int, swapA, swapB int) {
	tmp := array[swapA]
	array[swapA] = array[swapB]
	array[swapB] = tmp
}

func Quick(array []int) []int {
	defer timing.Timer(time.Now(), "Quick")

	quicksort(array, 0, len(array)-1)

	return array
}
