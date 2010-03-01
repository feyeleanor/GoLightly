//	TODO:	Rewrite Transform to run in parallel

package vm

type Program struct {
	code			[]*OpCode
	pc				int
}
func (p *Program) ValidPC() bool		{ return (p.pc > -1) && (p.pc < len(p.code)) }
func (p *Program) StepForward()			{ p.pc++ }
func (p *Program) StepBack()			{ p.pc-- }
func (p *Program) JumpTo(ops int)		{ p.pc = ops }
func (p *Program) JumpRelative(ops int)	{ p.pc += ops }
func (p *Program) Len() int				{ return len(p.code) }
func (p *Program) CheckCompatibility(i *InstructionSet) bool {
	return true
}
//func (p *Program) Symbol(i int) (s string, ok bool) {
//	if 0 <= i && i < p.symbols.Len() { return p.symbols.At(i), true }
//	return "", false
//}
func (p *Program) At(i int) *OpCode {
	if 0 <= i && i < len(p.code) { return p.code[i] }
	return nil
}
func (p *Program) I() *OpCode {
	if 0 <= p.pc && p.pc < len(p.code) { return p.code[p.pc] }
	return nil
}
func (p *Program) Transform(o, n *OpCode) int {
	replacements := 0
	for i, v := range p.code {
		if v.Identical(o) {
			p.code[i].Replace(n)
			replacements++
		}
	}
	return replacements
}
