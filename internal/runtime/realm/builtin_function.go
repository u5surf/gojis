package realm

import (
	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

func CreateBuiltinFunction(fn func(lang.Value, ...lang.Value) (lang.Value, errors.Error), realm *Realm, proto lang.Value, internalSlotsList ...lang.StringOrSymbol) *lang.Object {
	if realm == nil {
		realm = CurrentRealm()
	}
	if proto == nil {
		proto = realm.GetIntrinsicObject(IntrinsicNameFunctionPrototype)
	}
	fobj := lang.ObjectCreate(proto, internalSlotsList...)
	fobj.Call = fn
	fobj.Realm = realm
	fobj.Extensible = true
	fobj.ScriptOrModule = lang.Null
	return fobj
}
