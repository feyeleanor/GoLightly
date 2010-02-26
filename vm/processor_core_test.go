//	TODO: 	Refactor tests as specs
//	TODO:	Add tests for InstructionSet handling
//	TODO:	Add tests for ProcessorCore cloning
//	TODO:	Add tests for IOController handling

package vm
import "testing"

const BUFFER_ALLOCATION = 32

func defaultRegisterBlock() *RegisterBlock {
	r := new(RegisterBlock)
	r.Allocate(BUFFER_ALLOCATION)
	return r
}

func checkAllocatedVector(s *Vector, t *testing.T) {
	compareValues(s, t, s.Len(), BUFFER_ALLOCATION)
	compareValues(s, t, s.Cap(), BUFFER_ALLOCATION)
	for i := 0; i < s.Len(); i++ { compareValues(s, t, s.At(i), 0) }
}

func checkDefaultRegisterBlock(r *RegisterBlock, t *testing.T) {
	checkAllocatedVector(r.R, t)
	compareValues(r, t, r.M == nil, true)
	compareValues(r, t, r.I == nil, true)
	compareValues(r, t, r.PC, 0)
}

func TestRegisterBlock(t *testing.T) {
	r := defaultRegisterBlock()
	checkDefaultRegisterBlock(r, t)
	c := r.Clone()
	checkDefaultRegisterBlock(c, t)
	r.PC = 27
	r.M = new(Vector)
	r.M.Init(48)
	compareValues(r, t, r.PC, 27)
	compareValues(r, t, r.M.Cap(), 48)
	r.Replace(c)
	checkDefaultRegisterBlock(r, t)
	r.Clear()
	checkDefaultRegisterBlock(r, t)
}

func TestMMU(t *testing.T) {
	m := new(MMU)
	s := m.Allocate(BUFFER_ALLOCATION)
	checkAllocatedVector(s, t)
}

func defaultProgram(p *ProcessorCore) *[]*OpCode {
	cld := p.OpCode("cld", 0, 37, 0)								//	cld		0, 37
	inc := p.OpCode("inc", 0, 0, 0)									//	inc		0
	dec := p.OpCode("dec", 0, 0, 0)									//	dec		0
	ill := &OpCode{code: 99}										//	illegal operation
	return &[]*OpCode {cld, inc, inc, dec, inc, dec, dec, ill}
}

func advancedProgram(p *ProcessorCore) *[]*OpCode {
	return &[]*OpCode{
		p.OpCode("cld",		0,	1000,	0),			//	0	cld		R0, 1000
		p.OpCode("call",	5,	0,		0),			//	1	call	5
		p.OpCode("dec",		0,	0,		0),			//	2	dec		R0
		p.OpCode("jmpnz",	0,	-2,		0),			//	3	jmpnz	R0, -2
		p.OpCode("halt",	0,	0,		0),			//	4	halt
		p.OpCode("push",	0,	0,		0),			//	5	push	R0
		p.OpCode("inc",		1,	0,		0),			//	6	inc		R1
		p.OpCode("pop",		0,	0,		0),			//	7	pop		R0
		p.OpCode("ret",		0,	0,		0),			//	8	ret
	}
}

func checkProcessorInitialised(p *ProcessorCore, t *testing.T) {
	checkAllocatedVector(p.R, t)
	compareValues(p, t, p.Running, true)
	compareValues(p, t, p.PC, 0)
	compareValues(p, t, p.I == nil, true)
	compareValues(p, t, p.M == nil, true)	
}

func TestProcessorCoreCreation(t *testing.T) {
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

func checkFlowControl(p *ProcessorCore, t *testing.T, program *[]*OpCode) {
	checkInstruction(p, t, *program, 0)
	p.StepForward()
	checkInstruction(p, t, *program, 1)
	p.StepBack()
	checkInstruction(p, t, *program, 0)
	p.StepForward()
	checkInstruction(p, t, *program, 1)
	p.JumpTo(len(*program))
	compareValues(p, t, p.ValidPC(), false)
	resetProcessor(p, t)
	compareValues(p, t, p.ValidPC(), true)
}

func checkProgramExecution(p *ProcessorCore, t *testing.T) {
	compareValues(p, t, len(p.program), len(*defaultProgram(p)))
	p.LoadInstruction()
	p.Execute()															//	program[0]
	compareValues(p, t, p.R.At(0), 37)
	p.StepForward()
	p.Execute()
	compareValues(p, t, p.R.At(0), 38)									//	program[1]
	resetProcessor(p, t)
	p.StepForward()
	p.Execute()
	compareValues(p, t, p.R.At(0), 1)									//	program[1]
	resetProcessor(p, t)
	p.Run()																//	program[7] is an illegal instruction
	compareValues(p, t, p.Running, false)
	compareValues(p, t, p.Illegal_Operation, true)
	compareValues(p, t, p.PC, len(p.program) - 1)						//	terminated without executing program[7]
	compareValues(p, t, p.R.At(0), 37)
	p.program[7] = &OpCode{code: p.Code("cld"), data: []int{1, 100}}	//	replace program[7] with		cld	1, 100
	resetProcessor(p, t)
	p.Run()
	compareValues(p, t, p.Running, false)
	compareValues(p, t, p.Illegal_Operation, false)
	compareValues(p, t, p.PC, len(p.program))							//	terminated with all instructions executed
	compareValues(p, t, p.R.At(0), 37)
	compareValues(p, t, p.R.At(1), 100)
}

func TestProcessorCoreExecution(t *testing.T) {
	p := new(ProcessorCore)
	p.Init(BUFFER_ALLOCATION, nil)
	checkProcessorInitialised(p, t)
	compareValues(p, t, p.ValidPC(), false)
	program := defaultProgram(p)
	p.LoadProgram(program)
	checkFlowControl(p, t, program)
	checkProgramExecution(p, t)
	p.LoadProgram(advancedProgram(p))
	p.Run()
	compareValues(p, t, p.R.At(0), 0)
	compareValues(p, t, p.R.At(1), 1000)
}
