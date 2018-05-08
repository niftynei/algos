package sort

import (
	"github.com/niftynei/algos/timing"
	"time"
)

func mergesortWithConcurrency(array []int, c chan []int) {
	if len(array) == 1 {
		c <- array
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)
	middle := len(array) / 2
	go mergesortWithConcurrency(array[:middle], leftChan)
	go mergesortWithConcurrency(array[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	c <- mergeWithConcurrency(ldata, rdata)
}

func mergeWithConcurrency(left []int, right []int) []int {
	array := make([]int, len(left)+len(right))
	buffer1 := &queue{}
	buffer2 := &queue{}

	for _, item := range left {
		buffer1.enqueue(item)
	}
	for _, item := range right {
		buffer2.enqueue(item)
	}

	counter := 0
	for !(buffer1.isEmpty() || buffer2.isEmpty()) {
		if buffer1.head.value <= buffer2.head.value {
			array[counter] = buffer1.dequeue()
		} else {
			array[counter] = buffer2.dequeue()
		}
		counter += 1
	}

	for !buffer1.isEmpty() {
		array[counter] = buffer1.dequeue()
		counter += 1
	}
	for !buffer2.isEmpty() {
		array[counter] = buffer2.dequeue()
		counter += 1
	}
	return array
}

func MergeWithConcurrency(array []int) []int {
	defer timing.Timer(time.Now(), "Concurrent Merge")

	result := make(chan []int)
	go mergesortWithConcurrency(array, result)

	r := <-result
	response := make([]int, len(array))
	for _, v := range r {
		response = append(response, v)
	}
	close(result)
	return response
}
