//	TODO:	Further work on the memory page model and how that interacts with the register model
//	TODO:	Add more support for debugging to ProcessorCore
//	TODO:	Write tests

package golightly

import("container/vector")

type Value interface{}

type MemoryBlock []Value
func (m *MemoryBlock) Copy(destination *MemoryBlock) { for i, x := range m { destination[i] = x } }
func (m *MemoryBlock) Set(v Value) { for i, x := range m { m[i] = x } }
func (m *MemoryBlock) Clear() { m.Set(nil); }

type OpCode struct {
	code			int;
	a, b, c, d		Value;
}
type Program []OpCode

type ExecutionFlags struct {
	running			bool;
}

type ProcessorCore struct {
	PC				int;
	flags			ExecutionFlags;
	program			*Program;
	instructions	vector.Vector;
}
func (p *ProcessorCore) ValidPC() bool { return (p.PC < p.program.Len()) && p.Flags.running }
func (p *ProcessorCore) Load(program *[]OpCode) {
	p.program = program;
	p.PC = 0;
	p.Flags.running = false;
}
func (p *ProcessorCore) FetchInstruction() OpCode {
	if p.ValidPC { return p.program[p.PC] }
	else { return nil }
}
func (p *ProcessorCore) StepForward() { p.PC++ }
func (p *ProcessorCore) StepBack() { p.PC-- }
func (p *ProcessorCore) Init() {
	p.instructions = vector.New(64);
	p.instructions.Push(func (o OpCode) {});										//	NOOP
}
func (p *ProcessorCore) Execute(o OpCode) { p.instructions.At(o.code)(o) }
func (p *ProcessorCore) Run() {
	for i := p.FetchInstruction() {
		p.Execute(i);
		p.StepForward();
	}
}


type Processor struct {
	ProcessorCore;
	register_count	int;
	R				MemoryBlock;			//	registers
	MP				*MemoryBlock;			//	memory pointer
	call_stack		vector.Vector;
	data_stack		vector.Vector;
	features		map[string] bool;
}
func (p *Processor) AllocateMemory(n int) *MemoryBlock { return make(MemoryBlock, n) }
func (p *Processor) AllocateRegisters() { p.R = make(MemoryBlock, p.register_count) }
func (p *Processor) ResetRegisters() { p.R.Clear() } }
func (p *Processor) PushRegisters() {
	p.call_stack.Push(p.registers);
	p.call_stack.Push(p.PC);
	p.call_stack.Push(p.MP);
	p.AllocateRegisters();
}
func (p *Processor) PopRegisters() {
	if p.call_stack.Len() > 0 {
		p.MP = p.call_stack.Pop();
		p.PC = p.call_stack.Pop();
		p.R = p.call_stack.Pop();
	} else { p.Flags.running = false }
}
func (p *Processor) SupportFlowControl() {
	p.instructions.Push(func (o OpCode) { p.PC += o.a });							//	JUMP	n
	p.instructions.Push(func (o OpCode) { p.PC += p.R[o.a] });						//	IJUMP	r
	p.instructions.Push(func (o OpCode) { if p.R[o.a] == 0 { p.PC += o.b } });		//	ZJUMP	r, n
	p.instructions.Push(func (o OpCode) { p.PushRegisters(); p.PC = o.a });			//	CALL	n
	p.instructions.Push(func (o OpCode) { p.PushRegisters(); p.PC = R[o.a] });		//	ICALL	r
	p.instructions.Push(func (o OpCode) { p.PopRegisters() });						//	RET
	p.features["FlowControl"] = true;
}
func (p *Processor) SupportALU() {
	p.instructions.Push(func (o OpCode) { p.R[o.a]++ });							//	INC		r
	p.instructions.Push(func (o OpCode) { p.R[o.a]-- });							//	DEC		r
	p.instructions.Push(func (o OpCode) { p.R[o.a] = p.R[o.b] });					//	LD		r1, r2
	p.instructions.Push(func (o OpCode) { p.R[o.a] = o.b });						//	CLD		r, v
	p.instructions.Push(func (o OpCode) { p.R[o.a] += p.R[o.b] });					//	ADD		r1, r2
	p.instructions.Push(func (o OpCode) { p.R[o.a] -= p.R[o.b] });					//	SUB		r1, r2
	p.instructions.Push(func (o OpCode) { p.R[o.a] *= p.R[o.b] });					//	MUL		r1, r2
	p.instructions.Push(func (o OpCode) { p.R[o.a] /= p.R[o.b] });					//	DIV		r1, r2
	p.instructions.Push(func (o OpCode) { p.R[o.a] &= p.R[o.b] });					//	AND		r1, r2
	p.instructions.Push(func (o OpCode) { p.R[o.a] |= p.R[o.b] });					//	OR		r1, r2
	p.instructions.Push(func (o OpCode) { p.R[o.a] ^= p.R[o.b] });					//	XOR		r1, r2
	p.instructions.Push(func (o OpCode) { p.R[o.a] += o.b) });						//	CADD	r, v
	p.instructions.Push(func (o OpCode) { p.R[o.a] -= o.b });						//	CSUB	r, v
	p.instructions.Push(func (o OpCode) { p.R[o.a] *= o.b });						//	CMUL	r, v
	p.instructions.Push(func (o OpCode) { p.R[o.a] /= o.b });						//	CDIV	r, v
	p.instructions.Push(func (o OpCode) { p.R[o.a] &= o.b });						//	CAND	r, v
	p.instructions.Push(func (o OpCode) { p.R[o.a] |= o.b });						//	COR		r, v
	p.instructions.Push(func (o OpCode) { p.R[o.a] ^= o.b });						//	CXOR	r, v
	p.features["ALU"] = true;
}

