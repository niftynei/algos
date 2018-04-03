package sort

import (
	"time"

	"github.com/niftynei/algos/timing"
)

func Heap(array []int) (result []int) {
	defer timing.Timer(time.Now(), "Heap")
	result = make([]int, len(array))
	heap := makeHeep(array)
	for i := 0; i < len(array); i++ {
		result[i] = heap.extractMin()
	}
	return result
}

type Heep struct {
	heap  []int
	count int
}

func makeHeep(arr []int) *Heep {
	heep := &Heep{}
	heep.make(arr)
	return heep
}

func (h *Heep) make(arr []int) {
	h.init(len(arr))
	for _, elem := range arr {
		h.insert(elem)
	}
}

func (h *Heep) init(size int) {
	h.heap = make([]int, size+1)
	h.count = 0
}

func (h *Heep) extractMin() int {
	if h.count == 0 {
		return 0
	}

	top := h.heap[1]
	h.heap[1] = h.heap[h.count]
	h.count += -1
	h.bubbleDown(1)
	return top
}

func (h *Heep) bubbleDown(bubbleAt int) {
	childIndex := h.leftChildIndex(bubbleAt)
	minIndex := bubbleAt
	for i := 0; i <= 1; i++ {
		if childIndex+1 <= h.count {
			if h.heap[minIndex] > h.heap[childIndex+i] {
				minIndex = childIndex + i
			}
		}
	}
	if minIndex != bubbleAt {
		h.swap(bubbleAt, minIndex)
		h.bubbleDown(minIndex)
	}
}

func (h *Heep) bubbleUp(bubbleAt int) {
	parentIndex := h.parentIndex(bubbleAt)
	if parentIndex == -1 {
		return // heap root, no parent :'(
	}
	if h.heap[parentIndex] > h.heap[bubbleAt] {
		h.swap(bubbleAt, parentIndex)
		h.bubbleUp(parentIndex)
	}
}

func (h *Heep) swap(swapA, swapB int) {
	tmp := h.heap[swapA]
	h.heap[swapA] = h.heap[swapB]
	h.heap[swapB] = tmp
}

func (h *Heep) insert(i int) {
	// todo: check for overflow
	// note: we intentionally don't populate index-0
	h.count += 1
	h.heap[h.count] = i
	h.bubbleUp(h.count)
}

func (h *Heep) parentIndex(n int) int {
	if n == 1 {
		return -1
	}
	return n / 2
}

func (h *Heep) leftChildIndex(n int) int {
	return 2 * n
}

func (h *Heep) len() int {
	return h.count
}

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

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
