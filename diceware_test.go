package diceware

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLarge(t *testing.T) {
	w, err := Pick(LargeList)
	require.NoError(t, err)
	require.NotEmpty(t, w)
}

func TestShortList(t *testing.T) {
	w, err := Pick(ShortList)
	require.NoError(t, err)
	require.NotEmpty(t, w)
}

func TestShortList2(t *testing.T) {
	w, err := Pick(ShortList2)
	require.NoError(t, err)
	require.NotEmpty(t, w)
}

func BenchmarkDiceware(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pick(LargeList)
	}
}
