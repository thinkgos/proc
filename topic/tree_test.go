package topic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TreeAdd(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)

	assert.Equal(t, 1, tree.root.children["foo"].children["bar"].values[0])
}

func Test_TreeAddDuplicate(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)
	tree.Add("foo/bar", 1)

	assert.Equal(t, 1, len(tree.root.children["foo"].children["bar"].values))
}

func Test_TreeSet(t *testing.T) {
	tree := NewStandardTree()

	tree.Set("foo/bar", 1)

	assert.Equal(t, 1, tree.root.children["foo"].children["bar"].values[0])
}

func Test_TreeSetReplace(t *testing.T) {
	tree := NewStandardTree()

	tree.Set("foo/bar", 1)
	tree.Set("foo/bar", 2)

	assert.Equal(t, 2, tree.root.children["foo"].children["bar"].values[0])
}

func Test_TreeGet(t *testing.T) {
	tree := NewStandardTree()

	tree.Set("foo/#", 1)

	assert.Equal(t, 1, tree.Get("foo/#")[0])
}

func Test_TreeRemove(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)
	tree.Remove("foo/bar", 1)

	assert.Equal(t, 0, len(tree.root.children))
}

func Test_TreeRemoveMissing(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)
	tree.Remove("bar/baz", 1)

	assert.Equal(t, 1, len(tree.root.children))
}

func Test_TreeEmpty(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)
	tree.Add("foo/bar", 2)
	tree.Empty("foo/bar")

	assert.Equal(t, 0, len(tree.root.children))
}

func Test_TreeClear(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)
	tree.Add("foo/bar/baz", 1)
	tree.Clear(1)

	assert.Equal(t, 0, len(tree.root.children))
}

func Test_TreeMatchExact(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)

	assert.Equal(t, 1, tree.Match("foo/bar")[0])
}

func Test_TreeMatchWildcard1(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/+", 1)

	assert.Equal(t, 1, tree.Match("foo/bar")[0])
}

func Test_TreeMatchWildcard2(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/#", 1)

	assert.Equal(t, 1, tree.Match("foo/bar")[0])
}

func Test_TreeMatchWildcard3(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/#", 1)

	assert.Equal(t, 1, tree.Match("foo/bar/baz")[0])
}

func Test_TreeMatchWildcard4(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar/#", 1)

	assert.Equal(t, 1, tree.Match("foo/bar")[0])
}

func Test_TreeMatchWildcard5(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/#", 1)

	assert.Equal(t, 1, tree.Match("foo/bar/#")[0])
}

func Test_TreeMatchMultiple(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)
	tree.Add("foo/+", 2)
	tree.Add("foo/#", 3)

	assert.Equal(t, 3, len(tree.Match("foo/bar")))
}

func Test_TreeMatchNoDuplicates(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)
	tree.Add("foo/+", 1)
	tree.Add("foo/#", 1)

	assert.Equal(t, 1, len(tree.Match("foo/bar")))
}

func Test_TreeMatchFirst(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/+", 1)

	assert.Equal(t, 1, tree.MatchFirst("foo/bar"))
}

func Test_TreeMatchFirstNone(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/+", 1)

	assert.Nil(t, tree.MatchFirst("baz/qux"))
}

func Test_TreeSearchExact(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)

	assert.Equal(t, 1, tree.Search("foo/bar")[0])
}

func Test_TreeSearchWildcard1(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)

	assert.Equal(t, 1, tree.Search("foo/+")[0])
}

func Test_TreeSearchWildcard2(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)

	assert.Equal(t, 1, tree.Search("foo/#")[0])
}

func Test_TreeSearchWildcard3(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar/baz", 1)

	assert.Equal(t, 1, tree.Search("foo/#")[0])
}

func Test_TreeSearchWildcard4(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)

	assert.Equal(t, 1, tree.Search("foo/bar/#")[0])
}

func Test_TreeSearchWildcard5(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar/#", 1)

	assert.Equal(t, 1, tree.Search("foo/#")[0])
}

func Test_TreeSearchMultiple(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo", 1)
	tree.Add("foo/bar", 2)
	tree.Add("foo/bar/baz", 3)

	assert.Equal(t, 3, len(tree.Search("foo/#")))
}

func Test_TreeSearchNoDuplicates(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo", 1)
	tree.Add("foo/bar", 1)
	tree.Add("foo/bar/baz", 1)

	assert.Equal(t, 1, len(tree.Search("foo/#")))
}

func Test_TreeSearchFirst(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)

	assert.Equal(t, 1, tree.SearchFirst("foo/+"))
}

func Test_TreeSearchFirstNone(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)

	assert.Nil(t, tree.SearchFirst("baz/qux"))
}

func Test_TreeCount(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo", 1)
	tree.Add("foo/bar", 2)
	tree.Add("foo/bar/baz", 3)
	tree.Add("foo/bar/baz", 4)
	tree.Add("quz/bar/baz", 4)

	assert.Equal(t, 5, tree.Count())
}

func Test_TreeAll(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo", 1)
	tree.Add("foo/bar", 2)
	tree.Add("foo/bar/baz", 3)

	assert.Equal(t, 3, len(tree.All()))
}

func Test_TreeAllNoDuplicates(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo", 1)
	tree.Add("foo/bar", 1)
	tree.Add("foo/bar/baz", 1)

	assert.Equal(t, 1, len(tree.All()))
}

func Test_TreeReset(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("foo/bar", 1)
	tree.Reset()

	assert.Equal(t, 0, len(tree.root.children))
}

func Test_TreeString(t *testing.T) {
	tree := NewStandardTree()

	tree.Add("", 1)
	tree.Add("/foo", 7)
	tree.Add("/foo/bar", 42)

	assert.Equal(t, "topic.Tree:\n| '' => 1\n|   'foo' => 1\n|     'bar' => 1", tree.String())
}

func Benchmark_TreeAddSame(b *testing.B) {
	tree := NewStandardTree()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Add("foo/bar", 1)
	}
}

func Benchmark_TreeAddUnique(b *testing.B) {
	tree := NewStandardTree()

	strings := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {
		strings = append(strings, fmt.Sprintf("foo/%d", i))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Add(strings[i], 1)
	}
}

func Benchmark_TreeSetSame(b *testing.B) {
	tree := NewStandardTree()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Set("foo/bar", 1)
	}
}

func Benchmark_TreeSetUnique(b *testing.B) {
	tree := NewStandardTree()

	strings := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {
		strings = append(strings, fmt.Sprintf("foo/%d", i))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Set(strings[i], 1)
	}
}

func Benchmark_TreeMatchExact(b *testing.B) {
	tree := NewStandardTree()
	tree.Add("foo/bar", 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Match("foo/bar")
	}
}

func Benchmark_TreeMatchWildcardOne(b *testing.B) {
	tree := NewStandardTree()
	tree.Add("foo/+", 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Match("foo/bar")
	}
}

func Benchmark_TreeMatchWildcardSome(b *testing.B) {
	tree := NewStandardTree()
	tree.Add("#", 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Match("foo/bar")
	}
}

func Benchmark_TreeSearchExact(b *testing.B) {
	tree := NewStandardTree()
	tree.Add("foo/bar", 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Search("foo/bar")
	}
}

func Benchmark_TreeSearchWildcardOne(b *testing.B) {
	tree := NewStandardTree()
	tree.Add("foo/bar", 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Search("foo/+")
	}
}

func Benchmark_TreeSearchWildcardSome(b *testing.B) {
	tree := NewStandardTree()
	tree.Add("foo/bar", 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Search("#")
	}
}
