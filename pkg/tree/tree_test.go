package tree

import (
	"testing"
)

type anyType struct {
	Value int
}

func TestTree_Add(t *testing.T) {
	// Arrange
	tree := New[*anyType]()

	// Act
	tree.AddRoot(NewNode(0, &anyType{Value: 0}))
	tree.Add(0, NewNode(1, &anyType{Value: 0}))
	tree.Add(1, NewNode(3, &anyType{Value: 0}))
	tree.Add(1, NewNode(4, &anyType{Value: 0}))
	tree.Add(4, NewNode(5, &anyType{Value: 0}))

	nodes := tree.Backtrack()
	for _, node := range nodes {
		node.Print("label")
	}
}
