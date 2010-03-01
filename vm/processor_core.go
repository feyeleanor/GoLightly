//	TODO:	storing and retrieving pointers to memory buffers
//	TODO:	cloning should create a comms channel by which the parent and child cores can communicate
//	TODO:	should always have stdin, stdout and stderr channels

package vm

import "container/vector"
import "syscall"

type ExecutionFlags struct {
	Running				bool
	Illegal_Operation	bool
	Stack_Underflow		bool
	Segmentation_Error	bool
	Divide_by_Zero		bool
}
func (f *ExecutionFlags) Clear() {
	*f = ExecutionFlags{}
}

type ProcessorCore struct {
	ExecutionFlags
	*InstructionSet
	IOController
	R				*Vector
	M				*Vector
	CS				vector.IntVector
	DS				vector.IntVector
	*Program
}
func (p *ProcessorCore) Init(registers int, instructions *InstructionSet) {
	p.R = &Vector{make([]int, registers)}
	p.M = nil
	p.IOController.Init()
	if instructions == nil {
		p.InstructionSet = new(InstructionSet)
		p.InstructionSet.Init()
		p.DefineInstructions()
	} else {
		p.InstructionSet = instructions
	}
	p.ExecutionFlags.Clear()
	p.Program = new(Program)
}
func (p *ProcessorCore) ValidPC() bool {
	return p.Running && p.Program.ValidPC()
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
	p.Define("call",	func (o *Buffer) { p.Call((*o)[0]) })		//	CALL	n
	p.Define("ret",		func (o *Buffer) { p.Return() })			//	RET
	p.Define("push",	func (o *Buffer) {							//	PUSH	r
		p.DS.Push(p.R.At((*o)[0]))
	})
	p.Define("pop",		func (o *Buffer) {							//	POP		r
		p.R.Set((*o)[0], p.DS.Pop())
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
}
func (p *ProcessorCore) JumpTo(ops int) {
	p.pc = ops - 1
}
func (p *ProcessorCore) JumpRelative(ops int) {
	p.pc += ops - 1
}
func (p *ProcessorCore) Call(location int) {
	p.CS.Push(p.pc)
	p.pc = location - 1
}
func (p *ProcessorCore) Return() {
	if p.CS.Len() > 0 {
		p.pc = p.CS.Pop()
	} else {
		p.Running = false
		p.Stack_Underflow = true
	}
}
func (p *ProcessorCore) LoadProgram(program *Program) {
	p.Program = program
	p.R.ClearAll()
	p.M = nil
	p.ExecutionFlags.Clear()
	p.pc = 0
}
func (p *ProcessorCore) DivideByZero() {
	p.Divide_by_Zero = true
	p.Running = false
}
func (p *ProcessorCore) ResetState() {
	p.R.ClearAll()
	p.M = nil
	p.ExecutionFlags.Clear()
	p.pc = 0
}
func (p *ProcessorCore) Run() {
	p.Running = true
	for {
		if -1 < p.pc && p.pc < len(p.code) {
			if !p.Invoke(p.code[p.pc]) {
				p.Running = false
				p.Illegal_Operation = true
			}
		} else {
			p.Running = false
		}
		if !p.Running { break }
		p.pc += 1
	}
}
func (p *ProcessorCore) RunInline() {
	p.Running = true
	for {
		if -1 < p.pc && p.pc < len(p.code) {
			p.InlinedInstructions(p.code[p.pc])
		} else {
			p.Running = false
		}
		if !p.Running { break }
		p.pc += 1
	}
}
func (p *ProcessorCore) Execute() {
	if 0 <= p.pc && p.pc < len(p.code) {
		if !p.Invoke(p.code[p.pc]) {
			p.Running = false
			p.Illegal_Operation = true
		}
	} else {
		p.Running = false
	}
}
func (p *ProcessorCore) ExecuteInline() {
	if -1 < p.pc && p.pc < len(p.code) {
		p.InlinedInstructions(p.code[p.pc])
	} else {
		p.Running = false
	}
}
func (p *ProcessorCore) InlinedInstructions(o *OpCode) {
	switch o.code {
		case 0:
		case 1:			syscall.Sleep(int64(o.data[0]))
		case 2:			syscall.Sleep(int64(o.data[0]) << 32)
		case 3:			p.Running = false
		case 4:			p.pc += o.data[0] - 1
		case 5:			if p.R.Buffer.EqualsZero(o.data[0]) { p.pc += o.data[1] - 1 }
		case 6:			if !p.R.Buffer.EqualsZero(o.data[0]) { p.pc += o.data[1] - 1 }
		case 7:			p.CS.Push(p.pc)
						p.pc = o.data[0] - 1
		case 8:			if p.CS.Len() > 0 {
							p.pc = p.CS.Pop()
						} else {
							p.Running = false
							p.Stack_Underflow = true
						}
		case 9:			p.DS.Push(p.R.At(o.data[0]))
		case 10:		p.R.Set(o.data[0], p.DS.Pop())
		case 11:		p.R.Set(o.data[0], o.data[1])
		case 12:		p.IOController.Send(o.data[0], p.M)
		case 13:		p.M = p.IOController.Receive(o.data[0])
		case 14:		p.R.Buffer.Increment(o.data[0])
		case 15:		p.R.Buffer.Decrement(o.data[0])
		case 16:		p.R.Buffer.Add(o.data[0], o.data[1])
		case 17:		p.R.Buffer.Subtract(o.data[0], o.data[1])
		case 18:		p.R.Buffer.Multiply(o.data[0], o.data[1])
		case 19:		d := o.data[1]
						if d == 0 {
							p.Divide_by_Zero = true
							p.Running = false
						} else {
							p.R.Buffer.Divide(o.data[0], d)
						}
		case 20:		p.R.Buffer.And(o.data[0], o.data[1])
		case 21:		p.R.Buffer.Or(o.data[0], o.data[1])
		case 22:		p.R.Buffer.Xor(o.data[0], o.data[1])
		default:		if !p.Invoke(o) {
							p.Running = false
							p.Illegal_Operation = true
						}
	}
}
