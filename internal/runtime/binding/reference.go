package binding

import (
	"fmt"

	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

// Reference is a resolved name or property binding. A Reference consists of
// three components, the base value component, the referenced name component,
// and the Boolean-valued strict reference flag. The base value component is
// either undefined, an Object, a Boolean, a String, a Symbol, a Number, or an
// Environment Record. A base value component of undefined indicates that the
// Reference could not be resolved to a binding. The referenced name component
// is a String or Symbol value. A Super Reference is a Reference that is used to
// represent a name binding that was expressed using the super keyword. A Super
// Reference has an additional thisValue component, and its base value component
// will never be an Environment Record.
// Reference is specified in 6.2.4.
type Reference struct {
	baseComponent  lang.Value
	referencedName lang.StringOrSymbol
	thisValue      lang.Value
	strict         bool
}

// NewReference creates a new named reference with the given name and the given
// base value. The created reference is a strict reference if the strict
// argument is true.
func NewReference(n lang.StringOrSymbol, base lang.Value, strict bool) *Reference {
	r := new(Reference)
	r.referencedName = n
	r.baseComponent = base
	r.strict = strict
	return r
}

// NewSuperReference creates a new super reference with the given name and the
// given base and this value. The created super reference is a common reference
// with a set this value. The created reference is a strict reference if the
// strict argument is true.
func NewSuperReference(n lang.StringOrSymbol, base, this lang.Value, strict bool) *Reference {
	r := NewReference(n, base, strict)
	r.thisValue = this
	return r
}

// GetBase returns the base component value of this reference.
// GetBase is specified in 6.2.4.1.
func (r *Reference) GetBase() lang.Value {
	return r.baseComponent
}

// GetReferencedName returns the name of this reference.
// GetReferencedName is specified in 6.2.4.2.
func (r *Reference) GetReferencedName() lang.StringOrSymbol {
	return r.referencedName
}

// IsStrictReference is used to determine if this reference is a strict
// reference.
// IsStrictReference is specified in 6.2.4.3.
func (r *Reference) IsStrictReference() bool {
	return r.strict
}

// HasPrimitiveBase returns true if and only if the type of this references base
// component is one of [Boolean, String, Symbol, Number].
// HasPrimitiveBase is specified in 6.2.4.4.
func (r *Reference) HasPrimitiveBase() bool {
	switch r.baseComponent.Type() {
	case lang.TypeBoolean, lang.TypeString, lang.TypeSymbol, lang.TypeNumber:
		return true
	default:
		return false
	}
}

// IsPropertyReference returns true if and only if the type of this referecens
// base component is Object or this reference has a primitive base
// (HasPrimitiveBase).
// IsPropertyReference is specified in 6.2.4.5.
func (r *Reference) IsPropertyReference() bool {
	return r.baseComponent.Type() == lang.TypeObject ||
		r.HasPrimitiveBase()
}

// IsUnresolvableReference is used to determine if this references base
// component is Undefined.
// IsUnresolvableReference is specified in 6.2.4.6.
func (r *Reference) IsUnresolvableReference() bool {
	return r.baseComponent == lang.Undefined
}

// IsSuperReference is used to determine if this reference has a thisValue.
// IsSuperReference will return true even the thisValue is Null or Undefined,
// since the thisValue is present.
// IsSuperReference is specified in 6.2.4.7.
func (r *Reference) IsSuperReference() bool {
	return r.thisValue != nil
}

// GetValue returns the value of this reference.
// GetValue is specified in 6.2.4.8.
func (r *Reference) GetValue() (lang.Value, errors.Error) {
	base := r.GetBase()

	if r.IsUnresolvableReference() {
		return nil, errors.NewReferenceError(fmt.Sprintf("Unresolvable reference: '%v'", r.referencedName))
	}

	if r.IsPropertyReference() {
		if r.HasPrimitiveBase() {
			base = lang.ToObject(base)
		}
		panic("TODO: properties")
	}

	return base.Value().(Environment).GetBindingValue(r.GetReferencedName().String(), r.IsStrictReference())
}

// PutValue sets the value of this reference to the given one.
// PutValue is specified in 6.2.4.9.
func (r *Reference) PutValue(lang.Value) errors.Error {
	panic("TODO: 6.2.4.9 PutValue")
}

// GetThisValue returns the thisValue of this reference. If this reference is
// not a super reference and thus does not have a thisValue, the base of this
// reference is returned instead.
// GetThisValue is specified in 6.2.4.10.
func (r *Reference) GetThisValue() lang.Value {
	if r.IsSuperReference() {
		return r.thisValue
	}

	return r.GetBase()
}

// InitializeReferencedBinding initializes this reference in its environment
// with the given value.
// InitializeReferencedBinding is specified in 6.2.4.11.
func (r *Reference) InitializeReferencedBinding(val lang.Value) errors.Error {
	return r.GetBase().(Environment).InitializeBinding(r.GetReferencedName().String(), val)
}
