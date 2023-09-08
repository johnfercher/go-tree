package tree

import (
	"github.com/johnfercher/tree/pkg/node"
)

type Tree[T any] struct {
	root *node.Node[T]
}

func New[T any]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) AddRoot(node *node.Node[T]) (addedRoot bool) {
	if t.root == nil {
		t.root = node
		return true
	}

	return false
}

func (t *Tree[T]) GetRoot() (root *node.Node[T], hasRoot bool) {
	if t.root == nil {
		return nil, false
	}

	return t.root, true
}

func (t *Tree[T]) Add(parentID int, node *node.Node[T]) (addedNode bool) {
	if t.root == nil {
		return false
	}

	return t.add(parentID, t.root, node)
}

func (t *Tree[T]) Get(id int) (node *node.Node[T], found bool) {
	if t.root == nil {
		return nil, false
	}

	if t.root.ID == id {
		return t.root, true
	}

	return t.get(id, t.root)
}

func (t *Tree[T]) Backtrack(id int) ([]*node.Node[T], bool) {
	n, found := t.Get(id)
	if !found {
		return nil, found
	}

	return n.Backtrack(), true
}

func (t *Tree[T]) GetStructure() ([]string, bool) {
	if t.root == nil {
		return nil, false
	}

	return t.root.GetStructure(), true
}

func (t *Tree[T]) add(parentID int, parentNode *node.Node[T], newNode *node.Node[T]) bool {
	if parentID == parentNode.ID {
		parentNode.AddNext(newNode)
		return true
	}

	nexts := parentNode.GetNexts()
	for _, next := range nexts {
		added := t.add(parentID, next, newNode)
		if added {
			return true
		}
	}

	return false
}

func (t *Tree[T]) get(id int, parent *node.Node[T]) (*node.Node[T], bool) {
	nexts := parent.GetNexts()
	for _, next := range nexts {
		if next.ID == id {
			return next, true
		}

		node, found := t.get(id, next)
		if found {
			return node, true
		}
	}

	return nil, false
}
