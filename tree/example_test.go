package tree_test

import (
	"fmt"

	"github.com/johnfercher/go-tree/tree"

	"github.com/johnfercher/go-tree/node"
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
