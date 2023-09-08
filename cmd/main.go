package main

import "github.com/johnfercher/tree/pkg/tree"

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func main() {
	node := tree.NewNode(0, &Vector3D{
		X: 0,
		Y: 0,
		Z: 0,
	})

	v := node.GetContent()
}
