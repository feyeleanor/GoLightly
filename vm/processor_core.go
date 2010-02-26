//	TODO:	storing and retrieving pointers to memory buffers
//	TODO:	cloning should create a comms channel by which the parent and child cores can communicate
//	TODO:	should always have stdin, stdout and stderr channels

package vm

import "container/vector"
import "syscall"
//import "os"
//import "fmt"

type ExecutionFlags struct {
	Running				bool
	Illegal_Operation	bool
	Stack_Underflow		bool
	Segmentation_Error	bool
	Divide_by_Zero		bool
}
func (f *ExecutionFlags) Clear() {
	f.Running = false
	f.Illegal_Operation = false
	f.Stack_Underflow = false
	f.Segmentation_Error = false
}

type RegisterBlock struct {
	PC				int
	I				*OpCode
	R				*Vector
	M				*Vector
}
func (r *RegisterBlock) Allocate(count int) {
	r.PC = 0
	r.I = nil
	r.R = new(Vector)
	r.R.Init(count)
	r.M = nil
}
func (r *RegisterBlock) Clone() *RegisterBlock {
	return &RegisterBlock{ PC: r.PC, I: r.I, M: r.M, R: r.R.Clone() }
}
func (r *RegisterBlock) Replace(a *RegisterBlock) {
	r.PC, r.I, r.R, r.M = a.PC, a.I, a.R, a.M
}
func (r *RegisterBlock) Clear() {
	r.PC = 0
	r.I = nil
	r.R.ClearAll()
	r.M = nil
}

type MMU struct {}
func (m *MMU) Allocate(words int) *Vector {
	s := new(Vector)
	s.Init(words)
	return s
}

