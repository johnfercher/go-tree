package tree_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/go-tree/tree"

	"github.com/stretchr/testify/assert"
)

type anyType struct {
	Value string
}

func TestNewNode(t *testing.T) {
	// Arrange
	elements := []interface{}{
		true,
		42,
		3.14,
		"string1",
		&anyType{Value: "string2"},
	}

	types := []interface{}{
		"bool",
		"int",
		"float64",
		"string",
		"*tree_test.anyType",
	}

	for i, element := range elements {
		// Act
		sut := tree.NewNode(0, element)
		id, data := sut.Get()

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*tree.Node[interface {}]", fmt.Sprintf("%T", sut))
		assert.Equal(t, 0, id)
		assert.Equal(t, element, data)
		assert.Equal(t, types[i], fmt.Sprintf("%T", data))
	}
}

func TestNode_GetData(t *testing.T) {
	// Arrange
	sut := tree.NewNode(0, 42)

	// Act
	id, n := sut.Get()

	// Assert
	assert.Equal(t, 0, id)
	assert.Equal(t, 42, n)
}

func TestNode_IsRoot(t *testing.T) {
	// Arrange
	root := tree.NewNode(0, 42)
	anyNode := tree.NewNode(1, 43)
	leaf := tree.NewNode(2, 44)

	root.AddNext(anyNode)
	anyNode.AddNext(leaf)

	// Act
	rootTrue := root.IsRoot()
	rootFalse1 := anyNode.IsRoot()
	rootFalse2 := leaf.IsRoot()

	// Assert
	assert.True(t, rootTrue)
	assert.False(t, rootFalse1)
	assert.False(t, rootFalse2)
}

func TestNode_IsLeaf(t *testing.T) {
	// Arrange
	root := tree.NewNode(0, 42)
	anyNode := tree.NewNode(1, 43)
	leaf := tree.NewNode(2, 44)

	root.AddNext(anyNode)
	anyNode.AddNext(leaf)

	// Act
	leafFalse1 := root.IsLeaf()
	leafFalse2 := anyNode.IsLeaf()
	leafTrue := leaf.IsLeaf()

	// Assert
	assert.False(t, leafFalse1)
	assert.False(t, leafFalse2)
	assert.True(t, leafTrue)
}

func TestNode_Backtrack_WhenBuildManually_ShouldReturnCorrectly(t *testing.T) {
	// Arrange
	root := tree.NewNode(0, 42)
	anyNode := tree.NewNode(1, 43)
	leaf := tree.NewNode(2, 44)

	root.AddNext(anyNode)
	anyNode.AddNext(leaf)

	// Act
	arr1 := root.Backtrack()
	arr2 := anyNode.Backtrack()
	arr3 := leaf.Backtrack()

	// Assert
	assert.Equal(t, 1, len(arr1))
	assert.Equal(t, 2, len(arr2))
	assert.Equal(t, 3, len(arr3))
}

func TestNode_Backtrack_WhenBuildByTree_ShouldReturnCorrectly(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode(0, "0.0"))
	tr.Add(0, tree.NewNode(1, "1.0"))
	tr.Add(0, tree.NewNode(2, "1.1"))
	tr.Add(0, tree.NewNode(3, "1.2"))
	tr.Add(1, tree.NewNode(4, "2.0"))
	tr.Add(1, tree.NewNode(5, "2.1"))
	tr.Add(1, tree.NewNode(6, "2.2"))
	tr.Add(2, tree.NewNode(7, "2.0"))
	tr.Add(2, tree.NewNode(8, "2.1"))
	tr.Add(2, tree.NewNode(9, "2.2"))
	tr.Add(3, tree.NewNode(10, "3.0"))
	tr.Add(3, tree.NewNode(11, "3.1"))
	tr.Add(3, tree.NewNode(12, "3.2"))
	tr.Add(4, tree.NewNode(13, "4.0"))

	// Act & Assert
	// Case 0
	anyNode, _ := tr.Get(0)
	backtracked := anyNode.Backtrack()
	assert.Equal(t, 1, len(backtracked))
	id, data := backtracked[0].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 1
	anyNode, _ = tr.Get(1)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 1, id)
	assert.Equal(t, "1.0", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 2
	anyNode, _ = tr.Get(2)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 2, id)
	assert.Equal(t, "1.1", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 3
	anyNode, _ = tr.Get(3)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 3, id)
	assert.Equal(t, "1.2", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 4
	anyNode, _ = tr.Get(4)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 4, id)
	assert.Equal(t, "2.0", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 1, id)
	assert.Equal(t, "1.0", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 5
	anyNode, _ = tr.Get(5)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 5, id)
	assert.Equal(t, "2.1", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 1, id)
	assert.Equal(t, "1.0", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 6
	anyNode, _ = tr.Get(6)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 6, id)
	assert.Equal(t, "2.2", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 1, id)
	assert.Equal(t, "1.0", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 7
	anyNode, _ = tr.Get(7)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 7, id)
	assert.Equal(t, "2.0", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 2, id)
	assert.Equal(t, "1.1", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 8
	anyNode, _ = tr.Get(8)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 8, id)
	assert.Equal(t, "2.1", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 2, id)
	assert.Equal(t, "1.1", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 9
	anyNode, _ = tr.Get(9)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 9, id)
	assert.Equal(t, "2.2", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 2, id)
	assert.Equal(t, "1.1", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 10
	anyNode, _ = tr.Get(10)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 10, id)
	assert.Equal(t, "3.0", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 3, id)
	assert.Equal(t, "1.2", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 11
	anyNode, _ = tr.Get(11)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 11, id)
	assert.Equal(t, "3.1", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 3, id)
	assert.Equal(t, "1.2", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 12
	anyNode, _ = tr.Get(12)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 12, id)
	assert.Equal(t, "3.2", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 3, id)
	assert.Equal(t, "1.2", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)

	// Case 13
	anyNode, _ = tr.Get(13)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 4, len(backtracked))
	id, data = backtracked[0].Get()
	assert.Equal(t, 13, id)
	assert.Equal(t, "4.0", data)
	id, data = backtracked[1].Get()
	assert.Equal(t, 4, id)
	assert.Equal(t, "2.0", data)
	id, data = backtracked[2].Get()
	assert.Equal(t, 1, id)
	assert.Equal(t, "1.0", data)
	id, data = backtracked[3].Get()
	assert.Equal(t, 0, id)
	assert.Equal(t, "0.0", data)
}

func TestNode_GetStructure(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode(0, "0.0"))
	tr.Add(0, tree.NewNode(1, "1.0"))
	tr.Add(0, tree.NewNode(2, "1.1"))
	tr.Add(0, tree.NewNode(3, "1.2"))
	tr.Add(1, tree.NewNode(4, "2.0"))
	tr.Add(1, tree.NewNode(5, "2.1"))
	tr.Add(1, tree.NewNode(6, "2.2"))
	tr.Add(2, tree.NewNode(7, "2.0"))
	tr.Add(2, tree.NewNode(8, "2.1"))
	tr.Add(2, tree.NewNode(9, "2.2"))
	tr.Add(3, tree.NewNode(10, "3.0"))
	tr.Add(3, tree.NewNode(11, "3.1"))
	tr.Add(3, tree.NewNode(12, "3.2"))
	tr.Add(4, tree.NewNode(13, "4.0"))

	node, _ := tr.Get(0)

	// Act
	structure := node.GetStructure()

	// Assert
	assert.Equal(t, 14, len(structure))
	for _, str := range structure {
		assert.NotEmpty(t, str)
	}
}
