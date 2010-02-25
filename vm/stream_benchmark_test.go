package vm
import "testing"
//import "container/vector"

func oneIntegerStream() *Stream {
	s := new(Stream)
	s.Init(1)
	s.Buffer.Set(0, 100)
	return s
}

func oneFloatStream() *Stream {
	s := new(Stream)
	s.Init(1)
	s.Buffer.FSet(0, 100.00)
	return s
}

func twoIntegerStream() *Stream {
	s := new(Stream)
	s.Init(2)
	s.Buffer.Set(0, 100)
	s.Buffer.Set(1, 200)
	return s
}

func twoFloatStream() *Stream {
	s := new(Stream)
	s.Init(2)
	s.Buffer.Set(0, 100.00)
	s.Buffer.Set(1, 200.00)
	return s
}

func BenchmarkStreamAdd1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkStreamAdd2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkStreamAdd6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkStreamFAdd1(b *testing.B) {
	b.StopTimer()
	s1 := oneFloatStream()
	s2 := oneFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkStreamFAdd2(b *testing.B) {
	b.StopTimer()
	s1 := twoFloatStream()
	s2 := twoFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkStreamFAdd6(b *testing.B) {
	b.StopTimer()
	s1 := sixFloatStream()
	s2 := sixFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkStreamSubtract1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkStreamSubtract2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkStreamSubtract6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkStreamFSubtract1(b *testing.B) {
	b.StopTimer()
	s1 := oneFloatStream()
	s2 := oneFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkStreamFSubtract2(b *testing.B) {
	b.StopTimer()
	s1 := twoFloatStream()
	s2 := twoFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkStreamFSubtract6(b *testing.B) {
	b.StopTimer()
	s1 := sixFloatStream()
	s2 := sixFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkStreamMultiply1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkStreamMultiply2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkStreamMultiply6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkStreamFMultiply1(b *testing.B) {
	b.StopTimer()
	s1 := oneFloatStream()
	s2 := oneFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkStreamFMultiply2(b *testing.B) {
	b.StopTimer()
	s1 := twoFloatStream()
	s2 := twoFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkStreamFMultiply6(b *testing.B) {
	b.StopTimer()
	s1 := sixFloatStream()
	s2 := sixFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkStreamDivide1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s1.Set(0, 987654321)
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkStreamDivide2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s1.Set(0, 987654321)
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkStreamDivide6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s1.Set(0, 987654321)
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkStreamFDivide1(b *testing.B) {
	b.StopTimer()
	s1 := oneFloatStream()
	s2 := oneFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkStreamFDivide2(b *testing.B) {
	b.StopTimer()
	s1 := twoFloatStream()
	s2 := twoFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkStreamFDivide6(b *testing.B) {
	b.StopTimer()
	s1 := sixFloatStream()
	s2 := sixFloatStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkStreamAnd1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s1.Set(0, 987654321)
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkStreamAnd2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s1.Set(0, 987654321)
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkStreamAnd6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s1.Set(0, 987654321)
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkStreamOr1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s1.Set(0, 987654321)
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkStreamOr2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s1.Set(0, 987654321)
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkStreamOr6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s1.Set(0, 987654321)
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkStreamXor1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s1.Set(0, 987654321)
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkStreamXor2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s1.Set(0, 987654321)
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkStreamXor6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s1.Set(0, 987654321)
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkStreamIncrement1(b *testing.B) {
	b.StopTimer()
	s := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment(0, s.Len()) }
}

func BenchmarkStreamIncrement2(b *testing.B) {
	b.StopTimer()
	s := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment(0, s.Len()) }
}

func BenchmarkStreamIncrement6(b *testing.B) {
	b.StopTimer()
	s := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment(0, s.Len()) }
}

func BenchmarkStreamDecrement1(b *testing.B) {
	b.StopTimer()
	s := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement(0, s.Len()) }
}

func BenchmarkStreamDecrement2(b *testing.B) {
	b.StopTimer()
	s := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement(0, s.Len()) }
}

func BenchmarkStreamDecrement6(b *testing.B) {
	b.StopTimer()
	s := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement(0, s.Len()) }
}

func BenchmarkStreamNegate1(b *testing.B) {
	b.StopTimer()
	s := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate(0, s.Len()) }
}

func BenchmarkStreamNegate2(b *testing.B) {
	b.StopTimer()
	s := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate(0, s.Len()) }
}

func BenchmarkStreamNegate6(b *testing.B) {
	b.StopTimer()
	s := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate(0, s.Len()) }
}

func BenchmarkStreamShiftLeft1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkStreamShiftLeft2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkStreamShiftLeft6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkStreamShiftRight1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerStream()
	s2 := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkStreamShiftRight2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerStream()
	s2 := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkStreamShiftRight6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerStream()
	s2 := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkStreamInvert1(b *testing.B) {
	b.StopTimer()
	s := oneIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert(0, s.Len()) }
}

func BenchmarkStreamInvert2(b *testing.B) {
	b.StopTimer()
	s := twoIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert(0, s.Len()) }
}

func BenchmarkStreamInvert6(b *testing.B) {
	b.StopTimer()
	s := sixIntegerStream()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert(0, s.Len()) }
}
