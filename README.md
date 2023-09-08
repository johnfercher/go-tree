# Tree

[![GoDoc](https://godoc.org/github.com/johnfercher/tree?status.svg)](https://godoc.org/github.com/johnfercher/tree)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfercher/tree)](https://goreportcard.com/report/github.com/johnfercher/tree)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#template-engines)  
[![CI](https://github.com/johnfercher/tree/actions/workflows/goci.yml/badge.svg)](https://github.com/johnfercher/tree/actions/workflows/goci.yml)
[![Lint](https://github.com/johnfercher/tree/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/johnfercher/tree/actions/workflows/golangci-lint.yml)
[![Codecov](https://img.shields.io/codecov/c/github/johnfercher/tree)](https://codecov.io/gh/johnfercher/tree)


A generic unbalanced tree implementation, where you can define which node will be added to each node.

```golang
package main

import (
	"fmt"

	"github.com/johnfercher/tree/pkg/node"
	"github.com/johnfercher/tree/pkg/tree"
)

// nolint:gomnd,gocritic
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
	fmt.Println(ok)
	fmt.Println(root)

	node, ok := tr.Get(3)
	fmt.Println(ok)
	fmt.Println(node)

	structure, ok := tr.GetStructure()
	fmt.Println(ok)
	fmt.Println(structure)

	nodes, ok := tr.Backtrack(6)
	fmt.Println(ok)
	fmt.Println(nodes)
}
```
