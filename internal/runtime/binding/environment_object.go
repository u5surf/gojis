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

func (e *ObjectEnvironment) Outer() Environment {
	return e.outer
}

func (e *ObjectEnvironment) IsGlobalEnvironment() bool {
	panic("TODO")
}

func (e *ObjectEnvironment) IsModuleEnvironment() bool {
	panic("TODO")
}

func (e *ObjectEnvironment) HasBinding(n lang.String) bool {
	panic("TODO")
}

func (e *ObjectEnvironment) CreateMutableBinding(n lang.String, deletable bool) errors.Error {
	panic("TODO")
}

func (e *ObjectEnvironment) CreateImmutableBinding(n lang.String, strict bool) errors.Error {
	panic("TODO")
}

func (e *ObjectEnvironment) InitializeBinding(n lang.String, val lang.Value) errors.Error {
	panic("TODO")
}

func (e *ObjectEnvironment) SetMutableBinding(n lang.String, val lang.Value, strict bool) errors.Error {
	panic("TODO")
}

func (e *ObjectEnvironment) GetThisBinding() (lang.Value, errors.Error) {
	panic("TODO")
}

func (e *ObjectEnvironment) GetBindingValue(n lang.String, strict bool) (lang.Value, errors.Error) {
	panic("TODO")
}

func (e *ObjectEnvironment) DeleteBinding(n lang.String) bool {
	panic("TODO")
}

func (e *ObjectEnvironment) HasThisBinding() bool {
	panic("TODO")
}

func (e *ObjectEnvironment) HasSuperBinding() bool {
	panic("TODO")
}

func (e *ObjectEnvironment) WithBaseObject() lang.Value {
	panic("TODO")
}

func (e *ObjectEnvironment) Type() lang.Type    { return lang.TypeInternal }
func (e *ObjectEnvironment) Value() interface{} { return e }
