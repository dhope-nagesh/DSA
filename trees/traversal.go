package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.left)
	fmt.Printf("%d ", root.data)
	InOrderTraversal(root.right)
}

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	fmt.Printf("%d ", root.data)
}

func NewNode(data int) *Node {
	node := &Node{}
	node.data = data
	return node
}

func main() {
	// fmt.Println("Tree traversal")
	root := NewNode(1)
	root.left = NewNode(2)
	root.right = NewNode(3)
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)
	root.right.left = NewNode(6)
	root.right.right = NewNode(7)
	fmt.Println("Tree In Order traversal")
	InOrderTraversal(root)
	fmt.Println("")
	fmt.Println("Tree pre Order traversal")
	PreOrderTraversal(root)
	fmt.Println("")
	fmt.Println("Tree post Order traversal")
	PostOrderTraversal(root)
	fmt.Println("")
}
