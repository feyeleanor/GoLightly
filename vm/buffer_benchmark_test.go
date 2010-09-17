package vm

import "testing"

func BenchmarkBufferFloatsToInts1(b *testing.B) {
	b.StopTimer()
		a := make([]float, 1, 1)
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkBufferFloatsToInts10(b *testing.B) {
	b.StopTimer()
		a := make([]float, 10, 10)
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkBufferFloatsToInts100(b *testing.B) {
	b.StopTimer()
		a := make([]float, 100, 100)
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkBufferFloatsToInts1000(b *testing.B) {
	b.StopTimer()
		a := make([]float, 1000, 1000)
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkBufferFloatsFromInts1(b *testing.B) {
	b.StopTimer()
		a := make([]int, 1, 1)
	b.StartTimer()
    for i := 0; i < b.N; i++ { intsToFloats(a) }
}

func BenchmarkBufferFloatsFromInts10(b *testing.B) {
	b.StopTimer()
		a := make([]int, 10, 10)
	b.StartTimer()
    for i := 0; i < b.N; i++ { intsToFloats(a) }
}

func BenchmarkBufferFloatsFromInts100(b *testing.B) {
	b.StopTimer()
		a := make([]int, 100, 100)
	b.StartTimer()
    for i := 0; i < b.N; i++ { intsToFloats(a) }
}

func BenchmarkBufferFloatsFromInts1000(b *testing.B) {
	b.StopTimer()
		a := make([]int, 1000, 1000)
	b.StartTimer()
    for i := 0; i < b.N; i++ { intsToFloats(a) }
}

func BenchmarkBufferReplicate2x2(b *testing.B) {
	v := Buffer{0, 0}
	for i := 0; i < b.N; i++ { v.Replicate(2) }
}

func BenchmarkBufferReplicate2x10(b *testing.B) {
	v := Buffer{0, 0}
	for i := 0; i < b.N; i++ { v.Replicate(10) }
}

func BenchmarkBufferReplicate6x2(b *testing.B) {
	v := Buffer{0, 0, 0, 0, 0, 0}
	for i := 0; i < b.N; i++ { v.Replicate(2) }
}

func BenchmarkBufferReplicate6x10(b *testing.B) {
	v := Buffer{0, 0, 0, 0, 0, 0}
	for i := 0; i < b.N; i++ { v.Replicate(10) }
}

func BenchmarkBufferAt(b *testing.B) {
	v := Buffer{0}
    for i := 0; i < b.N; i++ { v.At(0) }
}

func BenchmarkBufferFAt(b *testing.B) {
	v := Buffer{0}
    for i := 0; i < b.N; i++ { v.FAt(0) }
}

func BenchmarkBufferSet(b *testing.B) {
	v := Buffer{0}
    for i := 0; i < b.N; i++ { v.Set(0, 1) }
}

func BenchmarkBufferFSet(b *testing.B) {
	v := Buffer{0}
    for i := 0; i < b.N; i++ { v.FSet(0, 1.099) }
}

func BenchmarkBufferAdd(b *testing.B) {
	v := Buffer{0, 1}
    for i := 0; i < b.N; i++ { v.Add(0, 1) }
}

func BenchmarkBufferFAdd(b *testing.B) {
	b.StopTimer()
		v := make(Buffer, 2)
		v.FSet(0, 0.0, 1.0)
	b.StartTimer()
    for i := 0; i < b.N; i++ { v.FAdd(0, 1) }
}

func BenchmarkBufferSubtract(b *testing.B) {
	v := Buffer{0, 1}
    for i := 0; i < b.N; i++ { v.Subtract(0, 1) }
}

func BenchmarkBufferFSubtract(b *testing.B) {
	b.StopTimer()
		v := make(Buffer, 2)
		v.FSet(0, 0.0, 1.0)
	b.StartTimer()
    for i := 0; i < b.N; i++ { v.FSubtract(0, 1) }
}

func BenchmarkBufferMultiply(b *testing.B) {
	v := Buffer{0, 2}
    for i := 0; i < b.N; i++ { v.Multiply(0, 1) }
}

func BenchmarkBufferFMultiply(b *testing.B) {
	b.StopTimer()
		v := make(Buffer, 2)
		v.FSet(0, 1.0, 1.0)
	b.StartTimer()
    for i := 0; i < b.N; i++ { v.FMultiply(0, 1) }
}

func BenchmarkBufferDivide(b *testing.B) {
	v := Buffer{987654321, 2}
    for i := 0; i < b.N; i++ { v.Divide(0, 1) }
}

func BenchmarkBufferFDivide(b *testing.B) {
	b.StopTimer()
		v := make(Buffer, 2)
		v.FSet(0, 987654321.0, 1.01)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.FDivide(0, 1) }
}

func BenchmarkBufferIncrement(b *testing.B) {
	v := Buffer{0}
    for i := 0; i < b.N; i++ { v.Increment(0) }
}

func BenchmarkBufferDecrement(b *testing.B) {
	v := Buffer{0}
    for i := 0; i < b.N; i++ { v.Decrement(0) }
}

func BenchmarkBufferNegate(b *testing.B) {
	v := Buffer{100}
    for i := 0; i < b.N; i++ { v.Negate(0) }
}

func BenchmarkBufferShiftLeft(b *testing.B) {
	v := Buffer{1, 1}
    for i := 0; i < b.N; i++ { v.ShiftLeft(0, 1) }
}

func BenchmarkBufferShiftRight(b *testing.B) {
	v := Buffer{987654321, 1}
    for i := 0; i < b.N; i++ { v.ShiftRight(0, 1) }
}

func BenchmarkBufferInvert(b *testing.B) {
	v := Buffer{100}
    for i := 0; i < b.N; i++ { v.Invert(0) }
}
