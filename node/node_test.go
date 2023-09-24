package node_test

import (
	"fmt"
	"github.com/johnfercher/go-tree/node"
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
		"*node_test.anyType",
	}

	for i, element := range elements {
		// Act
		sut := node.New(element)
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
	sut := node.New(42)

	// Act
	data := sut.GetData()

	// Assert
	assert.Equal(t, 42, data)
}

func TestNode_GetNexts(t *testing.T) {
	// Arrange
	sut := node.New(42)
	leaf := node.New(43)
	sut.AddNext(leaf)

	// Act
	nexts := sut.GetNexts()

	// Assert
	assert.Equal(t, 1, len(nexts))
	data := nexts[0].GetData()
	assert.Equal(t, 43, data)
}

func TestNode_GetPrevious(t *testing.T) {
	// Arrange
	root := node.New(42)
	sut := node.New(43)
	root.AddNext(sut)

	// Act
	previous := sut.GetPrevious()

	// Assert
	data := previous.GetData()
	assert.Equal(t, 42, data)
}

func TestNode_IsRoot(t *testing.T) {
	// Arrange
	root := node.New(42)
	anyNode := node.New(43)
	leaf := node.New(44)

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
	root := node.New(42)
	anyNode := node.New(43)
	leaf := node.New(44)

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
	root := node.New(42)
	anyNode := node.New(43)
	leaf := node.New(44)

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

	tr.AddRoot(node.New("0.0").WithID(0))
	tr.Add(0, node.New("1.0").WithID(1))
	tr.Add(0, node.New("1.1").WithID(2))
	tr.Add(0, node.New("1.2").WithID(3))
	tr.Add(1, node.New("2.0").WithID(4))
	tr.Add(1, node.New("2.1").WithID(5))
	tr.Add(1, node.New("2.2").WithID(6))
	tr.Add(2, node.New("2.0").WithID(7))
	tr.Add(2, node.New("2.1").WithID(8))
	tr.Add(2, node.New("2.2").WithID(9))
	tr.Add(3, node.New("3.0").WithID(10))
	tr.Add(3, node.New("3.1").WithID(11))
	tr.Add(3, node.New("3.2").WithID(12))
	tr.Add(4, node.New("4.0").WithID(13))

	// Act & Assert
	// Case 0
	anyNode, _ := tr.Get(0)
	backtracked := anyNode.Backtrack()
	assert.Equal(t, 1, len(backtracked))
	data := backtracked[0].GetData()
	assert.Equal(t, "0.0", data)

	// Case 1
	anyNode, _ = tr.Get(1)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "1.0", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "0.0", data)

	// Case 2
	anyNode, _ = tr.Get(2)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "1.1", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "0.0", data)

	// Case 3
	anyNode, _ = tr.Get(3)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 2, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "1.2", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "0.0", data)

	// Case 4
	anyNode, _ = tr.Get(4)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "2.0", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.0", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 5
	anyNode, _ = tr.Get(5)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "2.1", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.0", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 6
	anyNode, _ = tr.Get(6)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "2.2", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.0", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 7
	anyNode, _ = tr.Get(7)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "2.0", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.1", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 8
	anyNode, _ = tr.Get(8)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "2.1", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.1", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 9
	anyNode, _ = tr.Get(9)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "2.2", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.1", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 10
	anyNode, _ = tr.Get(10)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "3.0", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.2", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 11
	anyNode, _ = tr.Get(11)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "3.1", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.2", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 12
	anyNode, _ = tr.Get(12)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 3, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "3.2", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "1.2", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "0.0", data)

	// Case 13
	anyNode, _ = tr.Get(13)
	backtracked = anyNode.Backtrack()
	assert.Equal(t, 4, len(backtracked))
	data = backtracked[0].GetData()
	assert.Equal(t, "4.0", data)
	data = backtracked[1].GetData()
	assert.Equal(t, "2.0", data)
	data = backtracked[2].GetData()
	assert.Equal(t, "1.0", data)
	data = backtracked[3].GetData()
	assert.Equal(t, "0.0", data)
}

func TestNode_GetStructure(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(node.New("0.0").WithID(0))
	tr.Add(0, node.New("1.0").WithID(1))
	tr.Add(0, node.New("1.1").WithID(2))
	tr.Add(0, node.New("1.2").WithID(3))
	tr.Add(1, node.New("2.0").WithID(4))
	tr.Add(1, node.New("2.1").WithID(5))
	tr.Add(1, node.New("2.2").WithID(6))
	tr.Add(2, node.New("2.0").WithID(7))
	tr.Add(2, node.New("2.1").WithID(8))
	tr.Add(2, node.New("2.2").WithID(9))
	tr.Add(3, node.New("3.0").WithID(10))
	tr.Add(3, node.New("3.1").WithID(11))
	tr.Add(3, node.New("3.2").WithID(12))
	tr.Add(4, node.New("4.0").WithID(13))

	node, _ := tr.Get(0)

	// Act
	structure := node.GetStructure()

	// Assert
	assert.Equal(t, 14, len(structure))
	for _, str := range structure {
		assert.NotEmpty(t, str)
	}
}
