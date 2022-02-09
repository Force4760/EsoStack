package stack

import (
	"fmt"
	"strconv"
)

func (s *Stack) Input() {
	value := input()

	val, err := strconv.Atoi(value)
	if err != nil {
		val = len(value)
	}

	s.Push(float64(val))
}

func (s *Stack) PrintNum() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}
	val := s.PopNoTest()
	fmt.Println(val)

	return nil
}

func (s *Stack) Print() error {
	if s.length < 1 {
		return ErrNotEnoughtValuesOnTheStack
	}
	val := s.PopNoTest()
	fmt.Println(rune(int(val)))

	return nil
}