func (p *Processor) SupportMemoryPaging() {
	p.instructions.Push(func (o OpCode) { p.R[o.a] = p.AllocateMemory(o.b) });		//	MALLOC	r, n
	p.instructions.Push(func (o OpCode) { p.MP = p.R[o.a] });						//	SELECT	r
	p.instructions.Push(func (o OpCode) { p.MP = o.a });							//	PSELECT	p	
	p.instructions.Push(func (o OpCode) { p.R[o.a] = p.R[p.MP[o.b]]) });			//	ILD		r1, r2
	p.instructions.Push(func (o OpCode) { p.MP[p.R[o.b]] = p.R[o.a] });				//	ISTORE	r, m
	p.instructions.Push(func (o OpCode) { p.R[o.a] += p.MP[o.b] });					//	IADD	r, m
	p.instructions.Push(func (o OpCode) { p.R[o.a] -= p.MP[o.b] });					//	ISUB	r, m
	p.instructions.Push(func (o OpCode) { p.R[o.a] *= p.MP[o.b] });					//	IMUL	r, m
	p.instructions.Push(func (o OpCode) { p.R[o.a] /= p.MP[o.b] });					//	IDIV	r, m
	p.instructions.Push(func (o OpCode) { p.R[o.a] &= p.MP[o.b] });					//	IAND	r, m
	p.instructions.Push(func (o OpCode) { p.R[o.a] |= p.MP[o.b] });					//	IOR		r, m
	p.instructions.Push(func (o OpCode) { p.R[o.a] ^= p.MP[o.b] });					//	IXOR	r, m
	p.features["MemoryPaging"] = true;
}
func (p *Processor) SupportDataStack() {
	p.instructions.Push(func (o OpCode) { p.data_stack.Push(p.R[o.a]) });			//	PUSH	r
	p.instructions.Push(func (o OpCode) { p.data_stack.Push(o.a) });				//	CPUSH	v
	p.instructions.Push(func (o OpCode) { p.data_stack.Push(p.MP[o.a]) });			//	IPUSH	m
	p.instructions.Push(func (o OpCode) { p.R[o.a] = p.data_stack.Pop() });			//	POP		r
	p.instructions.Push(func (o OpCode) { p.MP[o.a] = p.data_stack.Pop() });		//	IPOP	m
	p.features["DataStack"] = true;
}
func (p *Processor) Init() {
	p.ProcessorCore.Init();
}

func NewProcessor(register_count uint) *Processor {
	p := Processor{call_stack: vector.New(16)};
	p.Init();
	p.SupportFlowControl();
	p.SupportALU();
	p.SupportMemoryPaging();
	p.SupportDataStack();
	return &p;
}