package main

import "fmt"

const ArraySize = 5

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

func Init() HashTable {
	hashTable := HashTable{}

	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}

	return hashTable
}

func (h *HashTable) Insert(k string, v int) {
	index := hash(k)
	h.array[index].insert(k, v)
}

func (h *HashTable) Delete(k string) {
	index := hash(k)
	h.array[index].delete(k)
}

func (h *HashTable) Search(k string) bool {
	index := hash(k)
	return h.array[index].search(k)
}

func (b *bucket) insert(k string, v int) {
	if !b.search(k) {
		newNode := &bucketNode{key: k, value: v}
		newNode.next = b.head

		b.head = newNode
	} else {
		err := fmt.Errorf("Node %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == k {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

func (b *bucket) delete(k string) {
	previousNode := b.head

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}

		previousNode = previousNode.next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {
	hashTable := Init()
	hashTable.Insert("Ana", 24)
	hashTable.Insert("Eric", 13)
	hashTable.Insert("Rebecca", 67)
	hashTable.Insert("Marie", 53)
	hashTable.Insert("D", 33)
}
