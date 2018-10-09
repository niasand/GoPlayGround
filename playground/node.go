package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func main() {
	tree := &Node{6, nil, nil}
	fmt.Println(tree.Key, tree.Right, tree.Left)
	fmt.Println(tree.Search(6))
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key {
		return n.Left.Search(key)
	}

	return true
}
