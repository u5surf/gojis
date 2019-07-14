package gojis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNull(t *testing.T) {
	t.Run("Lookup", func(t *testing.T) {
		require := require.New(t)

		require.Equal(Undefined, Null.Lookup(""))
		require.Equal(Undefined, Null.Lookup("foo"))
	})

	t.Run("SetFunction", func(t *testing.T) {
		// this should just not panic
		Null.SetFunction("some_func", func(Args) Object { return Undefined })
		Null.SetFunction("nothing", func(Args) Object { return Undefined })
	})

	t.Run("SetObject", func(t *testing.T) {
		// this should just not panic
		Null.SetObject("some_obj", Undefined)
		Null.SetObject("nothing", Undefined)
	})

	t.Run("IsXXX", func(t *testing.T) {
		require := require.New(t)

		require.False(Null.IsUndefined())
		require.True(Null.IsNull())
		require.False(Null.IsFunction())
	})

	t.Run("Type / Value", func(t *testing.T) {
		require := require.New(t)

		require.EqualValues(TypeNull, Null.Type())
		require.Nil(Null.Value())
	})

}
