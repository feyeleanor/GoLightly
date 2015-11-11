//	TODO:	bytecode optimisation
//	TODO:	JIT compilation
//	TODO:	AOT compilation

package vm

import (
	"github.com/feyeleanor/slices"
)

type Assembler interface {
	Assemble(name string, data interface{}) OpCode
}

type Instruction struct {
	op				int
	Movement	int
	params		string
}

type InstructionSet struct {
	ops					slices.Slice
	tokens			map[string] *Instruction
}

func (i *InstructionSet) Init() {
	i.tokens = make(map[string] *Instruction)
}

func (i *InstructionSet) Len() int {
	return i.ops.Len()
}

func (i *InstructionSet) Exists(name string) bool {
	_, ok := i.tokens[name]
	return ok
}

func (i *InstructionSet) Define(name, params string, Movement int, closure interface{}) (successful bool) {
	if _, ok := i.tokens[name]; !ok {
		i.ops.Append(closure)
		i.tokens[name] = &Instruction{op: i.ops.Len() - 1, params: params, Movement: Movement}
		successful = true
	}
	return
}

func (i *InstructionSet) Movement(name, params string, data interface{}) bool {
	return i.Define(name, params, 0, data)
}

func (i *InstructionSet) Operator(name, params string, data interface{}) bool {
	return i.Define(name, params, 1, data)
}

func (i *InstructionSet) Instruction(name string) *Instruction {
	if op, ok := i.tokens[name]; ok {
		return op
	}
	return nil
}

func (i *InstructionSet) Assemble(name string, data interface{}) OpCode {
	if op := i.Instruction(name); op != nil {
		return OpCode{Code: op.op, Movement: op.Movement, Data: data}
	}
	panic(name)
}

func (i *InstructionSet) Invoke(o *OpCode) {
	switch data := o.Data.(type) {
	case []int:
		i.ops.At(o.Code).(func (o []int))(data)
	}
}