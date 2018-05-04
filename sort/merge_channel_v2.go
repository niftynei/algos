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

func mergesort(array []int, c chan []int) {
	if len(array) == 1 {
		c <- array
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)
	middle := len(array) / 2
	go mergesort(array[:middle], leftChan)
	go mergesort(array[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	c <- merge(ldata, rdata)
}

func merge(left []int, right []int) ([]int) {
	array := make([]int, len(left) + len(right))
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

func Merge(array []int) []int {
	defer timing.Timer(time.Now(), "Merge")

	result := make(chan []int)
	go mergesort(array, result)

	r := <-result
	response := make([]int, len(array))
	for _, v := range r {
		response = append(response, v)
	}
	close(result)
	return response
}
