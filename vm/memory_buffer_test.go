//	TODO: 	Add tests for GetBuffer and PutBuffer

package vm
import "testing"
import "os"

var predicate_index int

func defaultBuffer() *Buffer {
	b := new(Buffer)
	b.Init(6)
	b.Set(0, 37)
	b.Set(1, int(byte("hello world"[1])))
	f := 3.7
	b.Set(2, int(f))
	return b
}

func valueTest(object interface{}, t *testing.T, value, target_value interface{}) {
	predicate_index += 1
	if value != target_value { t.Errorf("%T: test %d -> expected %v, got %v", object, predicate_index, target_value, value) }
}

func checkDefaultBuffer(b *Buffer, t *testing.T) {
	valueTest(b, t, b.Len(), 6)
	valueTest(b, t, b.Cap(), 6)
	valueTest(b, t, b.At(0), 37)
	valueTest(b, t, b.At(1), int(byte("e"[0])))
	valueTest(b, t, b.At(2), 3)
	valueTest(b, t, b.At(3), 0)
	valueTest(b, t, b.At(4), 0)
	valueTest(b, t, b.At(5), 0)
}

func TestCreateBuffer(t *testing.T) {
	os.Stdout.WriteString("Buffer Creation\n")
	checkDefaultBuffer(defaultBuffer(), t)
}

func TestClone(t *testing.T) {
	os.Stdout.WriteString("Cloning\n")
	checkDefaultBuffer(defaultBuffer().Clone(), t)
}

func TestSlice(t *testing.T) {
	os.Stdout.WriteString("Slicing\n")
	b := defaultBuffer().Slice(1, 3)
	valueTest(b, t, b.Len(), 2)
	valueTest(b, t, b.Cap(), 2)
	valueTest(b, t, b.At(0), int(byte("e"[0])))
	valueTest(b, t, b.At(1), 3)
}

func TestMaths(t *testing.T) {
	os.Stdout.WriteString("Maths\n")
	b := defaultBuffer()
	b.Increment(0)
	valueTest(b, t, b.At(0), 38)
	b.Decrement(0)
	valueTest(b, t, b.At(0), 37)
	b.Add(1, 5)
	valueTest(b, t, b.At(1), int(byte("j"[0])))
	b.Subtract(2, 4)
	valueTest(b, t, b.At(2), -1)
	b.Multiply(2, -4)
	valueTest(b, t, b.At(2), 4)
	b.Divide(2, 2)
	valueTest(b, t, b.At(2), 2)
	b.And(2, 10)
	valueTest(b, t, b.At(2), 2)
	b.Or(2, 10)
	valueTest(b, t, b.At(2), 10)
	b.Xor(2, 2)
	valueTest(b, t, b.At(2), 8)
}

func TestLogic(t *testing.T) {
	os.Stdout.WriteString("Logic\n")
	b := defaultBuffer()
	valueTest(b, t, b.LessThan(2, 3), false)
	valueTest(b, t, b.Equals(2, 3), true)
	valueTest(b, t, b.GreaterThan(2, 3), false)
	valueTest(b, t, b.LessThanZero(2), false)
	valueTest(b, t, b.EqualsZero(2), false)
	valueTest(b, t, b.GreaterThanZero(2), true)
	b.Copy(1, 2)
	valueTest(b, t, b.At(1), 3)
	valueTest(b, t, b.LessThan(1, 3), false)
	valueTest(b, t, b.Equals(1, 3), true)
	valueTest(b, t, b.GreaterThan(1, 3), false)
	valueTest(b, t, b.LessThanZero(1), false)
	valueTest(b, t, b.EqualsZero(1), false)
	valueTest(b, t, b.GreaterThanZero(1), true)
	b.Set(1, 0)
	valueTest(b, t, b.LessThan(1, 3), true)
	valueTest(b, t, b.Equals(1, 3), false)
	valueTest(b, t, b.GreaterThan(1, 3), false)
	valueTest(b, t, b.LessThanZero(1), false)
	valueTest(b, t, b.EqualsZero(1), true)
	valueTest(b, t, b.GreaterThanZero(1), false)
}
