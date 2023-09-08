package node

import (
	"fmt"
)

type Node[T any] struct {
	ID       int
	Data     T
	previous *Node[T]
	nexts    []*Node[T]
}

func New[T any](ID int, data T) *Node[T] {
	return &Node[T]{
		ID:   ID,
		Data: data,
	}
}

func (n *Node[T]) GetData() T {
	return n.Data
}

func (n *Node[T]) GetNexts() []*Node[T] {
	return n.nexts
}

func (n *Node[T]) IsRoot() bool {
	return n.previous == nil
}

func (n *Node[T]) IsLeaf() bool {
	return len(n.nexts) == 0
}

func (n *Node[T]) Backtrack() []*Node[T] {
	var nodes []*Node[T]

	current := n
	for current != nil {
		nodes = append(nodes, current)
		current = current.previous
	}

	return nodes
}

func (n *Node[T]) GetStructure() []string {
	var structure []string
	if n.previous == nil {
		structure = append(structure, fmt.Sprintf("(NULL) -> (%d)", n.ID))
	} else {
		structure = append(structure, fmt.Sprintf("(%d) -> (%d)", n.previous.ID, n.ID))
	}

	for _, next := range n.nexts {
		innerStructure := next.GetStructure()
		structure = append(structure, innerStructure...)
	}

	return structure
}

func (n *Node[T]) AddNext(node *Node[T]) {
	node.previous = n
	n.nexts = append(n.nexts, node)
}
