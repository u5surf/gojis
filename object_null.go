package gojis

const (
	// Null represents the Null ECMAScript language value.
	Null = null(0)
)

type null uint8

func (u null) Lookup(objName string) Object                  { return Undefined }
func (u null) SetFunction(name string, fn func(Args) Object) { /* no-op */ }
func (u null) CallWithArgs(args ...interface{}) (Object, error) {
	panic("TODO: return API error 'not callable'")
}
func (u null) SetObject(name string, obj Object) { /* no-op */ }
func (u null) IsUndefined() bool                 { return false }
func (u null) IsNull() bool                      { return true }
func (u null) IsFunction() bool                  { return false }
func (u null) Type() Type                        { return TypeNull }
func (u null) Value() interface{}                { return nil }
