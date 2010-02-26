package vm
import "testing"
import "unsafe"

func floatArray() []float {
	return []float{1.0, 2.0, 3.0, 4.5, 5.4, 6.0, 7.0, 8.0, 9.0, 10.0}
}

func intsToFloats(i []int) []float { return *(*[]float)(unsafe.Pointer(&i)) }
func floatsToInts(f []float) []int { return *(*[]int)(unsafe.Pointer(&f)) }

func BenchmarkBufferFloatsToInts(b *testing.B) {
	b.StopTimer()
	a := floatArray()
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkBufferFloatsFromInts(b *testing.B) {
	b.StopTimer()
	a := floatsToInts(floatArray())
	b.StartTimer()
    for i := 0; i < b.N; i++ { f := intsToFloats(a); f[0] += 2.0 }
}

func BenchmarkBufferReplicate2x2(b *testing.B) {
	b.StopTimer()
	buffer := twoIntegerBuffer()
	b.StartTimer()
	for i := 0; i < b.N; i++ { buffer.Replicate(2) }
}

func BenchmarkBufferReplicate2x10(b *testing.B) {
	b.StopTimer()
	buffer := twoIntegerBuffer()
	b.StartTimer()
	for i := 0; i < b.N; i++ { buffer.Replicate(10) }
}

func BenchmarkBufferReplicate6x2(b *testing.B) {
	b.StopTimer()
	buffer := sixIntegerBuffer()
	b.StartTimer()
	for i := 0; i < b.N; i++ { buffer.Replicate(2) }
}

func BenchmarkBufferReplicate6x10(b *testing.B) {
	b.StopTimer()
	buffer := sixIntegerBuffer()
	b.StartTimer()
	for i := 0; i < b.N; i++ { buffer.Replicate(10) }
}

func BenchmarkBufferAt(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.At(0) }
}

func BenchmarkBufferFAt(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FAt(0) }
}

func BenchmarkBufferSet(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Set(0, 1) }
}

func BenchmarkBufferFSet(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FSet(0, 1.099) }
}

func BenchmarkBufferAdd(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Add(0, 1) }
}

func BenchmarkBufferFAdd(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.FSet(0, 0.0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FAdd(0, 1.0) }
}

func BenchmarkBufferSubtract(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Subtract(0, 1) }
}

func BenchmarkBufferFSubtract(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.FSet(0, 0.0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FSubtract(0, 1.0) }
}

func BenchmarkBufferMultiply(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.Set(0, 2)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Multiply(0, 1) }
}

func BenchmarkBufferFMultiply(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.FSet(0, 1.0)
	buffer.FSet(1, 1.0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FMultiply(0, 1.0) }
}

func BenchmarkBufferDivide(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.Set(0, 987654321)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Divide(0, 1) }
}

func BenchmarkBufferFDivide(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.FSet(0, 987654321.0)
	buffer.FSet(1, 2.0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FDivide(0, 1.0) }
}

func BenchmarkBufferIncrement(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.Set(0, 0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Decrement(0) }
}

func BenchmarkBufferDecrement(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.Set(0, 0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Decrement(0) }
}

func BenchmarkBufferNegate(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Negate(0) }
}

func BenchmarkBufferShiftLeft(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.ShiftLeft(0, 1) }
}

func BenchmarkBufferShiftRight(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.ShiftRight(0, 1) }
}

func BenchmarkBufferInvert(b *testing.B) {
    b.StopTimer()
	buffer := sixIntegerBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Invert(0) }
}
