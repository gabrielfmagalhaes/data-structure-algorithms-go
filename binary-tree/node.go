package main

import "fmt"

type Node struct {
	Key       int
	LeftNode  *Node
	RightNode *Node
}

func (n *Node) Insert(key int) {
	if key < n.Key {
		if n.LeftNode == nil {
			n.LeftNode = &Node{Key: key}
		} else {
			n.LeftNode.Insert(key)
		}
	} else if key > n.Key {
		if n.RightNode == nil {
			n.RightNode = &Node{Key: key}
		} else {
			n.RightNode.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if key < n.Key {
		return n.LeftNode.Search(key)
	} else if key > n.Key {
		return n.RightNode.Search(key)
	}

	return true
}

func main() {
	tree := &Node{Key: 50}

	tree.Insert(47)
	tree.Insert(20)
	tree.Insert(13)
	tree.Insert(23)
	tree.Insert(49)

	tree.Insert(65)
	tree.Insert(78)
	tree.Insert(70)
	tree.Insert(54)
	tree.Insert(89)

	fmt.Println(tree)
	fmt.Printf("\nValue 65 exists in binary tree? %v", tree.Search(65))
	fmt.Printf("\nValue 33 exists in binary tree? %v", tree.Search(33))
}
