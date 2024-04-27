package confuse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntR(t *testing.T) {
	require.Panics(t, func() {
		IntR(5, 1)
	})
	got := IntR(2, 20)
	require.LessOrEqual(t, 2, got)
	require.LessOrEqual(t, got, 20)
	require.Equal(t, 10, IntR(10, 10))
}

func TestInt32R(t *testing.T) {
	require.Panics(t, func() {
		Int32R(5, 1)
	})
	got := Int32R(2, 20)
	require.LessOrEqual(t, int32(2), got)
	require.LessOrEqual(t, got, int32(20))
	require.Equal(t, int32(10), Int32R(10, 10))
}

func TestInt64R(t *testing.T) {
	require.Panics(t, func() {
		Int64R(5, 1)
	})
	got := Int64R(2, 20)
	require.LessOrEqual(t, int64(2), got)
	require.LessOrEqual(t, got, int64(20))
	require.Equal(t, int64(10), Int64R(10, 10))
}

func TestFloat64R(t *testing.T) {
	require.Panics(t, func() {
		Float64R(5.1, 1.1)
	})
	got := Float64R(2.1, 20.1)
	require.LessOrEqual(t, 2.1, got)
	require.LessOrEqual(t, got, 20.1)
	require.Equal(t, 10.1, Float64R(10.1, 10.1))
}

func TestUintR(t *testing.T) {
	require.Panics(t, func() {
		UintR(5, 1)
	})
	got := UintR(2, 20)
	require.LessOrEqual(t, uint(2), got)
	require.LessOrEqual(t, got, uint(20))
	require.Equal(t, uint(10), UintR(10, 10))
}

func TestUint32R(t *testing.T) {
	require.Panics(t, func() {
		Uint32R(5, 1)
	})
	got := Uint32R(2, 20)
	require.LessOrEqual(t, uint32(2), got)
	require.LessOrEqual(t, got, uint32(20))
	require.Equal(t, uint32(10), Uint32R(10, 10))
}

func TestUint64R(t *testing.T) {
	require.Panics(t, func() {
		Uint64R(5, 1)
	})
	got := Uint64R(2, 20)
	require.LessOrEqual(t, uint64(2), got)
	require.LessOrEqual(t, got, uint64(20))
	require.Equal(t, uint64(10), Uint64R(10, 10))
}

func TestNR(t *testing.T) {
	require.Panics(t, func() {
		NR(5, 1)
	})
	got := NR(2, 20)
	require.LessOrEqual(t, 2, got)
	require.LessOrEqual(t, got, 20)
	require.Equal(t, 10, NR(10, 10))
}
