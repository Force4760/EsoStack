package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	t.Run("Abs with positive", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Abs()

		want := NewStack([]float64{1, 2, 3})

		assert.Equal(t, got, want)
	})
	t.Run("Abs with negative", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -3})

		got.Abs()

		want := NewStack([]float64{1, 2, 3})

		assert.Equal(t, got, want)
	})
	t.Run("Abs with 0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		got.Abs()

		want := NewStack([]float64{1, 2, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Abs without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Abs()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestFact(t *testing.T) {
	t.Run("Fact with positive", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 5})

		got.Fact()

		want := NewStack([]float64{1, 2, 120})

		assert.Equal(t, got, want)
	})
	t.Run("Fact with positive 2", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 4})

		got.Fact()

		want := NewStack([]float64{1, 2, 24})

		assert.Equal(t, got, want)
	})
	t.Run("Fact with positive 3", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Fact()

		want := NewStack([]float64{1, 2, 6})

		assert.Equal(t, got, want)
	})
	t.Run("Fact with negative", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -3})

		got.Fact()

		want := NewStack([]float64{1, 2, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Fact with 0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		got.Fact()

		want := NewStack([]float64{1, 2, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Fact without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Fact()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestSqrt(t *testing.T) {
	t.Run("Sqrt with positive", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 9})

		got.Sqrt()

		want := NewStack([]float64{1, 2, 3})

		assert.Equal(t, got, want)
	})
	t.Run("Sqrt with negative", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -3})

		err := got.Sqrt()

		want := ErrNotValid

		assert.Equal(t, err, want)
	})
	t.Run("Sqrt with 0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		got.Sqrt()

		want := NewStack([]float64{1, 2, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Sqrt without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Sqrt()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestFloor(t *testing.T) {
	t.Run("Floor with positive", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 9.5})

		got.Floor()

		want := NewStack([]float64{1, 2, 9})

		assert.Equal(t, got, want)
	})
	t.Run("Floor with negative", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -3.2})

		got.Floor()

		want := NewStack([]float64{1, 2, -4})

		assert.Equal(t, got, want)
	})
	t.Run("Floor with 0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		got.Floor()

		want := NewStack([]float64{1, 2, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Floor without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Floor()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestCeil(t *testing.T) {
	t.Run("Ceil with positive", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 9.5})

		got.Ceil()

		want := NewStack([]float64{1, 2, 10})

		assert.Equal(t, got, want)
	})
	t.Run("Ceil with negative", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -3.2})

		got.Ceil()

		want := NewStack([]float64{1, 2, -3})

		assert.Equal(t, got, want)
	})
	t.Run("Ceil with 0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		got.Ceil()

		want := NewStack([]float64{1, 2, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Ceil without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Ceil()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestMax(t *testing.T) {
	t.Run("Max", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -1})

		got.Max()

		want := NewStack([]float64{1, 2})

		assert.Equal(t, got, want)
	})
	t.Run("Max without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Max()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestMin(t *testing.T) {
	t.Run("Min", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -1})

		got.Min()

		want := NewStack([]float64{1, -1})

		assert.Equal(t, got, want)
	})
	t.Run("Min without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Min()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestSimetric(t *testing.T) {
	t.Run("Simetric with 0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		got.Simetric()

		want := NewStack([]float64{1, 2, 0})

		assert.Equal(t, got, want)
	})
	t.Run("Simetric with negative", func(t *testing.T) {
		got := NewStack([]float64{1, 2, -10})

		got.Simetric()

		want := NewStack([]float64{1, 2, 10})

		assert.Equal(t, got, want)
	})
	t.Run("Simetric with positive", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 10})

		got.Simetric()

		want := NewStack([]float64{1, 2, -10})

		assert.Equal(t, got, want)
	})
	t.Run("Simetric without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Simetric()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}
