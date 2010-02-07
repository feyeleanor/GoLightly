package vm

type OpCode struct {
	code, a, b, c	int;
}

func (o *OpCode) Equals(p *OpCode) bool {
	return o.code == p.code && o.a == p.a && o.b == p.b && o.c == p.c
}
