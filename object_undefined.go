package gojis

const (
	// Undefined represents the Undefined ECMAScript language value.
	Undefined = undefined(0)
)

type undefined uint8

func (u undefined) Lookup(objName string) Object                  { return Undefined }
func (u undefined) SetFunction(name string, fn func(Args) Object) { /* no-op */ }
func (u undefined) CallWithArgs(args ...interface{}) (Object, error) {
	panic("TODO: return API error not callable")
}
func (u undefined) SetObject(name string, obj Object) { /* no-op */ }
func (u undefined) IsUndefined() bool                 { return true }
func (u undefined) IsNull() bool                      { return false }
func (u undefined) IsFunction() bool                  { return false }
func (u undefined) Type() Type                        { return TypeUndefined }
func (u undefined) Value() interface{}                { return nil }
