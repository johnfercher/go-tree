package main

import (
	"fmt"

	"github.com/johnfercher/tree/pkg/node"
	"github.com/johnfercher/tree/pkg/tree"
)

func main() {
	tr := tree.New[string]()

	tr.AddRoot(node.New(0, "0.0"))

	tr.Add(0, node.New(1, "0.1"))
	tr.Add(0, node.New(2, "0.2"))

	tr.Add(1, node.New(3, "1.3"))
	tr.Add(1, node.New(4, "1.4"))

	tr.Add(2, node.New(5, "2.5"))
	tr.Add(2, node.New(6, "2.6"))

	root, ok := tr.GetRoot()
	node, ok := tr.Get(3)
	structure, ok := tr.GetStructure()
	nodes, ok := tr.Backtrack(6)

	fmt.Println(ok)
	fmt.Println(root)
	fmt.Println(node)
	fmt.Println(structure)
	fmt.Println(nodes)
}
