package main

import "fmt"

// Node in a binary Tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewNode creates an initial BinaryTree
func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

// Insert new Node
func (n *Node) Insert(value int) *Node {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}

	return n
}

// PreOrder print
func (n Node) PreOrder() {
	fmt.Printf("%d ", n.Value)
	if n.Left != nil {
		n.Left.PreOrder()
	}

	if n.Right != nil {
		n.Right.PreOrder()
	}
}

// InOrder print
func (n Node) InOrder() {
	if n.Left != nil {
		n.Left.InOrder()
	}
	fmt.Printf("%d ", n.Value)

	if n.Right != nil {
		n.Right.InOrder()
	}
}

// PostOrder Print
func (n Node) PostOrder() {
	if n.Left != nil {
		n.Left.PostOrder()
	}

	if n.Right != nil {
		n.Right.PostOrder()
	}

	fmt.Printf("%d ", n.Value)
}

// Search if there is a node with the value of v
func (n *Node) Search(v int) bool {
	if n == nil {
		return false
	}

	if v < n.Value {
		return n.Left.Search(v)
	} else if v > n.Value {
		return n.Right.Search(v)
	}

	return true
}

func main() {
	node := NewNode(80)

	node.Insert(40).
		Insert(50).
		Insert(12).
		Insert(100).
		Insert(200).
		Insert(130).
		Insert(5).
		Insert(10).
		Insert(250)

	fmt.Println("PreOrder Print")
	node.PreOrder()
	fmt.Println()

	fmt.Println("InOrder Print")
	node.InOrder()
	fmt.Println()

	fmt.Println("PostOrder Print")
	node.PostOrder()

	fmt.Printf("\n Is there a Node with value 5? %t \n", node.Search(5))
}
