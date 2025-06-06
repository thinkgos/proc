package ring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Ring(t *testing.T) {
	r := New[int](5)

	// less
	for i := range 3 {
		r.Push(i)
	}
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 3, r.Len())
	require.False(t, r.IsFull())
	require.Equal(t, []int{0, 1, 2}, r.CollectValues())

	// full
	for i := range 2 {
		r.Push(i + 3)
	}
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 5, r.Len())
	require.True(t, r.IsFull())
	require.Equal(t, []int{0, 1, 2, 3, 4}, r.CollectValues())

	// over
	for i := range 2 {
		r.Push(i + 5)
	}
	require.Equal(t, 5, r.Capacity())
	require.Equal(t, 5, r.Len())
	require.True(t, r.IsFull())
	require.Equal(t, []int{2, 3, 4, 5, 6}, r.CollectValues())
}
