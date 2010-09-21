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

func BenchmarkIntBufferCollect1(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { v.Collect(func(x int) int { return x }) }
}

func BenchmarkIntBufferCollect10(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Collect(func(x int) int { return x }) }
}

func BenchmarkIntBufferCollect100(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(10)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Collect(func(x int) int { return x }) }
}

func BenchmarkIntBufferCollect1000(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Collect(func(x int) int { return x }) }
}

func BenchmarkIntBufferInject1(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { v.Inject(0, func(memo, x int) int { return memo }) }
}

func BenchmarkIntBufferInject10(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Inject(0, func(memo, x int) int { return memo }) }
}

func BenchmarkIntBufferInject100(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(10)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Inject(0, func(memo, x int) int { return memo }) }
}

func BenchmarkIntBufferInject1000(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Inject(0, func(memo, x int) int { return memo }) }
}

func BenchmarkIntBufferCycle10x1(b *testing.B) {
	v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ { v.Cycle(1, func(i, x int) {}) }
}

func BenchmarkIntBufferCycle10x10(b *testing.B) {
	v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ { v.Cycle(10, func(i, x int) {}) }
}

func BenchmarkIntBufferCycle10x100(b *testing.B) {
	v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ { v.Cycle(100, func(i, x int) {}) }
}

func BenchmarkIntBufferCycle10x1000(b *testing.B) {
	v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ { v.Cycle(1000, func(i, x int) {}) }
}

func BenchmarkIntBufferCycle1000x1(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Cycle(1, func(i, x int) {}) }
}

func BenchmarkIntBufferCycle1000x10(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Cycle(10, func(i, x int) {}) }
}

func BenchmarkIntBufferCycle1000x100(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Cycle(100, func(i, x int) {}) }
}

func BenchmarkIntBufferCycle1000x1000(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Cycle(1000, func(i, x int) {}) }
}

func BenchmarkIntBufferCombine1(b *testing.B) {
	v1 := IntBuffer{0}
	v2 := IntBuffer{0}
	for i := 0; i < b.N; i++ { v1.Combine(v2, func(x, y int) int { return x }) }
}

func BenchmarkIntBufferCombine10(b *testing.B) {
	b.StopTimer()
		v1 := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		v2 := v1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { v1.Combine(v2, func(x, y int) int { return x }) }
}

func BenchmarkIntBufferCombine100(b *testing.B) {
	b.StopTimer()
		v1 := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(10)
		v2 := v1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { v1.Combine(v2, func(x, y int) int { return x }) }
}

func BenchmarkIntBufferCombine1000(b *testing.B) {
	b.StopTimer()
		v1 := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
		v2 := v1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { v1.Combine(v2, func(x, y int) int { return x }) }
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

func BenchmarkIntBufferCount1(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { v.Count(func(x int) bool { return x == 0 }) }
}

func BenchmarkIntBufferCount10(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Count(func(x int) bool { return x == 0 }) }
}

func BenchmarkIntBufferCount100(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(10)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Count(func(x int) bool { return x == 0 }) }
}

func BenchmarkIntBufferCount1000(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Count(func(x int) bool { return x == 0 }) }
}

func BenchmarkIntBufferAny1Fail(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { v.Any(func(x int) bool { return x == 99 }) }
}

func BenchmarkIntBufferAny10Fail(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Any(func(x int) bool { return x == 99 }) }
}

func BenchmarkIntBufferAny100Fail(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(10)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Any(func(x int) bool { return x == 99 }) }
}

func BenchmarkIntBufferAny1000Fail(b *testing.B) {
	b.StopTimer()
		v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Replicate(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.Any(func(x int) bool { return x == 99 }) }
}

func BenchmarkIntBufferAll1Succeed(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { v.All(func(x int) bool { return x == 0 }) }
}

func BenchmarkIntBufferAll10Succeed(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 10)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.All(func(x int) bool { return x == 0 }) }
}

func BenchmarkIntBufferAll100Succeed(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.All(func(x int) bool { return x == 0 }) }
}

func BenchmarkIntBufferAll1000Succeed(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.All(func(x int) bool { return x == 0 }) }
}

func BenchmarkIntBufferNone1Succeed(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { v.None(func(x int) bool { return x == 1 }) }
}

func BenchmarkIntBufferNone10Succeed(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 10)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.None(func(x int) bool { return x == 1 }) }
}

func BenchmarkIntBufferNone100Succeed(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.None(func(x int) bool { return x == 1 }) }
}

func BenchmarkIntBufferNone1000Succeed(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.None(func(x int) bool { return x == 1 }) }
}


func BenchmarkIntBufferOne1Full(b *testing.B) {
	v := IntBuffer{1}
	for i := 0; i < b.N; i++ { v.One(func(x int) bool { return x == 1 }) }
}

func BenchmarkIntBufferOne10Full(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 10)
		v[9] = 1
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.One(func(x int) bool { return x == 1 }) }
}

func BenchmarkIntBufferOne100Full(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 100)
		v[99] = 1
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.One(func(x int) bool { return x == 1 }) }
}

func BenchmarkIntBufferOne1000Full(b *testing.B) {
	b.StopTimer()
		v := make(IntBuffer, 1000)
		v[999] = 1
	b.StartTimer()
	for i := 0; i < b.N; i++ { v.None(func(x int) bool { return x == 1 }) }
}

func BenchmarkIntBufferConvertByteSlice1(b *testing.B) {
	v := IntBuffer{0}
	for i := 0; i < b.N; i++ { AsByteSlice(v) }
}

func BenchmarkIntBufferConvertByteSlice10(b *testing.B) {
	v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ { AsByteSlice(v) }
}

func BenchmarkIntBufferConvertByteSlice100(b *testing.B) {
	v := (IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}).Replicate(10)
	for i := 0; i < b.N; i++ { AsByteSlice(v) }
}

//func BenchmarkIntBufferConvertFloatBuffer1(b *testing.B) {
//	v := IntBuffer{0}
//	for i := 0; i < b.N; i++ { AsFloatBuffer(v) }
//}

//func BenchmarkIntBufferConvertFloatBuffer10(b *testing.B) {
//	v := IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	for i := 0; i < b.N; i++ { AsFloatBuffer(v) }
//}
