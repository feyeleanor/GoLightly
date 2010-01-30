package vm

import "container/vector"

type ExecutionFlags struct {
	running				bool
	illegal_operation	bool
	stack_underflow		bool
	segmentation_error	bool
}
func (f *ExecutionFlags) Clear() {
	f.running = false
	f.illegal_operation = false
	f.stack_underflow = false
	f.segmentation_error = false
}

type RegisterBlock struct {
	PC				int;
	I				*OpCode;
	R				*Buffer;
	M				*Buffer;
}
func (r *RegisterBlock) Allocate(count int) {
	r.PC = 0
	r.I = nil
	r.R = new(Buffer)
	r.R.Init(count)
	r.M = nil
}
func (r *RegisterBlock) Clone() *RegisterBlock {
	return &RegisterBlock{ PC: r.PC, I: r.I, M: r.M, R: r.R.Clone() }
}
func (r *RegisterBlock) Replace(a *RegisterBlock) {
	r.PC, r.I, r.R, r.M = a.PC, a.I, a.R, a.M
}

type MMU struct {}
func (m *MMU) Allocate(words int) *Buffer {
	b := new(Buffer)
	b.Init(words)
	return b
}

type ProcessorCore struct {
	ExecutionFlags;
	InstructionSet;
	MMU;
	RegisterBlock;
	program			[]OpCode;
	call_stack		vector.Vector;
}
func (p *ProcessorCore) Init(registers int) {
	p.InstructionSet.Init()
	p.ExecutionFlags.Clear()
	p.RegisterBlock.Allocate(registers)
	p.Define("jump",	func (o *OpCode)	{ p.Jump(o.a) })															//	JUMP	n
	p.Define("call",	func (o *OpCode)	{ p.Call(o.a) })															//	CALL	n
	p.Define("ret",		func (o *OpCode)	{ p.Return() })																//	RET
	p.Define("cld",		func (o *OpCode)	{ p.R.Set(o.a, o.b) })														//	CLD		r, v
	p.Define("inc",		func (o *OpCode)	{ p.R.Increment(o.a) })														//	INC		r
	p.Define("dec",		func (o *OpCode)	{ p.R.Decrement(o.a) });													//	DEC		r
	p.Define("cadd",	func (o *OpCode)	{ p.R.Add(o.a, o.b) });														//	CADD	r, v
	p.Define("csub",	func (o *OpCode)	{ p.R.Subtract(o.a, o.b) });												//	CSUB	r, v
	p.Define("cmul",	func (o *OpCode)	{ p.R.Multiply(o.a, o.b) });												//	CMUL	r, v
	p.Define("cdiv",	func (o *OpCode)	{ p.R.Divide(o.a, o.b) });													//	CDIV	r, v
	p.Define("cand",	func (o *OpCode)	{ p.R.And(o.a, o.b) });														//	CAND	r, v
	p.Define("cor",		func (o *OpCode)	{ p.R.Or(o.a, o.b) });														//	COR		r, v
	p.Define("cxor",	func (o *OpCode)	{ p.R.Xor(o.a, o.b) });														//	CXOR	r, v
	p.Define("iadd",	func (o *OpCode)	{ p.R.Add(o.a, p.M.At(o.b)) })												//	IADD	r, m
	p.Define("isub",	func (o *OpCode)	{ p.R.Subtract(o.a, p.M.At(o.b)) })											//	ISUB	r, m
	p.Define("imul",	func (o *OpCode)	{ p.R.Multiply(o.a, p.M.At(o.b)) })											//	IMUL	r, m
	p.Define("idiv",	func (o *OpCode)	{ p.R.Divide(o.a, p.M.At(o.b)) })											//	IDIV	r, m
	p.Define("iand",	func (o *OpCode)	{ p.R.And(o.a, p.M.At(o.b)) })												//	IAND	r, m
	p.Define("ior",		func (o *OpCode)	{ p.R.Or(o.a, p.M.At(o.b)) })												//	IOR		r, m
	p.Define("ixor",	func (o *OpCode)	{ p.R.Xor(o.a, p.M.At(o.b)) })												//	IXOR	r, m
	p.Define("malloc",	func (o *OpCode)	{ p.R.PutBuffer(o.a, p.MMU.Allocate(o.b)) })								//	MALLOC	r, n
	p.Define("select",	func (o *OpCode)	{ p.M = p.R.GetBuffer(o.a) })												//	SELECT	r
}
func (p *ProcessorCore) ValidPC() bool {
	return (p.PC < len(p.program)) && p.running
}
func (p *ProcessorCore) Call(location int) {
	p.call_stack.Push(p.RegisterBlock.Clone());
	p.RegisterBlock.Allocate(p.R.Len());
	p.PC = location;
}
func (p *ProcessorCore) Return() {
	if r := p.call_stack.Pop().(*RegisterBlock); r != nil {
		p.RegisterBlock.Replace(r)
	} else {
		p.running = false
		p.stack_underflow = true
	}
}
func (p ProcessorCore) LoadProgram(program *[]OpCode) {
	p.program = *program;
	p.PC = 0;
	p.ExecutionFlags.Clear();
}
func (p ProcessorCore) LoadInstruction() {
	if p.ValidPC() {
		p.I = &p.program[p.PC];
	} else {
		p.segmentation_error = true
		p.running = false
	}
}
func (p ProcessorCore) StepForward() {
	p.PC++
}
func (p ProcessorCore) StepBack() {
	p.PC--
}
func (p ProcessorCore) Jump(ops int) {
	p.PC += ops
}
func (p ProcessorCore) Execute() {
	if !p.InstructionSet.Invoke(p.I) {
		p.running = false
		p.illegal_operation = true
	}
}
func (p ProcessorCore) Run() {
	for {
		p.LoadInstruction()
		if p.running {
			p.Execute()
			p.StepForward()
		} else {
			break
		}
	}
}
