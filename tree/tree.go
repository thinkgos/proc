package tree

import (
	"cmp"
)

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

type NodeTree[I cmp.Ordered, T any] interface {
	GetId() I
	GetPid() I
	AppendChildren(T)
}

type Node[I cmp.Ordered, U NodeTree[I, U]] interface {
	MapId() I
	MapTree() U
}

// IntoTree 列表转树, 切片无children, map转换
// 两次循环就可以获取列表转树
func IntoTree[I cmp.Ordered, E Node[I, U], U NodeTree[I, U]](rows []E, rootPid I) []U {
	nodes := make([]U, 0, len(rows))
	nodeMaps := make(map[I]U)
	for _, v := range rows {
		vv := v.MapTree()
		nodes = append(nodes, vv)
		nodeMaps[v.MapId()] = vv
	}
	return intoTree(nodeMaps, nodes, rootPid)
}

// IntoTree 列表转树
// 两次循环就可以获取列表转树
func IntoTree2[T cmp.Ordered, E NodeTree[T, E]](rows []E, rootPid T) []E {
	nodeMaps := make(map[T]E)
	for _, v := range rows {
		vv := v
		nodeMaps[v.GetId()] = vv
	}
	return intoTree(nodeMaps, rows, rootPid)
}

func intoTree[T cmp.Ordered, E NodeTree[T, E]](nodeMaps map[T]E, rows []E, rootPid T) []E {
	var root []E
	for _, v := range rows {
		node := v
		pid := node.GetPid()
		if pid == rootPid {
			root = append(root, node)
		} else if parent, exists := nodeMaps[pid]; exists {
			parent.AppendChildren(node)
		}
	}
	return root
}
