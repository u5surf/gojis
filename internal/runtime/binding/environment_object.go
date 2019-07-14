package binding

import (
	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

var _ Environment = (*ObjectEnvironment)(nil)

// ObjectEnvironment is associated with an object called its binding object. An
// object Environment Record binds the set of string identifier names that
// directly correspond to the property names of its binding object. Property
// keys that are not strings in the form of an IdentifierName are not included
// in the set of bound identifiers. Both own and inherited properties are
// included in the set regardless of the setting of their [[Enumerable]]
// attribute. Because properties can be dynamically added and deleted from
// objects, the set of identifiers bound by an object Environment Record may
// potentially change as a side-effect of any operation that adds or deletes
// properties. Any bindings that are created as a result of such a side-effect
// are considered to be a mutable binding even if the Writable attribute of the
// corresponding property has the value false. Immutable bindings do not exist
// for object Environment Records. Object Environment Records created for
// withwith statements (13.11) can provide their binding object as an implicit
// this value for use in function calls. The capability is controlled by a
// withEnvironment Boolean value that is associated with each object Environment
// Record. By default, the value of withEnvironment is false for any object
// Environment Record.
// ObjectEnvironment is specified in 8.1.1.2.
type ObjectEnvironment struct {
	outer Environment

	bindingObject *lang.Object
}

// Outer returns the outer environment of this object environment.
func (e *ObjectEnvironment) Outer() Environment {
	return e.outer
}

// IsGlobalEnvironment returns false.
func (e *ObjectEnvironment) IsGlobalEnvironment() bool {
	return false
}

// IsModuleEnvironment returns false.
func (e *ObjectEnvironment) IsModuleEnvironment() bool {
	return false
}

// HasBinding determines if its associated binding object has a property whose
// name is the value of the argument n.
// HasBinding is specified in 8.1.1.2.1.
func (e *ObjectEnvironment) HasBinding(n lang.String) bool {
	panic("TODO")
}

// CreateMutableBinding creates in an Environment Record's associated binding
// object a property whose name is the String value and initializes it to the
// value undefined. If Boolean argument D has the value true the new property's
// [[Configurable]] attribute is set to true; otherwise it is set to false.
// CreateMutableBinding is specified in 8.1.1.2.2.
func (e *ObjectEnvironment) CreateMutableBinding(n lang.String, deletable bool) errors.Error {
	panic("TODO")
}

// CreateImmutableBinding is not described by the specification and will panic
// upon call.
// CreateImmutableBinding is not specified in 8.1.1.2.3.
func (e *ObjectEnvironment) CreateImmutableBinding(n lang.String, strict bool) errors.Error {
	panic("Not described by specification, should not be called.")
}

// InitializeBinding is used to set the bound
// value of the current binding of the identifier whose name is the value of the
// argument n to the value of argument val. An uninitialized binding for n must
// already exist.
// InitializeBinding is specified in 8.1.1.2.4.
func (e *ObjectEnvironment) InitializeBinding(n lang.String, val lang.Value) errors.Error {
	panic("TODO")
}

// SetMutableBinding attempts to set the value of the Environment Record's
// associated binding object's property whose name is the value of the argument
// n to the value of argument val. A property named n normally already exists but
// if it does not or is not currently writable, error handling is determined by
// the value of strict.
// SetMutableBinding is specified in 8.1.1.2.5.
func (e *ObjectEnvironment) SetMutableBinding(n lang.String, val lang.Value, strict bool) errors.Error {
	panic("TODO")
}

// GetThisBinding is not specified for ObjectEnvironment.
func (e *ObjectEnvironment) GetThisBinding() (lang.Value, errors.Error) {
	panic("Not specified for ObjectEnvironment, must not be called.")
}

// GetBindingValue returns the value of its associated binding object's property
// whose name is the String value of the argument identifier n. The property
// should already exist but if it does not the result depends upon the value of
// the strict argument.
// GetBindingValue is specified in 8.1.1.2.6.
func (e *ObjectEnvironment) GetBindingValue(n lang.String, strict bool) (lang.Value, errors.Error) {
	panic("TODO")
}

// DeleteBinding can only delete bindings that correspond to properties of the
// environment object whose [[Configurable]] attribute have the value true.
// DeleteBinding is specified in 8.1.1.2.7.
func (e *ObjectEnvironment) DeleteBinding(n lang.String) bool {
	panic("TODO")
}

// HasThisBinding returns false.
// HasThisBinding is specified in 8.1.1.2.8.
func (e *ObjectEnvironment) HasThisBinding() bool {
	return false
}

// HasSuperBinding returns false.
// HasSuperBinding is specified in 8.1.1.2.9.
func (e *ObjectEnvironment) HasSuperBinding() bool {
	return false
}

// WithBaseObject return Undefined as their WithBaseObject unless their
// withEnvironment flag is true.
// WithBaseObject is specified in 8.1.1.2.10.
func (e *ObjectEnvironment) WithBaseObject() lang.Value {
	panic("TODO")
}

// Type returns TypeInternal.
func (e *ObjectEnvironment) Type() lang.Type { return lang.TypeInternal }

// Value returns the ObjectEnvironment itself.
func (e *ObjectEnvironment) Value() interface{} { return e }
