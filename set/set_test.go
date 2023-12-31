package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(1)
	s.Add(1)

	require.Equal(t, 1, s.Count())
	require.ElementsMatch(t, []int{1}, s.Items())
	require.Equal(t, true, s.Has(1))
	require.Equal(t, false, s.Has(2))
	require.Equal(t, false, s.Has(3))

	s.Add(2)
	s.Add(3)
	s.Add(1)

	require.Equal(t, 3, s.Count())
	require.ElementsMatch(t, []int{1, 2, 3}, s.Items())
	require.Equal(t, true, s.Has(1))
	require.Equal(t, true, s.Has(2))
	require.Equal(t, true, s.Has(3))

	s.Remove(1)

	require.Equal(t, 2, s.Count())
	require.ElementsMatch(t, []int{2, 3}, s.Items())
	require.Equal(t, false, s.Has(1))
	require.Equal(t, true, s.Has(2))
	require.Equal(t, true, s.Has(3))

	s.Clear()

	require.Equal(t, 0, s.Count())
	require.ElementsMatch(t, []int{}, s.Items())
	require.Equal(t, false, s.Has(1))
	require.Equal(t, false, s.Has(2))
	require.Equal(t, false, s.Has(3))
}
