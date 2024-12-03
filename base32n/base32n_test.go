package base32n

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_ValidString(t *testing.T) {
	want := time.Now().UnixMicro()
	s := StdEncoding.Encode(want)

	got, err := StdEncoding.Decode(s)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func Test_IllegalString(t *testing.T) {
	_, err := StdEncoding.Decode("bril03vgrwtu")
	require.Error(t, err)
}
