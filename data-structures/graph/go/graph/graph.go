package graph

type Node struct {
	Value int
	Edges []*Node
}

type Graph []*Node
