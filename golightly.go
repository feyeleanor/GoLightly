//	TODO:	Further work on the memory page model and how that interacts with the register model
//	TODO:	Add more support for debugging to ProcessorCore
//	TODO:	Consider alternatives for VLIW
//				1	pool of channels all working against current processor state
//				2	duplicated processors each feeding back results
//				synchoronous or asynchronous operation?

package golightly

import . "golightly/vm"

type Processor struct {
	ProcessorCore
	data_stack []int
}

func (p *Processor) Init(registers int) {
	p.ProcessorCore.Init(registers, new(InstructionSet))

	p.Movement("ijump", "r", func(o *OpCode) {
		p.PC = p.R.At(o.Data.(int)).(int)
	})

	p.Movement("zjump", "r, n", func(o *OpCode) {
		if p.R.Equal(o.a, 0) {
			p.PC = o.b
		}
	})

	p.Movement("icall", "r", func(o *OpCode) {
		p.Call(p.R.At(o.a).(int))
	})

	p.Operator("ld", "r1, r2", func(o *OpCode) {
		p.R.Copy(o.a, o.b)
	})

	p.Operator("push", "r", func(o *OpCode) {
		p.data_stack = append(p.data_stack, p.R[o.ia])
	})

	p.Operator("cpush", "v", func(o *OpCode) {
		p.data_stack.Append(o.ia)
	})

	p.Operator("ipush", "m", func(o *OpCode) {
		p.data_stack.Append(p.MP[o.ia])
	})

	p.Operator("pop", "r", func(o *OpCode) {
		p.R.Set(o.a, p.data_stack.Pop())
	})

	p.Operator("ipop", "m", func(o *OpCode) {
		p.MP.Set(o.a, p.data_stack.Pop())
	})

	/*
		p.Operator("pselect", "p", func(o *OpCode) {
			p.MP = IntBuffer(o.a)
		})
	*/

	p.Operator("ild", "r1, r2", func(o *OpCode) {
		p.R.Set(o.a, p.R.At(p.MP.At(o.b)))
	})

	p.Operator("istore", "r, m", func(o *OpCode) {
		p.MP.Set(p.R.At(o.b), p.R.At(o.a))
	})
}
