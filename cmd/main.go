package main

import (
	"fmt"

	"github.com/johnfercher/tree/pkg/tree"
)

// nolint:gomnd,gocritic
func main() {
	tr := tree.New[string]()

	tr.AddRoot(tree.NewNode(0, "0.0"))

	tr.Add(0, tree.NewNode(1, "0.1"))
	tr.Add(0, tree.NewNode(2, "0.2"))

	tr.Add(1, tree.NewNode(3, "1.3"))
	tr.Add(1, tree.NewNode(4, "1.4"))

	tr.Add(2, tree.NewNode(5, "2.5"))
	tr.Add(2, tree.NewNode(6, "2.6"))

	root, ok := tr.GetRoot()
	fmt.Println(ok)         // true
	fmt.Println(root.Get()) // 0, 0.0

	node, ok := tr.Get(3)
	fmt.Println(ok)         // true
	fmt.Println(node.Get()) // 3, 1.3

	structure, ok := tr.GetStructure()
	fmt.Println(ok)        // true
	fmt.Println(structure) // (NULL) -> (0), (0) -> (1),  (1) -> (3),  (1) -> (4),  (0) -> (2),  (2) -> (5),  (2) -> (6)

	nodes, ok := tr.Backtrack(6)
	fmt.Println(ok) // true
	for _, node := range nodes {
		fmt.Println(node.Get()) // 6, 2.6; 2, 0.2; 0, 0.0
	}
}
