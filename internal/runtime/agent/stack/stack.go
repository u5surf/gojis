package stack

// Stack is a LIFO data structure.
type Stack interface {
	// Push adds a new element on top of the stack.
	Push(interface{})
	// Pop removes the topmost element from the stack and returns it.
	Pop() interface{}
	// Peek returns the topmost element of the stack without removing it.
	Peek() interface{}
}
