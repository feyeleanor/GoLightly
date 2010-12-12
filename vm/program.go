//	TODO:	Rewrite Transform to run in parallel

package vm

type Program []OpCode
func (p Program) CheckCompatibility(i *InstructionSet) bool {
	return true
}
//func (p *Program) Symbol(i int) (s string, ok bool) {
//	if 0 <= i && i < p.symbols.Len() { return p.symbols.At(i), true }
//	return "", false
//}
func (p Program) Transform(o, n *OpCode) (replacements int) {
	for i, v := range p {
		if v.Identical(*o) {
			p[i].Replace(n)
			replacements++
		}
	}
	return replacements
}