type ProcessorCore struct {
	ExecutionFlags
	MMU
	RegisterBlock
	IOController
	*InstructionSet
	program			[]*OpCode
	call_stack		vector.IntVector
	data_stack		vector.IntVector
}
func (p *ProcessorCore) Init(registers int, instructions *InstructionSet) {
	p.RegisterBlock.Allocate(registers)
	p.IOController.Init()
	if instructions == nil {
		p.InstructionSet = new(InstructionSet)
		p.InstructionSet.Init()
		p.DefineInstructions()
	} else {
		p.InstructionSet = instructions
	}
	p.ResetState()
}
//	Make a copy of the current processor, binding it to the current processor with
//	the supplied io channel
func (p *ProcessorCore) Clone(c chan *Vector) (q *ProcessorCore, i int) {
	q = new(ProcessorCore)
	q.Init(p.R.Len(), p.InstructionSet)
	q.IOController.Open(c)
	p.IOController.Open(c)
	i = p.IOController.Len() - 1
	return
}
func (p *ProcessorCore) DefineInstructions() {
	p.Define("noop",	func (o *Buffer) {})						//	NOOP
	p.Define("nsleep",	func (o *Buffer) {							//	NSLEEP	n
		syscall.Sleep(int64((*o)[0]))
	})
	p.Define("sleep",	func (o *Buffer) {							//	SLEEP	n
		syscall.Sleep(int64((*o)[0]) << 32)
	})
	p.Define("halt",	func (o *Buffer) { p.Running = false })		//	HALT
	p.Define("jmp",		func (o *Buffer) {							//	JMP		n
		p.JumpRelative((*o)[0])
	})
	p.Define("jmpz",	func (o *Buffer) {							//	JMPZ	r, n
		if p.R.Buffer.EqualsZero((*o)[0]) { p.JumpRelative((*o)[1]) }
	})
	p.Define("jmpnz",	func (o *Buffer) {							//	JMPNZ	r, n
		if !p.R.Buffer.EqualsZero((*o)[0]) { p.JumpRelative((*o)[1]) }
	})
	p.Define("call",	func (o *Buffer) {							//	CALL	n
		p.Call((*o)[0])
	})
	p.Define("ret",		func (o *Buffer) { p.Return() })			//	RET
	p.Define("push",	func (o *Buffer) {							//	PUSH	r
		p.data_stack.Push(p.R.At((*o)[0]))
	})
	p.Define("pop",		func (o *Buffer) {							//	POP		r
		p.R.Set((*o)[0], p.data_stack.Pop())
	})
	p.Define("cld",		func (o *Buffer) {							//	CLD		r, v
		p.R.Set((*o)[0], (*o)[1])
	})
	p.Define("send",	func (o *Buffer) {							//	SEND	c
		p.IOController.Send((*o)[0], p.M)
	})
	p.Define("recv",	func (o *Buffer) {							//	RECV	c
		p.M = p.IOController.Receive((*o)[0])
	})
	p.Define("inc",		func (o *Buffer) {							//	INC		r
		p.R.Buffer.Increment((*o)[0])
	})
	p.Define("dec",		func (o *Buffer) {							//	DEC		r
		p.R.Buffer.Decrement((*o)[0])
	})
	p.Define("add",		func (o *Buffer) {							//	ADD		r1, r2
		p.R.Buffer.Add((*o)[0], (*o)[1])
	})
	p.Define("sub",		func (o *Buffer) {							//	SUB		r1, r2
		p.R.Buffer.Subtract((*o)[0], (*o)[1])
	})
	p.Define("mul",		func (o *Buffer) {							//	MUL		r1, r2
		p.R.Buffer.Multiply((*o)[0], (*o)[1])
	})
	p.Define("div",		func (o *Buffer) {							//	DIV		r1, r2
		if p.R.At((*o)[1]) == 0 {
			p.DivideByZero()
		} else {
			p.R.Buffer.Divide((*o)[0], (*o)[1])
		}
	})
	p.Define("and",		func (o *Buffer) {							//	AND		r1, r2
		p.R.Buffer.And((*o)[0], (*o)[1])
	})
	p.Define("or",		func (o *Buffer) {							//	OR		r1, r2
		p.R.Buffer.Or((*o)[0], (*o)[1])
	})
	p.Define("xor",		func (o *Buffer) {							//	XOR		r1, r2
		p.R.Buffer.Xor((*o)[0], (*o)[1])
	})
//	p.Define("malloc",	func (o *Buffer) { p.R.PutBuffer((*o)[0], p.MMU.Allocate((*o)[1])) })				//	MALLOC	r, n
//	p.Define("select",	func (o *Buffer) { p.M = p.R.GetBuffer((*o)[0]) })									//	SELECT	r
}
func (p *ProcessorCore) ValidPC() bool {
	return (p.PC > -1) && (p.PC < len(p.program)) && p.Running
}
func (p *ProcessorCore) Call(location int) {
	p.call_stack.Push(p.PC + 1)
	p.PC = location
	p.LoadInstruction()
	p.Execute()
}
func (p *ProcessorCore) Return() {
	if p.call_stack.Len() > 0 {
		p.PC = p.call_stack.Pop()
		p.LoadInstruction()
		p.Execute()
	} else {
		p.Running = false
		p.Stack_Underflow = true
	}
}
func (p *ProcessorCore) LoadProgram(program *[]*OpCode) {
	p.program = *program
	p.ResetState()
	p.LoadInstruction()
}
func (p *ProcessorCore) LoadInstruction() {
	if p.ValidPC() {
		p.I = p.program[p.PC]
	} else {
		p.Segmentation_Error = true
		p.Running = false
	}
}
func (p *ProcessorCore) DivideByZero() {
	p.Divide_by_Zero = true
	p.Running = false
}
func (p *ProcessorCore) ResetState() {
	p.RegisterBlock.Clear()
	p.ExecutionFlags.Clear()
	p.Running = true
}
func (p *ProcessorCore) StepForward() {
	p.PC++
	p.LoadInstruction()
}
func (p *ProcessorCore) StepBack() {
	p.PC--
	p.LoadInstruction()
}
func (p *ProcessorCore) JumpTo(ops int) {
	p.PC = ops
	p.LoadInstruction()
	p.Execute()
}
func (p *ProcessorCore) JumpRelative(ops int) {
	p.PC = p.PC + ops
	p.LoadInstruction()
	p.Execute()
}
func (p *ProcessorCore) Execute() {
	if !p.Invoke(p.I) {
		p.Running = false
		p.Illegal_Operation = true
	}
}
func (p *ProcessorCore) Run() {
	if p.PC > 0 { p.ResetState() }
	p.LoadInstruction()
	for {
		p.Execute()
		if p.Running {
			p.StepForward()
		} else {
			break
		}
	}
}
func (p *ProcessorCore) RunInline() {
	if p.PC > 0 || !p.ValidPC() { p.ResetState() }
	for {
		if p.ValidPC() {
			p.I = p.program[p.PC]
		} else {
			p.Segmentation_Error = true
			p.Running = false
			break
		}
		if !p.InlinedInstructions() { break }
		p.PC++
	}
}
func (p *ProcessorCore) InlinedInstructions() bool {
	switch p.I.code {
		case 0:
		case 1:			syscall.Sleep(int64(p.I.data[0]))
		case 2:			syscall.Sleep(int64(p.I.data[0]) << 32)
		case 3:			p.Running = false
		case 4:			p.JumpRelative(p.I.data[0])
		case 5:			if p.R.Buffer.EqualsZero(p.I.data[0]) { p.JumpRelative(p.I.data[1]) }
		case 6:			if !p.R.Buffer.EqualsZero(p.I.data[0]) { p.JumpRelative(p.I.data[1]) }
		case 7:			p.Call(p.I.data[0])
		case 8:			p.Return()
		case 9:			p.data_stack.Push(p.R.At(p.I.data[0]))
		case 10:		p.R.Set(p.I.data[0], p.data_stack.Pop())
		case 11:		p.R.Set(p.I.data[0], p.I.data[1])
		case 12:		p.IOController.Send(p.I.data[0], p.M)
		case 13:		p.M = p.IOController.Receive(p.I.data[0])
		case 14:		p.R.Buffer.Increment(p.I.data[0])
		case 15:		p.R.Buffer.Decrement(p.I.data[0])
		case 16:		p.R.Buffer.Add(p.I.data[0], p.I.data[1])
		case 17:		p.R.Buffer.Subtract(p.I.data[0], p.I.data[1])
		case 18:		p.R.Buffer.Multiply(p.I.data[0], p.I.data[1])
		case 19:		d := p.I.data[1]
						if d == 0 {
							p.Divide_by_Zero = true
							p.Running = false
						} else {
							p.R.Buffer.Divide(p.I.data[0], d)
						}
		case 20:		p.R.Buffer.And(p.I.data[0], p.I.data[1])
		case 21:		p.R.Buffer.Or(p.I.data[0], p.I.data[1])
		case 22:		p.R.Buffer.Xor(p.I.data[0], p.I.data[1])
		default:		if !p.Invoke(p.I) {
							p.Running = false
							p.Illegal_Operation = true
							return false
						}
	}
	return true
}
