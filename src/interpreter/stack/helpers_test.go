package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolToFloat(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		got := boolToFloat(true)

		want := 1.

		assert.Equal(t, got, want)
	})
	t.Run("False", func(t *testing.T) {
		got := boolToFloat(false)

		want := 0.

		assert.Equal(t, got, want)
	})
}

func TestFloatToBool(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		got := FloatToBool(10)

		want := true

		assert.Equal(t, got, want)
	})
	t.Run("negative", func(t *testing.T) {
		got := FloatToBool(-9)

		want := true

		assert.Equal(t, got, want)
	})
	t.Run("null", func(t *testing.T) {
		got := FloatToBool(0)

		want := false

		assert.Equal(t, got, want)
	})
	t.Run("almost null", func(t *testing.T) {
		got := FloatToBool(0.001)

		want := false

		assert.Equal(t, got, want)
	})
}

func TestIsEqual(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		got := isEqual(
			float64(0.1)+float64(0.2),
			float64(0.3),
		)

		want := true

		assert.Equal(t, got, want)
	})
	t.Run("Diff", func(t *testing.T) {
		got := isEqual(
			float64(0.2)+float64(0.2),
			float64(0.3),
		)

		want := false

		assert.Equal(t, got, want)
	})
}

func TestFactorial(t *testing.T) {
	t.Run("-10", func(t *testing.T) {
		got := factorial(-10)

		want := 1

		assert.Equal(t, got, want)
	})
	t.Run("0", func(t *testing.T) {
		got := factorial(0)

		want := 1

		assert.Equal(t, got, want)
	})
	t.Run("1", func(t *testing.T) {
		got := factorial(1)

		want := 1

		assert.Equal(t, got, want)
	})
	t.Run("2", func(t *testing.T) {
		got := factorial(2)

		want := 2

		assert.Equal(t, got, want)
	})
	t.Run("3", func(t *testing.T) {
		got := factorial(3)

		want := 6

		assert.Equal(t, got, want)
	})
	t.Run("4", func(t *testing.T) {
		got := factorial(4)

		want := 24

		assert.Equal(t, got, want)
	})
	t.Run("5", func(t *testing.T) {
		got := factorial(5)

		want := 120

		assert.Equal(t, got, want)
	})
}
