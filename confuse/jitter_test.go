package confuse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Jitter(t *testing.T) {
	val := Jitter2(100, 0.5)
	require.GreaterOrEqual(t, val, 50)
	require.LessOrEqual(t, val, 150)
	t.Log("jitter val:", val)
}
