package vm

import "testing"
import "fmt"
import . "golightly/storage"
//import "container/vector"

func BenchmarkCreateOpCode(b *testing.B) {
	var op *OpCode
	for i := 0; i < b.N; i++ { op = new(OpCode) }
	op.code = 1
}

func BenchmarkIdenticalOpCode(b *testing.B) {
	op1 := OpCode{1, 0, nil}
	op2 := OpCode{2, 0, nil}
	for i := 0; i < b.N; i++ { op1.Identical(op2) }
}

func BenchmarkReplaceOpCode(b *testing.B) {
	op1 := OpCode{1, 0, nil}
	op2 := OpCode{2, 0, nil}
	for i := 0; i < b.N; i++ { op1.Replace(&op2) }
}

func BenchmarkInitInstructionSet(b *testing.B) {
	for i := 0; i < b.N; i++ { new(InstructionSet).Init() }
}

func BenchmarkDefineInstruction(b *testing.B) {
	b.StopTimer()
		in := new(InstructionSet)
		in.Init()
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.Define(fmt.Sprintf("%v", i), 0, func(o *IntBuffer) {}) }
}

func BenchmarkDefineOperator(b *testing.B) {
	b.StopTimer()
		in := new(InstructionSet)
		in.Init()
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.Operator(fmt.Sprintf("%v", i), func(o *IntBuffer) {}) }
}

func BenchmarkDefineMovement(b *testing.B) {
	b.StopTimer()
		in := new(InstructionSet)
		in.Init()
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.Movement(fmt.Sprintf("%v", i), func(o *IntBuffer) {}) }
}

func BenchmarkNameToInstruction(b *testing.B) {
	b.StopTimer()
		in := new(InstructionSet)
		in.Init()
		for i := 0; i < 1000; i++ {
			in.Operator(fmt.Sprintf("%v", i), func(o *IntBuffer) {})
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.Instruction("3") }
}

func BenchmarkAssemble(b *testing.B) {
	b.StopTimer()
		in := new(InstructionSet)
		in.Init()
		for i := 0; i < 1000; i++ {
			in.Operator(fmt.Sprintf("%v", i), func(o *IntBuffer) {})
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.Assemble("3", &IntBuffer{1, 2, 3}) }
}

func BenchmarkInstructionInvocation(b *testing.B) {
	b.StopTimer()
		in := new(InstructionSet)
		in.Init()
		for i := 0; i < 1000; i++ {
			in.Operator(fmt.Sprintf("%v", i), func(o *IntBuffer) {})
		}
		op := in.Assemble("3", nil)
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.Invoke(&op) }
}
