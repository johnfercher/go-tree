package tree_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/tree/pkg/tree"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := tree.New[int]()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*tree.Tree[int]", fmt.Sprintf("%T", sut))
}

func TestTree_AddRoot_WhenTreeIsEmpty_ShouldReturnTrue(t *testing.T) {
	// Arrange
	sut := tree.New[int]()

	// Act
	added := sut.AddRoot(tree.NewNode(0, 42))

	// Assert
	assert.True(t, added)
}

func TestTree_AddRoot_WhenTreeIsNotEmpty_ShouldReturnFalse(t *testing.T) {
	// Arrange
	sut := tree.New[int]()

	// Act
	_ = sut.AddRoot(tree.NewNode(0, 42))
	added := sut.AddRoot(tree.NewNode(0, 43))

	// Assert
	assert.False(t, added)
}

func TestTree_GetRoot_WhenThereIsNotRoot_ShouldReturnFalse(t *testing.T) {
	// Arrange
	sut := tree.New[int]()

	// Act
	root, hasRoot := sut.GetRoot()

	// Assert
	assert.Nil(t, root)
	assert.False(t, hasRoot)
}

func TestTree_GetRoot_WhenThereIsRoot_ShouldReturnTrue(t *testing.T) {
	// Arrange
	sut := tree.New[int]()
	sut.AddRoot(tree.NewNode(0, 42))

	// Act
	root, hasRoot := sut.GetRoot()

	// Assert
	assert.NotNil(t, root)
	assert.True(t, hasRoot)
}

func TestTree_Add_WhenThereIsNoRoot_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[int]()

	// Act
	added := tr.Add(0, tree.NewNode(0, 42))

	// Assert
	assert.False(t, added)
}

func TestTree_Add_WhenRootIsNotRight_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[int]()

	// Act
	_ = tr.AddRoot(tree.NewNode(0, 42))
	added := tr.Add(3, tree.NewNode(1, 42))

	// Assert
	assert.False(t, added)
}

func TestTree_Add_WhenThereIsRootAndRootIsRight_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[int]()

	// Act
	_ = tr.AddRoot(tree.NewNode(0, 42))
	added := tr.Add(0, tree.NewNode(1, 42))

	// Assert
	assert.True(t, added)
}

func TestTree_Get_WhenThereIsNoRoot_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	// Act
	n, found := tr.Get(8)

	// Assert
	assert.Nil(t, n)
	assert.False(t, found)
}

func TestTree_Get_WhenThereIsNoId_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode(0, "0"))
	tr.Add(0, tree.NewNode(1, "1.0"))

	// Act
	node, found := tr.Get(8)

	// Assert
	assert.Nil(t, node)
	assert.False(t, found)
}

func TestTree_Get_WhenThereIsIdOnRoot_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode(0, "0"))

	// Act
	node, found := tr.Get(0)

	// Assert
	assert.NotNil(t, node)
	assert.True(t, found)
}

func TestTree_Get_WhenThereIsId_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode(0, "0"))
	tr.Add(0, tree.NewNode(1, "1.0"))

	// Act
	node, found := tr.Get(1)

	// Assert
	assert.NotNil(t, node)
	assert.True(t, found)
}

func TestTree_Backtrack_WhenIdNotFound_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode(0, "0.0"))

	// Act
	n, found := tr.Backtrack(1)

	// Assert
	assert.Nil(t, n)
	assert.False(t, found)
}

func TestTree_Backtrack_WhenIdFound_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode(0, "0.0"))
	tr.Add(0, tree.NewNode(1, "1.0"))
	tr.Add(1, tree.NewNode(2, "2.0"))
	tr.Add(2, tree.NewNode(3, "3.0"))

	// Act
	n, found := tr.Backtrack(3)

	// Assert
	assert.NotNil(t, n)
	assert.True(t, found)
}

func TestTree_GetStructure_WhenThereIsNoRoot_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	// Act
	structure, found := tr.GetStructure()

	// Assert
	assert.Nil(t, structure)
	assert.False(t, found)
}

func TestTree_GetStructure_WhenThereIsRoot_ShouldReturnTrue(t *testing.T) {
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

	// Act
	structure, found := tr.GetStructure()

	// Assert
	assert.NotNil(t, structure)
	assert.True(t, found)
	for _, str := range structure {
		assert.NotEmpty(t, str)
	}
}
