//	TODO: 	Refactor tests as specs
//	TODO:	Add tests for InstructionSet handling
//	TODO:	Add tests for ProcessorCore cloning

package vm
import "testing"
import "os"

const BUFFER_ALLOCATION = 32

func defaultRegisterBlock() *RegisterBlock {
	r := new(RegisterBlock)
	r.Allocate(BUFFER_ALLOCATION)
	return r
}

func checkAllocatedBuffer(b *Buffer, t *testing.T) {
	compareValues(b, t, b.Len(), BUFFER_ALLOCATION)
	compareValues(b, t, b.Cap(), BUFFER_ALLOCATION)
	for i := 0; i < b.Len(); i++ { compareValues(b, t, b.At(i), 0) }
}

func checkDefaultRegisterBlock(r *RegisterBlock, t *testing.T) {
	checkAllocatedBuffer(r.R, t)
	compareValues(r, t, r.M == nil, true)
	compareValues(r, t, r.I == nil, true)
	compareValues(r, t, r.PC, 0)
}

func TestRegisterBlock(t *testing.T) {
	os.Stdout.WriteString("Register Block Creation\n")
	r := defaultRegisterBlock()
	checkDefaultRegisterBlock(r, t)
	os.Stdout.WriteString("Register Block Cloning\n")
	c := r.Clone()
	checkDefaultRegisterBlock(c, t)
	os.Stdout.WriteString("Register Block Replacement\n")
	r.PC = 27
	r.M = new(Buffer)
	r.M.Init(48)
	compareValues(r, t, r.PC, 27)
	compareValues(r, t, r.M.Cap(), 48)
	r.Replace(c)
	checkDefaultRegisterBlock(r, t)
	r.Clear()
	checkDefaultRegisterBlock(r, t)
}

func TestMMU(t *testing.T) {
	os.Stdout.WriteString("MMU Allocation\n")
	m := new(MMU)
	b := m.Allocate(BUFFER_ALLOCATION)
	checkAllocatedBuffer(b, t)
}

func defaultProgram() *[]*OpCode {
	cld := &OpCode{code: 4, a: 0, b: 37}					//	cld		0, 37
	inc := &OpCode{code: 5, a: 0}							//	inc		0
	dec := &OpCode{code: 6, a: 0}							//	dec		0
	ill := &OpCode{code: 99}								//	illegal operation
	return &[]*OpCode {cld, inc, inc, dec, inc, dec, dec, ill}
}

func checkProcessorInitialised(p *ProcessorCore, t *testing.T) {
	checkAllocatedBuffer(p.R, t)
	compareValues(p, t, p.Running, true)
	compareValues(p, t, p.PC, 0)
	compareValues(p, t, p.I == nil, true)
	compareValues(p, t, p.M == nil, true)	
}

func TestProcessorCoreCreation(t *testing.T) {
	os.Stdout.WriteString("Processor Core Creation\n")
	p := new(ProcessorCore)
	p.Init(BUFFER_ALLOCATION, nil)
	checkProcessorInitialised(p, t)
}

func checkInstruction(p *ProcessorCore, t *testing.T, program []*OpCode, pc int) {
	compareValues(p, t, p.I.Identical(program[pc]), true)
	compareValues(p, t, p.PC, pc)
}

func resetProcessor(p *ProcessorCore, t *testing.T) {
	p.ResetState()
	checkProcessorInitialised(p, t)
}

func TestProcessorCoreExecution(t *testing.T) {
	os.Stdout.WriteString("Processor Core Program Execution\n")
	p := new(ProcessorCore)
	p.Init(BUFFER_ALLOCATION, nil)
	checkProcessorInitialised(p, t)
	compareValues(p, t, p.ValidPC(), false)
	program := defaultProgram()
	p.LoadProgram(program)
	checkInstruction(p, t, *program, 0)
	p.StepForward()
	checkInstruction(p, t, *program, 1)
	p.StepBack()
	checkInstruction(p, t, *program, 0)
	p.StepForward()
	checkInstruction(p, t, *program, 1)
	p.Jump(len(*program))
	compareValues(p, t, p.ValidPC(), false)
	resetProcessor(p, t)
	compareValues(p, t, p.ValidPC(), true)
	compareValues(p, t, len(p.program), len(*defaultProgram()))
	p.LoadInstruction()
	p.Execute()													//	program[0]
	compareValues(p, t, p.R.At(0), 37)
	p.StepForward()
	p.Execute()
	compareValues(p, t, p.R.At(0), 38)							//	program[1]
	resetProcessor(p, t)
	p.StepForward()
	p.Execute()
	compareValues(p, t, p.R.At(0), 1)							//	program[1]
	resetProcessor(p, t)
	p.Run()														//	program[7] is an illegal instruction
	compareValues(p, t, p.Running, false)
	compareValues(p, t, p.Illegal_Operation, true)
	compareValues(p, t, p.PC, len(p.program) - 1)				//	terminated without executing program[7]
	compareValues(p, t, p.R.At(0), 37)
	p.program[7] = &OpCode{code: 4, a: 1, b: 100}				//	replace program[7] with		cld	1, 100
	resetProcessor(p, t)
	p.Run()
	compareValues(p, t, p.Running, false)
	compareValues(p, t, p.Illegal_Operation, false)
	compareValues(p, t, p.PC, len(p.program))					//	terminated with all instructions executed
	compareValues(p, t, p.R.At(0), 37)
	compareValues(p, t, p.R.At(1), 100)
}
