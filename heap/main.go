package main

import (
	"errors"
	"fmt"
)

type Heap struct {
	size  int
	items []interface{}
}

func Init(size int) Heap {
	return Heap{
		size:  size,
		items: make([]interface{}, size),
	}
}

func (h *Heap) Push(element interface{}) error {
	if !isHeapFull(h.items) {
		for index := range h.items {
			if h.items[index] == nil {
				h.items[index] = element
				return nil
			}
		}
	}

	return errors.New("Heap is full")
}

func (h *Heap) Pop() (interface{}, error) {
	if !isHeapEmpty(h.items) {
		for i := len(h.items) - 1; i >= 0; i-- {
			if h.items[i] != nil {
				poppedValue := h.items[i]
				h.items[i] = nil

				return poppedValue, nil
			}
		}
	}

	return nil, errors.New("Heap is empty")
}

func isHeapFull(items []interface{}) bool {
	return items[len(items)-1] != nil
}

func isHeapEmpty(items []interface{}) bool {
	return items[0] == nil
}

func main() {
	heap := Init(10)
	printPushHeap(heap, 10)
	printPushHeap(heap, 20)
	printPushHeap(heap, 30)

	printPopHeap(heap)
	printPopHeap(heap)

	printPushHeap(heap, "Hello")
	printPushHeap(heap, 43)

	printPopHeap(heap)
	printPopHeap(heap)

	printPushHeap(heap, "World")
	printPushHeap(heap, "!")
	printPushHeap(heap, 86)
	printPushHeap(heap, 45)
	printPushHeap(heap, 15)
	printPushHeap(heap, 75)

	printPopHeap(heap)
	printPopHeap(heap)
	printPopHeap(heap)

}

func printPushHeap(heap Heap, value interface{}) {
	err := heap.Push(value)

	if err != nil {
		fmt.Printf("Cannot push value: %v to heap. %v\n", value, err)
	} else {
		fmt.Printf("Value pushed to heap: %v\n", value)
	}
}

func printPopHeap(heap Heap) {
	popValue, err := heap.Pop()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Removed value from heap: %v\n", popValue)
	}
}
