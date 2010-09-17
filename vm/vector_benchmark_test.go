package vm
import "testing"

func BenchmarkVectorAdd1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkVectorAdd2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkVectorAdd6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkVectorFAdd1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 1)}
		s1.FSet(0, 100.00)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkVectorFAdd2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 2)}
		s1.FSet(0, 100.00, 200.0)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkVectorFAdd6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 6)}
		s1.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FAdd(0, s2) }
}

func BenchmarkVectorSubtract1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkVectorSubtract2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkVectorSubtract6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkVectorFSubtract1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 1)}
		s1.FSet(0, 100.00)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkVectorFSubtract2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 2)}
		s1.FSet(0, 100.00, 200.0)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkVectorFSubtract6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 6)}
		s1.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
}

func BenchmarkVectorMultiply1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkVectorMultiply2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkVectorMultiply6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkVectorFMultiply1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 1)}
		s1.FSet(0, 100.00)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkVectorFMultiply2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 2)}
		s1.FSet(0, 100.00, 200.0)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkVectorFMultiply6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{ make(Buffer, 6)}
		s1.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
}

func BenchmarkVectorDivide1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkVectorDivide2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkVectorDivide6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkVectorFDivide1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 1)}
		s1.FSet(0, 100.00)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkVectorFDivide2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 2)}
		s1.FSet(0, 100.00, 200.0)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkVectorFDivide6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{make(Buffer, 6)}
		s1.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
}

func BenchmarkVectorAnd1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkVectorAnd2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkVectorAnd6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkVectorOr1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkVectorOr2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkVectorOr6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkVectorXor1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkVectorXor2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkVectorXor6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{987654321, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkVectorIncrement1(b *testing.B) {
	s := &Vector{Buffer{100}}
	for i := 0; i < b.N; i++ { s.Increment(0, 1) }
}

func BenchmarkVectorIncrement2(b *testing.B) {
	s := &Vector{Buffer{100, 200}}
	for i := 0; i < b.N; i++ { s.Increment(0, 2) }
}

func BenchmarkVectorIncrement6(b *testing.B) {
	s := &Vector{Buffer{37, 101, 3, 5, 2, 2}}
	for i := 0; i < b.N; i++ { s.Increment(0, 6) }
}

func BenchmarkVectorDecrement1(b *testing.B) {
	s := &Vector{Buffer{100}}
	for i := 0; i < b.N; i++ { s.Decrement(0, 1) }
}

func BenchmarkVectorDecrement2(b *testing.B) {
	s := &Vector{Buffer{100, 200}}
	for i := 0; i < b.N; i++ { s.Decrement(0, 2) }
}

func BenchmarkVectorDecrement6(b *testing.B) {
	s := &Vector{Buffer{37, 101, 3, 5, 2, 2}}
	for i := 0; i < b.N; i++ { s.Decrement(0, 6) }
}

func BenchmarkVectorNegate1(b *testing.B) {
	s := &Vector{Buffer{100}}
	for i := 0; i < b.N; i++ { s.Negate(0, 1) }
}

func BenchmarkVectorNegate2(b *testing.B) {
	s := &Vector{Buffer{100, 200}}
	for i := 0; i < b.N; i++ { s.Negate(0, 2) }
}

func BenchmarkVectorNegate6(b *testing.B) {
	s := &Vector{Buffer{37, 101, 3, 5, 2, 2}}
	for i := 0; i < b.N; i++ { s.Negate(0, 6) }
}

func BenchmarkVectorShiftLeft1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkVectorShiftLeft2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkVectorShiftLeft6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{37, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkVectorShiftRight1(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkVectorShiftRight2(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{100, 200}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkVectorShiftRight6(b *testing.B) {
	b.StopTimer()
		s1 := &Vector{Buffer{37, 101, 3, 5, 2, 2}}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkVectorInvert1(b *testing.B) {
	s := &Vector{Buffer{100}}
	for i := 0; i < b.N; i++ { s.Invert(0, 1) }
}

func BenchmarkVectorInvert2(b *testing.B) {
	s := &Vector{Buffer{100, 200}}
	for i := 0; i < b.N; i++ { s.Invert(0, 2) }
}

func BenchmarkVectorInvert6(b *testing.B) {
	s := &Vector{Buffer{37, 101, 3, 5, 2, 2}}
	for i := 0; i < b.N; i++ { s.Invert(0, 6) }
}
