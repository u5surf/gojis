package agent

import (
	"github.com/gojisvm/gojis/internal/runtime/binding"
	"github.com/gojisvm/gojis/internal/runtime/lang"
	"github.com/gojisvm/gojis/internal/runtime/realm"
)

// ExecutionContext is a specification device that is used to track the runtime
// evaluation of code by an ECMAScript implementation. At any point in time,
// there is at most one execution context per agent that is actually executing
// code. This is known as the agent's running execution context. All references
// to the running execution context in this specification denote the running
// execution context of the surrounding agent.
// ExecutionContext is specified in 8.3.
type ExecutionContext struct {
	Function       lang.Value // Object or Null
	Realm          *realm.Realm
	ScriptOrModule lang.InternalValue

	LexicalEnvironment  binding.Environment
	VariableEnvironment binding.Environment

	Generator interface{} // TODO: Table 23, GeneratorObject
}
