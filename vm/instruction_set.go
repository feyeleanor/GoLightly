//	TODO:	bytecode optimisation
//	TODO:	JIT compilation
//	TODO:	AOT compilation

package vm

import "container/vector"
import "fmt"
import . "golightly/storage"
import "reflect"

type OpCode struct {
	code	int
	data	IntBuffer
}
func (o *OpCode) Similar(p *OpCode) bool {
	return o.code == p.code && o.data.Len() == p.data.Len()
}
func (o *OpCode) Identical(p *OpCode) bool {
	return reflect.DeepEqual(o.data, &p.data)
}
func (o *OpCode) Replace(p *OpCode) {
	o.code = p.code
	o.data = p.data.Clone()
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
func (i *InstructionSet) Exists(name string) bool {
	_, ok := i.tokens[name]
	return ok
}
func (i *InstructionSet) Define(name string, closure func (o *IntBuffer)) bool {
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
func (i *InstructionSet) OpCode(name string, data *IntBuffer) (r *OpCode) {
	if op := i.Code(name); op != -1 {
		if data != nil {
			r = &OpCode{op, *data}
		} else {
			r = &OpCode{op, IntBuffer{}}
		}
	}
	return
}
func (i *InstructionSet) Invoke(o *OpCode) bool {
	if o.code < 0 || o.code >= i.Len() {
		return false
	}
	i.ops.At(o.code).(func (o *IntBuffer))(&o.data)
	return true
}
