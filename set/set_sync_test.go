package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSyncSet(t *testing.T) {
	s := NewSync[int]()
	s.Add(1)
	s.Add(1)
	s.Add(1)

	require.Equal(t, 1, s.Count())
	require.ElementsMatch(t, []int{1}, s.Items())

	s.Add(2)
	s.Add(3)
	s.Add(1)

	require.Equal(t, 3, s.Count())
	require.ElementsMatch(t, []int{1, 2, 3}, s.Items())
}