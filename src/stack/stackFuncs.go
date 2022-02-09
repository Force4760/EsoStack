package stack

func (s *Stack) Swap() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a, b := s.Pop2NoTest()
	s.Push(a)
	s.Push(b)
	return nil
}

func (s *Stack) Drop() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}
	s.values = s.values[:s.length-1]
	s.length -= 1
	return nil
}

func (s *Stack) Dup() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a := s.PopNoTest()
	s.Push(a)
	s.Push(a)
	return nil
}

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

func (s *Stack) Push(val float64) {
	s.values = append(s.values, val)
	s.length += 1
}

func (s *Stack) Pop() (float64, error) {
	if s.length < 1 {
		return 0, ErrNotEnoughtValuesOnTheStack
	}
	val := s.values[s.length-1]
	s.Drop()
	return val, nil
}

func (s *Stack) Pop2() (float64, float64, error) {
	if s.length < 2 {
		return 0, 0, ErrNotEnoughtValuesOnTheStack
	}
	a := s.PopNoTest()
	b := s.PopNoTest()
	return a, b, nil
}

func (s *Stack) PopNoTest() float64 {
	val := s.values[s.length-1]
	s.Drop()
	return val
}

func (s *Stack) Pop2NoTest() (float64, float64) {
	a := s.PopNoTest()
	b := s.PopNoTest()
	return a, b
}

func (s *Stack) IsTrue() bool {
	if s.length < 1 {
		return false
	}

	a := s.PopNoTest()
	s.Push(a)

	return a != 0
}