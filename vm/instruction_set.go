//	TODO:	bytecode optimisation
//	TODO:	JIT compilation
//	TODO:	AOT compilation

package vm

import "container/vector"
import "fmt"

type OpCode struct {
	code	int
	data	Buffer
}
func (o *OpCode) Similar(p *OpCode) bool {
	return o.code == p.code && o.data.Len() == p.data.Len()
}
func (o *OpCode) Identical(p *OpCode) bool {
	return o.Similar(p) && o.data.Identical(&p.data)
}
func (o *OpCode) Replace(p *OpCode) {
	o.code = p.code
	o.data = *p.data.Clone()
}
func (o *OpCode) String() string {
	return fmt.Sprintf("%v: %v", o.code, o.data)
}

type InstructionSet struct {
	ops				vector.Vector
	tokens			map[string]int
}
func (i *InstructionSet) Init() {
	i.tokens = make(map[string]int)
}
func (i *InstructionSet) Len() int {
	return i.ops.Len()
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
func (i *InstructionSet) Code(name string) int {
	if op, ok := i.tokens[name]; ok {
		return op
	}
	return -1
}
func (i *InstructionSet) OpCode(name string, a, b, c int) *OpCode {
	if op := i.Code(name); op != -1 {
		return &OpCode{code: op, data: []int{a, b, c}}
	}
	return nil
}
func (i *InstructionSet) Invoke(o *OpCode) bool {
	if o.code < 0 || o.code >= i.ops.Len() { return false }
	i.ops.At(o.code).(func (o *OpCode))(o)
	return true
}
