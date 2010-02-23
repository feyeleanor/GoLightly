//	TODO:	Improve tests for Stream-level operations
//	TODO: 	Add tests for GetBuffer, PutBuffer and Clear

package vm
import "testing"
import "os"

func defaultStream() *Stream {
	s := new(Stream)
	s.Init(6)
	s.Buffer.Set(0, 37)
	s.Buffer.Set(1, int(byte("hello world"[1])))
	f := 3.7
	s.Buffer.Set(2, int(f))
	s.Buffer.Set(3, 5)
	s.Buffer.Set(4, 2)
	s.Buffer.Set(5, 2)
	return s
}

func checkDefaultStream(s *Stream, t *testing.T, value bool) {
	compareValues(s, t, s.Identical(defaultStream()), value)
}

func TestCreateStream(t *testing.T) {
	os.Stdout.WriteString("Stream Creation\n")
	checkDefaultStream(defaultStream(), t, true)
}

func TestCloneStream(t *testing.T) {
	os.Stdout.WriteString("Stream Cloning\n")
	checkDefaultStream(defaultStream().Clone(), t, true)
}

func TestSliceStream(t *testing.T) {
	os.Stdout.WriteString("Stream Slicing\n")
	s := defaultStream().Slice(1, 3)
	compareValues(s, t, s.Len(), 2)
	compareValues(s, t, s.Cap(), 2)
	compareValues(s, t, s.At(0), int(byte("e"[0])))
	compareValues(s, t, s.At(1), 3)
}

func TestStreamMaths(t *testing.T) {
	os.Stdout.WriteString("Stream Maths\n")
	s1 := defaultStream()
	s1.Buffer.Increment(0)											//	b[0] == 38
	compareValues(s1, t, s1.At(0), 38)
	s1.Buffer.Decrement(0)											//	b[0] == 37
	compareValues(s1, t, s1.At(0), 37)
	s1.Buffer.Add(1, 3)												//	b[1] == 'j'
	compareValues(s1, t, s1.At(1), int(byte("j"[0])))
	s1.Buffer.Subtract(2, 3)										//	b[2] == -2
	compareValues(s1, t, s1.At(2), -2)
	s1.Buffer.Negate(4)												//	b[4] == -2
	compareValues(s1, t, s1.At(4), -2)
	s1.Buffer.Multiply(2, 4)										//	b[2] == 4
	compareValues(s1, t, s1.At(2), 4)
	s1.Buffer.Divide(2, 5)											//	b[2] == 2
	compareValues(s1, t, s1.At(2), 2)
	s1.Buffer.Multiply(5, 3)										//	b[5] == 10
	s1.Buffer.And(2, 5)												//	b[2] == 2
	compareValues(s1, t, s1.At(2), 2)
	s1.Buffer.Or(2, 5)												//	b[2] == 10
	compareValues(s1, t, s1.At(2), 10)
	s1.Buffer.Negate(4)												//	b[4] == 2
	compareValues(s1, t, s1.At(4), 2)
	s1.Buffer.Xor(2, 4)												//	b[2] == 8
	compareValues(s1, t, s1.At(2), 8)
	s1 = defaultStream()
	s2 := defaultStream()
	s1.Add(s2)
	compareValues(s1, t, s1.At(0), 74)
	compareValues(s1, t, s1.At(1), int(byte("e"[0])) * 2)
	compareValues(s1, t, s1.At(2), 6)
	compareValues(s1, t, s1.At(3), 10)
	compareValues(s1, t, s1.At(4), 4)
	compareValues(s1, t, s1.At(5), 4)
	s1 = defaultStream()
	s1.Subtract(s2)
	compareValues(s1, t, s1.At(0), 0)
	compareValues(s1, t, s1.At(1), 0)
	compareValues(s1, t, s1.At(2), 0)
	compareValues(s1, t, s1.At(3), 0)
	compareValues(s1, t, s1.At(4), 0)
	compareValues(s1, t, s1.At(5), 0)
	s1 = defaultStream()
	s1.Multiply(s2)
	compareValues(s1, t, s1.At(0), 37 * 37)
	compareValues(s1, t, s1.At(1), int(byte("e"[0])) * int(byte("e"[0])))
	compareValues(s1, t, s1.At(2), 9)
	compareValues(s1, t, s1.At(3), 25)
	compareValues(s1, t, s1.At(4), 4)
	compareValues(s1, t, s1.At(5), 4)
	s1 = defaultStream()
	s1.Divide(s2)
	compareValues(s1, t, s1.At(0), 1)
	compareValues(s1, t, s1.At(1), 1)
	compareValues(s1, t, s1.At(2), 1)
	compareValues(s1, t, s1.At(3), 1)
	compareValues(s1, t, s1.At(4), 1)
	compareValues(s1, t, s1.At(5), 1)
	s1 = defaultStream()
	s1.Negate()
	compareValues(s1, t, s1.At(0), -37)
	compareValues(s1, t, s1.At(1), -int(byte("e"[0])))
	compareValues(s1, t, s1.At(2), -3)
	compareValues(s1, t, s1.At(3), -5)
	compareValues(s1, t, s1.At(4), -2)
	compareValues(s1, t, s1.At(5), -2)
	s1 = defaultStream()
	s1.Increment()
	compareValues(s1, t, s1.At(0), 38)
	compareValues(s1, t, s1.At(1), int(byte("e"[0])) + 1)
	compareValues(s1, t, s1.At(2), 4)
	compareValues(s1, t, s1.At(3), 6)
	compareValues(s1, t, s1.At(4), 3)
	compareValues(s1, t, s1.At(5), 3)
	s1 = defaultStream()
	s1.Decrement()
	compareValues(s1, t, s1.At(0), 36)
	compareValues(s1, t, s1.At(1), int(byte("e"[0])) - 1)
	compareValues(s1, t, s1.At(2), 2)
	compareValues(s1, t, s1.At(3), 4)
	compareValues(s1, t, s1.At(4), 1)
	compareValues(s1, t, s1.At(5), 1)
}

