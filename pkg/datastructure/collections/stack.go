// Package collections
package collections

type Stack struct {
	top  uint
	data []any
}

func NewStack(c int) Stack {
	st := Stack{0, make([]any, c)}
	return st
}

func (s *Stack) Push(v any) {
	if int(s.top) == len(s.data) {
		temp := make([]any, len(s.data)*2+1)
		copy(temp, s.data)
		s.data = temp
	}

	s.data[s.top] = v
	s.top++
}

func (s *Stack) Pop() (v any) {
	v = nil
	if s.top == 0 {
		return
	}

	s.top--
	v = s.data[s.top]

	return
}

func (s *Stack) Empty() bool {
	return s.top == 0
}
