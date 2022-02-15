package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	size  int
	items []interface{}
}

func Init(size int) Queue {
	return Queue{
		size:  size,
		items: make([]interface{}, size),
	}
}

func (q *Queue) Enqueue(element interface{}) error {
	if !isQueueFull(q.items) {
		for index := range q.items {
			if q.items[index] == nil {
				q.items[index] = element
				return nil
			}
		}
	}

	return errors.New("Queue is full")
}

func (h *Queue) Dequeue() (interface{}, error) {
	result := h.items[0]

	if !isQueueEmpty(h.items) {
		for i := range h.items {
			if i != len(h.items)-1 {
				h.items[i] = h.items[i+1]
			}
		}

		return result, nil
	}

	return nil, errors.New("Queue is empty")
}

func isQueueFull(items []interface{}) bool {
	return items[len(items)-1] != nil
}

func isQueueEmpty(items []interface{}) bool {
	return items[0] == nil
}

func main() {
	queue := Init(10)
	printEnqueueQueue(queue, 10)
	printEnqueueQueue(queue, 20)
	printEnqueueQueue(queue, 30)

	printDequeueQueue(queue)
	printDequeueQueue(queue)

	printEnqueueQueue(queue, "Hello")
	printEnqueueQueue(queue, 43)

	printDequeueQueue(queue)
	printDequeueQueue(queue)

	printEnqueueQueue(queue, "World")
	printEnqueueQueue(queue, "!")
	printEnqueueQueue(queue, 86)
	printEnqueueQueue(queue, 45)
	printEnqueueQueue(queue, 15)
	printEnqueueQueue(queue, 75)

	printDequeueQueue(queue)
	printDequeueQueue(queue)
	printDequeueQueue(queue)

}

func printEnqueueQueue(queue Queue, value interface{}) {
	err := queue.Enqueue(value)

	if err != nil {
		fmt.Printf("Cannot push value: %v to queue. %v\n", value, err)
	} else {
		fmt.Printf("Value pushed to queue: %v\n", value)
	}
}

func printDequeueQueue(queue Queue) {
	popValue, err := queue.Dequeue()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Removed value from queue: %v\n", popValue)
	}
}
