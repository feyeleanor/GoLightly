//	TODO: 	Write tests :)

package golightly
import (
	"testing"
)

type TestCall func(int, float) (bool)

func populate_block(m MemoryBlock) {
	m[0] = "hello world";
	m[1] = 37;
	m[2] = true;
	m[3] = 3.7;
	m[4] = func(a int, b float) bool { return float(a) / 10 > b };
	m[5] = TestCall(m[4])(m[1], m[2]);
}

func Test_CreateBlock(t *testing.T) {
	block := make(MemoryBlock, 10);
	populate_block(block);
}

func Test_ConvertToMemoryBlock(t *testing.T) {
	block := MemoryBlock(make([]Value, 10));
	populate_block(block);
}