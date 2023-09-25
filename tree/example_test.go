package tree_test

import (
	"fmt"
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/go-tree/tree"
)

// ExampleNew demonstrates how to create tree.
func ExampleNew() {
	tr := tree.New[string]()

	// Add nodes do tree
	tr.AddRoot(node.New("root"))

	// Do more things
}

// ExampleTree_AddRoot demonstrates how to add root node to tree.
func ExampleTree_AddRoot() {
	tr := tree.New[int]()

	tr.AddRoot(node.New(42))

	// Do more things
}

// ExampleTree_GetRoot demonstrates how to retrieve root node from tree.
func ExampleTree_GetRoot() {
	tr := tree.New[float64]()
	tr.AddRoot(node.New(3.14))

	node, ok := tr.GetRoot()
	if !ok {
		return
	}
	fmt.Println(node.GetData())

	// Do more things
}

// ExampleTree_Add demonstrates how to add node to tree.
func ExampleTree_Add() {
	tr := tree.New[bool]()
	tr.AddRoot(node.New(true))

	tr.Add(0, node.New(false))

	// Do more things
}

// ExampleTree_Get demonstrates how to retrieve node from tree.
func ExampleTree_Get() {
	tr := tree.New[uint]()
	tr.AddRoot(node.New(uint(42)))

	node, ok := tr.Get(0)
	if !ok {
		return
	}
	fmt.Println(node.GetData())

	// Do more things
}

// ExampleTree_Backtrack demonstrates how to retrieve path of nodes from node to root.
func ExampleTree_Backtrack() {
	tr := tree.New[string]()
	tr.AddRoot(node.New("root"))
	tr.Add(0, node.New("level1"))
	tr.Add(1, node.New("level2"))
	tr.Add(2, node.New("leaf"))

	nodes, ok := tr.Backtrack(3)
	if !ok {
		return
	}
	for _, node := range nodes {
		fmt.Println(node.GetData())
	}

	// Do more things
}

// ExampleTree_GetStructure demonstrates how to retrieve tree structure.
func ExampleTree_GetStructure() {
	tr := tree.New[string]()
	tr.AddRoot(node.New("root"))
	tr.Add(0, node.New("level1"))
	tr.Add(1, node.New("level2"))
	tr.Add(2, node.New("leaf"))

	structure, ok := tr.GetStructure()
	if !ok {
		return
	}
	for _, str := range structure {
		fmt.Println(str)
	}

	// Do more things
}

// New demonstrates how to create a node with ID.
func ExampleNewNode() {
	n := node.New("node")

	n.GetData()

	// Do more things
}

// ExampleNode_GetData demonstrates how to retrieve data from node.
func ExampleNode_GetData() {
	n := node.New(3.14)

	data := n.GetData()
	fmt.Println(data)

	// Do more things
}

// ExampleNode_GetID demonstrates how to retrieve id from node.
func ExampleNode_GetID() {
	n := node.New(3.14).WithID(1)

	id := n.GetID()
	fmt.Println(id)

	// Do more things
}

// ExampleNode_GetNexts demonstrates how to retrieve next nodes from node.
func ExampleNode_GetNexts() {
	root := node.New("root")
	leaf := node.New("leaf")

	root.AddNext(leaf)
	nexts := root.GetNexts()
	fmt.Println(len(nexts))

	// Do more things
}

// ExampleNode_GetPrevious demonstrates how to retrieve next nodes from node.
func ExampleNode_GetPrevious() {
	root := node.New("root")
	leaf := node.New("leaf")

	root.AddNext(leaf)
	previous := leaf.GetPrevious()
	fmt.Println(previous.GetData())

	// Do more things
}

// ExampleNode_IsRoot demonstrates how to retrieve info if node is root.
func ExampleNode_IsRoot() {
	n := node.New('b')

	root := n.IsRoot()
	fmt.Println(root)

	// Do more things
}

// ExampleNode_IsLeaf demonstrates how to retrieve info if node is leaf.
func ExampleNode_IsLeaf() {
	n1 := node.New('a')
	n2 := node.New('b')

	n1.AddNext(n2)

	leaf := n2.IsLeaf()
	fmt.Println(leaf)

	// Do more things
}

// ExampleNode_Backtrack demonstrates how to retrieve the path between node to root.
func ExampleNode_Backtrack() {
	n1 := node.New('a')
	n2 := node.New('b')
	n3 := node.New('c')

	n1.AddNext(n2)
	n2.AddNext(n3)

	nodes := n3.Backtrack()
	for _, node := range nodes {
		fmt.Println(node.GetData())
	}

	// Do more things
}

// ExampleNode_GetStructure demonstrates how to retrieve the tree structure from node.
func ExampleNode_GetStructure() {
	n1 := node.New('a')
	n2 := node.New('b')
	n3 := node.New('c')

	n1.AddNext(n2)
	n2.AddNext(n3)

	structure := n3.GetStructure()
	for _, str := range structure {
		fmt.Println(str)
	}

	// Do more things
}

// ExampleNode_AddNext demonstrates how to add a node to a parent.
func ExampleNode_AddNext() {
	n1 := node.New('a')
	n2 := node.New('b')

	n1.AddNext(n2)

	// Do more things
}
