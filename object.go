package gojis

var (
	undefined *Object
	null      *Object
)

// Object represents any ECMAScript language value. This can be a String or a
// Number as well as Null or Undefined. To check if the object represents Null,
// use Object#IsNull. To check if the object represents Undefined, use
// Object#IsUndefined.
type Object struct{}

// Lookup returns a property of this object with the given name. If no such
// property exists, Undefined will be returned. This method never returns nil.
func (o *Object) Lookup(objName string) *Object { panic("TODO") }

// SetFunction adds a property (more specific, a function object) to this
// object. When the function object is invoked, the given function will be
// executed, with all parameters wrapped into the Args object.
func (o *Object) SetFunction(name string, fn func(Args)) { panic("TODO") }

// CallWithArgs attempts to invoke this objects 'Call' property. Note that this
// property is only set if this object is a function or constructor object. If
// this object does not have a 'Call' property that is callable, an error will
// be returned.
func (o *Object) CallWithArgs(args ...interface{}) (*Object, error) { panic("TODO") }

// SetObject adds a property with the given name to this object. The propertie's
// value will be the given Object.
func (o *Object) SetObject(name string, obj *Object) { panic("TODO") }

// IsUndefined is used to determine whether this Object represents the Undefined
// value.
func (o *Object) IsUndefined() bool { panic("TODO") }

// IsNull is used to determine whether this Object represents the Null value.
func (o *Object) IsNull() bool { panic("TODO") }

// IsFunction is used to determine whether this Object can be invoked. If this
// returns true, CallWithArgs will not return an error.
func (o *Object) IsFunction() bool { panic("TODO") }

// Type returns the ECMAScript language type of this object.
func (o *Object) Type() Type { panic("TODO") }

// Value returns the Go value corresponding to the ECMAScript language value of
// this object.
func (o *Object) Value() interface{} { panic("TODO") }
