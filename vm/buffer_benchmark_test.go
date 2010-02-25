package vm
import "testing"
import "container/vector"
//import "unsafe"

func BenchmarkBufferFloatsToInts(b *testing.B) {
	b.StopTimer()
	a := floatBuffer()
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkBufferFloatsFromInts(b *testing.B) {
	b.StopTimer()
	a := floatsToInts(floatBuffer())
	b.StartTimer()
    for i := 0; i < b.N; i++ { f := intsToFloats(a); f[0] += 2.0 }
}

func BenchmarkBufferAt(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.At(0) }
}

func BenchmarkBufferFAt(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FAt(0) }
}

func BenchmarkBufferSet(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Set(0, 1) }
}

func BenchmarkBufferFSet(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FSet(0, 1.099) }
}

func BenchmarkBufferAdd(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Add(0, 1) }
}

func BenchmarkBufferFAdd(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.FSet(0, 0.0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FAdd(0, 1.0) }
}

func BenchmarkBufferSubtract(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Subtract(0, 1) }
}

func BenchmarkBufferFSubtract(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.FSet(0, 0.0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FSubtract(0, 1.0) }
}

func BenchmarkBufferMultiply(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 2)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Multiply(0, 1) }
}

func BenchmarkBufferFMultiply(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.FSet(0, 1.0)
	buffer.FSet(1, 1.0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FMultiply(0, 1.0) }
}

func BenchmarkBufferDivide(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 987654321)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Divide(0, 1) }
}

func BenchmarkBufferFDivide(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.FSet(0, 987654321.0)
	buffer.FSet(1, 2.0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FDivide(0, 1.0) }
}

func BenchmarkBufferIncrement(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Decrement(0) }
}

func BenchmarkBufferDecrement(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Decrement(0) }
}

func BenchmarkBufferNegate(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Negate(0) }
}

func BenchmarkBufferShiftLeft(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.ShiftLeft(0, 1) }
}

func BenchmarkBufferShiftRight(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.ShiftRight(0, 1) }
}

func BenchmarkBufferInvert(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Invert(0) }
}

func BenchmarkValueAt(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.At(0) }
}

func BenchmarkValueSet(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.Set(0, 1) }
}

func BenchmarkValueAdd(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
	a := *v
    b.StartTimer()
//    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) + 1) }
    for i := 0; i < b.N; i++ { a[0] = a[0].(int) + 1 }
}

func BenchmarkValueSubtract(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
	a := *v
    b.StartTimer()
//    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) - 1) }
    for i := 0; i < b.N; i++ { a[0] = a[0].(int) - 1 }
}

func BenchmarkValueMultiply(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(2)
	a := *v
    b.StartTimer()
//    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) * 27) }
    for i := 0; i < b.N; i++ { a[0] = a[0].(int) * 27 }
}

func BenchmarkValueDivide(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(987654321)
	a := *v
    b.StartTimer()
//    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) / 4) }
    for i := 0; i < b.N; i++ { a[0] = a[0].(int) / 1 }
}
