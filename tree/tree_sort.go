package tree

import "cmp"

type NodeSlice[I cmp.Ordered, E any] []NodeTree[I, E]

func (nodes NodeSlice[T, E]) Len() int {
	return len(nodes)
}
func (nodes NodeSlice[T, E]) Less(i, j int) bool {
	return nodes[i].GetId() < nodes[j].GetId()
}
func (nodes NodeSlice[T, E]) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}
