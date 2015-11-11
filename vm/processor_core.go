//	TODO:	storing and retrieving pointers to memory buffers
//	TODO:	cloning should create a comms channel by which the parent and child cores can communicate
//	TODO:	should always have stdin, stdout and stderr channels

package vm

import (
	"github.com/feyeleanor/slices"
	"time"
)

type IsExecutable interface {
	Map(f interface{}) interface{}
	Reduce(f interface{}) interface{}
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
	p.Operator("noop", "", func () {})

	p.Operator("nsleep", "n",	func (n time.Duration) {
		time.Sleep(n)
	})

	p.Operator("sleep", "n", func (n time.Duration) {
		time.Sleep(n << 32)
	})

	p.Movement("halt", "", func () {
		p.Running = false
	})

	p.Movement("jmp", "n", func (n int) {
		p.PC += n
	})

	p.Movement("jmpz", "r, n", func (o []int) {
		if p.R.ZeroSameAs(o[0]) {
			p.PC += o[1]
		}
	})

	p.Movement("jmpnz", "r, n", func (o []int) {
		if !p.R.ZeroSameAs(o[0]) {
			p.PC += o[1]
		}
	})

	p.Movement("call", "n", func (n int) {
		p.Call(n)
	})

	p.Movement("ret", "", func () {
		p.Return()
	})

	p.Operator("push", "r", func (r int) {
		p.DS.Append(p.R[r])
	})

	p.Operator("pop", "r", func (r int) {
		p.R[r], _ = p.DS.Pop()
	})

	p.Operator("cld", "r, v", func (o []int) {
		p.R[o[0]] = o[1]
	})

	p.Operator("send", "c", func (c int) {
		p.IOController.Send(c, p.M)
	})

	p.Operator("recv", "c", func (c int) {
		p.M = p.IOController.Receive(c)
	})

	p.Operator("inc", "r", func (r int) {
		p.R.Increment(r)
	})

	p.Operator("dec", "r", func (r int) {
		p.R.Decrement(r)
	})

	p.Operator("add", "r1, r2", func (o []int) {
		p.R.Add(o[0], o[1])
	})

	p.Operator("sub", "r1, r2", func (o []int) {
		p.R.Subtract(o[0], o[1])
	})

	p.Operator("mul", "r1, r2", func (o []int) {
		p.R.Multiply(o[0], o[1])
	})

	p.Operator("div", "r1, r2", func (o []int) {
		p.R.Divide(o[0], o[1])
	})

	p.Operator("and", "r1, r2", func (o []int) {
		p.R.And(o[0], o[1])
	})

	p.Operator("or", "r1, r2", func (o []int) {
		p.R.Or(o[0], o[1])
	})

	p.Operator("xor", "r1, r2", func (o []int) {
		p.R.Xor(o[0], o[1])
	})
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
	switch data := o.Data.(type) {
	case int:
		p.ops[o.Code].(func (int))(data)
	case []int:
		p.ops[o.Code].(func ([]int))(data)
	default:
		panic(nil)
	}
	p.PC += o.Movement
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