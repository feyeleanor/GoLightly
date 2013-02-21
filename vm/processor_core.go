//	TODO:	storing and retrieving pointers to memory buffers
//	TODO:	cloning should create a comms channel by which the parent and child cores can communicate
//	TODO:	should always have stdin, stdout and stderr channels

package vm

import "github.com/feyeleanor/slices"
import "time"

type IsExecutable interface {
	Map(f interface{}) interface{}
	Reduce(f interface{}) interface{}
}

type Thread struct {
	Running			bool
	R				slices.ISlice
	M				slices.ISlice
	CS				*slices.ISlice
	DS				slices.ISlice
	PC				int
	Program
}
func (t *Thread) I() OpCode							{ return t.Program[t.PC] }
func (t *Thread) ValidPC() bool						{ return t.PC > -1 && t.PC < len(t.Program) }
func (t *Thread) Call(location int) {
	t.CS.Append(t.PC)
	t.PC = location
}
func (t *Thread) Return() {
	if address, ok := t.CS.Pop(); ok {
		t.PC = address + 1
	} else {
		panic(t)
	}
}


type ProcessorCore struct {
	*InstructionSet
	IOController
	Thread
}
func (p *ProcessorCore) Init(registers int, instructions *InstructionSet) {
	p.Thread = Thread{ R: make(slices.ISlice, registers) }
	if instructions == nil {
		p.InstructionSet = new(InstructionSet)
		p.InstructionSet.Init()
		p.DefineInstructions()
	} else {
		p.InstructionSet = instructions
	}
}

//	Make a copy of the current processor, binding it to the current processor with
//	the supplied io channel
func (p *ProcessorCore) Clone(c chan slices.ISlice) (q *ProcessorCore, i int) {
	q = new(ProcessorCore)
	q.Init(len(p.R), p.InstructionSet)
	q.IOController = append(q.IOController, c)
	p.IOController = append(p.IOController, c)
	i = len(p.IOController) - 1
	return
}
func (p *ProcessorCore) DefineInstructions() {
	p.Operator("noop",		func () {})																							//	NOOP
	p.Operator("nsleep",	func (n time.Duration) { time.Sleep(n) })																//	NSLEEP	n
	p.Operator("sleep",		func (n time.Duration) { time.Sleep(n << 32) })															//	SLEEP	n
	p.Movement("halt",		func () { p.Running = false })																		//	HALT
	p.Movement("jmp",		func (n int) { p.PC += n })																			//	JMP		n
	p.Movement("jmpz",		func (o []int) { if p.R.ZeroSameAs(o[0]) { p.PC += o[1] } })										//	JMPZ	r, n
	p.Movement("jmpnz",		func (o []int) { if !p.R.ZeroSameAs(o[0]) { p.PC += o[1] } })										//	JMPNZ	r, n
	p.Movement("call",		func (n int) { p.Call(n) })																			//	CALL	n
	p.Movement("ret",		func () { p.Return() })																				//	RET
	p.Operator("push",		func (r int) { p.DS.Append(p.R[r]) })																	//	PUSH	r
	p.Operator("pop",		func (r int) { p.R[r], _ = p.DS.Pop() })																//	POP		r
	p.Operator("cld",		func (o []int) { p.R[o[0]] = o[1] })																//	CLD		r, v
	p.Operator("send",		func (c int) { p.IOController.Send(c, p.M) })														//	SEND	c
	p.Operator("recv",		func (c int) { p.M = p.IOController.Receive(c) })													//	RECV	c
	p.Operator("inc",		func (r int) { p.R.Increment(r) })																	//	INC		r
	p.Operator("dec",		func (r int) { p.R.Decrement(r) })																	//	DEC		r
	p.Operator("add",		func (o []int) { p.R.Add(o[0], o[1]) })																//	ADD		r1, r2
	p.Operator("sub",		func (o []int) { p.R.Subtract(o[0], o[1]) })														//	SUB		r1, r2
	p.Operator("mul",		func (o []int) { p.R.Multiply(o[0], o[1]) })														//	MUL		r1, r2
	p.Operator("div",		func (o []int) { p.R.Divide(o[0], o[1]) })															//	DIV		r1, r2
	p.Operator("and",		func (o []int) { p.R.And(o[0], o[1]) })																//	AND		r1, r2
	p.Operator("or",		func (o []int) { p.R.Or(o[0], o[1]) })																//	OR		r1, r2
	p.Operator("xor",		func (o []int) { p.R.Xor(o[0], o[1]) })																//	XOR		r1, r2
}
func (p *ProcessorCore) LoadProgram(program Program) {
	p.Program = make(Program, len(program))
	copy(p.Program, program)
	slices.ClearAll(p.R)
	p.M = nil
	p.PC = 0
}
func (p *ProcessorCore) ResetState() {
	slices.ClearAll(p.R)
	p.M = nil
	p.PC = 0
}
func (p *ProcessorCore) Execute() {
	o := p.Program[p.PC]
	switch data := o.data.(type) {
	case int:
		p.ops[o.code].(func (int))(data)
	case []int:
		p.ops[o.code].(func ([]int))(data)
	default:
		panic(nil)
	}
	p.PC += o.movement
}
func (p *ProcessorCore) Run() {
	defer func() {
		if x := recover(); x != nil {
			p.Running = false
		}
	}()
	p.Running = true
	for p.Running {
		p.Execute()
	}
}

type InlinedProcessorCore struct {
	ProcessorCore
}
func (p *InlinedProcessorCore) Execute() {
	o := p.Program[p.PC]
	switch data := o.data.(type) {
	case int:
		switch o.code {
		case 4:			p.PC += data
		case 7:			p.CS.Append(p.PC)
						p.PC = data
		case 9:			p.DS.Append(p.R[data])
		case 10:		p.R[data], _ = p.DS.Pop()
		case 12:		p.IOController.Send(data, p.M)
		case 13:		p.M = p.IOController.Receive(data)
		case 14:		p.R.Increment(data)
		case 15:		p.R.Decrement(data)
		default:		p.ProcessorCore.Execute()
		}
	case time.Duration:
		switch o.code {
		case 1:			time.Sleep(data)
		case 2:			time.Sleep(data << 32)
		default:		p.ProcessorCore.Execute()
		}
	case []int:
		switch o.code {
		case 5:			if p.R.ZeroSameAs(data[0]) { p.PC += data[1] } else { p.PC++ }
		case 6:			if !p.R.ZeroSameAs(data[0]) { p.PC += data[1] } else { p.PC++ }
		case 11:		p.R[data[0]] = data[1]
		case 16:		p.R.Add(data[0], data[1])
		case 17:		p.R.Subtract(data[0], data[1])
		case 18:		p.R.Multiply(data[0], data[1])
		case 19:		p.R.Divide(data[0], data[1])
		case 20:		p.R.And(data[0], data[1])
		case 21:		p.R.Or(data[0], data[1])
		case 22:		p.R.Xor(data[0], data[1])
		default:		p.ProcessorCore.Execute()
		}
	default:
		switch o.code {
		case 0:
		case 3:			p.Running = false
		case 8:			p.PC, _ = p.CS.Pop(); p.PC++
		default:		p.ProcessorCore.Execute()
		}
	}
	p.PC += o.movement
}
func (p *InlinedProcessorCore) Run() {
	defer func() {
		if x := recover(); x != nil {
			p.Running = false
		}
	}()
	p.Running = true
	for p.Running {
		p.Execute()
	}
}