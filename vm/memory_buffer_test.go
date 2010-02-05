//	TODO: 	Write tests :)

package vm
import "testing"
import "os"

func populate(b *Buffer) {
	b.Set(0, 37)
	b.Set(1, int(byte("hello world"[1])))
	f := 3.7
	b.Set(2, int(f))
}

func checkDefaultBuffer(b *Buffer, t *testing.T) {
	if b.Len() != 6 { t.Errorf("%T: B1) expected 6, got %d", b, b.Len()) }
	if b.Cap() != 6 { t.Errorf("%T: B2) expected 6, got %d", b, b.Cap()) }
	if b.At(0) != 37 { t.Errorf("%T: B3) expected 37, got %d", b, b.At(0)) }
	if b.At(1) != int(byte("e"[0])) { t.Errorf("%T: B4) expected 'e', got %d", b, b.At(1)) }
	if b.At(2) != 3 { t.Errorf("%T: B5) expected 3, got %d", b, b.At(2)) }
}

func TestCreateBuffer(t *testing.T) {
	os.Stdout.WriteString("Buffer Creation\n")
	b := new(Buffer)
	b.Init(6)
	populate(b)
	checkDefaultBuffer(b, t)
}

func checkBufferSlice(b *Buffer, t *testing.T) {
	if b.Len() != 2 { t.Errorf("%T: C1) expected 6, got %d", b, b.Len()) }
	if b.Cap() != 2 { t.Errorf("%T: C2) expected 6, got %d", b, b.Cap()) }
	if b.At(0) != int(byte("e"[0])) { t.Errorf("%T: C3) expected 'e', got %d", b, b.At(0)) }
	if b.At(1) != 3 { t.Errorf("%T: C4) expected 3, got %d", b, b.At(1)) }
}

func TestSlice(t *testing.T) {
	os.Stdout.WriteString("Slicing\n")
	b := new(Buffer)
	b.Init(6)
	populate(b)
	checkBufferSlice(b.Slice(1, 3), t)
}

func TestClone(t *testing.T) {
	os.Stdout.WriteString("Cloning\n")
	b := new(Buffer)
	b.Init(6)
	populate(b)
	checkDefaultBuffer(b.Clone(), t)
}

func TestMaths(t *testing.T) {
	os.Stdout.WriteString("Maths\n")
	b := new(Buffer)
	b.Init(6)
	populate(b)
	b.Increment(0)
	if b.At(0) != 38 { t.Errorf("%T: D1) expected 38, got %d", b, b.At(0)) }
	b.Decrement(0)
	if b.At(0) != 37 { t.Errorf("%T: D2) expected 37, got %d", b, b.At(0)) }
	b.Add(1, 5)
	if b.At(1) != int(byte("j"[0])) { t.Errorf("%T: D3) expected 'j', got %d", b, b.At(1)) }
	b.Subtract(2, 4)
	if b.At(2) != -1 { t.Errorf("%T: D4) expected -1, got %d", b, b.At(2)) }
	b.Multiply(2, -2)
	if b.At(2) != 2 { t.Errorf("%T: D5) expected 2, got %d", b, b.At(2)) }
}
