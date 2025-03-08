package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Crypt(t *testing.T) {
	org := "hahaha"

	dst, err := GenerateFromPassword(org)
	require.Nil(t, err)
	require.Nil(t, CompareHashAndPassword(dst, org))
}

func Benchmark_Crypt_GenerateFromPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateFromPassword("hahaha")
	}
}

func Benchmark_Crypt_CompareHashAndPassword(b *testing.B) {
	org := "hahaha"
	dst, _ := GenerateFromPassword(org)

	for i := 0; i < b.N; i++ {
		_ = CompareHashAndPassword(dst, org)
	}
}
