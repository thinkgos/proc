package ring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Ring(t *testing.T) {
	r := New[int](5)
	_, ok := r.PeekLatest()
	require.False(t, ok)

	// less
	for i := range 3 {
		r.Push(i)
	}
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 3, r.Len())
	require.False(t, r.IsFull())
	require.Equal(t, []int{0, 1, 2}, r.CollectValues())
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
	latest, ok = r.PeekLatest()
	require.True(t, ok)
	require.Equal(t, 4, latest)

	// over
	for i := range 2 {
		r.Push(i + 5)
	}
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 5, r.Len())
	require.True(t, r.IsFull())
	require.Equal(t, []int{2, 3, 4, 5, 6}, r.CollectValues())
	latest, ok = r.PeekLatest()
	require.True(t, ok)
	require.Equal(t, 6, latest)
}
