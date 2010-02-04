//	TODO:	Further work on the memory page model and how that interacts with the register model
//	TODO:	Add more support for debugging to ProcessorCore

package golightly

import(
	"feyeleanor/golightly/vm";
	"container/vector";
)

type Processor struct {
	ProcessorCore;
	data_stack		vector.IntVector;
}
func (p *Processor) Init(registers int) {
	p.ProcessorCore.Init(registers)
	p.Define("ijump",	func (o *OpCode) { p.Jump(p.R.At(o.ia)) })											//	IJUMP	r
	p.Define("zjump",	func (o *OpCode) { if p.R.Equal(o.a, 0) { p.Jump(o.b) } })							//	ZJUMP	r, n
	p.Define("icall",	func (o *OpCode) { p.Call(p.R.At(o.a)) })											//	ICALL	r
	p.Define("ld",		func (o *OpCode) { p.R.Copy(o.a, o.b) })											//	LD		r1, r2
	p.Define("push",	func (o *OpCode) { p.data_stack.Push(p.R[o.ia]) })									//	PUSH	r
	p.Define("cpush",	func (o *OpCode) { p.data_stack.Push(o.ia) })										//	CPUSH	v
	p.Define("ipush",	func (o *OpCode) { p.data_stack.Push(p.MP[o.ia]) })									//	IPUSH	m
	p.Define("pop",		func (o *OpCode) { p.R.Set(o.a, p.data_stack.Pop()) })								//	POP		r
	p.Define("ipop",	func (o *OpCode) { p.MP.Set(o.a, p.data_stack.Pop()) })								//	IPOP	m
//	p.Define("pselect",	func (o *OpCode) { p.MP = Buffer(o.a) })											//	PSELECT	p	
	p.Define("ild",		func (o *OpCode) { p.R.Set(o.a, p.R.At(p.MP.At(o.b))) })							//	ILD		r1, r2
	p.Define("istore",	func (o *OpCode) { p.MP.Set(p.R.At(o.b), p.R.At(o.a)) })							//	ISTORE	r, m
}