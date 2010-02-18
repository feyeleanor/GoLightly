package vm
import "testing"
import "os"

var REGISTER int

func opcode(i int) *OpCode {
	return &OpCode{code: i}
}

func defaultInstructionSet() *InstructionSet {
	i := new(InstructionSet)
	i.Init()
	i.Define("zero",	func (o *OpCode)	{ REGISTER = 0 })
	i.Define("one",		func (o *OpCode)	{ REGISTER = 1 })
	i.Define("two",		func (o *OpCode)	{ REGISTER = 2 })
	i.Define("three",	func (o *OpCode)	{ REGISTER = 3 })
	i.Define("four",	func (o *OpCode)	{ REGISTER = 4 })
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
		f.(func (o *OpCode))(opcode(j))
		compareValues(i, t, REGISTER, j)
	}
	for j := 0; j < i.Len(); j++ {
		i.Invoke(opcode(j))
		compareValues(i, t, REGISTER, j)
	}
}

func checkInstructionSearch(i *InstructionSet, t *testing.T) {
	compareValues(i, t, i.OpCode("zero"), 0)
	compareValues(i, t, i.OpCode("two"), 2)
	compareValues(i, t, i.OpCode("five"), -1)
}

func TestInstructionSetCreation(t *testing.T) {
	os.Stdout.WriteString("Instruction Set Creation\n")
	i := defaultInstructionSet()
	compareValues(i, t, i.Len(), 5)
	checkDefaultInstructionSet(i, t)
	checkInstructionInvocation(i, t)
	checkInstructionSearch(i, t)
}
