package tree

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

var _ Node[int, *DeptTree] = (*Dept)(nil)
var _ NodeTree[int, *DeptTree] = (*DeptTree)(nil)

type Dept struct {
	Id   int
	Pid  int
	Name string
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
	Children []*DeptTree
}

// GetId implements NodeTree.
func (d *DeptTree) GetId() int { return d.Id }

// GetPid implements NodeTree.
func (d *DeptTree) GetPid() int { return d.Pid }

// AppendChildren implements NodeTree.
func (d *DeptTree) AppendChildren(v *DeptTree) {
	d.Children = append(d.Children, v)
}

var arr = []*Dept{
	{1, 0, "超然科技"},
	{8, 3, "bb"},
	{6, 2, "研究院"},
	{2, 0, "低速科技"},
	{3, 1, "科研中心"},
	{4, 1, "运营中心"},
	{5, 2, "吃喝院"},
	{7, 3, "aa"},
	{9, 4, "cc"},
	{10, 5, "dd"},
	{11, 6, "ee"},
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
}
