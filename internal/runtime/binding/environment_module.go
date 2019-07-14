package binding

import "github.com/gojisvm/gojis/internal/runtime/lang"

// ModuleEnvironment is a declarative environment that is used to represent the
// outer scope of an ECMAScript Module. In addition to normal mutable and
// immutable bindings, module environments also provide immutable import
// bindings which are bindings that provide indirect access to a target binding
// that exists in another environment. ModuleEnvironment is specified in
// 8.1.1.5.
type ModuleEnvironment struct {
	*DeclarativeEnvironment
}

// CreateImportBinding creates a new initialized immutable indirect binding for
// the name n. A binding must not already exist in this environment for the
// given name. CreateImportBinding is specified in 8.1.1.5.5.
func (e *ModuleEnvironment) CreateImportBinding(n lang.String, m interface{}, n2 lang.String) {
	panic("TODO: modules")
}

// GetThisBinding returns Undefined.
// GetThisBinding is specified in 8.1.1.5.4.
func (e *ModuleEnvironment) GetThisBinding() lang.Value {
	return lang.Undefined
}
