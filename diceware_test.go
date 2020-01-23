package diceware

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPassphrase(t *testing.T) {
	phrase, err := Passphrase(4, ShortList2)
	require.NoError(t, err)
	require.NotEmpty(t, phrase)
	chunks := strings.Split(phrase, " ")
	assert.Equal(t, 4, len(chunks))
}

func BenchmarkPickLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pick(LargeList)
	}
}
