package stack

// Swap the top 2 elements of the stack
// a/b/c -> b/a/c
//
// Needs 2 elements
func (s *Stack) Swap() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()
	s.Push(a)
	s.Push(b)

	return nil
}

// Drop the top element of the stack
// a/b/c -> b/c
//
// Needs 1 element
func (s *Stack) Drop() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	s.values = s.values[:s.length-1]
	s.length -= 1

	return nil
}

// Duplicate the top element of the stack
// a/b/c -> a/a/b/c
//
// Needs 1 element
func (s *Stack) Dup() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()
	s.Push(a)
	s.Push(a)

	return nil
}

// Duplicate the top 2 elements of the stack
// a/b/c -> a/b/a/b/c
//
// Needs 2 elements
func (s *Stack) Dup2() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()
	s.Push(b)
	s.Push(a)
	s.Push(b)
	s.Push(a)

	return nil
}

// Copy The 2nd element to the top
// a/b/c -> b/a/b/c
//
// Needs 2 elements
func (s *Stack) Over() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a, b := s.Pop2NoTest()
	s.Push(b)
	s.Push(a)
	s.Push(b)

	return nil
}

// Rotate the top 3 elements of the stack
// a/b/c -> c/a/b
//
// Needs 3 elements
func (s *Stack) Rot() error {
	if s.length < 3 {
		return ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()
	b := s.PopNoTest()
	c := s.PopNoTest()

	s.Push(b)
	s.Push(a)
	s.Push(c)

	return nil
}

// Push an element `val` to the stack
// a/b/c -> val/a/b/c
//
// Needs 0 elements
func (s *Stack) Push(val float64) {
	s.values = append(s.values, val)
	s.length += 1
}

// Remove and return the top element of the stack
// a/b/c -> b/c
//
// Needs 1 elements
func (s *Stack) Pop() (float64, error) {
	if s.length < 1 {
		return 0, ErrNotEnoughtValuesOnTheStack
	}

	val := s.values[s.length-1]
	s.Drop()

	return val, nil
}

// Remove and return the top 2 elements of the stack
// a/b/c -> c
//
// Needs 2 elements
func (s *Stack) Pop2() (float64, float64, error) {
	if s.length < 2 {
		return 0, 0, ErrNotEnoughtValuesOnTheStack
	}

	a := s.PopNoTest()
	b := s.PopNoTest()

	return a, b, nil
}

// Pop, but does not test for enough elements
func (s *Stack) PopNoTest() float64 {
	val := s.values[s.length-1]
	s.Drop()

	return val
}

// Pop2, but does not test for enough elements
func (s *Stack) Pop2NoTest() (float64, float64) {
	a := s.PopNoTest()
	b := s.PopNoTest()

	return a, b
}

// Predicate for checking if the top element of the stack is different from 0
// Does not remove this value from the stack
//
// Needs 1 element
func (s *Stack) IsTrue() bool {
	if s.length < 1 {
		return false
	}

	a := s.PopNoTest()
	s.Push(a)

	return int(a) != 0
}
