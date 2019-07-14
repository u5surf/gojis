package stack

import (
	"unsafe"
)

type unsafeStack struct {
	s internalSliceStack
}

// NewUnsafeStack returns a new stack that uses a slice as underlying data
// structure. An unsafe stack does not store elements, but pointers to the
// elements. This allows for constant 4-byte elements in the slice. Do not use
// it though.
func NewUnsafeStack() Stack {
	s := new(unsafeStack)
	s.s = internalSliceStack{}
	return s
}

func (s *unsafeStack) Push(v interface{}) {
	ptr := unsafe.Pointer(&v) // #nosec
	s.s = s.s.push(ptr)
}

func (s *unsafeStack) Pop() (elem interface{}) {
	s.s, elem, _ = s.s.pop()
	if elem == nil {
		return
	}

	elem = *(*interface{})(elem.(unsafe.Pointer))
	return
}

func (s *unsafeStack) Peek() (elem interface{}) {
	elem, _ = s.s.peek()
	if elem == nil {
		return
	}

	elem = *(*interface{})(elem.(unsafe.Pointer))
	return
}
