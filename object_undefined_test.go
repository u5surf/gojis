package gojis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUndefined(t *testing.T) {
	t.Run("Lookup", func(t *testing.T) {
		require := require.New(t)

		require.Equal(Undefined, Undefined.Lookup(""))
		require.Equal(Undefined, Undefined.Lookup("foo"))
	})

	t.Run("SetFunction", func(t *testing.T) {
		// this should just not panic
		Undefined.SetFunction("some_func", func(Args) Object { return Null })
		Undefined.SetFunction("nothing", func(Args) Object { return Null })
	})

	t.Run("SetObject", func(t *testing.T) {
		// this should just not panic
		Undefined.SetObject("some_obj", Null)
		Undefined.SetObject("nothing", Null)
	})

	t.Run("IsXXX", func(t *testing.T) {
		require := require.New(t)

		require.True(Undefined.IsUndefined())
		require.False(Undefined.IsNull())
		require.False(Undefined.IsFunction())
	})

	t.Run("Type / Value", func(t *testing.T) {
		require := require.New(t)

		require.EqualValues(TypeUndefined, Undefined.Type())
		require.Nil(Undefined.Value())
	})

}