func TestStreamBitOperators(t *testing.T) {
	os.Stdout.WriteString("Buffer Bit Manipulation\n")
	s := defaultStream()									//	b[0] == 37, b[5] == 2
	s.Buffer.ShiftRight(0, 5)
	compareValues(s, t, s.At(0), 148)
	s.Buffer.ShiftLeft(0, 5)
	compareValues(s, t, s.At(0), 37)
	s.Buffer.Invert(0)
	compareValues(s, t, s.At(0), ^37)
}

//func TestLogic(t *testing.T) {
//	os.Stdout.WriteString("Buffer Logic\n")
//	b := defaultBuffer()
//	checkDefaultBuffer(b, t, true)
//	compareValues(b, t, b.LessThan(2, 3), true)				//	b[2] == 3, b[3] == 5
//	compareValues(b, t, b.Equals(2, 3), false)
//	compareValues(b, t, b.GreaterThan(2, 3), false)
//	compareValues(b, t, b.LessThanZero(2), false)
//	compareValues(b, t, b.EqualsZero(2), false)
//	compareValues(b, t, b.GreaterThanZero(2), true)
//	b.Copy(1, 2)											//	b[1] == 3
//	checkDefaultBuffer(b, t, false)
//	compareValues(b, t, b.At(1), 3)
//	compareValues(b, t, b.LessThan(1, 3), true)				//	b[1] == 3, b[3] == 5
//	compareValues(b, t, b.Equals(1, 2), true)				//	b[1] == 3, b[2] == 3
//	compareValues(b, t, b.GreaterThan(1, 3), false)
//	compareValues(b, t, b.LessThanZero(1), false)
//	compareValues(b, t, b.EqualsZero(1), false)
//	compareValues(b, t, b.GreaterThanZero(1), true)
//	b.Set(1, 0)												//	b[1] == 0, b[3] == 5
//	checkDefaultBuffer(b, t, false)
//	compareValues(b, t, b.LessThan(1, 3), true)
//	compareValues(b, t, b.Equals(1, 3), false)
//	compareValues(b, t, b.GreaterThan(1, 3), false)
//	compareValues(b, t, b.LessThanZero(1), false)
//	compareValues(b, t, b.EqualsZero(1), true)
//	compareValues(b, t, b.GreaterThanZero(1), false)
//}
