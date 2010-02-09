//	TODO:	Rewrite Transform to run in parallel

package vm

import "container/vector"

type Program struct {
	code			[]*OpCode
	symbols			*vector.StringVector
	labels			map[string]int
	breakpoints		*vector.IntVector
}
func (p *Program) Init() {
	p.symbols = new(vector.StringVector)
	p.breakpoints = new(vector.IntVector)
	p.labels = make(map[string]int)
}
func (p *Program) CheckCompatibility(i *InstructionSet) bool {
	return true
}
func (p *Program) Symbol(i int) (s string, ok bool) {
	if 0 <= i && i < p.symbols.Len() {
		return p.symbols.At(i), true
	}
	return "", false
}
func (p *Program) At(i int) *OpCode {
	if 0 <= i && i < len(p.code) {
		return p.code[i]
	}
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
