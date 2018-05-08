package sort

import (
	"time"

	"github.com/niftynei/algos/timing"
)

type queue struct {
	head *queueNode
	tail *queueNode
}

type queueNode struct {
	value int
	child *queueNode
}

func (Q *queue) dequeue() int {
	result := Q.head.value

	if Q.head == Q.tail {
		Q.head, Q.tail = nil, nil
		return result
	}
	Q.head = Q.head.child

	return result
}

func (Q *queue) enqueue(value int) {
	newQueueNode := &queueNode{value: value}

	if Q.head == nil {
		Q.head = newQueueNode
	}

	if Q.tail != nil {
		Q.tail.child = newQueueNode
	}
	Q.tail = newQueueNode
}

func (Q *queue) isEmpty() bool {
	return Q.head == nil
}

func mergesort(array []int, low, high int) {
	if low < high {
		middle := (low + high) / 2
		mergesort(array, low, middle)
		mergesort(array, middle+1, high)
		merge(array, low, middle, high)
	}
}

func merge(array []int, low, middle, high int) {
	buffer1 := &queue{}
	buffer2 := &queue{}

	for i := low; i <= middle; i++ {
		buffer1.enqueue(array[i])
	}

	for i := middle + 1; i <= high; i++ {
		buffer2.enqueue(array[i])
	}

	counter := low
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
}

func Merge(array []int) []int {
	defer timing.Timer(time.Now(), "Merge")

	mergesort(array, 0, len(array)-1)
	return array
}
