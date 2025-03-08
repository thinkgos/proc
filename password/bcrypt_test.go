package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BCrypt(t *testing.T) {
	org := "hahaha"
	cpt := NewBcrypt(0)

	dst, err := cpt.GenerateFromPassword(org)
	require.Nil(t, err)
	require.Nil(t, cpt.CompareHashAndPassword(dst, org))
}

func Benchmark_BCrypt_GenerateFromPassword(b *testing.B) {
	cpt := NewBcrypt(0)

	for i := 0; i < b.N; i++ {
		_, _ = cpt.GenerateFromPassword("hahaha")
	}
}

func Benchmark_BCrypt_CompareHashAndPassword(b *testing.B) {
	org := "hahaha"
	cpt := NewBcrypt(0)
	dst, _ := cpt.GenerateFromPassword(org)

	for i := 0; i < b.N; i++ {
		_ = cpt.CompareHashAndPassword(dst, org)
	}
}
