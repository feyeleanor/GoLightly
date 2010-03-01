package vm
import "testing"

func oneIntegerVector() *Vector {
	return &Vector{Buffer{100}}
}

func oneFloatVector() *Vector {
	s := new(Vector)
	s.Resize(1)
	s.Buffer.FSet(0, 100.00)
	return s
}

func twoIntegerVector() *Vector {
	return &Vector{*twoIntegerBuffer()}
}

func twoFloatVector() *Vector {
	s := new(Vector)
	s.Resize(2)
	s.Buffer.Set(0, 100.00)
	s.Buffer.Set(1, 200.00)
	return s
}

func BenchmarkVectorAdd1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkVectorAdd2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkVectorAdd6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkVectorFAdd1(b *testing.B) {
	b.StopTimer()
	s1 := oneFloatVector()
	s2 := oneFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkVectorFAdd2(b *testing.B) {
	b.StopTimer()
	s1 := twoFloatVector()
	s2 := twoFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkVectorFAdd6(b *testing.B) {
	b.StopTimer()
	s1 := sixFloatVector()
	s2 := sixFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkVectorSubtract1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkVectorSubtract2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkVectorSubtract6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkVectorFSubtract1(b *testing.B) {
	b.StopTimer()
	s1 := oneFloatVector()
	s2 := oneFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkVectorFSubtract2(b *testing.B) {
	b.StopTimer()
	s1 := twoFloatVector()
	s2 := twoFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkVectorFSubtract6(b *testing.B) {
	b.StopTimer()
	s1 := sixFloatVector()
	s2 := sixFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkVectorMultiply1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkVectorMultiply2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkVectorMultiply6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkVectorFMultiply1(b *testing.B) {
	b.StopTimer()
	s1 := oneFloatVector()
	s2 := oneFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkVectorFMultiply2(b *testing.B) {
	b.StopTimer()
	s1 := twoFloatVector()
	s2 := twoFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkVectorFMultiply6(b *testing.B) {
	b.StopTimer()
	s1 := sixFloatVector()
	s2 := sixFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkVectorDivide1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s1.Set(0, 987654321)
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkVectorDivide2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s1.Set(0, 987654321)
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkVectorDivide6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s1.Set(0, 987654321)
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkVectorFDivide1(b *testing.B) {
	b.StopTimer()
	s1 := oneFloatVector()
	s2 := oneFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkVectorFDivide2(b *testing.B) {
	b.StopTimer()
	s1 := twoFloatVector()
	s2 := twoFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkVectorFDivide6(b *testing.B) {
	b.StopTimer()
	s1 := sixFloatVector()
	s2 := sixFloatVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkVectorAnd1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s1.Set(0, 987654321)
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkVectorAnd2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s1.Set(0, 987654321)
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkVectorAnd6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s1.Set(0, 987654321)
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkVectorOr1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s1.Set(0, 987654321)
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkVectorOr2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s1.Set(0, 987654321)
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkVectorOr6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s1.Set(0, 987654321)
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkVectorXor1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s1.Set(0, 987654321)
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkVectorXor2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s1.Set(0, 987654321)
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkVectorXor6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s1.Set(0, 987654321)
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkVectorIncrement1(b *testing.B) {
	b.StopTimer()
	s := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment(0, s.Len()) }
}

func BenchmarkVectorIncrement2(b *testing.B) {
	b.StopTimer()
	s := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment(0, s.Len()) }
}

func BenchmarkVectorIncrement6(b *testing.B) {
	b.StopTimer()
	s := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Increment(0, s.Len()) }
}

func BenchmarkVectorDecrement1(b *testing.B) {
	b.StopTimer()
	s := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement(0, s.Len()) }
}

func BenchmarkVectorDecrement2(b *testing.B) {
	b.StopTimer()
	s := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement(0, s.Len()) }
}

func BenchmarkVectorDecrement6(b *testing.B) {
	b.StopTimer()
	s := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Decrement(0, s.Len()) }
}

func BenchmarkVectorNegate1(b *testing.B) {
	b.StopTimer()
	s := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate(0, s.Len()) }
}

func BenchmarkVectorNegate2(b *testing.B) {
	b.StopTimer()
	s := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate(0, s.Len()) }
}

func BenchmarkVectorNegate6(b *testing.B) {
	b.StopTimer()
	s := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Negate(0, s.Len()) }
}

func BenchmarkVectorShiftLeft1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkVectorShiftLeft2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkVectorShiftLeft6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkVectorShiftRight1(b *testing.B) {
	b.StopTimer()
	s1 := oneIntegerVector()
	s2 := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkVectorShiftRight2(b *testing.B) {
	b.StopTimer()
	s1 := twoIntegerVector()
	s2 := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkVectorShiftRight6(b *testing.B) {
	b.StopTimer()
	s1 := sixIntegerVector()
	s2 := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkVectorInvert1(b *testing.B) {
	b.StopTimer()
	s := oneIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert(0, s.Len()) }
}

func BenchmarkVectorInvert2(b *testing.B) {
	b.StopTimer()
	s := twoIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert(0, s.Len()) }
}

func BenchmarkVectorInvert6(b *testing.B) {
	b.StopTimer()
	s := sixIntegerVector()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s.Invert(0, s.Len()) }
}
