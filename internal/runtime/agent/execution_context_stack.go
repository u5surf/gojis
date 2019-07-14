package agent

import "github.com/gojisvm/gojis/internal/runtime/agent/stack"

// ExecutionContextStack is a stack that is used to keep track of the currently
// executing and soon to execute ExecutionContexts.
// ExecutionContextStack is mentioned in 8.3.
type ExecutionContextStack struct {
	stack stack.Stack
}

// NewExecutionContextStack creates a new ExecutionContextStack that is using a SliceStack.
func NewExecutionContextStack() *ExecutionContextStack {
	s := new(ExecutionContextStack)
	s.stack = stack.NewSliceStack()
	return s
}

// IsEmpty is used to determine if there are any ExecutionContexts on the stack.
func (s ExecutionContextStack) IsEmpty() bool {
	return s.stack.Peek() == nil
}

// Push adds a new ExecutionContext to the stack.
func (s ExecutionContextStack) Push(ctx *ExecutionContext) {
	s.stack.Push(ctx)
}

// Pop removes the topmost ExecutionContext from the stack and returns it.
func (s ExecutionContextStack) Pop() *ExecutionContext {
	elem := s.stack.Pop()
	if elem == nil {
		return nil
	}

	return elem.(*ExecutionContext)
}

// Peek returns the topmost ExecutionContext without removing it.
func (s ExecutionContextStack) Peek() *ExecutionContext {
	elem := s.stack.Peek()
	if elem == nil {
		return nil
	}

	return elem.(*ExecutionContext)
}
