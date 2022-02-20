package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	t.Run("Equal equal", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 2})

		got.Equal()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Equal diferent", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Equal()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Equal without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Equal()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestDiferent(t *testing.T) {
	t.Run("Different equal", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 2})

		got.Diferent()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Different different", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Diferent()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Different without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Diferent()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestLesser(t *testing.T) {
	t.Run("Lesser equal", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 2})

		got.Lesser()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Lesser lesser", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 1})

		got.Lesser()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Lesser greater", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 5})

		got.Lesser()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Lesser without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Lesser()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestGreater(t *testing.T) {
	t.Run("Greater equal", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 2})

		got.Greater()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Greater lesser", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 1})

		got.Greater()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Greater greater", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 5})

		got.Greater()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Greater without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Greater()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestLesserEq(t *testing.T) {
	t.Run("LesserEq equal", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 2})

		got.LesserEq()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("LesserEq lesser", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 1})

		got.LesserEq()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("LesserEq greater", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 5})

		got.LesserEq()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("LesserEq without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.LesserEq()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}
func TestGreaterEq(t *testing.T) {
	t.Run("GreaterEq equal", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 2})

		got.GreaterEq()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("GreaterEq lesser", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 1})

		got.GreaterEq()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("GreaterEq greater", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 5})

		got.GreaterEq()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("GreaterEq without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.GreaterEq()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}
