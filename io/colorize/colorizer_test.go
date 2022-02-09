package colorize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorize(t *testing.T) {
	t.Run("Test RED", func(t *testing.T) {
		got := Colorize("Testing the colorizer", RED)
		want := "\033[0;31mTesting the colorizer\033[0m"

		assert.Equal(t, want, got)
	})
	t.Run("Test BOLD", func(t *testing.T) {
		got := Colorize("Testing the colorizer", BOLD)
		want := "\033[1mTesting the colorizer\033[0m"

		assert.Equal(t, want, got)
	})
}

func TestColorizeWord(t *testing.T) {
	t.Run("Test Word Middle", func(t *testing.T) {

		got := ColorizeWord("Testing the colorizer", "the", RED)
		want := "Testing \033[0;31mthe\033[0m colorizer"
		assert.Equal(t, want, got)
	})
	t.Run("Test Word Start", func(t *testing.T) {

		got := ColorizeWord("Testing the colorizer", "Testing", RED)
		want := "\033[0;31mTesting\033[0m the colorizer"
		assert.Equal(t, want, got)
	})
	t.Run("Test Word End", func(t *testing.T) {

		got := ColorizeWord("Testing the colorizer", "colorizer", RED)
		want := "Testing the \033[0;31mcolorizer\033[0m"
		assert.Equal(t, want, got)
	})
	t.Run("Test Word Whole", func(t *testing.T) {

		got := ColorizeWord("Testing the colorizer", "Testing the colorizer", RED)
		want := "\033[0;31mTesting the colorizer\033[0m"
		assert.Equal(t, want, got)
	})
}
