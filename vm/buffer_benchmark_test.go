package vm
import "testing"
import "container/vector"
import "big"

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
    for i := 0; i < b.N; i++ { buffer.Multiply(0, 27) }
}

func BenchmarkBufferDivide(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 987654321)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Divide(0, 4) }
}

func BenchmarkBigAdd(b *testing.B) {
    b.StopTimer()
	i0 := big.NewInt(0)
	i1 := big.NewInt(1)
    b.StartTimer()
    for i := 0; i < b.N; i++ { i0.Add(i0, i1) }
}

func BenchmarkBigSubtract(b *testing.B) {
    b.StopTimer()
	i0 := big.NewInt(0)
	i1 := big.NewInt(1)
    b.StartTimer()
    for i := 0; i < b.N; i++ { i0.Sub(i0, i1) }
}

func BenchmarkBigMultiply(b *testing.B) {
    b.StopTimer()
	i0 := big.NewInt(2)
	i1 := big.NewInt(27)
    b.StartTimer()
    for i := 0; i < b.N; i++ { i1.Mul(i0, i1) }
}

func BenchmarkBigDivide(b *testing.B) {
    b.StopTimer()
	i0 := big.NewInt(987654321)
	i1 := big.NewInt(4)
    b.StartTimer()
    for i := 0; i < b.N; i++ { i1.Div(i0, i1) }
}

func BenchmarkValueAdd(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) + 1) }
}

func BenchmarkValueSubtract(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) - 1) }
}

func BenchmarkValueMultiply(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(2)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) * 27) }
}

func BenchmarkValueDivide(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(987654321)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) / 4) }
}
