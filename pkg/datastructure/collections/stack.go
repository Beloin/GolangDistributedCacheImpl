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
	if len(s.data) == cap(s.data) {
		temp := make([]any, len(s.data)*2+1)
		copy(temp, s.data)
		s.data = temp
	}

	s.data[s.top] = v
	s.top++
}

func (s *Stack) Pop() (v any) {
	if s.top == 0 {
		return nil
	}

	v = s.data[s.top]
	s.top--

	return
}
