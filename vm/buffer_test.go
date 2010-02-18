//	TODO: 	Add tests for GetBuffer, PutBuffer and Clear

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

func compareValues(object interface{}, t *testing.T, value, target_value interface{}) {
	predicate_index += 1
	if value != target_value { t.Errorf("%T: test %d -> expected %v, got %v", object, predicate_index, target_value, value) }
}

func checkDefaultBuffer(b *Buffer, t *testing.T) {
	compareValues(b, t, b.Len(), 6)
	compareValues(b, t, b.Cap(), 6)
	compareValues(b, t, b.At(0), 37)
	compareValues(b, t, b.At(1), int(byte("e"[0])))
	compareValues(b, t, b.At(2), 3)
	compareValues(b, t, b.At(3), 0)
	compareValues(b, t, b.At(4), 0)
	compareValues(b, t, b.At(5), 0)
}

func TestCreateBuffer(t *testing.T) {
	os.Stdout.WriteString("Buffer Creation\n")
	checkDefaultBuffer(defaultBuffer(), t)
}

func TestClone(t *testing.T) {
	os.Stdout.WriteString("Buffer Cloning\n")
	checkDefaultBuffer(defaultBuffer().Clone(), t)
}

func TestSlice(t *testing.T) {
	os.Stdout.WriteString("Buffer Slicing\n")
	b := defaultBuffer().Slice(1, 3)
	compareValues(b, t, b.Len(), 2)
	compareValues(b, t, b.Cap(), 2)
	compareValues(b, t, b.At(0), int(byte("e"[0])))
	compareValues(b, t, b.At(1), 3)
}

func TestMaths(t *testing.T) {
	os.Stdout.WriteString("Buffer Maths\n")
	b := defaultBuffer()
	b.Increment(0)
	compareValues(b, t, b.At(0), 38)
	b.Decrement(0)
	compareValues(b, t, b.At(0), 37)
	b.Add(1, 5)
	compareValues(b, t, b.At(1), int(byte("j"[0])))
	b.Subtract(2, 4)
	compareValues(b, t, b.At(2), -1)
	b.Multiply(2, -4)
	compareValues(b, t, b.At(2), 4)
	b.Divide(2, 2)
	compareValues(b, t, b.At(2), 2)
	b.And(2, 10)
	compareValues(b, t, b.At(2), 2)
	b.Or(2, 10)
	compareValues(b, t, b.At(2), 10)
	b.Xor(2, 2)
	compareValues(b, t, b.At(2), 8)
}

func TestLogic(t *testing.T) {
	os.Stdout.WriteString("Buffer Logic\n")
	b := defaultBuffer()
	compareValues(b, t, b.LessThan(2, 3), false)
	compareValues(b, t, b.Equals(2, 3), true)
	compareValues(b, t, b.GreaterThan(2, 3), false)
	compareValues(b, t, b.LessThanZero(2), false)
	compareValues(b, t, b.EqualsZero(2), false)
	compareValues(b, t, b.GreaterThanZero(2), true)
	b.Copy(1, 2)
	compareValues(b, t, b.At(1), 3)
	compareValues(b, t, b.LessThan(1, 3), false)
	compareValues(b, t, b.Equals(1, 3), true)
	compareValues(b, t, b.GreaterThan(1, 3), false)
	compareValues(b, t, b.LessThanZero(1), false)
	compareValues(b, t, b.EqualsZero(1), false)
	compareValues(b, t, b.GreaterThanZero(1), true)
	b.Set(1, 0)
	compareValues(b, t, b.LessThan(1, 3), true)
	compareValues(b, t, b.Equals(1, 3), false)
	compareValues(b, t, b.GreaterThan(1, 3), false)
	compareValues(b, t, b.LessThanZero(1), false)
	compareValues(b, t, b.EqualsZero(1), true)
	compareValues(b, t, b.GreaterThanZero(1), false)
}
