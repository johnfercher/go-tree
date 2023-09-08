package tree

import (
	"fmt"
)

type Node[T any] struct {
	ID       int
	Data     T
	Previous *Node[T]
	Nexts    []*Node[T]
}

func NewNode[T any](ID int, data T) *Node[T] {
	return &Node[T]{
		ID:   ID,
		Data: data,
	}
}

func (n *Node[T]) GetContent() T {
	return n.Data
}

func (n *Node[T]) IsLeaf() bool {
	if n == nil {
		return false
	}

	return len(n.Nexts) == 0
}

func (n *Node[T]) Backtrack() []*Node[T] {
	var nodes []*Node[T]

	current := n
	for current != nil {
		nodes = append(nodes, current)
		current = current.Previous
	}

	return nodes
}

func (n *Node[T]) Print(label string) {
	if n.Previous == nil {
		fmt.Printf("%s - ID(%d), NextSize(%d), HasPrevious(false)\n", label, n.ID, len(n.Nexts))
	} else {
		fmt.Printf("%s - ID(%d), NextSize(%d), Previous(%d)\n", label, n.ID, len(n.Nexts), n.Previous.ID)
	}
}
