package vm
import "testing"
//import "container/vector"

func twoElementStream() *Stream {
	s := new(Stream)
	s.Init(2)
	s.Buffer.Set(0, 100)
	s.Buffer.Set(1, 200)
	return s
}

func BenchmarkStreamAt(b *testing.B) {
    b.StopTimer()
	s := defaultStream()
    b.StartTimer()
    for i := 0; i < b.N; i++ { s.At(0) }
}

func BenchmarkStreamSet(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Set(0, 1) }
}

func BenchmarkStreamBufferAdd(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Add(0, 1) }
}

func BenchmarkStreamAdd(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkStreamAdd2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkStreamBufferSubtract(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Subtract(0, 1) }
}

func BenchmarkStreamSubtract(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkStreamSubtract2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkStreamBufferMultiply(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Multiply(0, 1) }
}

func BenchmarkStreamMultiply(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkStreamMultiply2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkStreamBufferDivide(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	s.Set(0, 987654321)
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Divide(0, 1) }
}

func BenchmarkStreamDivide(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s1.Set(0, 987654321)
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkStreamDivide2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s1.Set(0, 987654321)
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkStreamBufferAnd(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	s.Set(0, 987654321)
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.And(0, 1) }
}

func BenchmarkStreamAnd(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s1.Set(0, 987654321)
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkStreamAnd2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s1.Set(0, 987654321)
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkStreamBufferOr(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	s.Set(0, 987654321)
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Or(0, 1) }
}

func BenchmarkStreamOr(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s1.Set(0, 987654321)
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkStreamOr2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s1.Set(0, 987654321)
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkStreamBufferXor(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	s.Set(0, 987654321)
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Xor(0, 1) }
}

func BenchmarkStreamXor(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s1.Set(0, 987654321)
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkStreamXor2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s1.Set(0, 987654321)
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkStreamBufferIncrement(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Increment(0) }
}

func BenchmarkStreamIncrement(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment(0, s.Len()) }
}

func BenchmarkStreamIncrement2(b *testing.B) {
	b.StopTimer()
	s := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment(0, s.Len()) }
}

func BenchmarkStreamBufferDecrement(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Decrement(0) }
}

func BenchmarkStreamDecrement(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement(0, s.Len()) }
}

func BenchmarkStreamDecrement2(b *testing.B) {
	b.StopTimer()
	s := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement(0, s.Len()) }
}

func BenchmarkStreamBufferNegate(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Negate(0) }
}

func BenchmarkStreamNegate(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate(0, s.Len()) }
}

func BenchmarkStreamNegate2(b *testing.B) {
	b.StopTimer()
	s := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate(0, s.Len()) }
}

func BenchmarkStreamBufferShiftLeft(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.ShiftLeft(0, 1) }
}

func BenchmarkStreamShiftLeft(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkStreamShiftLeft2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkStreamBufferShiftRight(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.ShiftRight(0, 1) }
}

func BenchmarkStreamShiftRight(b *testing.B) {
	b.StopTimer()
	s1 := defaultStream()
	s2 := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkStreamShiftRight2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkStreamBufferInvert(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Buffer.Invert(0) }
}

func BenchmarkStreamInvert(b *testing.B) {
	b.StopTimer()
	s := defaultStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert(0, s.Len()) }
}

func BenchmarkStreamInvert2(b *testing.B) {
	b.StopTimer()
	s := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert(0, s.Len()) }
}
