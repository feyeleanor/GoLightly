//	TODO: 	Add tests for GetBuffer, PutBuffer and Clear

package vm
import "testing"
import "math"

var predicate_index int

func twoIntegerBuffer() *Buffer {
	return &Buffer{100, 200}
}

func sixIntegerBuffer() *Buffer {
	c := "hello world"[1]
	f := 3.7
	return &Buffer{37, int(byte(c)), int(f), 5, 2, 2}
}

func sixFloatBuffer() *Buffer {
	b := new(Buffer)
	b.Resize(6)
	b.FSet(0, 37.0)
	b.FSet(1, float("e"[0]))					//	ASCII == 101.0
	b.FSet(2, 3.7)
	b.FSet(3, 5.0)
	b.FSet(4, 2.0)
	b.FSet(5, 2.0)
	return b
}

func compareValues(object interface{}, t *testing.T, value, target_value interface{}) {
	predicate_index += 1
	if value != target_value { t.Errorf("%T: test %d -> expected %v, got %v", object, predicate_index, target_value, value) }
}

func compareFloatValues(object interface{}, t *testing.T, value, target_value, tolerance float) {
	compareValues(object, t, math.Fabs(float64(value - target_value)) < float64(tolerance), true)
}

func checkIntegerBuffer(b, o *Buffer, t *testing.T, value bool) {
	compareValues(b, t, b.Identical(o), value)
}

func checkFloatBuffer(b, o *Buffer, tolerance float, t *testing.T, value bool) {
	compareValues(b, t, b.FIdentical(o, tolerance), value)
}

func TestBufferCreate(t *testing.T) {
	checkIntegerBuffer(sixIntegerBuffer(), sixIntegerBuffer(), t, true)
}

func TestBufferClone(t *testing.T) {
	checkIntegerBuffer(sixIntegerBuffer().Clone(), sixIntegerBuffer(), t, true)
}

func TestBufferReplication(t *testing.T) {
	b1 := twoIntegerBuffer()
	b2 := b1.Replicate(3)
	compareValues(b2, t, b2.Len(), 6)
	compareValues(b2, t, b2.Cap(), 6)
	compareValues(b2, t, b2.At(0) == b1.At(0), true)
	compareValues(b2, t, b2.At(1) == b1.At(1), true)
	compareValues(b2, t, b2.At(2) == b1.At(0), true)
	compareValues(b2, t, b2.At(3) == b1.At(1), true)
	compareValues(b2, t, b2.At(4) == b1.At(0), true)
	compareValues(b2, t, b2.At(5) == b1.At(1), true)
}

func TestBufferSlice(t *testing.T) {
	b := sixIntegerBuffer().Slice(1, 3)
	compareValues(b, t, b.Len(), 2)
	compareValues(b, t, b.Cap(), 2)
	compareValues(b, t, b.At(0), int(byte("e"[0])))
	compareValues(b, t, b.At(1), 3)
}

func TestBufferMaths(t *testing.T) {
	b := sixIntegerBuffer()
	b.Increment(0)											//	b[0] == 38
	compareValues(b, t, b.At(0), 38)
	b.Decrement(0)											//	b[0] == 37
	compareValues(b, t, b.At(0), 37)
	b.Add(1, 3)												//	b[1] == 'j'
	compareValues(b, t, b.At(1), int(byte("j"[0])))
	b.Subtract(2, 3)										//	b[2] == -2
	compareValues(b, t, b.At(2), -2)
	b.Negate(4)												//	b[4] == -2
	compareValues(b, t, b.At(4), -2)
	b.Multiply(2, 4)										//	b[2] == 4
	compareValues(b, t, b.At(2), 4)
	b.Divide(2, 5)											//	b[2] == 2
	compareValues(b, t, b.At(2), 2)
	b.Multiply(5, 3)										//	b[5] == 10
	b.And(2, 5)												//	b[2] == 2
	compareValues(b, t, b.At(2), 2)
	b.Or(2, 5)												//	b[2] == 10
	compareValues(b, t, b.At(2), 10)
	b.Negate(4)												//	b[4] == 2
	compareValues(b, t, b.At(4), 2)
	b.Xor(2, 4)												//	b[2] == 8
	compareValues(b, t, b.At(2), 8)
}

