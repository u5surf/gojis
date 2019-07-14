package binding

import (
	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

// Status represents the status of a binding.
// Can either be 'lexical', 'initialized', 'uninitialized'.
type Status string

// Available binding statuses.
const (
	StatusLexical       Status = "lexical"
	StatusInitialized   Status = "initialized"
	StatusUninitialized Status = "uninitialized"
)

var _ Environment = (*FunctionEnvironment)(nil)

// FunctionEnvironment is a declarative environment that is
// ised to represent the top-level scope of a function and,
// if the function is not an arrow function, provides a this
// binding. If a function is not an arrow function and references
// super, its function environment also contains the state that is
// ised to perform super method invocations from within the function.
// FunctionEnvironment is specified in 8.1.1.3.
type FunctionEnvironment struct {
	*DeclarativeEnvironment

	// Used as 'this' for invocation of the function.
	ThisValue lang.Value
	// If the value is StatusLexical, this function is an
	// arrow function and does not have a local this value.
	ThisBindingStatus Status
	// The function object whose invocation caused this
	// environment to be created.
	FunctionObject *lang.Object
	// If the associated function has 'super' property
	// accesses and is not an arrow function, HomeObject
	// is the object that the function is bound to as a
	// method. The default value for HomeObject is Undefined.
	HomeObject lang.Value // Object or Undefined
	// If this environment was created by the Construct-internal
	// method, NewTarget is the value of the Construct-function's
	// newTarget parameter. Otherwise, this value is Undefined.
	NewTarget lang.Value // Object or Undefined
}

// BindThisValue sets ThisValue and records that it has been initialized.
// BindThisValue is specified in 8.1.1.3.1.
func (e *FunctionEnvironment) BindThisValue(val lang.Value) (lang.Value, errors.Error) {
	if e.ThisBindingStatus == StatusInitialized {
		return nil, errors.NewReferenceError("'This' has already been initialized")
	}

	e.ThisValue = val
	e.ThisBindingStatus = StatusInitialized
	return val, nil
}

// HasThisBinding returns true if and only if ThisBindingStatus is not StatusLexical.
// HasThisBinding is specified in 8.1.1.3.2.
func (e *FunctionEnvironment) HasThisBinding() bool {
	return e.ThisBindingStatus != StatusLexical
}

// HasSuperBinding returns true if and only if this HomeObject is not Undefined.
// HasSuperBinding is specified in 8.1.1.3.3.
func (e *FunctionEnvironment) HasSuperBinding() bool {
	if e.ThisBindingStatus == StatusLexical {
		return false
	}

	return e.HomeObject != lang.Undefined
}

// GetThisBinding returns the ThisValue of this environment.
// If it has not been initialized yet, a reference error will be returned.
// GetThisBinding is specified in 8.1.1.3.4.
func (e *FunctionEnvironment) GetThisBinding() (lang.Value, errors.Error) {
	if e.ThisBindingStatus == StatusUninitialized {
		return nil, errors.NewReferenceError("'This' has not been initialized yet")
	}

	return e.ThisValue, nil
}

// GetSuperBase returns the prototype object of the super binding's value.
// GetSuperBase is specified in 8.1.1.3.5.
func (e *FunctionEnvironment) GetSuperBase() lang.Value {
	if e.HomeObject == lang.Undefined {
		return lang.Undefined
	}

	return e.HomeObject.(*lang.Object).GetPrototypeOf()
}
