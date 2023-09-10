package tree_test

import (
	"fmt"

	"github.com/johnfercher/go-tree/tree"
)

// ExampleNew demonstrates how to create tree.
func ExampleNew() {
	tr := tree.New[string]()

	// Add nodes do tree
	tr.AddRoot(tree.NewNode(0, "root"))

	// Do more things
}

// ExampleTree_AddRoot demonstrates how to add root node to tree.
func ExampleTree_AddRoot() {
	tr := tree.New[int]()

	tr.AddRoot(tree.NewNode(0, 42))

	// Do more things
}

// ExampleTree_GetRoot demonstrates how to retrieve root node from tree.
func ExampleTree_GetRoot() {
	tr := tree.New[float64]()
	tr.AddRoot(tree.NewNode(0, 3.14))

	node, ok := tr.GetRoot()
	if !ok {
		return
	}
	fmt.Println(node.Get())

	// Do more things
}

// ExampleTree_Add demonstrates how to add node to tree.
func ExampleTree_Add() {
	tr := tree.New[bool]()
	tr.AddRoot(tree.NewNode(0, true))

	tr.Add(0, tree.NewNode(1, false))

	// Do more things
}

// ExampleTree_Get demonstrates how to retrieve node from tree.
func ExampleTree_Get() {
	tr := tree.New[uint]()
	tr.AddRoot(tree.NewNode(0, uint(42)))

	node, ok := tr.Get(0)
	if !ok {
		return
	}
	fmt.Println(node.Get())

	// Do more things
}

// ExampleTree_Backtrack demonstrates how to retrieve path of nodes from node to root.
func ExampleTree_Backtrack() {
	tr := tree.New[string]()
	tr.AddRoot(tree.NewNode(0, "root"))
	tr.Add(0, tree.NewNode(1, "level1"))
	tr.Add(1, tree.NewNode(2, "level2"))
	tr.Add(2, tree.NewNode(3, "leaf"))

	nodes, ok := tr.Backtrack(3)
	if !ok {
		return
	}
	for _, node := range nodes {
		fmt.Println(node.Get())
	}

	// Do more things
}

// ExampleTree_GetStructure demonstrates how to retrieve tree structure.
func ExampleTree_GetStructure() {
	tr := tree.New[string]()
	tr.AddRoot(tree.NewNode(0, "root"))
	tr.Add(0, tree.NewNode(1, "level1"))
	tr.Add(1, tree.NewNode(2, "level2"))
	tr.Add(2, tree.NewNode(3, "leaf"))

	structure, ok := tr.GetStructure()
	if !ok {
		return
	}
	for _, str := range structure {
		fmt.Println(str)
	}

	// Do more things
}

// ExampleNewNode demonstrates how to create a node.
func ExampleNewNode() {
	n := tree.NewNode(0, "node")

	n.Get()

	// Do more things
}

// ExampleNode_Get demonstrates how to retrieve id and data from node.
func ExampleNode_Get() {
	n := tree.NewNode(0, 3.14)

	id, data := n.Get()
	fmt.Println(id)
	fmt.Println(data)

	// Do more things
}

// ExampleNode_GetNexts demonstrates how to retrieve next nodes from node.
func ExampleNode_GetNexts() {
	root := tree.NewNode(0, "root")
	leaf := tree.NewNode(1, "leaf")

	root.AddNext(leaf)
	nexts := root.GetNexts()
	fmt.Println(len(nexts))

	// Do more things
}

// ExampleNode_GetPrevious demonstrates how to retrieve next nodes from node.
func ExampleNode_GetPrevious() {
	root := tree.NewNode(0, "root")
	leaf := tree.NewNode(1, "leaf")

	root.AddNext(leaf)
	previous := leaf.GetPrevious()
	fmt.Println(previous.Get())

	// Do more things
}

// ExampleNode_IsRoot demonstrates how to retrieve info if node is root.
func ExampleNode_IsRoot() {
	n := tree.NewNode(0, 'b')

	root := n.IsRoot()
	fmt.Println(root)

	// Do more things
}

// ExampleNode_IsLeaf demonstrates how to retrieve info if node is leaf.
func ExampleNode_IsLeaf() {
	n1 := tree.NewNode(0, 'a')
	n2 := tree.NewNode(0, 'b')

	n1.AddNext(n2)

	leaf := n2.IsLeaf()
	fmt.Println(leaf)

	// Do more things
}

// ExampleNode_Backtrack demonstrates how to retrieve the path between node to root.
func ExampleNode_Backtrack() {
	n1 := tree.NewNode(0, 'a')
	n2 := tree.NewNode(0, 'b')
	n3 := tree.NewNode(0, 'c')

	n1.AddNext(n2)
	n2.AddNext(n3)

	nodes := n3.Backtrack()
	for _, node := range nodes {
		fmt.Println(node.Get())
	}

	// Do more things
}

// ExampleNode_GetStructure demonstrates how to retrieve the tree structure from node.
func ExampleNode_GetStructure() {
	n1 := tree.NewNode(0, 'a')
	n2 := tree.NewNode(0, 'b')
	n3 := tree.NewNode(0, 'c')

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
	n1 := tree.NewNode(0, 'a')
	n2 := tree.NewNode(0, 'b')

	n1.AddNext(n2)

	// Do more things
}
