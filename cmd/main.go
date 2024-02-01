package main

import (
	"fmt"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/go-tree/tree"
)

// nolint:gomnd,gocritic
func main() {
	tr := tree.New[string]()

	tr.AddRoot(node.New("0.0").WithID(0))

	tr.Add(0, node.New("0.1").WithID(1))
	tr.Add(0, node.New("0.2").WithID(2))

	tr.Add(1, node.New("1.3").WithID(3))
	tr.Add(1, node.New("1.4").WithID(4))

	tr.Add(2, node.New("2.5").WithID(5))
	tr.Add(2, node.New("2.6").WithID(6))

	root, ok := tr.GetRoot()
	fmt.Println(ok)             // true
	fmt.Println(root.GetData()) // 0.0

	found, ok := tr.Get(3)
	fmt.Println(ok)              // true
	fmt.Println(found.GetData()) // 1.3

	structure, ok := tr.GetStructure()
	fmt.Println(ok)        // true
	fmt.Println(structure) // (NULL) -> (0), (0) -> (1),  (1) -> (3),  (1) -> (4),  (0) -> (2),  (2) -> (5),  (2) -> (6)

	nodes, ok := tr.Backtrack(6)
	fmt.Println(ok) // true
	for _, n := range nodes {
		fmt.Println(n.GetData()) // 2.6; 0.2; 0.0
	}
}
