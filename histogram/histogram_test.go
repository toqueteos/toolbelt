package histogram

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHistogram(t *testing.T) {
	s := New[string]()
	s.Add("a")
	s.Add("a")
	s.Add("b")

	require.Equal(t, 2, s.Count())
	require.Equal(t, map[string]int{"a": 2, "b": 1}, s.Items())

	s.Add("b")
	s.Add("c")
	s.Add("a")

	require.Equal(t, 3, s.Count())
	require.Equal(t, map[string]int{"a": 3, "b": 2, "c": 1}, s.Items())
}
