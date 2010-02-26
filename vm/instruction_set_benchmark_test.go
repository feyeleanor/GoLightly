package vm
import "testing"
import "fmt"
//import "container/vector"

func BenchmarkCreateOpCode(b *testing.B) {
	var op *OpCode
	for i := 0; i < b.N; i++ { op = new(OpCode) }
	op.code = 1
}

func BenchmarkSimilarOpCode(b *testing.B) {
	op1 := &OpCode{code: 1}
	op2 := &OpCode{code: 2}
	for i := 0; i < b.N; i++ { op1.Similar(op2) }
}

func BenchmarkIdenticalOpCode(b *testing.B) {
	op1 := &OpCode{code: 1}
	op2 := &OpCode{code: 2}
	for i := 0; i < b.N; i++ { op1.Identical(op2) }
}

func BenchmarkReplaceOpCode(b *testing.B) {
	op1 := &OpCode{code: 1}
	op2 := &OpCode{code: 2}
	for i := 0; i < b.N; i++ { op1.Replace(op2) }
}

func BenchmarkInitInstructionSet(b *testing.B) {
	for i := 0; i < b.N; i++ { new(InstructionSet).Init() }
}

func BenchmarkDefineInstruction(b *testing.B) {
	in := new(InstructionSet)
	in.Init()
	for i := 0; i < b.N; i++ { in.Define(fmt.Sprintf("%v", i), func(o *OpCode) {}) }
}

func defaultBMInstructionSet() *InstructionSet {
	in := new(InstructionSet)
	in.Init()
	for i := 0; i < 1000; i++ { in.Define(fmt.Sprintf("%v", i), func(o *OpCode) {}) }
	return in
}

func BenchmarkInstructionToCode(b *testing.B) {
	b.StopTimer()
	in := defaultBMInstructionSet()
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.Code("3") }
}

func BenchmarkInstructionToOpCode(b *testing.B) {
	b.StopTimer()
	in := defaultBMInstructionSet()
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.OpCode("3", 1, 2, 3) }
}

func BenchmarkInstructionInvocation(b *testing.B) {
	b.StopTimer()
	in := defaultBMInstructionSet()
	op := &OpCode{3, []int{0, 0, 0}}
	b.StartTimer()
	for i := 0; i < b.N; i++ { in.Invoke(op) }
}
