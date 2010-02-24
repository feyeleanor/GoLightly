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
	for i := 0; i < b.N; i++ { s1.Add(s2) }
}

func BenchmarkStreamAdd2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(s2) }
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
	for i := 0; i < b.N; i++ { s1.Subtract(s2) }
}

func BenchmarkStreamSubtract2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(s2) }
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
	for i := 0; i < b.N; i++ { s1.Multiply(s2) }
}

func BenchmarkStreamMultiply2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(s2) }
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
	for i := 0; i < b.N; i++ { s1.Divide(s2) }
}

func BenchmarkStreamDivide2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s1.Set(0, 987654321)
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(s2) }
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
	for i := 0; i < b.N; i++ { s.Increment() }
}

func BenchmarkStreamIncrement2(b *testing.B) {
	b.StopTimer()
	s := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment() }
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
	for i := 0; i < b.N; i++ { s.Decrement() }
}

func BenchmarkStreamDecrement2(b *testing.B) {
	b.StopTimer()
	s := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement() }
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
	for i := 0; i < b.N; i++ { s.Negate() }
}

func BenchmarkStreamNegate2(b *testing.B) {
	b.StopTimer()
	s := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate() }
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
	for i := 0; i < b.N; i++ { s1.ShiftLeft(s2) }
}

func BenchmarkStreamShiftLeft2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(s2) }
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
	for i := 0; i < b.N; i++ { s1.ShiftRight(s2) }
}

func BenchmarkStreamShiftRight2(b *testing.B) {
	b.StopTimer()
	s1 := twoElementStream()
	s2 := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(s2) }
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
	for i := 0; i < b.N; i++ { s.Invert() }
}

func BenchmarkStreamInvert2(b *testing.B) {
	b.StopTimer()
	s := twoElementStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert() }
}
