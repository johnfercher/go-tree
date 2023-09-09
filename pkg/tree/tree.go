package tree

// nolint:structcheck,gocritic
type Tree[T any] struct {
	root *Node[T]
}

func New[T any]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) AddRoot(node *Node[T]) (addedRoot bool) {
	if t.root == nil {
		t.root = node
		return true
	}

	return false
}

func (t *Tree[T]) GetRoot() (root *Node[T], hasRoot bool) {
	if t.root == nil {
		return nil, false
	}

	return t.root, true
}

func (t *Tree[T]) Add(parentID int, node *Node[T]) (addedNode bool) {
	if t.root == nil {
		return false
	}

	return t.add(parentID, t.root, node)
}

func (t *Tree[T]) Get(id int) (node *Node[T], found bool) {
	if t.root == nil {
		return nil, false
	}

	if t.root.id == id {
		return t.root, true
	}

	return t.get(id, t.root)
}

func (t *Tree[T]) Backtrack(id int) ([]*Node[T], bool) {
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

func (t *Tree[T]) add(parentID int, parentNode *Node[T], newNode *Node[T]) bool {
	if parentID == parentNode.id {
		parentNode.AddNext(newNode)
		return true
	}

	for _, next := range parentNode.nexts {
		added := t.add(parentID, next, newNode)
		if added {
			return true
		}
	}

	return false
}

func (t *Tree[T]) get(id int, parent *Node[T]) (*Node[T], bool) {
	for _, next := range parent.nexts {
		if next.id == id {
			return next, true
		}

		node, found := t.get(id, next)
		if found {
			return node, true
		}
	}

	return nil, false
}
