package binding

import (
	"fmt"

	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

var _ Environment = (*GlobalEnvironment)(nil)

// GlobalEnvironment is used to represent the outer most scope that is shared by all of the ECMAScript Script
// elements that are processed in a common realm. A global Environment Record provides the bindings for built-in
// globals (clause 18), properties of the global object, and for all top-level declarations (13.2.8, 13.2.10) that occur within a
// Script.
// A global Environment Record is logically a single record but it is specified as a composite encapsulating an object
// Environment Record and a declarative Environment Record. The object Environment Record has as its base object the
// global object of the associated Realm Record. This global object is the value returned by the global Environment
// Record's GetThisBinding concrete method. The object Environment Record component of a global Environment
// Record contains the bindings for all built-in globals (clause 18) and all bindings introduced by a FunctionDeclaration,
// GeneratorDeclaration, AsyncFunctionDeclaration, AsyncGeneratorDeclaration, or VariableStatement contained in global
// code. The bindings for all other ECMAScript declarations in global code are contained in the declarative Environment
// Record component of the global Environment Record.
// Properties may be created directly on a global object. Hence, the object Environment Record component of a global
// Environment Record may contain both bindings created explicitly by FunctionDeclaration, GeneratorDeclaration,
// AsyncFunctionDeclaration, AsyncGeneratorDeclaration, or VariableDeclaration declarations and bindings created
// implicitly as properties of the global object. In order to identify which bindings were explicitly created using
// declarations, a global Environment Record maintains a list of the names bound using its CreateGlobalVarBinding and
// CreateGlobalFunctionBinding concrete methods.
// GlobalEnvironment is specified in 8.1.1.4.
type GlobalEnvironment struct {
	ObjectRecord      *ObjectEnvironment
	GlobalThisValue   *lang.Object
	DeclarativeRecord *DeclarativeEnvironment
	VarNames          []string
}

// Outer returns nil.
func (e *GlobalEnvironment) Outer() Environment {
	return nil
}

// GetThisBinding returns the GlobalThisValue.
// GetThisBinding is specified in 8.1.1.4.11.
func (e *GlobalEnvironment) GetThisBinding() (lang.Value, errors.Error) {
	return e.GlobalThisValue, nil
}

// HasVarDeclaration determines if the argument identifier has a binding
// in this record that was created using a VariableStatement or a FunctionDeclaration.
// HasVarDeclaration is specified in 8.1.1.4.12.
func (e *GlobalEnvironment) HasVarDeclaration(n lang.String) bool {
	nVal := n.Value().(string)
	for _, varName := range e.VarNames {
		if varName == nVal {
			return true
		}
	}
	return false
}

// HasLexicalDeclaration determines if the argument identifier has a
// binding in this record that was created using a lexical declaration
// such as a LexicalDeclaration or a ClassDeclaration.
// HasLexicalDeclaration is specified in 8.1.1.4.13.
func (e *GlobalEnvironment) HasLexicalDeclaration(n lang.String) bool {
	return e.DeclarativeRecord.HasBinding(n)
}

func (e *GlobalEnvironment) HasRestrictedGlobalProperty(n lang.String) bool {
	existingProp := e.ObjectRecord.bindingObject.GetOwnProperty(lang.NewStringOrSymbol(n)).Value()
	if existingProp == lang.Undefined {
		return false
	}

	panic("TODO: properties")
}

func (e *GlobalEnvironment) CanDeclareGlobalVar(n lang.String) bool {
	globalObj := e.ObjectRecord.bindingObject

	if lang.HasOwnProperty(globalObj, lang.NewStringOrSymbol(n)) {
		return true
	}

	return lang.InternalIsExtensible(globalObj)
}

func (e *GlobalEnvironment) CanDeclareGlobalFunction(n lang.String) {
	panic("TODO: properties")
}

func (e *GlobalEnvironment) CreateGlobalVarBinding(n lang.String, deletable bool) {
	globalObj := e.ObjectRecord.bindingObject

	hasProperty := lang.HasOwnProperty(globalObj, lang.NewStringOrSymbol(n))
	extensible := lang.InternalIsExtensible(globalObj)
	if !hasProperty.Value().(bool) && extensible {
		e.ObjectRecord.CreateMutableBinding(n, deletable)
		e.ObjectRecord.InitializeBinding(n, lang.Undefined)
	}

	if !e.HasVarDeclaration(n) {
		e.VarNames = append(e.VarNames, n.Value().(string))
	}
}

func (e *GlobalEnvironment) CreateGlobalFunctionBinding(n lang.String, val lang.Value, deletable bool) {
	panic("TODO: properties")
}

/* -- implements Environment -- */

func (e *GlobalEnvironment) HasBinding(n lang.String) bool {
	return e.DeclarativeRecord.HasBinding(n) || e.ObjectRecord.HasBinding(n)
}

func (e *GlobalEnvironment) CreateMutableBinding(n lang.String, deletable bool) errors.Error {
	if e.DeclarativeRecord.HasBinding(n) {
		return errors.NewTypeError(fmt.Sprintf("Declarative environment record already has a binding for '%v'", n))
	}

	return e.DeclarativeRecord.CreateMutableBinding(n, deletable)
}

func (e *GlobalEnvironment) CreateImmutableBinding(n lang.String, strict bool) errors.Error {
	if e.DeclarativeRecord.HasBinding(n) {
		return errors.NewTypeError(fmt.Sprintf("Declarative environment record already has a binding for '%v'", n))
	}

	return e.DeclarativeRecord.CreateImmutableBinding(n, strict)
}

func (e *GlobalEnvironment) InitializeBinding(n lang.String, val lang.Value) errors.Error {
	if e.DeclarativeRecord.HasBinding(n) {
		return e.DeclarativeRecord.InitializeBinding(n, val)
	}

	return e.ObjectRecord.InitializeBinding(n, val)
}

func (e *GlobalEnvironment) SetMutableBinding(n lang.String, val lang.Value, strict bool) errors.Error {
	if e.DeclarativeRecord.HasBinding(n) {
		return e.DeclarativeRecord.SetMutableBinding(n, val, strict)
	}

	return e.ObjectRecord.SetMutableBinding(n, val, strict)
}

func (e *GlobalEnvironment) GetBindingValue(n lang.String, strict bool) (lang.Value, errors.Error) {
	if e.DeclarativeRecord.HasBinding(n) {
		return e.DeclarativeRecord.GetBindingValue(n, strict)
	}

	return e.ObjectRecord.GetBindingValue(n, strict)
}

func (e *GlobalEnvironment) DeleteBinding(n lang.String) bool {
	if e.DeclarativeRecord.HasBinding(n) {
		return e.DeclarativeRecord.DeleteBinding(n)
	}

	if lang.HasOwnProperty(e.ObjectRecord.bindingObject, lang.NewStringOrSymbol(n)) {
		status := e.ObjectRecord.DeleteBinding(n)
		if status {
			nVal := n.Value().(string)
			for i, varName := range e.VarNames {
				if varName == nVal {
					e.VarNames[i] = e.VarNames[len(e.VarNames)-1]
					e.VarNames = e.VarNames[:len(e.VarNames)-1]
					break
				}
			}
		}

		return status
	}

	return true
}

// HasThisBinding returns true.
// HasThisBinding is specified in 8.1.1.4.8.
func (e *GlobalEnvironment) HasThisBinding() bool {
	return true
}

// HasSuperBinding returns false.
// HasSuperBinding is specified in 8.1.1.4.9.
func (e *GlobalEnvironment) HasSuperBinding() bool {
	return false
}

// WithBaseObject returns Undefined.
// WithBaseObject is specified in 8.1.1.4.10.
func (e *GlobalEnvironment) WithBaseObject() lang.Value {
	return lang.Undefined
}

// Type returns TypeInternal.
func (e *GlobalEnvironment) Type() lang.Type { return lang.TypeInternal }

// Value returns the GlobalEnvironment itself.
func (e *GlobalEnvironment) Value() interface{} { return e }
