package vm
import "testing"
import "container/vector"
import "unsafe"

func floatBuffer() []float {
	return []float{1.0, 2.0, 3.0, 4.5, 5.4, 6.0, 7.0, 8.0, 9.0, 10.0}
}

func intsToFloats(i []int) []float { return *(*[]float)(unsafe.Pointer(&i)) }
func floatsToInts(f []float) []int { return *(*[]int)(unsafe.Pointer(&f)) }

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

func BenchmarkBufferFloatAt(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.FloatAt(0) }
}

func BenchmarkBufferFloatAdds(b *testing.B) {
	b.StopTimer()
	a := []float{ 3.2, 4.7, 0.0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ { a[2] = a[0] + a[1] }
}

func BenchmarkBufferSet(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Set(0, 1) }
}

func BenchmarkBufferFloatSet(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { a := 1.099; buffer.Set(0, *(*int)(unsafe.Pointer(&a))) }
}

func BenchmarkBufferAdd(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Add(0, 1) }
}

func BenchmarkBufferSubtract(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Subtract(0, 1) }
}

func BenchmarkBufferMultiply(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 2)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Multiply(0, 1) }
}

func BenchmarkBufferDivide(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 987654321)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Divide(0, 1) }
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
