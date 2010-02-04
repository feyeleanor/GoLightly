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

func TestCreateBuffer(t *testing.T) {
	os.Stdout.WriteString("Testing Buffer Creation\n")
	b := new(Buffer)
	b.Init(6)
	populate(b)
	if b.Len() != 6 {
		t.Errorf("%T: B1) expected 6, got %d", b, b.Len())
	}
	if b.Cap() != 6 {
		t.Errorf("%T: B2) expected 6, got %d", b, b.Cap())
	}
	if b.At(0) != 37 {
		t.Errorf("%T: B3) expected 37, got %d", b, b.At(0))
	}
	if b.At(1) != int(byte("e"[0])) {
		t.Errorf("%T: B4) expected 'e', got %d", b, b.At(1))
	}
	if b.At(2) != 3 {
		t.Errorf("%T: B5) expected 3, got %d", b, b.At(2))
	}
}
