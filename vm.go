package gojis

import "io"

// Type represents the ECMAScript language types.
type Type uint8

// Available types. TypeUnknown must not be used.
// If it appears, this indicates, that somewhere the type was not set
// (it is the default value for Type).
const (
	TypeUnknown = iota
	TypeUndefined
	TypeNull
	TypeString
)

var (
	undefined *Object
)

// VM represents an instance of the GojisVM.
// It can be used to evaluate ECMAScript code.
type VM struct {
	*Object // the global object
}

// NewVM creates a new, initialized VM that is ready to use.
func NewVM() *VM { panic("TODO") }

// Eval evaluates the given ECMAScript code, and returns an Object, representing the
// result of the evaluation. The result may be Null or Undefined.
//
// Internally, this function directly delegates to the method 'eval' that is specified
// by the ECMAScript language specification.
func (vm *VM) Eval(script string) *Object { return vm.Lookup("eval") }

// SetConsole is used to change the console of the VM.
// Calls like 'console.log' will be written to the given writer.
func (vm *VM) SetConsole(w io.Writer) { panic("TODO") }

// Args represent the arguments that are passed to a function call.
// The arguments can be retrieved using Args#Get(int), and can be used
// in the function.
// The arguments are Objects, which can be Null or Undefined.
// If 3 arguments are passed to the function, args.Get(5) will
// return Undefined, NOT nil.
type Args struct {
	o []*Object
}

// Get returns the argument at the given index.
// If there is no such argument, Undefined will be returned.
// This method never returns nil.
func (a *Args) Get(index int) *Object {
	if index >= len(a.o) {
		return undefined
	}
	return a.o[index]
}

// Len returns the amount of arguments.
// For example, if 3 arguments were passed (indices 0, 1 and 2),
// this method will return 3.
// Please note, that args.Get(3) and higher will return Undefined,
// whereas args.Get(0), 1 and 2 will return the respective arguments.
func (a *Args) Len() int {
	return len(a.o)
}

// Object represents any ECMAScript language value.
// This can be a String or a Number as well as Null or Undefined.
// To check if the object represents Null, use Object#IsNull.
// To check if the object represents Undefined, use Object#IsUndefined.
type Object struct{}

// Lookup returns a property of this object with the given name.
// If no such property exists, Undefined will be returned.
// This method never returns nil.
func (o *Object) Lookup(objName string) *Object { panic("TODO") }

// SetFunction adds a property (more specific, a function object) to this
// object. When the function object is invoked, the given function will be
// executed, with all parameters wrapped into the Args object.
func (o *Object) SetFunction(name string, fn func(Args)) { panic("TODO") }

// CallWithArgs attempts to invoke this objects 'Call' property. Note that this property
// is only set if this object is a function or constructor object.
// If this object does not have a 'Call' property that is callable, an error will be returned.
func (o *Object) CallWithArgs(args ...interface{}) (*Object, error) { panic("TODO") }

// SetObject adds a property with the given name to this object.
// The propertie's value will be the given Object.
func (o *Object) SetObject(name string, obj *Object) { panic("TODO") }

// IsUndefined is used to determine whether this Object represents the Undefined value.
func (o *Object) IsUndefined() bool { panic("TODO") }

// IsNull is used to determine whether this Object represents the Null value.
func (o *Object) IsNull() bool { panic("TODO") }

// IsFunction is used to determine whether this Object can be invoked.
// If this returns true, CallWithArgs will not return an error.
func (o *Object) IsFunction() bool { panic("TODO") }

// Type returns the ECMAScript language type of this object.
func (o *Object) Type() Type { panic("TODO") }

// Value returns the Go value corresponding to the ECMAScript language value
// of this object.
func (o *Object) Value() interface{} { panic("TODO") }
