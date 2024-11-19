package tree

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type Dept struct {
	Id   int
	Pid  int
	Name string
}

func (d *Dept) MapId() int { return d.Id }
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

func (d *DeptTree) GetId() int  { return d.Id }
func (d *DeptTree) GetPid() int { return d.Pid }

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
	vv := IntoTree(arr, 0)
	v, err := json.MarshalIndent(vv, " ", "  ")
	require.NoError(t, err)
	t.Log(string(v))
}
