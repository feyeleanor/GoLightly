package vm
import "testing"

var REGISTER int

func opcode(i int) *OpCode {
	return &OpCode{code: i}
}

func defaultInstructionSet() *InstructionSet {
	i := new(InstructionSet)
	i.Init()
	i.Define("zero",	func (o *Buffer) { REGISTER = 0 })
	i.Define("one",		func (o *Buffer) { REGISTER = 1 })
	i.Define("two",		func (o *Buffer) { REGISTER = 2 })
	i.Define("three",	func (o *Buffer) { REGISTER = 3 })
	i.Define("four",	func (o *Buffer) { REGISTER = 4 })
	return i
}

func checkInstructionExists(name string, i *InstructionSet, t *testing.T) {
	_, ok := i.tokens[name]
	compareValues(i, t, ok, true)
	compareValues(i, t, ok, true)
}

func checkDefaultInstructionSet(i *InstructionSet, t *testing.T) {
	checkInstructionExists("zero", i, t)
	checkInstructionExists("one", i, t)
	checkInstructionExists("two", i, t)
	checkInstructionExists("three", i, t)
	checkInstructionExists("four", i, t)
}

func checkInstructionInvocation(i *InstructionSet, t *testing.T) {
	for j, f := range i.ops {
		f.(func (o *Buffer))(&opcode(j).data)
		compareValues(i, t, REGISTER, j)
	}
	for j := 0; j < i.Len(); j++ {
		i.Invoke(opcode(j))
		compareValues(i, t, REGISTER, j)
	}
}

func checkInstructionSearch(i *InstructionSet, t *testing.T) {
	compareValues(i, t, i.Code("zero"), 0)
	compareValues(i, t, i.Code("two"), 2)
	compareValues(i, t, i.Code("five"), -1)
}

func checkInstructionCompilation(i *InstructionSet, t *testing.T) {
	zero, one := i.Code("zero"), i.Code("one")
	compareValues(i, t, i.OpCode("zero", 0, 0, 0).Identical(&OpCode{code: zero, data: []int{0, 0, 0}}), true)
	compareValues(i, t, i.OpCode("zero", 1, 0, 0).Identical(&OpCode{code: zero, data: []int{0, 0, 0}}), false)
	compareValues(i, t, i.OpCode("one", 0, 0, 0).Similar(&OpCode{code: one, data: []int{1, 0, 0}}), true)
	compareValues(i, t, i.OpCode("zero", 1, 0, 0).Similar(&OpCode{code: one, data: []int{1, 0, 0}}), false)
}

func TestInstructionSetCreation(t *testing.T) {
	i := defaultInstructionSet()
	compareValues(i, t, i.Len(), 5)
	checkDefaultInstructionSet(i, t)
	checkInstructionInvocation(i, t)
	checkInstructionSearch(i, t)
	checkInstructionCompilation(i, t)
}
