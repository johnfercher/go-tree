package tree

import (
	"fmt"
)

type Tree[T any] struct {
	root      *Node[T]
	lastAdded *Node[T]
}

func New[T any]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Root() (*Node[T], bool) {
	if t.root == nil {
		return nil, false
	}

	return t.root, true
}

func (t *Tree[T]) AddRoot(node *Node[T]) bool {
	if t.root == nil {
		t.root = node
		return true
	}

	return false
}

func (t *Tree[T]) Add(parentID int, node *Node[T]) bool {
	if t.root == nil {
		return false
	}

	return t.add(parentID, t.root, node)
}

func (t *Tree[T]) Backtrack() []*Node[T] {
	return t.lastAdded.Backtrack()
}

func (t *Tree[T]) Print() {
	if t.root == nil {
		fmt.Println("empty tree")
		return
	}
}

func (t *Tree[T]) add(parentID int, node *Node[T], newNode *Node[T]) bool {
	if parentID == node.ID {
		newNode.Previous = node
		t.lastAdded = newNode
		node.Nexts = append(node.Nexts, newNode)
		return true
	}

	for _, next := range node.Nexts {
		added := t.add(parentID, next, newNode)
		if added {
			return true
		}
	}

	return false
}

func (t *Tree[T]) findNode(current *Node[T], parentID int, depth int) (*Node[T], *Node[T], bool) {
	if current.ID == parentID {
		return current, nil, true
	}

	for _, next := range current.Nexts {
		if next.ID == parentID {
			return current, next, true
		}
	}

	for _, next := range current.Nexts {
		parent, node, found := t.findNode(next, parentID, depth+1)
		if found {
			return parent, node, true
		}
	}

	return nil, nil, false
}
