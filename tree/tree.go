package tree

import (
	"cmp"
)



type NodeTree[I cmp.Ordered, T any] interface {
	GetId() I
	GetPid() I
	AppendChildren(T)
}

type Node[I cmp.Ordered, U NodeTree[I, U]] interface {
	MapId() I
	MapTree() U
}

// IntoTree 列表转树, 切片无children
// 两次循环就可以获取列表转树
func IntoTree[I cmp.Ordered, E Node[I, U], U NodeTree[I, U]](rows []E, rootPid I) []U {
	nodes := make([]U, 0, len(rows))
	nodeMaps := make(map[I]U)
	for _, v := range rows {
		e := v.MapTree()
		nodes = append(nodes, e)
		nodeMaps[v.MapId()] = e
	}
	return intoTree(nodeMaps, nodes, rootPid)
}

// IntoTree 列表转树, 切片有children
// 两次循环就可以获取列表转树
func IntoTree2[T cmp.Ordered, E NodeTree[T, E]](rows []E, rootPid T) []E {
	nodeMaps := make(map[T]E)
	for _, e := range rows {
		nodeMaps[e.GetId()] = e
	}
	return intoTree(nodeMaps, rows, rootPid)
}

func intoTree[T cmp.Ordered, E NodeTree[T, E]](nodeMaps map[T]E, rows []E, rootPid T) []E {
	var root []E
	for _, e := range rows {
		pid := e.GetPid()
		if pid == rootPid {
			root = append(root, e)
		} else if parent, exists := nodeMaps[pid]; exists {
			parent.AppendChildren(e)
		}
	}
	return root
}
