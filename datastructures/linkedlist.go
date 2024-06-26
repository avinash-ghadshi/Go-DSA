package datastructures

import "fmt"

type Node struct {
	Next *Node
	Data interface{}
}

func NewNode(data interface{}) *Node {
	return &Node{
		Next: nil,
		Data: data,
	}
}

func (n *Node) Add(tail *Node) *Node {
	if tail != nil {
		tail.Next = n
	}
	tail = n
	return tail
}

func Traverse(head *Node) {
	for head.Next != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
	fmt.Println(head.Data)
}
