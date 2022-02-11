package stack

// a == b
//
// Needs 2 elements
func (s *Stack) Equal() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	c := boolToInt(
		isEqual(a, b),
	)
	s.Push(c)

	return nil
}

// a < b
//
// Needs 2 elements
func (s *Stack) Lesser() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	c := boolToInt(a < b)
	s.Push(c)

	return nil
}

// a > b
//
// Needs 2 elements
func (s *Stack) Greater() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	c := boolToInt(a > b)
	s.Push(c)

	return nil
}

// a <= b
//
// Needs 2 elements
func (s *Stack) LesserEq() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	c := boolToInt(a <= b)
	s.Push(c)

	return nil
}

// a >= b
//
// Needs 2 elements
func (s *Stack) GreaterEq() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	c := boolToInt(a >= b)
	s.Push(c)

	return nil
}

// a != b
//
// Needs 2 elements
func (s *Stack) Diferent() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	c := boolToInt(
		!isEqual(a, b),
	)
	s.Push(c)

	return nil
}
