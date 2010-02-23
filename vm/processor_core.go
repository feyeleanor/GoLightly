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
	R				*Stream
	M				*Stream
}
func (r *RegisterBlock) Allocate(count int) {
	r.PC = 0
	r.I = nil
	r.R = new(Stream)
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
	r.R.Clear()
	r.M = nil
}

type MMU struct {}
func (m *MMU) Allocate(words int) *Stream {
	s := new(Stream)
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
func (p *ProcessorCore) Clone(c chan *Stream) (q *ProcessorCore, i int) {
	q = new(ProcessorCore)
	q.Init(p.R.Len(), p.InstructionSet)
	q.IOController.Open(c)
	p.IOController.Open(c)
	i = p.IOController.Len() - 1
	return
}
func (p *ProcessorCore) DefineInstructions() {
	p.Define("noop",	func (o *OpCode)	{})																//	NOOP
	p.Define("nsleep",	func (o *OpCode)	{ syscall.Sleep(int64(o.a)) })									//	NSLEEP	n
	p.Define("sleep",	func (o *OpCode)	{ syscall.Sleep(int64(o.a) << 32) })							//	SLEEP	n
	p.Define("halt",	func (o *OpCode)	{ p.Running = false })											//	HALT
	p.Define("jmp",		func (o *OpCode)	{ p.JumpRelative(o.a) })										//	JMP		n
	p.Define("jmpz",	func (o *OpCode)	{ if p.R.Buffer.EqualsZero(o.a) { p.JumpRelative(o.b) } })		//	JMPZ	r, n
	p.Define("jmpnz",	func (o *OpCode)	{ if !p.R.Buffer.EqualsZero(o.a) { p.JumpRelative(o.b) } })		//	JMPNZ	r, n
	p.Define("call",	func (o *OpCode)	{ p.Call(o.a) })												//	CALL	n
	p.Define("ret",		func (o *OpCode)	{ p.Return() })													//	RET
	p.Define("push",	func (o *OpCode)	{ p.data_stack.Push(p.R.At(o.a)) })								//	PUSH	r
	p.Define("pop",		func (o *OpCode)	{ p.R.Set(o.a, p.data_stack.Pop()) })							//	POP		r
	p.Define("cld",		func (o *OpCode)	{ p.R.Set(o.a, o.b) })											//	CLD		r, v
	p.Define("send",	func (o *OpCode)	{ p.IOController.Send(o.a, p.M) })								//	SEND	c
	p.Define("recv",	func (o *OpCode)	{ p.M = p.IOController.Receive(o.a) })							//	RECV	c
	p.Define("inc",		func (o *OpCode)	{ p.R.Buffer.Increment(o.a) })									//	INC		r
	p.Define("dec",		func (o *OpCode)	{ p.R.Buffer.Decrement(o.a) })									//	DEC		r
	p.Define("add",		func (o *OpCode)	{ p.R.Buffer.Add(o.a, o.b) })									//	ADD		r1, r2
	p.Define("sub",		func (o *OpCode)	{ p.R.Buffer.Subtract(o.a, o.b) })								//	SUB		r1, r2
	p.Define("mul",		func (o *OpCode)	{ p.R.Buffer.Multiply(o.a, o.b) })								//	MUL		r1, r2
	p.Define("div",		func (o *OpCode)	{																//	DIV		r1, r2
		if p.R.At(o.b) == 0 {
			p.DivideByZero()
		} else {
			p.R.Buffer.Divide(o.a, o.b)
		}
	})
	p.Define("and",		func (o *OpCode)	{ p.R.Buffer.And(o.a, o.b) })									//	AND		r1, r2
	p.Define("or",		func (o *OpCode)	{ p.R.Buffer.Or(o.a, o.b) })									//	OR		r1, r2
	p.Define("xor",		func (o *OpCode)	{ p.R.Buffer.Xor(o.a, o.b) })									//	XOR		r1, r2
//	p.Define("iadd",	func (o *OpCode)	{ p.R.Buffer.Add(o.a, p.M.At(o.b)) })							//	IADD	r, m
//	p.Define("isub",	func (o *OpCode)	{ p.R.Buffer.Subtract(o.a, p.M.At(o.b)) })						//	ISUB	r, m
//	p.Define("imul",	func (o *OpCode)	{ p.R.Buffer.Multiply(o.a, p.M.At(o.b)) })						//	IMUL	r, m
//	p.Define("idiv",	func (o *OpCode)	{																//	IDIV	r, m
//		if p.R.At(o.b) == 0 {
//			p.DivideByZero()
//		} else {
//			p.R.Buffer.Divide(o.a, p.M.At(o.b))
//		}
//	})
//	p.Define("iand",	func (o *OpCode)	{ p.R.Buffer.And(o.a, p.M.At(o.b)) })							//	IAND	r, m
//	p.Define("ior",		func (o *OpCode)	{ p.R.Buffer.Or(o.a, p.M.At(o.b)) })							//	IOR		r, m
//	p.Define("ixor",	func (o *OpCode)	{ p.R.Buffer.Xor(o.a, p.M.At(o.b)) })							//	IXOR	r, m
//	p.Define("malloc",	func (o *OpCode)	{ p.R.PutBuffer(o.a, p.MMU.Allocate(o.b)) })					//	MALLOC	r, n
//	p.Define("select",	func (o *OpCode)	{ p.M = p.R.GetBuffer(o.a) })									//	SELECT	r
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
		case 1:			syscall.Sleep(int64(p.I.a))
		case 2:			syscall.Sleep(int64(p.I.a) << 32)
		case 3:			p.Running = false
		case 4:			p.JumpRelative(p.I.a)
		case 5:			if p.R.Buffer.EqualsZero(p.I.a) { p.JumpRelative(p.I.b) }
		case 6:			if !p.R.Buffer.EqualsZero(p.I.a) { p.JumpRelative(p.I.b) }
		case 7:			p.Call(p.I.a)
		case 8:			p.Return()
		case 9:			p.data_stack.Push(p.R.At(p.I.a))
		case 10:		p.R.Set(p.I.a, p.data_stack.Pop())
		case 11:		p.R.Set(p.I.a, p.I.b)
		case 12:		p.IOController.Send(p.I.a, p.M)
		case 13:		p.M = p.IOController.Receive(p.I.a)
		case 14:		p.R.Buffer.Increment(p.I.a)
		case 15:		p.R.Buffer.Decrement(p.I.a)
		case 16:		p.R.Buffer.Add(p.I.a, p.I.b)
		case 17:		p.R.Buffer.Subtract(p.I.a, p.I.b)
		case 18:		p.R.Buffer.Multiply(p.I.a, p.I.b)
		case 19:		d := p.I.b
						if d == 0 {
							p.Divide_by_Zero = true
							p.Running = false
						} else {
							p.R.Buffer.Divide(p.I.a, d)
						}
		case 20:		p.R.Buffer.And(p.I.a, p.I.b)
		case 21:		p.R.Buffer.Or(p.I.a, p.I.b)
		case 22:		p.R.Buffer.Xor(p.I.a, p.I.b)
//		case 23:		p.R.Buffer.Add(p.I.a, p.M.At(p.I.b))
//		case 24:		p.R.Buffer.Subtract(p.I.a, p.M.At(p.I.b))
//		case 25:		p.R.Buffer.Multiply(p.I.a, p.M.At(p.I.b))
//		case 26:		d := p.M.At(p.I.b)
//						if d == 0 {
//							p.Divide_by_Zero = true
//							p.Running = false
//						} else {
//							p.R.Buffer.Divide(p.I.a, d)
//						}
//		case 27:		p.R.Buffer.And(p.I.a, p.M.At(p.I.b))
//		case 28:		p.R.Buffer.Or(p.I.a, p.M.At(p.I.b))
//		case 29:		p.R.Buffer.Xor(p.I.a, p.M.At(p.I.b))
		default:		if !p.Invoke(p.I) {
							p.Running = false
							p.Illegal_Operation = true
							return false
						}
	}
	return true
}
