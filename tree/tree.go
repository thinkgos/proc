package tree

import (
	"cmp"
	"slices"
)

type Node[I cmp.Ordered, U NodeTree[I, U]] interface {
	MapId() I
	MapTree() U
}

type NodeTree[I cmp.Ordered, T any] interface {
	GetId() I
	GetPid() I
	AppendChildren(T)
}

type NodeSort[T any] interface {
	SortChildren(cmp func(a, b T) int)
}

// Map implement Node map to NodeTree
func Map[I cmp.Ordered, E Node[I, U], U NodeTree[I, U]](rows []E) []U {
	nodes := make([]U, 0, len(rows))
	for _, v := range rows {
		nodes = append(nodes, v.MapTree())
	}
	return nodes
}

// IntoTree 列表转树, 切片无children
// 元素顺序由x本身顺序决定, 可提前排序, 然后转树(或使用 SortFunc)
func IntoTree[T cmp.Ordered, E Node[T, U], U NodeTree[T, U]](x []E, rootPid T) []U {
	nodeMaps, nodes := intoMapTree(x)
	return intoTree(nodeMaps, nodes, rootPid)
}

// IntoTree 列表转树, 切片有children, 顺序由x本身顺序决定
// 元素顺序由x本身顺序决定, 可提前排序, 然后转树(或使用 SortFunc)
func IntoTree2[T cmp.Ordered, E NodeTree[T, E]](x []E, rootPid T) []E {
	nodeMaps := intoMap(x)
	return intoTree(nodeMaps, x, rootPid)
}

// SortFunc 树排序
func SortFunc[T NodeSort[T]](x []T, cmp func(a, b T) int) {
	if len(x) == 0 {
		return
	}
	slices.SortFunc(x, cmp)
	for _, v := range x {
		v.SortChildren(cmp)
	}
}

// T -> U 的映射
// E -> U 转换
func intoMapTree[T cmp.Ordered, E Node[T, U], U NodeTree[T, U]](x []E) (map[T]U, []U) {
	nodes := make([]U, 0, len(x))
	nodeMaps := make(map[T]U)
	for _, v := range x {
		e := v.MapTree()
		nodes = append(nodes, e)
		nodeMaps[v.MapId()] = e
	}
	return nodeMaps, nodes
}

// T -> E 映射
func intoMap[T cmp.Ordered, E NodeTree[T, E]](x []E) map[T]E {
	// T -> E 映射
	nodeMaps := make(map[T]E)
	for _, e := range x {
		nodeMaps[e.GetId()] = e
	}
	return nodeMaps
}

// 转树
func intoTree[T cmp.Ordered, E NodeTree[T, E]](nodeMaps map[T]E, x []E, rootPid T) []E {
	var root []E
	for _, e := range x {
		pid := e.GetPid()
		if pid == rootPid {
			root = append(root, e)
		} else if parent, exists := nodeMaps[pid]; exists {
			parent.AppendChildren(e)
		}
	}
	return root
}
