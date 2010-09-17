package vm

import "testing"
import . "golightly/test"

const BUFFER_ALLOCATION = 32

func simpleProgram(p *ProcessorCore) *Program {
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

func TestProcessorCore(t *testing.T) {
	NewTest(t).
	Run("Creation", func(TC *Test) {
		p := new(ProcessorCore)
		p.Init(BUFFER_ALLOCATION, nil)
		for i := 0; i < p.R.Len(); i++ {
			TC.	Identical(p.R.At(i), 0)
		}
		TC.	Identical(BUFFER_ALLOCATION, p.R.Len(), p.R.Cap()).
			Refute(p.Running).
			Refute(p.ValidPC()).
			Identical(p.pc, 0).
			Confirm(p.M == nil)
	}).
	Run("Basic Instruction Stepping", func(TC *Test) {
		p := new(ProcessorCore)
		p.Init(BUFFER_ALLOCATION, nil)

		program := simpleProgram(p)
		p.LoadProgram(program)
		TC.	Identical(p.I(), program.code[0]).
			Identical(p.pc, 0).
			Confirm(p.Program.ValidPC()).
			Refute(p.ValidPC())

		p.StepForward()
		TC.	Identical(p.I(), program.code[1]).
			Identical(p.pc, 1).
			Confirm(p.Program.ValidPC()).
			Refute(p.ValidPC())

		p.StepBack()
		TC.	Identical(p.I(), program.code[0]).
			Identical(p.pc, 0).
			Confirm(p.Program.ValidPC()).
			Refute(p.ValidPC())

		p.JumpTo(program.Len())
		TC.	Confirm(p.Program.ValidPC()).
			Identical(program.Len() - 1, p.pc).
			Refute(p.ValidPC())

		p.StepForward()
		TC.	Refute(p.Program.ValidPC()).
			Refute(p.ValidPC())

		p.ResetState()
		for i := 0; i < p.R.Len(); i++ {
			TC.	Identical(p.R.At(i), 0)
		}
		TC.	Identical(BUFFER_ALLOCATION, p.R.Len(), p.R.Cap()).
			Refute(p.Running).
			Identical(p.pc, 0).
			Confirm(p.M == nil).
			Confirm(p.Program.ValidPC()).
			Refute(p.ValidPC()).
			Identical(p.Program.Len(), program.Len())
	}).
	Run("Execution", func(TC *Test) {
		p := new(ProcessorCore)
		p.Init(BUFFER_ALLOCATION, nil)

		program := simpleProgram(p)
		p.LoadProgram(program)
		p.Execute()												//	program[0]
		TC.	Identical(p.R.At(0), 37)

		p.StepForward()
		p.Execute()
		TC.	Identical(p.R.At(0), 38)							//	program[1]

		p.ResetState()
		p.StepForward()											//	skip program[0]
		p.Execute()
		TC.	Identical(p.R.At(0), 1)								//	program[1]

		p.ResetState()
		p.Run()													//	program[7] is an illegal instruction
		TC.	Refute(p.Running).
			Confirm(p.Illegal_Operation).
			Identical(p.pc, p.Program.Len() - 1).				//	terminated without executing program[7]
			Identical(p.R.At(0), 37)

		p.Program.code[7] = &OpCode{							//	patch program[7] with a valid instruction
			code: p.Code("cld"),
			data: Buffer{1, 100},
		}
		p.ResetState()
		p.Run()
		TC.	Refute(p.Running).
			Refute(p.Illegal_Operation).
			Identical(p.pc, p.Program.Len()).					//	terminated with all instructions executed
			Identical(p.R.At(0), 37).
			Identical(p.R.At(1), 100)

		p.LoadProgram(advancedProgram(p))
		p.Run()
		TC.	Identical(p.R.At(0), 0).
			Identical(p.R.At(1), 1000)
	}).
	Run("To Do", func(TC *Test) {
		TC.	Untested("InstructionSet handling").
			Untested("ProcessorCore cloning").
			Untested("IOController handling")
	})
}
