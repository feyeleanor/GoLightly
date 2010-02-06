//	TODO: 	Add tests

package vm
import "testing"
import "os"

const BUFFER_ALLOCATION = 16

func defaultRegisterBlock() *RegisterBlock {
	r := new(RegisterBlock)
	r.Allocate(BUFFER_ALLOCATION)
	return r
}

func checkAllocatedBuffer(b *Buffer, t *testing.T) {
	valueTest(b, t, b.Len(), BUFFER_ALLOCATION)
	valueTest(b, t, b.Cap(), BUFFER_ALLOCATION)
	valueTest(b, t, b.At(0), 0)
	valueTest(b, t, b.At(15), 0)
}

func checkDefaultRegisterBlock(r *RegisterBlock, t *testing.T) {
	checkAllocatedBuffer(r.R, t)
	valueTest(r, t, r.M == nil, true)
	valueTest(r, t, r.I == nil, true)
	valueTest(r, t, r.PC, 0)
}

func TestRegisterBlock(t *testing.T) {
	os.Stdout.WriteString("Register Block Creation\n")
	r := defaultRegisterBlock()
	checkDefaultRegisterBlock(r, t)
	os.Stdout.WriteString("Register Block Cloning\n")
	c := r.Clone()
	checkDefaultRegisterBlock(c, t)
	os.Stdout.WriteString("Register Block Replacement\n")
	r.PC = 27
	r.M = new(Buffer)
	r.M.Init(48)
	valueTest(r, t, r.PC, 27)
	valueTest(r, t, r.M.Cap(), 48)
	r.Replace(c)
	checkDefaultRegisterBlock(r, t)
}

func TestMMU(t *testing.T) {
	os.Stdout.WriteString("MMU Allocation\n")
	m := new(MMU)
	b := m.Allocate(BUFFER_ALLOCATION)
	checkAllocatedBuffer(b, t)
}

func TestProcessorCoreCreation(t *testing.T) {
	os.Stdout.WriteString("Processor Core Creation\n")
	p := new(ProcessorCore)
	p.Init(BUFFER_ALLOCATION)
//	checkAllocatedBuffer(p.RegisterBlock.R, t)
//	valueTest(p, t, p.RegisterBlock.M == nil, true)
//	valueTest(p, t, p.RegisterBlock.I == nil, true)
//	valueTest(p, t, p.RegisterBlock.PC, 0)
}
