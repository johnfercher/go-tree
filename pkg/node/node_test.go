package node_test

import (
	"fmt"
	"github.com/johnfercher/tree/pkg/node"
	"github.com/johnfercher/tree/pkg/tree"
	"github.com/stretchr/testify/assert"
	"testing"
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
		"*node_test.anyType",
	}

	for i, element := range elements {
		// Act
		sut := node.New(0, element)
		data := sut.GetData()

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*node.Node[interface {}]", fmt.Sprintf("%T", sut))
		assert.Equal(t, element, data)
		assert.Equal(t, types[i], fmt.Sprintf("%T", data))
	}
}

func TestNode_GetData(t *testing.T) {
	// Arrange
	sut := node.New(0, 42)

	// Act
	n := sut.GetData()

	// Assert
	assert.Equal(t, 42, n)
}

func TestNode_IsRoot(t *testing.T) {
	// Arrange
	root := node.New(0, 42)
	anyNode := node.New(1, 43)
	leaf := node.New(2, 44)

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
	root := node.New(0, 42)
	anyNode := node.New(1, 43)
	leaf := node.New(2, 44)

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
	root := node.New(0, 42)
	anyNode := node.New(1, 43)
	leaf := node.New(2, 44)

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

	tr.AddRoot(node.New(0, "0.0"))
	tr.Add(0, node.New(1, "1.0"))
	tr.Add(0, node.New(2, "1.1"))
	tr.Add(0, node.New(3, "1.2"))
	tr.Add(1, node.New(4, "2.0"))
	tr.Add(1, node.New(5, "2.1"))
	tr.Add(1, node.New(6, "2.2"))
	tr.Add(2, node.New(7, "2.0"))
	tr.Add(2, node.New(8, "2.1"))
	tr.Add(2, node.New(9, "2.2"))
	tr.Add(3, node.New(10, "3.0"))
	tr.Add(3, node.New(11, "3.1"))
	tr.Add(3, node.New(12, "3.2"))
	tr.Add(4, node.New(13, "4.0"))

	// Act & Assert
	anyNode, _ := tr.Get(0)
	backtracked := anyNode.Backtrack()
	assert.Equal(t, 1, len(backtracked))
	assert.Equal(t, "0.0", backtracked[0].GetData())

	anyNode, _ = tr.Get(1)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	assert.Equal(t, "1.0", backtracked[0].GetData())
	assert.Equal(t, "0.0", backtracked[1].GetData())

	anyNode, _ = tr.Get(2)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	assert.Equal(t, "1.1", backtracked[0].GetData())
	assert.Equal(t, "0.0", backtracked[1].GetData())

	anyNode, _ = tr.Get(3)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	assert.Equal(t, "1.2", backtracked[0].GetData())
	assert.Equal(t, "0.0", backtracked[1].GetData())

	anyNode, _ = tr.Get(4)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "2.0", backtracked[0].GetData())
	assert.Equal(t, "1.0", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(5)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "2.1", backtracked[0].GetData())
	assert.Equal(t, "1.0", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(6)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "2.2", backtracked[0].GetData())
	assert.Equal(t, "1.0", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(7)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "2.0", backtracked[0].GetData())
	assert.Equal(t, "1.1", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(8)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "2.1", backtracked[0].GetData())
	assert.Equal(t, "1.1", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(9)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "2.2", backtracked[0].GetData())
	assert.Equal(t, "1.1", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(10)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "3.0", backtracked[0].GetData())
	assert.Equal(t, "1.2", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(11)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "3.1", backtracked[0].GetData())
	assert.Equal(t, "1.2", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(12)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	assert.Equal(t, "3.2", backtracked[0].GetData())
	assert.Equal(t, "1.2", backtracked[1].GetData())
	assert.Equal(t, "0.0", backtracked[2].GetData())

	anyNode, _ = tr.Get(13)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 4, len(backtracked))
	assert.Equal(t, "4.0", backtracked[0].GetData())
	assert.Equal(t, "2.0", backtracked[1].GetData())
	assert.Equal(t, "1.0", backtracked[2].GetData())
	assert.Equal(t, "0.0", backtracked[3].GetData())
}

func TestNode_GetStructure(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(node.New(0, "0.0"))
	tr.Add(0, node.New(1, "1.0"))
	tr.Add(0, node.New(2, "1.1"))
	tr.Add(0, node.New(3, "1.2"))
	tr.Add(1, node.New(4, "2.0"))
	tr.Add(1, node.New(5, "2.1"))
	tr.Add(1, node.New(6, "2.2"))
	tr.Add(2, node.New(7, "2.0"))
	tr.Add(2, node.New(8, "2.1"))
	tr.Add(2, node.New(9, "2.2"))
	tr.Add(3, node.New(10, "3.0"))
	tr.Add(3, node.New(11, "3.1"))
	tr.Add(3, node.New(12, "3.2"))
	tr.Add(4, node.New(13, "4.0"))

	node, _ := tr.Get(0)

	// Act
	structure := node.GetStructure()

	// Assert
	assert.Equal(t, 14, len(structure))
	for _, str := range structure {
		assert.NotEmpty(t, str)
	}
}
