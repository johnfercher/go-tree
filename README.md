# go-tree

[![GoDoc](https://godoc.org/github.com/johnfercher/go-tree?status.svg)](https://godoc.org/github.com/johnfercher/go-tree)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfercher/go-tree)](https://goreportcard.com/report/github.com/johnfercher/go-tree)
[![CI](https://github.com/johnfercher/go-tree/actions/workflows/goci.yml/badge.svg)](https://github.com/johnfercher/go-tree/actions/workflows/goci.yml)
[![Lint](https://github.com/johnfercher/go-tree/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/johnfercher/go-tree/actions/workflows/golangci-lint.yml)
[![Codecov](https://codecov.io/gh/johnfercher/go-tree/branch/main/graph/badge.svg)](https://codecov.io/gh/johnfercher/go-tree)
[![Visits Badge](https://badges.pufler.dev/visits/johnfercher/go-tree)](https://badges.pufler.dev)


A generic unbalanced tree implementation, where you can define which node will be added to each node.

## Installation

* With `go get`:

```bash
go get -u github.com/johnfercher/go-tree
```

## Contributing

| Command         | Description                                      | Dependencies                                                  |
|-----------------|--------------------------------------------------|---------------------------------------------------------------|
| `make build`    | Build project                                    | `go`                                                          |
| `make test`     | Run unit tests                                   | `go`                                                          |
| `make fmt`      | Format files                                     | `gofmt`, `gofumpt` and `goimports`                            |
| `make lint`     | Check files                                      | `golangci-lint` and `goreportcard-cli`                        |
| `make dod`      | (Definition of Done) Format files and check files | Same as `make build`, `make test`, `make fmt` and `make lint` | 
| `make install`  | Install all dependencies                         | `go`, `curl` and `git`                                        |
| `make examples` | Run all examples                                 | `go`                                                          |

## Features
### Node
* [NewNode](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#NewNode)
* [AddNext](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.AddNext)
* [Backtrack](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.Backtrack)
* [Get](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.Get)
* [GetNexts](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.GetNexts)
* [GetPrevious](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.GetPrevious)
* [GetStructure](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.GetStructure)
* [IsLeaf](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.IsLeaf)
* [IsRoot](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.IsLeaf)
* [Filter](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Node.Filter)

### Tree
* [New](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#New)
* [Add](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Tree.Add)
* [AddRoot](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Tree.AddRoot)
* [Backtrack](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Tree.Backtrack)
* [Get](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Tree.Get)
* [GetRoot](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Tree.GetRoot)
* [GetStructure](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Tree.GetStructure)
* [Filter](https://pkg.go.dev/github.com/johnfercher/go-tree/tree#Tree.Filter)

## Example

```golang
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
```
