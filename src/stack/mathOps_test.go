package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	t.Run("Swap with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Plus()

		want := NewStack([]float64{1, 5})

		assert.Equal(t, got, want)
	})
	t.Run("Swap with enought elements 2", func(t *testing.T) {
		got := NewStack([]float64{1, -2, 2})

		got.Plus()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Swap without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Plus()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestMinus(t *testing.T) {
	t.Run("Minus with enought elements 1", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Minus()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Minus with enought elements 2", func(t *testing.T) {
		got := NewStack([]float64{1, -2, -2})

		got.Minus()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Minus with enought elements 3", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -2})

		got.Minus()

		want := NewStack([]float64{1, -4})

		assert.Equal(t, got, want)
	})
	t.Run("Minus without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Minus()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestTimes(t *testing.T) {
	t.Run("Times with enought elements 1", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Times()

		want := NewStack([]float64{1, 6})

		assert.Equal(t, got, want)
	})
	t.Run("Times with enought elements 2", func(t *testing.T) {
		got := NewStack([]float64{1, -2, 3})

		got.Times()

		want := NewStack([]float64{1, -6})

		assert.Equal(t, got, want)
	})
	t.Run("Times without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Times()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestDiv(t *testing.T) {
	t.Run("Div with enought elements 1", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 2})

		got.Div()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Div with enought elements 2", func(t *testing.T) {
		got := NewStack([]float64{1, 4, 2})

		got.Div()

		want := NewStack([]float64{1, 0.5})

		assert.Equal(t, got, want)
	})
	t.Run("Div with enought elements 3", func(t *testing.T) {
		got := NewStack([]float64{1, 4, 0})

		got.Div()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Divide by 0", func(t *testing.T) {
		got := NewStack([]float64{0, 1})

		err := got.Div()

		assert.Equal(t, err, ErrDivideBy0)
	})
	t.Run("Dup2 without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Div()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestModulus(t *testing.T) {
	t.Run("Modulus with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Modulus()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Modulus with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 3.10, 5.20})

		got.Modulus()

		want := NewStack([]float64{1, 2})

		assert.Equal(t, got, want)
	})
	t.Run("Modulus with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 0, 5.20})

		err := got.Modulus()

		want := ErrDivideBy0

		assert.Equal(t, err, want)
	})
	t.Run("Over without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Modulus()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestPower(t *testing.T) {
	t.Run("Power with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Power()

		want := NewStack([]float64{1, 9})

		assert.Equal(t, got, want)
	})
	t.Run("Power with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		got.Power()

		want := NewStack([]float64{1, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Power with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 0, 10})

		got.Power()

		want := NewStack([]float64{1, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Power with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 0.5, 9})

		got.Power()

		want := NewStack([]float64{1, 3})

		assert.Equal(t, got, want)
	})
	t.Run("Power without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Power()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}
