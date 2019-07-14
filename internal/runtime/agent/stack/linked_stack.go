package stack

type linkedStack struct {
	top    *node
	length int
}

type node struct {
	value interface{}
	prev  *node // TODO: use a sync.Pool for the nodes?
}

// NewLinkedStack returns a new stack, using a single linked list as underlying
// data structure.
func NewLinkedStack() Stack {
	return &linkedStack{nil, 0}
}

func (s *linkedStack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *linkedStack) Pop() (elem interface{}) {
	if s.length == 0 {
		return
	}

	elem = s.top.value
	s.top = s.top.prev
	s.length--
	return
}

func (s *linkedStack) Push(v interface{}) {
	s.top = &node{v, s.top}
	s.length++
}
