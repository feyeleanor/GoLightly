package vm

import "testing"
import . "golightly/test"

const BUFFER_ALLOCATION = 32

func simpleProgram(a Assembler) Program {
	return Program{ a.Assemble("cld", []int{0, 37}),				//	cld		R0, 37
					a.Assemble("inc", 0),							//	inc		R0
					a.Assemble("inc", 0),							//	inc		R0
					a.Assemble("dec", 0),							//	dec		R0
					a.Assemble("inc", 0),							//	inc		R0
					a.Assemble("dec", 0),							//	dec		R0
					a.Assemble("dec", 0),							//	dec		R0
					OpCode{99,	1,	nil},							//	illegal operation
	}
}

func intermediateProgram(a Assembler) Program {
	return Program{ a.Assemble("cld", []int{0, 1}),					//	0	cld		R0, 1
					a.Assemble("call", 4),							//	1	call	5
					a.Assemble("jmpnz",	[]int{0, -1}),				//	2	jmpnz	R0, -1
					a.Assemble("halt", nil),						//	3	halt
					a.Assemble("dec", 0),							//	4	dec		R0
					a.Assemble("inc", 1),							//	5	inc		R1
					a.Assemble("ret", nil),							//	6	ret
	}
}

func advancedProgram(a Assembler) Program {
	program := intermediateProgram(a)
	program[0].Data = []int{0, 1000}
	return program
}

func TestProcessorCore(t *testing.T) {
	NewTest(t).
	Run("Creation", func(TC *Test) {
		p := new(ProcessorCore)
		p.Init(BUFFER_ALLOCATION, nil)
		for _, v := range p.R {
			TC.	Identical(v, 0)
		}
		TC.	Identical(BUFFER_ALLOCATION, len(p.R), cap(p.R)).
			Refute(p.Running).
			Refute(p.ValidPC()).
			Identical(p.PC, 0).
			Confirm(p.M == nil)
	}).
	Run("Basic Instruction Stepping", func(TC *Test) {
		p := new(ProcessorCore)
		p.Init(BUFFER_ALLOCATION, nil)

		I := func() { p.I() }

		program := simpleProgram(p)
		TC.	Identical(len(program), 8)

		p.LoadProgram(program)
		TC.	Identical(len(program), len(p.Program)).
			Safe(I).
			Identical(p.I(), program[0]).
			Identical(p.PC, 0).
			Confirm(p.ValidPC()).
			Refute(p.Running)

		p.PC++
		TC.	Safe(I).
			Identical(p.I(), program[1]).
			Identical(p.PC, 1).
			Confirm(p.ValidPC()).
			Refute(p.Running)

		p.PC--
		TC.	Safe(I).
			Identical(p.I(), program[0]).
			Identical(p.PC, 0).
			Confirm(p.ValidPC()).
			Refute(p.Running)

		p.PC = len(program)
		TC.	Erroneous(I).
			Refute(p.ValidPC()).
			Refute(p.Running)

		p.PC++
		TC.	Erroneous(I).
			Identical(p.PC, len(p.Program) + 1).
			Refute(p.ValidPC()).
			Refute(p.Running)

		p.ResetState()
		for _, v := range p.R {
			TC.	Identical(v, 0)
		}
		TC.	Identical(BUFFER_ALLOCATION, len(p.R), cap(p.R)).
			Refute(p.Running).
			Identical(p.PC, 0).
			Confirm(p.M == nil).
			Confirm(p.ValidPC()).
			Refute(p.Running)
	}).
	Run("Execution", func(TC *Test) {
		p := new(ProcessorCore)
		p.Init(BUFFER_ALLOCATION, nil)

		program := simpleProgram(p)
		p.LoadProgram(program)
		p.Execute()												//	program[0]
		TC.	Identical(p.R[0], 37)

		p.PC++
		p.Execute()
		TC.	Identical(p.R[0], 38)								//	program[1]

		p.ResetState()
		p.PC++											//	skip program[0]
		p.Execute()												//	execute program[1]
		TC.	Identical(p.R[0], 1).
			Identical(p.PC, 2)

		p.ResetState()
		p.Run()													//	program[7] is an illegal instruction
		TC.	Refute(p.Running).
			Identical(p.PC, 7).
			Identical(p.R[0], 37)

/*		p.Program.Code[7] = OpCode{								//	patch program[7] with a valid instruction
			Code: p.Instruction("cld").op,
			Data: []int{1, 100},
		}
		p.ResetState()
		p.Run()
		TC.	Refute(p.Running).
			Identical(p.PC, p.Program.Len()).					//	terminated with all instructions executed
			Identical(p.R[0], 37).
			Identical(p.R[1], 100)

		p.LoadProgram(advancedProgram(p))
		p.Run()
		TC.	Identical(p.R[0], 0).
			Identical(p.R[1], 1000)
*/	}).
	Run("To Do", func(TC *Test) {
		TC.	Untested("InstructionSet handling").
			Untested("ProcessorCore cloning").
			Untested("IOController handling")
	})
}