package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	t.Run("Swap with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Swap()

		want := NewStack([]float64{1, 3, 2})

		assert.Equal(t, got, want)
	})
	t.Run("Swap without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Swap()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestDrop(t *testing.T) {
	t.Run("Drop with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Drop()

		want := NewStack([]float64{1, 2})

		assert.Equal(t, got, want)
	})
	t.Run("Drop without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Drop()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestDup(t *testing.T) {
	t.Run("Dup with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Dup()

		want := NewStack([]float64{1, 2, 3, 3})

		assert.Equal(t, got, want)
	})
	t.Run("Dup without enought elements", func(t *testing.T) {
		got := NewStack([]float64{})

		err := got.Dup()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestDup2(t *testing.T) {
	t.Run("Dup2 with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Dup2()

		want := NewStack([]float64{1, 2, 3, 2, 3})

		assert.Equal(t, got, want)
	})
	t.Run("Dup2 without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Dup2()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestOver(t *testing.T) {
	t.Run("Over with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Over()

		want := NewStack([]float64{1, 2, 3, 2})

		assert.Equal(t, got, want)
	})
	t.Run("Over without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Over()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestRot(t *testing.T) {
	t.Run("Rot with enought elements", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Rot()

		want := NewStack([]float64{2, 3, 1})

		assert.Equal(t, got, want)
	})
	t.Run("Rot without enought elements", func(t *testing.T) {
		got := NewStack([]float64{1})

		err := got.Rot()

		assert.Equal(t, err, ErrNotEnoughtValuesOnTheStack)
	})
}

func TestPush(t *testing.T) {
	t.Run("Push non-empty", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		got.Push(10)

		want := NewStack([]float64{1, 2, 3, 10})

		assert.Equal(t, got, want)
	})
}

func TestPop(t *testing.T) {
	t.Run("Pop non-empty", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		_ = got.PopNoTest()

		want := NewStack([]float64{1, 2})

		assert.Equal(t, got, want)
	})
	t.Run("Pop non-empty 2", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		a := got.PopNoTest()

		want := 3.0

		assert.Equal(t, a, want)

	})
}

func TestPop2(t *testing.T) {
	t.Run("Pop2 non-empty", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		_, _ = got.Pop2NoTest()

		want := NewStack([]float64{1})

		assert.Equal(t, got, want)
	})
	t.Run("Pop2 non-empty 2", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		a, _ := got.Pop2NoTest()

		want := 3.0

		assert.Equal(t, a, want)

	})
	t.Run("Pop2 non-empty 3", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		_, b := got.Pop2NoTest()

		want := 2.0

		assert.Equal(t, b, want)

	})
}

func TestIsTrue(t *testing.T) {
	t.Run("IsTrue non-empty 0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		a, _ := got.IsTrueNotRemove()

		want := false

		assert.Equal(t, a, want)
	})
	t.Run("IsTrue non-empty non-0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		a, _ := got.IsTrueNotRemove()

		want := true

		assert.Equal(t, a, want)

	})
	t.Run("IsTrue empty", func(t *testing.T) {
		got := NewStack([]float64{})

		_, a := got.IsTrueNotRemove()

		want := ErrNotEnoughtValuesOnTheStack

		assert.Equal(t, a, want)

	})

	t.Run("IsTrue non-empty 0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 0})

		a, _ := got.IsTrue()

		want := false

		assert.Equal(t, a, want)
	})
	t.Run("IsTrue non-empty non-0", func(t *testing.T) {
		got := NewStack([]float64{1, 2, 3})

		a, _ := got.IsTrue()

		want := true

		assert.Equal(t, a, want)

	})
	t.Run("IsTrue empty", func(t *testing.T) {
		got := NewStack([]float64{})

		_, a := got.IsTrue()

		want := ErrNotEnoughtValuesOnTheStack

		assert.Equal(t, a, want)

	})
}
