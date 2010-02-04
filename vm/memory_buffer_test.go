//	TODO: 	Write tests :)

package vm
import (
	"testing"
)

func populate(b *Buffer) {
	b.Set(0, 37)
	b.Set(1, int(byte("hello world"[1])))
	b.Set(2, int(3.7))
}

func Test_CreateBuffer(t *testing.T) {
	b := new(Buffer)
	b.Init(6)
	b.Len()
	b.Cap()
	populate(b)
}
