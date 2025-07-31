package ring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ConcurrentRing(t *testing.T) {
	r := NewConcurrentRing[int](5)
	_, ok := r.Peek()
	require.False(t, ok)
	_, ok = r.PeekLatest()
	require.False(t, ok)

	// less
	for i := range 3 {
		r.Push(i)
	}
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 3, r.Len())
	require.False(t, r.IsFull())
	require.Equal(t, []int{0, 1, 2}, r.CollectValues())
	oldest, ok := r.Peek()
	require.True(t, ok)
	require.Equal(t, 0, oldest)
	latest, ok := r.PeekLatest()
	require.True(t, ok)
	require.Equal(t, 2, latest)

	// full
	for i := range 2 {
		r.Push(i + 3)
	}
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 5, r.Len())
	require.True(t, r.IsFull())
	require.Equal(t, []int{0, 1, 2, 3, 4}, r.CollectValues())
	oldest, ok = r.Peek()
	require.True(t, ok)
	require.Equal(t, 0, oldest)
	latest, ok = r.PeekLatest()
	require.True(t, ok)
	require.Equal(t, 4, latest)

	// overflow then overwrite the old element
	for i := range 2 {
		r.Push(i + 5)
	}
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 5, r.Len())
	require.True(t, r.IsFull())
	require.Equal(t, []int{2, 3, 4, 5, 6}, r.CollectValues())
	oldest, ok = r.Peek()
	require.True(t, ok)
	require.Equal(t, 2, oldest)
	latest, ok = r.PeekLatest()
	require.True(t, ok)
	require.Equal(t, 6, latest)

	// Pop one element
	val, ok := r.Pop()
	require.True(t, ok)
	require.Equal(t, 2, val)
	require.Equal(t, 4, r.Len())
	require.False(t, r.IsFull())
	require.Equal(t, []int{3, 4, 5, 6}, r.CollectValues())
	oldest, ok = r.Peek()
	require.True(t, ok)
	require.Equal(t, 3, oldest)
	latest, ok = r.PeekLatest()
	require.True(t, ok)
	require.Equal(t, 6, latest)

	// Push again
	r.Push(7)
	r.Push(8)
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 5, r.Len())
	require.True(t, r.IsFull())
	require.Equal(t, []int{4, 5, 6, 7, 8}, r.CollectValues())
	oldest, ok = r.Peek()
	require.True(t, ok)
	require.Equal(t, 4, oldest)
	latest, ok = r.PeekLatest()
	require.True(t, ok)
	require.Equal(t, 8, latest)

	// pop all
	vals := r.PopWithin(-1)
	require.Equal(t, 5, len(vals))
	require.Equal(t, []int{4, 5, 6, 7, 8}, vals)
	require.False(t, r.IsFull())
	require.Equal(t, []int(nil), r.CollectValues())
	oldest, ok = r.Peek()
	require.False(t, ok)
	latest, ok = r.PeekLatest()
	require.False(t, ok)

	// pop empty ring
	val, ok = r.Pop()
	require.False(t, ok)
	vals = r.PopWithin(10)
	require.Equal(t, 0, len(vals))
	require.Equal(t, []int{}, vals)
	require.False(t, r.IsFull())
	require.Equal(t, []int(nil), r.CollectValues())
	oldest, ok = r.Peek()
	require.False(t, ok)
	latest, ok = r.PeekLatest()
	require.False(t, ok)
}
