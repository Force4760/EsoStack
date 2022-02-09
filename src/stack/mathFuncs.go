package stack

import (
	"math"
)

func (s *Stack) Abs() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a := s.PopNoTest()
	s.Push(math.Abs(a))
	return nil
}

func (s *Stack) Sqrt() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	if a < 0 {
		return ErrNotValid
	}

	s.Push(math.Sqrt(a))
	return nil
}

func (s *Stack) Fact() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := int(s.PopNoTest())

	s.Push(float64(factorial(a)))
	return nil
}

func factorial(num int) int {
	var result int = 1
	if num < 0 {
	} else {
		for i := 1; i <= num; i++ {
			result *= i
		}
	}
	return result
}

func (s *Stack) Floor() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	s.Push(math.Floor(a))
	return nil
}

func (s *Stack) Ceil() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	s.Push(math.Ceil(a))
	return nil
}

func (s *Stack) Exp() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	s.Push(math.Exp(a))
	return nil
}

func (s *Stack) Log() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	if a < 0 {
		return ErrNotValid
	}

	s.Push(math.Log(a))
	return nil
}

func (s *Stack) Max() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	s.Push(math.Max(a, b))
	return nil
}

func (s *Stack) Min() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	s.Push(math.Min(a, b))
	return nil
}

func (s *Stack) Simetric() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	s.Push(-a)
	return nil
}

// trignometric functions

func (s *Stack) Tan() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	s.Push(math.Tan(a))
	return nil
}

func (s *Stack) Sin() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	s.Push(math.Sin(a))
	return nil
}

func (s *Stack) Cos() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	s.Push(math.Cos(a))
	return nil
}
