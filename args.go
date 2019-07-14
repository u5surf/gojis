package gojis

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

// Len returns the amount of arguments. For example, if 3 arguments were passed
// (indices 0, 1 and 2), this method will return 3. Please note, that in this
// example, args.Get(3) and higher will return Undefined, whereas args.Get(0), 1
// and 2 will return the respective arguments.
func (a *Args) Len() int {
	return len(a.o)
}
