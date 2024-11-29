package topic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	tests := map[string]string{
		"topic/hello":         "topic/hello",
		"topic//hello":        "topic/hello",
		"topic///hello":       "topic/hello",
		"/topic":              "/topic",
		"//topic":             "/topic",
		"///topic":            "/topic",
		"topic/":              "topic",
		"topic//":             "topic",
		"topic///":            "topic",
		"topic///cool//hello": "topic/cool/hello",
		"topic//cool///hello": "topic/cool/hello",
	}

	for str, result := range tests {
		str, err := Parse(str, true)
		require.Equal(t, result, str)
		require.NoError(t, err, str)
	}
}

func Test_ParseZeroLengthError(t *testing.T) {
	_, err := Parse("", true)
	require.Equal(t, ErrZeroLength, err)

	_, err = Parse("/", true)
	require.Equal(t, ErrZeroLength, err)

	_, err = Parse("//", true)
	require.Equal(t, ErrZeroLength, err)
}

func Test_ParseDisallowWildcards(t *testing.T) {
	tests := map[string]bool{
		"topic":            true,
		"topic/hello":      true,
		"topic/cool/hello": true,
		"+":                false,
		"#":                false,
		"topic/+":          false,
		"topic/#":          false,
	}

	for str, result := range tests {
		_, err := Parse(str, false)

		if result {
			require.NoError(t, err, str)
		} else {
			require.Error(t, err, str)
		}
	}
}

func Test_ParseAllowWildcards(t *testing.T) {
	tests := map[string]bool{
		"topic":            true,
		"topic/hello":      true,
		"topic/cool/hello": true,
		"+":                true,
		"#":                true,
		"topic/+":          true,
		"topic/#":          true,
		"topic/+/hello":    true,
		"topic/cool/+":     true,
		"topic/cool/#":     true,
		"+/cool/#":         true,
		"+/+/#":            true,
		"":                 false,
		"++":               false,
		"##":               false,
		"#/+":              false,
		"#/#":              false,
	}

	for str, result := range tests {
		_, err := Parse(str, true)

		if result {
			require.NoError(t, err, str)
		} else {
			require.Error(t, err, str)
		}
	}
}

func Test_ContainsWildcards(t *testing.T) {
	require.True(t, ContainsWildcards("topic/+"))
	require.True(t, ContainsWildcards("topic/#"))
	require.False(t, ContainsWildcards("topic/hello"))
}

func Benchmark_Parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Parse("foo", true)
		if err != nil {
			panic(err)
		}
	}
}
