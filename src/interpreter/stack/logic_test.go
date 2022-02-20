package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	t.Run("T && T", func(t *testing.T) {
		got := NewStack([]float64{1, 5})

		got.And()

		want := NewStack([]float64{1})

		assert.Equal(t, got, want)
	})
	t.Run("T && F", func(t *testing.T) {
		got := NewStack([]float64{1, 0})

		got.And()

		want := NewStack([]float64{0})

		assert.Equal(t, got, want)
	})
	t.Run("F && T", func(t *testing.T) {
		got := NewStack([]float64{0, 1})

		got.And()

		want := NewStack([]float64{0})

		assert.Equal(t, got, want)
	})
	t.Run("F && F", func(t *testing.T) {
		got := NewStack([]float64{0, 0})

		got.And()

		want := NewStack([]float64{0})

		assert.Equal(t, got, want)
	})
	t.Run("Error", func(t *testing.T) {
		got := NewStack([]float64{0})

		err := got.And()

		want := ErrNotEnoughtValuesOnTheStack

		assert.Equal(t, err, want)
	})
}

func TestOr(t *testing.T) {
	t.Run("T || T", func(t *testing.T) {
		got := NewStack([]float64{1, 5})

		got.Or()

		want := NewStack([]float64{1})

		assert.Equal(t, got, want)
	})
	t.Run("T || F", func(t *testing.T) {
		got := NewStack([]float64{1, 0})

		got.Or()

		want := NewStack([]float64{1})

		assert.Equal(t, got, want)
	})
	t.Run("F || T", func(t *testing.T) {
		got := NewStack([]float64{0, 1})

		got.Or()

		want := NewStack([]float64{1})

		assert.Equal(t, got, want)
	})
	t.Run("F || F", func(t *testing.T) {
		got := NewStack([]float64{0, 0})

		got.Or()

		want := NewStack([]float64{0})

		assert.Equal(t, got, want)
	})
	t.Run("Error", func(t *testing.T) {
		got := NewStack([]float64{0})

		err := got.Or()

		want := ErrNotEnoughtValuesOnTheStack

		assert.Equal(t, err, want)
	})
}

func TestXor(t *testing.T) {
	t.Run("T XX T", func(t *testing.T) {
		got := NewStack([]float64{1, 5})

		got.Xor()

		want := NewStack([]float64{0})

		assert.Equal(t, got, want)
	})
	t.Run("T XX F", func(t *testing.T) {
		got := NewStack([]float64{1, 0})

		got.Xor()

		want := NewStack([]float64{1})

		assert.Equal(t, got, want)
	})
	t.Run("F XX T", func(t *testing.T) {
		got := NewStack([]float64{0, 1})

		got.Xor()

		want := NewStack([]float64{1})

		assert.Equal(t, got, want)
	})
	t.Run("F XX F", func(t *testing.T) {
		got := NewStack([]float64{0, 0})

		got.Xor()

		want := NewStack([]float64{0})

		assert.Equal(t, got, want)
	})
	t.Run("Error", func(t *testing.T) {
		got := NewStack([]float64{0})

		err := got.Xor()

		want := ErrNotEnoughtValuesOnTheStack

		assert.Equal(t, err, want)
	})
}

func TestNot(t *testing.T) {
	t.Run("!T", func(t *testing.T) {
		got := NewStack([]float64{5})

		got.Not()

		want := NewStack([]float64{0})

		assert.Equal(t, got, want)
	})
	t.Run("!F", func(t *testing.T) {
		got := NewStack([]float64{0})

		got.Not()

		want := NewStack([]float64{1})

		assert.Equal(t, got, want)
	})
	t.Run("Error", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Not()

		want := ErrNotEnoughtValuesOnTheStack

		assert.Equal(t, err, want)
	})
}
