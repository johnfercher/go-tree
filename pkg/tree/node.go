package tree

import (
	"fmt"
)

// nolint:structcheck,gocritic
type Node[T any] struct {
	id       int
	data     T
	previous *Node[T]
	nexts    []*Node[T]
}

func NewNode[T any](id int, data T) *Node[T] {
	return &Node[T]{
		id:   id,
		data: data,
	}
}

func (n *Node[T]) Get() (int, T) {
	return n.id, n.data
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
		structure = append(structure, fmt.Sprintf("(NULL) -> (%d)", n.id))
	} else {
		structure = append(structure, fmt.Sprintf("(%d) -> (%d)", n.previous.id, n.id))
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
