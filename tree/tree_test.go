package tree

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

var _ Node[int, *DeptTree] = (*Dept)(nil)
var _ NodeTree[int, *DeptTree] = (*DeptTree)(nil)
var _ NodeSort[*DeptTree] = (*DeptTree)(nil)

type Dept struct {
	Id   int    `json:"id"`
	Pid  int    `json:"pid"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
}

// MapId implements Node.
func (d *Dept) MapId() int { return d.Id }

// MapTree implements Node.
func (d *Dept) MapTree() *DeptTree {
	return &DeptTree{
		Dept:     d,
		Children: nil,
	}
}

type DeptTree struct {
	*Dept
	Children []*DeptTree `json:"children"`
}

// GetId implements NodeTree.
func (d *DeptTree) GetId() int { return d.Id }

// GetPid implements NodeTree.
func (d *DeptTree) GetPid() int { return d.Pid }

// AppendChildren implements NodeTree.
func (d *DeptTree) AppendChildren(v *DeptTree) {
	d.Children = append(d.Children, v)
}

// SortChildren implements NodeSort.
func (d *DeptTree) SortChildren(cmp func(a, b *DeptTree) int) {
	SortFunc(d.Children, cmp)
}

func CmpDept(a, b *DeptTree) int {
	return a.Sort - b.Sort
}

var arr = []*Dept{
	{Id: 1, Pid: 0, Name: "超然科技", Sort: 1},
	{Id: 8, Pid: 3, Name: "bb", Sort: 2},
	{Id: 6, Pid: 2, Name: "研究院", Sort: 4},
	{Id: 2, Pid: 0, Name: "低速科技", Sort: 5},
	{Id: 3, Pid: 1, Name: "科研中心", Sort: 4},
	{Id: 4, Pid: 1, Name: "运营中心", Sort: 2},
	{Id: 5, Pid: 2, Name: "吃喝院", Sort: 3},
	{Id: 7, Pid: 3, Name: "aa", Sort: 4},
	{Id: 9, Pid: 4, Name: "cc", Sort: 2},
	{Id: 10, Pid: 5, Name: "dd", Sort: 3},
	{Id: 11, Pid: 6, Name: "ee", Sort: 4},
}

func TestTree(t *testing.T) {
	gotTree1 := IntoTree(arr, 0)
	tree1, err := json.MarshalIndent(gotTree1, " ", "  ")
	require.NoError(t, err)

	arrTree := Map(arr)
	gotTree2 := IntoTree2(arrTree, 0)
	tree2, err := json.MarshalIndent(gotTree2, " ", "  ")
	require.NoError(t, err)

	require.Equal(t, tree2, tree1)

	t.Log(string(tree1))
	SortFunc(gotTree1, CmpDept)
	tree11, err := json.MarshalIndent(gotTree1, " ", "  ")
	require.NoError(t, err)
	t.Log(string(tree11))
}
