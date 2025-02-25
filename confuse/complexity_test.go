package confuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComplexity_IsComplexEnough(t *testing.T) {
	testlist := []struct {
		complexity  *Complexity
		truevalues  []string
		falsevalues []string
	}{
		{
			NewComplexity(),
			[]string{"1", "-", "a", "A", "ñ", "日本語"},
			[]string{},
		},
		{
			NewComplexity(WithLower(), WithMeet()),
			[]string{"abc", "abc!"},
			[]string{"ABC", "123", "=!$", ""},
		},
		{
			NewComplexity(WithUpper(), WithMeet()),
			[]string{"ABC"},
			[]string{"abc", "123", "=!$", "abc!", ""},
		},
		{
			NewComplexity(WithDigit(), WithMeet()),
			[]string{"123"},
			[]string{"abc", "ABC", "=!$", "abc!", ""},
		},
		{
			NewComplexity(WithSpec(), WithMeet()),
			[]string{"=!$", "abc!"},
			[]string{"abc", "ABC", "123", ""},
		},
		{
			NewComplexity(WithLower(), WithSpec(), WithMeet()),
			[]string{"abc!"},
			[]string{"abc", "ABC", "123", "=!$", "abcABC123", ""},
		},
		{
			NewComplexity(WithLowerUpperDigit(), WithMeet()),
			[]string{"abcABC123"},
			[]string{"abc", "ABC", "123", "=!$", "abc!", ""},
		},
		{
			NewComplexity(WithAll()),
			[]string{"abcABC123!"},
			[]string{"abc", "ABC", "123", "=!$", "abc!", ""},
		},
	}

	for i, test := range testlist {
		for ii, val := range test.truevalues {
			assert.Truef(t, test.complexity.IsComplexEnough(val), "true, index: %d idx: %d", i, ii)
		}
		for ii, val := range test.falsevalues {
			assert.Falsef(t, test.complexity.IsComplexEnough(val), "false, index: %d idx: %d", i, ii)
		}
	}
}

func TestComplexity_Generate(t *testing.T) {
	const maxCount = 50
	const pwdLen = 50

	test := func(t *testing.T, opt ...Option) {
		c := NewComplexity(opt...)
		for range maxCount {
			pwd := c.Generate(pwdLen)
			assert.Equal(t, pwdLen, len(pwd))
			assert.True(t, c.IsComplexEnough(pwd), "Failed complexities for generated: %s", pwd)
		}
	}

	test(t, WithLower(), WithMeet())
	test(t, WithUpper(), WithMeet())
	test(t, WithLowerUpper(), WithSpec(), WithMeet())
	test(t, WithMeet())
}

func TestDefaultComplexity(t *testing.T) {
	s := Generate(10)
	require.True(t, IsComplexEnough(s))
}
