package tree

import (
	"cmp"
	"slices"
)

// Node 节点为不包含的children字段的结构体, 需要转换成包含children的结构体U.
// 其中:
// - I: id/pid类型
// - U: 带children的结构体, 受 NodeTree[I, U] 约束.
type Node[I cmp.Ordered, U NodeTree[I, U]] interface {
	MapId() I
	MapTree() U
}

// NodeTree 节点为包含children字段的结构体.
// 其中:
// - I: id/pid的类型
// - T: 实际使用应为NodeTree本身
type NodeTree[I cmp.Ordered, T any] interface {
	GetId() I
	GetPid() I
	AppendChildren(T)
}

// NodeSort 节点排序
type NodeSort[T any] interface {
	SortChildren(cmp func(a, b T) int)
}

// Map 将Node转换为NodeTree
func Map[I cmp.Ordered, E Node[I, U], U NodeTree[I, U]](x []E) []U {
	s := make([]U, 0, len(x))
	for _, v := range x {
		s = append(s, v.MapTree())
	}
	return s
}

// IntoTree 列表转树, 其中, x的元素为Node[T,U]节点(无children字段)
// NOTE: 元素顺序由x本身顺序决定, 可以提前排序, 然后转树(或使用 SortFunc 对树进行排序)
func IntoTree[I cmp.Ordered, E Node[I, U], U NodeTree[I, U]](x []E, rootPid I) []U {
	return IntoTreeFunc(x, rootPid, dummy)
}

// IntoTree2 列表转树, 其中, x的元素为NodeTree[T,U]节点(有children字段)
// NOTE: 元素顺序由x本身顺序决定, 可以提前排序, 然后转树(或使用 SortFunc 对树进行排序)
func IntoTree2[I cmp.Ordered, U NodeTree[I, U]](x []U, rootPid I) []U {
	return IntoTree2Func(x, rootPid, dummy)
}

// IntoTree 列表转树, 其中, x的元素为Node[T,U]节点(无children字段), f 在加入父节点前的回调, 当为根节点时, parent = cur
// NOTE: 元素顺序由x本身顺序决定, 可以提前排序, 然后转树(或使用 SortFunc 对树进行排序)
func IntoTreeFunc[I cmp.Ordered, E Node[I, U], U NodeTree[I, U]](x []E, rootPid I, f func(parent, cur U) U) []U {
	nodeMaps, nodes := intoMapTree(x)
	return intoTree(nodeMaps, nodes, rootPid, f)
}

// IntoTree2 列表转树, 其中, x的元素为NodeTree[T,U]节点(有children字段), f 在加入父节点前的回调, 当为根节点时,  parent = cur
// NOTE: 元素顺序由x本身顺序决定, 可以提前排序, 然后转树(或使用 SortFunc 对树进行排序)
func IntoTree2Func[I cmp.Ordered, U NodeTree[I, U]](x []U, rootPid I, f func(parent, cur U) U) []U {
	nodeMaps := intoMap(x)
	return intoTree(nodeMaps, x, rootPid, f)
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

// I -> U 映射
// E -> U 转换
func intoMapTree[I cmp.Ordered, E Node[I, U], U NodeTree[I, U]](x []E) (map[I]U, []U) {
	nodeMaps := make(map[I]U)
	nodes := make([]U, 0, len(x))
	for _, v := range x {
		e := v.MapTree()
		nodeMaps[v.MapId()] = e
		nodes = append(nodes, e)
	}
	return nodeMaps, nodes
}

// I -> U 映射
func intoMap[I cmp.Ordered, U NodeTree[I, U]](x []U) map[I]U {
	nodeMaps := make(map[I]U)
	for _, e := range x {
		nodeMaps[e.GetId()] = e
	}
	return nodeMaps
}

// 转树
func intoTree[I cmp.Ordered, U NodeTree[I, U]](nodeMaps map[I]U, x []U, rootPid I, f func(parent, cur U) U) []U {
	var root []U

	for _, e := range x {
		pid := e.GetPid()
		if pid == rootPid {
			root = append(root, f(e, e))
		} else if parent, exists := nodeMaps[pid]; exists {
			parent.AppendChildren(f(parent, e))
		}
	}
	return root
}

func dummy[T any](_, cur T) T { return cur }
