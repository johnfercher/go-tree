package main

import (
	"fmt"

	"github.com/johnfercher/go-tree/tree"
)

// nolint:gomnd,gocritic
func main() {
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode("0.0").WithID(0))

	tr.Add(0, tree.NewNode("0.1").WithID(1))
	tr.Add(0, tree.NewNode("0.2").WithID(2))

	tr.Add(1, tree.NewNode("1.3").WithID(3))
	tr.Add(1, tree.NewNode("1.4").WithID(4))

	tr.Add(2, tree.NewNode("2.5").WithID(5))
	tr.Add(2, tree.NewNode("2.6").WithID(6))

	root, ok := tr.GetRoot()
	fmt.Println(ok)             // true
	fmt.Println(root.GetData()) // 0.0

	node, ok := tr.Get(3)
	fmt.Println(ok)             // true
	fmt.Println(node.GetData()) // 1.3

	structure, ok := tr.GetStructure()
	fmt.Println(ok)        // true
	fmt.Println(structure) // (NULL) -> (0), (0) -> (1),  (1) -> (3),  (1) -> (4),  (0) -> (2),  (2) -> (5),  (2) -> (6)

	nodes, ok := tr.Backtrack(6)
	fmt.Println(ok) // true
	for _, node := range nodes {
		fmt.Println(node.GetData()) // 2.6; 0.2; 0.0
	}
}
