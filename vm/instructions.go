//	TODO:	bytecode optimisation
//	TODO:	JIT compilation
//	TODO:	AOT compilation

package vm

import "container/vector"

type InstructionSet struct {
	ops				vector.Vector
	tokens			map[string]int
}
func (i *InstructionSet) Init() {
	i.tokens = make(map[string]int)
	i.Define("noop", func (o *OpCode) {})											//	NOOP
}
func (i *InstructionSet) Define(name string, closure func (o *OpCode)) bool {
	// Ensure instruction token hasn't yet been defined
	if _, ok := i.tokens[name]; !ok {
		i.ops.Push(closure)
		i.tokens[name] = i.ops.Len() - 1
		return true
	}
	return false
}
func (i *InstructionSet) Find(name string) *func (o *OpCode) {
	if op, error := i.tokens[name]; !error {
		return i.ops.At(op).(*func (o *OpCode))
	}
	return nil
}
func (i *InstructionSet) Invoke(o *OpCode) bool {
	if o.code < 0 || o.code >= i.ops.Len() { return false }
	i.ops.At(o.code).(func (o *OpCode))(o)
	return true
}
