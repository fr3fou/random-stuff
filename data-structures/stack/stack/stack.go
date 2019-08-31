package stack

// Stack is a LIFO data structur
type Stack struct {
	array []int
}

// New returns a new instance of a stack
func New() *Stack {
	return &Stack{}
}

// Push puts an element at the top of the stack
func (s *Stack) Push(n int) {
	s.array = append(s.array, n)
}

// Pop returns the top element of the stack and removes it
func (s *Stack) Pop() int {
	ln := len(s.array)
	el := s.array[ln-1]
	s.array = s.array[:ln]
	return el
}

// Peek returns the top most element of the stack without removing it
func (s *Stack) Peek() int {
	return s.array[len(s.array)-1]
}
