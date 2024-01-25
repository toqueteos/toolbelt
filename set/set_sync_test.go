package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSyncSet(t *testing.T) {
	s1 := NewSync[int]()
	s1.Add(1)
	s1.Add(1)
	s1.Add(1)
	require.Equal(t, 1, s1.Count())
	require.ElementsMatch(t, []int{1}, s1.Items())
	require.Equal(t, true, s1.Has(1))
	require.Equal(t, false, s1.Has(2))
	require.Equal(t, false, s1.Has(3))

	s1.Add(2)
	s1.Add(3)
	s1.Add(1)
	require.Equal(t, 3, s1.Count())
	require.ElementsMatch(t, []int{1, 2, 3}, s1.Items())
	require.Equal(t, true, s1.Has(1))
	require.Equal(t, true, s1.Has(2))
	require.Equal(t, true, s1.Has(3))

	s1.Remove(1)
	require.Equal(t, 2, s1.Count())
	require.ElementsMatch(t, []int{2, 3}, s1.Items())
	require.Equal(t, false, s1.Has(1))
	require.Equal(t, true, s1.Has(2))
	require.Equal(t, true, s1.Has(3))

	s1.Clear()
	require.Equal(t, 0, s1.Count())
	require.ElementsMatch(t, []int{}, s1.Items())
	require.Equal(t, false, s1.Has(1))
	require.Equal(t, false, s1.Has(2))
	require.Equal(t, false, s1.Has(3))

	s2 := NewSyncFrom([]int{1, 1, 1})
	require.Equal(t, 1, s2.Count())
	require.ElementsMatch(t, []int{1}, s2.Items())
}
