package tree

import "github.com/johnfercher/go-tree/node"

// nolint:structcheck,gocritic
// Tree represents the main entity of the package.
type Tree[T any] struct {
	root *node.Node[T]
}

// New creates a new Tree.
func New[T any]() *Tree[T] {
	return &Tree[T]{}
}

// AddRoot adds a root node to Tree.
func (t *Tree[T]) AddRoot(n *node.Node[T]) (addedRoot bool) {
	if t.root == nil {
		t.root = n
		return true
	}

	return false
}

// GetRoot retrieves the root node from Tree.
func (t *Tree[T]) GetRoot() (root *node.Node[T], hasRoot bool) {
	if t.root == nil {
		return nil, false
	}

	return t.root, true
}

// Add adds a node into a parent node.
func (t *Tree[T]) Add(parentID int, node *node.Node[T]) (addedNode bool) {
	if t.root == nil {
		return false
	}

	return t.add(parentID, t.root, node)
}

// Get retrieves node from Tree.
func (t *Tree[T]) Get(id int) (node *node.Node[T], found bool) {
	if t.root == nil {
		return nil, false
	}

	if t.root.GetID() == id {
		return t.root, true
	}

	return t.get(id, t.root)
}

// Backtrack retrieves a path from node to root.
func (t *Tree[T]) Backtrack(id int) ([]*node.Node[T], bool) {
	n, found := t.Get(id)
	if !found {
		return nil, found
	}

	return n.Backtrack(), true
}

// GetStructure retrieves Tree structure.
func (t *Tree[T]) GetStructure() ([]string, bool) {
	if t.root == nil {
		return nil, false
	}

	return t.root.GetStructure(), true
}

// Filter remove all sub-nodes that doesn´t respect a rule.
func (t *Tree[T]) Filter(filterFunc func(obj T) bool) (*Tree[T], bool) {
	if t.root == nil {
		return nil, false
	}

	newRoot, ok := t.root.Filter(filterFunc)
	if !ok {
		return nil, false
	}

	newTree := New[T]()
	newTree.AddRoot(newRoot)

	return newTree, true
}

func (t *Tree[T]) add(parentID int, parentNode *node.Node[T], newNode *node.Node[T]) bool {
	if parentID == parentNode.GetID() {
		parentNode.AddNext(newNode)
		return true
	}

	for _, next := range parentNode.GetNexts() {
		added := t.add(parentID, next, newNode)
		if added {
			return true
		}
	}

	return false
}

func (t *Tree[T]) get(id int, parent *node.Node[T]) (*node.Node[T], bool) {
	for _, next := range parent.GetNexts() {
		if next.GetID() == id {
			return next, true
		}

		node, found := t.get(id, next)
		if found {
			return node, true
		}
	}

	return nil, false
}
