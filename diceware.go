package diceware

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

var max = big.NewInt(6)

// Pick picks a word from the given word list
func Pick(words WordList) (string, error) {
	var rs RollSet
	rs.RollDice(words.Dice())
	return words.At(rs), nil
}

// Passphrase generates a string composed of `nbWord` words from the given WordList, separated by a space
func Passphrase(nbWord int, words WordList) (string, error) {
	return PassphraseWithSep(nbWord, words, " ")
}

// PassphraseWithSep generates a string composed of `nbWord` words from the given WordList, separated by the given separator
func PassphraseWithSep(nbWord int, words WordList, sep string) (string, error) {
	var builder strings.Builder
	for i := 0; i < nbWord-1; i++ {
		w, err := Pick(words)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&builder, "%s%s", w, sep)
	}
	w, err := Pick(words)
	if err != nil {
		return "", err
	}
	builder.WriteString(w)
	return builder.String(), nil
}

// RollSet is a set die rolls
// Each roll takes 3 bits
type RollSet uint16

// RollDice rolls the given number of dice
func (r *RollSet) RollDice(nb int) error {
	*r = 0
	for i := 0; i < nb; i++ {
		roll, err := rand.Int(rand.Reader, max)
		if err != nil {
			return err
		}
		*r = *r | (RollSet(roll.Int64()+1) << (i * 3))
	}
	return nil
}

// WordList is a list of word
type WordList interface {
	// Dice returns the number of dice used to pick a word
	Dice() int
	// At returns the word at the given RollSet
	At(rs RollSet) string
}

type wordList map[RollSet]string

type listWith4Dice wordList

func (wl listWith4Dice) At(rs RollSet) string {
	return wl[rs]
}

func (listWith4Dice) Dice() int {
	return 4
}

type listWith5Dice wordList

func (listWith5Dice) Dice() int {
	return 5
}
func (wl listWith5Dice) At(rs RollSet) string {
	return wl[rs]
}
