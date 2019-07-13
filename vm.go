package gojis

import "io"

type Type uint8

const (
	TypeUnknown = iota
	TypeUndefined
	TypeNull
	TypeString
)

var (
	undefined *Object
)

type VM struct {
	*Object // the global object
}

func NewVM() *VM { panic("TODO") }

func (vm *VM) Eval(script string) *Object { return vm.Lookup("eval") }

func (vm *VM) SetConsole(w io.Writer) { panic("TODO") }

type Args struct {
	o []*Object
}

func (a *Args) Get(index int) *Object {
	if index >= len(a.o) {
		return undefined
	}
	return a.o[index]
}

func (a *Args) Len() int {
	return len(a.o)
}

type Object struct{}

func (o *Object) Lookup(objName string) *Object            { panic("TODO") }
func (o *Object) SetFunction(name string, fn func(Args))   { panic("TODO") }
func (o *Object) CallWithArgs(args ...interface{}) *Object { panic("TODO") }
func (o *Object) SetObject(name string, obj *Object)       { panic("TODO") }

func (o *Object) IsUndefined() bool   { panic("TODO") }
func (o *Object) IsNull() bool        { panic("TODO") }
func (o *Object) IsFunction() bool    { panic("TODO") }
func (o *Object) IsConstructor() bool { panic("TODO") }

func (o *Object) Type() Type         { panic("TODO") }
func (o *Object) Value() interface{} { panic("TODO") }
