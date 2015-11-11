package vm

import "time"

type InlinedProcessorCore struct {
	ProcessorCore
}

func (p *InlinedProcessorCore) Execute() {
	o := p.Program[p.PC]
	switch data := o.Data.(type) {
	case int:
		switch o.Code {
		case 4:
			p.PC += data
		case 7:
			p.CS.Append(p.PC)
			p.PC = data
		case 9:
			p.DS.Append(p.R[data])
		case 10:
			p.R[data], _ = p.DS.Pop()
		case 12:
			p.IOController.Send(data, p.M)
		case 13:
			p.M = p.IOController.Receive(data)
		case 14:
			p.R.Increment(data)
		case 15:
			p.R.Decrement(data)
		default:
			p.ProcessorCore.Execute()
		}
	case time.Duration:
		switch o.Code {
		case 1:
			time.Sleep(data)
		case 2:
			time.Sleep(data << 32)
		default:
			p.ProcessorCore.Execute()
		}
	case []int:
		switch o.Code {
		case 5:
			if p.R.ZeroSameAs(data[0]) {
				p.PC += data[1]
			} else {
				p.PC++
			}
		case 6:
			if !p.R.ZeroSameAs(data[0]) {
				p.PC += data[1]
			} else {
				p.PC++
			}
		case 11:
			p.R[data[0]] = data[1]
		case 16:
			p.R.Add(data[0], data[1])
		case 17:
			p.R.Subtract(data[0], data[1])
		case 18:
			p.R.Multiply(data[0], data[1])
		case 19:
			p.R.Divide(data[0], data[1])
		case 20:
			p.R.And(data[0], data[1])
		case 21:
			p.R.Or(data[0], data[1])
		case 22:
			p.R.Xor(data[0], data[1])
		default:
			p.ProcessorCore.Execute()
		}
	default:
		switch o.Code {
		case 0:
		case 3:
			p.Running = false
		case 8:
			p.PC, _ = p.CS.Pop(); p.PC++
		default:
			p.ProcessorCore.Execute()
		}
	}
	p.PC += o.Movement
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