package vm

import . "github.com/feyeleanor/slices"

type Thread struct {
	Running		bool
	R					ISlice
	M					ISlice
	CS				*ISlice
	DS				ISlice
	PC				int
	Program
}

func (t *Thread) I() OpCode {
	return t.Program[t.PC]
}

func (t *Thread) ValidPC() bool {
	return t.PC > -1 && t.PC < len(t.Program)
}

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