func TestBufferFloatingPointMaths(t *testing.T) {
	b := sixFloatBuffer()
	compareValues(b, t, b.FAt(0), 37.0)
	compareValues(b, t, b.FAt(1), 101.0)
	b.FAdd(0, 1)
	compareFloatValues(b, t, b.FAt(0), 138.0, 0.001)
	b.FSubtract(0, 1)
	compareFloatValues(b, t, b.FAt(0), 37.0, 0.001)
	b.FMultiply(0, 1)
	compareFloatValues(b, t, b.FAt(0), 3737.0, 0.001)
	b.FDivide(0, 1)
	compareFloatValues(b, t, b.FAt(0), 37.0, 0.001)
	b.FNegate(0)
	compareFloatValues(b, t, b.FAt(0), -37.0, 0.001)
}

func TestBufferBitOperators(t *testing.T) {
	b := sixIntegerBuffer()									//	b[0] == 37, b[5] == 2
	b.ShiftRight(0, 5)
	compareValues(b, t, b.At(0), 148)
	b.ShiftLeft(0, 5)
	compareValues(b, t, b.At(0), 37)
	b.Invert(0)
	compareValues(b, t, b.At(0), ^37)
}

func TestBufferIntegerLogic(t *testing.T) {
	b := sixIntegerBuffer()
	checkIntegerBuffer(b, sixIntegerBuffer(), t, true)
	compareValues(b, t, b.LessThan(2, 3), true)				//	b[2] == 3, b[3] == 5
	compareValues(b, t, b.Equals(2, 3), false)
	compareValues(b, t, b.GreaterThan(2, 3), false)
	compareValues(b, t, b.LessThanZero(2), false)
	compareValues(b, t, b.EqualsZero(2), false)
	compareValues(b, t, b.GreaterThanZero(2), true)
	b.Copy(1, 2)											//	b[1] == 3
	checkIntegerBuffer(b, sixIntegerBuffer(), t, false)
	compareValues(b, t, b.At(1), 3)
	compareValues(b, t, b.LessThan(1, 3), true)				//	b[1] == 3, b[3] == 5
	compareValues(b, t, b.Equals(1, 2), true)				//	b[1] == 3, b[2] == 3
	compareValues(b, t, b.GreaterThan(1, 3), false)
	compareValues(b, t, b.LessThanZero(1), false)
	compareValues(b, t, b.EqualsZero(1), false)
	compareValues(b, t, b.GreaterThanZero(1), true)
	b.Set(1, 0)												//	b[1] == 0, b[3] == 5
	checkIntegerBuffer(b, sixIntegerBuffer(), t, false)
	compareValues(b, t, b.LessThan(1, 3), true)
	compareValues(b, t, b.Equals(1, 3), false)
	compareValues(b, t, b.GreaterThan(1, 3), false)
	compareValues(b, t, b.LessThanZero(1), false)
	compareValues(b, t, b.EqualsZero(1), true)
	compareValues(b, t, b.GreaterThanZero(1), false)
}

func TestBufferFloatingPointLogic(t *testing.T) {
	b := sixFloatBuffer()
	checkFloatBuffer(b, sixFloatBuffer(), 0.001, t, true)
	compareValues(b, t, b.FLessThan(0, 1), true)
	compareValues(b, t, b.FEquals(0, 1, 0.001), false)
	compareValues(b, t, b.FGreaterThan(0, 1), false)
	compareValues(b, t, b.FLessThanZero(0), false)
	compareValues(b, t, b.FEqualsZero(0, 0.001), false)
	compareValues(b, t, b.FGreaterThanZero(0), true)
}
