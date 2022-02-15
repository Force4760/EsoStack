package stack

// a && b -> Logical And
// T && T = T    |    T && F = F
// F && T = F    |    F && F = F
//
// Needs 2 element
func (s *Stack) And() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	aNb := boolToFloat(
		FloatToBool(a) && FloatToBool(b),
	)

	s.Push(aNb)
	return nil
}

// a || b -> Logical Or
// T || T = T    |    T || F = T
// F || T = T    |    F || F = F
//
// Needs 2 element
func (s *Stack) Or() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	aNb := boolToFloat(
		FloatToBool(a) || FloatToBool(b),
	)

	s.Push(aNb)
	return nil
}

// a XX b -> Logical Exclusive Or (XOR)
// T XX T = F    |    T XX F = T
// F XX T = T    |    F XX F = F
//
// Needs 2 element
func (s *Stack) Xor() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()

	// Go does not have an XOR operator
	// But != acts in the same way
	aNb := boolToFloat(
		FloatToBool(a) != FloatToBool(b),
	)

	s.Push(aNb)
	return nil
}

// !a -> Logical Negation (NOT)
// !T = F        |    !F = T
//
// Needs 1 element
func (s *Stack) Not() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()

	aNb := boolToFloat(!FloatToBool(a))

	s.Push(aNb)
	return nil
}
