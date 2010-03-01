//	TODO: 	Refactor tests as specs
//	TODO:	Add tests for InstructionSet handling
//	TODO:	Add tests for ProcessorCore cloning
//	TODO:	Add tests for IOController handling

package vm
import "testing"

const BUFFER_ALLOCATION = 32

func checkAllocatedVector(s *Vector, t *testing.T) {
	compareValues(s, t, s.Len(), BUFFER_ALLOCATION)
	compareValues(s, t, s.Cap(), BUFFER_ALLOCATION)
	for i := 0; i < s.Len(); i++ { compareValues(s, t, s.At(i), 0) }
}

func defaultProgram(p *ProcessorCore) *Program {
	code := []*OpCode{
		p.OpCode("cld", &Buffer{0, 37, 0}),				//	cld		0, 37
		p.OpCode("inc", &Buffer{0, 0, 0}),				//	inc		0
		p.OpCode("inc", &Buffer{0, 0, 0}),				//	inc		0
		p.OpCode("dec", &Buffer{0, 0, 0}),				//	dec		0
		p.OpCode("inc", &Buffer{0, 0, 0}),				//	inc		0
		p.OpCode("dec", &Buffer{0, 0, 0}),				//	dec		0
		p.OpCode("dec", &Buffer{0, 0, 0}),				//	dec		0
		&OpCode{99, Buffer{0, 0, 0}},					//	illegal operation
	}
	return &Program{code: code}
}

func advancedProgram(p *ProcessorCore) *Program {
	code := []*OpCode{
		p.OpCode("cld",		&Buffer{0, 1000,	0}),		//	0	cld		R0, 1000
		p.OpCode("call",	&Buffer{5, 0,	0}),			//	1	call	5
		p.OpCode("dec",		&Buffer{0, 0,	0}),			//	2	dec		R0
		p.OpCode("jmpnz",	&Buffer{0, -2,	0}),			//	3	jmpnz	R0, -2
		p.OpCode("halt",	&Buffer{0, 0,	0}),			//	4	halt
		p.OpCode("push",	&Buffer{0, 0,	0}),			//	5	push	R0
		p.OpCode("inc",		&Buffer{1, 0,	0}),			//	6	inc		R1
		p.OpCode("pop",		&Buffer{0, 0,	0}),			//	7	pop		R0
		p.OpCode("ret",		&Buffer{0, 0,	0}),			//	8	ret
	}
	return &Program{code: code}
}

func checkProcessorInitialised(p *ProcessorCore, t *testing.T) {
	checkAllocatedVector(p.R, t)
	compareValues(p, t, p.Running, false)
	compareValues(p, t, p.pc, 0)
	compareValues(p, t, p.M == nil, true)	
}

func TestProcessorCoreCreation(t *testing.T) {
	p := new(ProcessorCore)
	p.Init(BUFFER_ALLOCATION, nil)
	checkProcessorInitialised(p, t)
}

func checkInstruction(p *ProcessorCore, t *testing.T, program *Program, pc int) {
	compareValues(p, t, p.I().Identical(program.code[pc]), true)
	compareValues(p, t, p.pc, pc)
}

func resetProcessor(p *ProcessorCore, t *testing.T) {
	p.ResetState()
	checkProcessorInitialised(p, t)
}

func checkFlowControl(p *ProcessorCore, t *testing.T, program *Program) {
	checkInstruction(p, t, program, 0)
	p.StepForward()
	checkInstruction(p, t, program, 1)
	p.StepBack()
	checkInstruction(p, t, program, 0)
	p.StepForward()
	checkInstruction(p, t, program, 1)
	p.JumpTo(program.Len())
	compareValues(p, t, p.Program.ValidPC(), true)
	compareValues(p, t, p.ValidPC(), false)
	p.StepForward()
	compareValues(p, t, p.Program.ValidPC(), false)
	compareValues(p, t, p.ValidPC(), false)
	resetProcessor(p, t)
	compareValues(p, t, p.Program.ValidPC(), true)
	compareValues(p, t, p.ValidPC(), false)
}

func checkProgramExecution(p *ProcessorCore, t *testing.T) {
	compareValues(p, t, p.Program.Len(), defaultProgram(p).Len())
	p.Execute()																//	program[0]
	compareValues(p, t, p.R.At(0), 37)
	p.StepForward()
	p.Execute()
	compareValues(p, t, p.R.At(0), 38)										//	program[1]
	resetProcessor(p, t)
	p.StepForward()
	p.Execute()
	compareValues(p, t, p.R.At(0), 1)										//	program[1]
	resetProcessor(p, t)
	p.Run()																	//	program[7] is an illegal instruction
	compareValues(p, t, p.Running, false)
	compareValues(p, t, p.Illegal_Operation, true)
	compareValues(p, t, p.pc, p.Program.Len() - 1)							//	terminated without executing program[7]
	compareValues(p, t, p.R.At(0), 37)
	p.Program.code[7] = &OpCode{code: p.Code("cld"), data: Buffer{1, 100}}	//	replace program[7] with cld 1, 100
	resetProcessor(p, t)
	p.Run()
	compareValues(p, t, p.Running, false)
	compareValues(p, t, p.Illegal_Operation, false)
	compareValues(p, t, p.pc, p.Program.Len())								//	terminated with all instructions executed
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
