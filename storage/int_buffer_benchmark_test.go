package storage

import "testing"

func BenchmarkIntBufferReadIndex(b *testing.B) {
	x := 0
	v := IntBuffer{x}
    for i := 0; i < b.N; i++ { x = v[0] }
}

func BenchmarkIntBufferWriteIndex(b *testing.B) {
	x := 0
	v := IntBuffer{x}
    for i := 0; i < b.N; i++ { v[0] = x }
}

func BenchmarkIntBufferWriteIndex10(b *testing.B) {
	v := make(IntBuffer, 10)
    for i := 0; i < b.N; i++ {
		for j := 0; j < len(v); j++ { v[j] = 1 }
	}
}

func BenchmarkIntBufferWriteIndex25(b *testing.B) {
	v := make(IntBuffer, 25)
    for i := 0; i < b.N; i++ {
		for j := 0; j < len(v); j++ { v[j] = 1 }
	}
}

func BenchmarkIntBufferOverwrite1(b *testing.B) {
	v := IntBuffer{0}
	d := IntBuffer{1}
    for i := 0; i < b.N; i++ { copy(v, d) }
}

func BenchmarkIntBufferOverwrite10(b *testing.B) {
	v := make(IntBuffer, 10)
	d := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    for i := 0; i < b.N; i++ { copy(v, d) }
}

func BenchmarkIntBufferOverwrite25(b *testing.B) {
	v := make(IntBuffer, 25)
	d := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4}
    for i := 0; i < b.N; i++ { copy(v, d) }
}

func BenchmarkIntBufferAdd(b *testing.B) {
	v := IntBuffer{0, 1}
    for i := 0; i < b.N; i++ { v.Add(0, 1) }
}

func BenchmarkIntBufferSubtract(b *testing.B) {
	v := IntBuffer{0, 1}
    for i := 0; i < b.N; i++ { v.Subtract(0, 1) }
}

func BenchmarkIntBufferMultiply(b *testing.B) {
	v := IntBuffer{0, 2}
    for i := 0; i < b.N; i++ { v.Multiply(0, 1) }
}

func BenchmarkIntBufferDivide(b *testing.B) {
	v := IntBuffer{987654321, 2}
    for i := 0; i < b.N; i++ { v.Divide(0, 1) }
}

func BenchmarkIntBufferIncrement(b *testing.B) {
	v := IntBuffer{0}
    for i := 0; i < b.N; i++ { v.Increment(0) }
}

func BenchmarkIntBufferDecrement(b *testing.B) {
	v := IntBuffer{0}
    for i := 0; i < b.N; i++ { v.Decrement(0) }
}

func BenchmarkIntBufferNegate(b *testing.B) {
	v := IntBuffer{100}
    for i := 0; i < b.N; i++ { v.Negate(0) }
}

func BenchmarkIntBufferShiftLeft(b *testing.B) {
	v := IntBuffer{1, 1}
    for i := 0; i < b.N; i++ { v.ShiftLeft(0, 1) }
}

func BenchmarkIntBufferShiftRight(b *testing.B) {
	v := IntBuffer{987654321, 1}
    for i := 0; i < b.N; i++ { v.ShiftRight(0, 1) }
}

func BenchmarkIntBufferInvert(b *testing.B) {
	v := IntBuffer{100}
    for i := 0; i < b.N; i++ { v.Invert(0) }
}

func BenchmarkIntBufferConvertByteSlice1(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { v.ByteSlice() }
}

func BenchmarkIntBufferConvertByteSlice10(b *testing.B) {
	v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ { v.ByteSlice() }
}

func BenchmarkIntBufferConvertFloatBuffer1(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { v.FloatBuffer() }
}

func BenchmarkIntBufferConvertFloatBuffer10(b *testing.B) {
	v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ { v.FloatBuffer() }
}
