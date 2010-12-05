package storage

import "testing"

func BenchmarkIntVectorAdd1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkIntVectorAdd2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

func BenchmarkIntVectorAdd6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Add(0, s2) }
}

//func BenchmarkIntVectorFAdd1(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 1)
//		f := s1.FloatBuffer()
//		f[0] = 100.00
//		f2 := f.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { f.Add(0, f2) }
//}

//func BenchmarkIntVectorFAdd2(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 2)
//		f := s1.FloatBuffer()
//		f[0] = 100.00
//		f[1] = 200.0
//		f2 := f.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { f.Add(0, f2) }
//}

//func BenchmarkIntVectorFAdd6(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 6)
//		f := s1.FloatBuffer()
//		f[0] = 37.0
//		f[1] = 101.0
//		f[2] = 3.7
//		f[3] = 5.0
//		f[4] = 2.0
//		f[5] = 2.0
//		f2 := f.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { f.Add(0, f2) }
//}

func BenchmarkIntVectorSubtract1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkIntVectorSubtract2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

func BenchmarkIntVectorSubtract6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Subtract(0, s2) }
}

//func BenchmarkIntVectorFSubtract1(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 1)
//		f := s1.FloatBuffer()
//		f[0] = 100.00
//		f2 := f.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { f.Subtract(0, f2) }
//}

//func BenchmarkIntVectorFSubtract2(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 2)
//		f := s1.FloatBuffer()
//		f[0] = 100.00
//		f[1] = 200.0
//		f := f.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { f.Subtract(0, f2) }
//}

//func BenchmarkIntVectorFSubtract6(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 6)
//		s1.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
//		s2 := s1.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { s1.FSubtract(0, s2) }
//}

func BenchmarkIntVectorMultiply1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkIntVectorMultiply2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

func BenchmarkIntVectorMultiply6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Multiply(0, s2) }
}

//func BenchmarkIntVectorFMultiply1(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 1)
//		s1.FSet(0, 100.00)
//		s2 := s1.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
//}

//func BenchmarkIntVectorFMultiply2(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 2)
//		s1.FSet(0, 100.00, 200.0)
//		s2 := s1.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
//}

//func BenchmarkIntVectorFMultiply6(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 6)
//		s1.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
//		s2 := s1.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { s1.FMultiply(0, s2) }
//}

func BenchmarkIntVectorDivide1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkIntVectorDivide2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

func BenchmarkIntVectorDivide6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Divide(0, s2) }
}

//func BenchmarkIntVectorFDivide1(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 1)
//		s1.FSet(0, 100.00)
//		s2 := s1.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
//}

//func BenchmarkIntVectorFDivide2(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 2)
//		s1.FSet(0, 100.00, 200.0)
//		s2 := s1.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
//}

//func BenchmarkIntVectorFDivide6(b *testing.B) {
//	b.StopTimer()
//		s1 := make(IntVector, 6)
//		s1.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
//		s2 := s1.Clone()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ { s1.FDivide(0, s2) }
//}

func BenchmarkIntVectorAnd1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkIntVectorAnd2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkIntVectorAnd6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.And(0, s2) }
}

func BenchmarkIntVectorOr1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkIntVectorOr2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkIntVectorOr6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Or(0, s2) }
}

func BenchmarkIntVectorXor1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkIntVectorXor2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkIntVectorXor6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{987654321, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.Xor(0, s2) }
}

func BenchmarkIntVectorIncrement1(b *testing.B) {
	s := IntVector{100}
	for i := 0; i < b.N; i++ { s.Increment(0, 1) }
}

func BenchmarkIntVectorIncrement2(b *testing.B) {
	s := IntVector{100, 200}
	for i := 0; i < b.N; i++ { s.Increment(0, 2) }
}

func BenchmarkIntVectorIncrement6(b *testing.B) {
	s := IntVector{37, 101, 3, 5, 2, 2}
	for i := 0; i < b.N; i++ { s.Increment(0, 6) }
}

func BenchmarkIntVectorDecrement1(b *testing.B) {
	s := IntVector{100}
	for i := 0; i < b.N; i++ { s.Decrement(0, 1) }
}

func BenchmarkIntVectorDecrement2(b *testing.B) {
	s := IntVector{100, 200}
	for i := 0; i < b.N; i++ { s.Decrement(0, 2) }
}

func BenchmarkIntVectorDecrement6(b *testing.B) {
	s := IntVector{37, 101, 3, 5, 2, 2}
	for i := 0; i < b.N; i++ { s.Decrement(0, 6) }
}

func BenchmarkIntVectorNegate1(b *testing.B) {
	s := IntVector{100}
	for i := 0; i < b.N; i++ { s.Negate(0, 1) }
}

func BenchmarkIntVectorNegate2(b *testing.B) {
	s := IntVector{100, 200}
	for i := 0; i < b.N; i++ { s.Negate(0, 2) }
}

func BenchmarkIntVectorNegate6(b *testing.B) {
	s := IntVector{37, 101, 3, 5, 2, 2}
	for i := 0; i < b.N; i++ { s.Negate(0, 6) }
}

func BenchmarkIntVectorShiftLeft1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkIntVectorShiftLeft2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkIntVectorShiftLeft6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{37, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftLeft(0, s2) }
}

func BenchmarkIntVectorShiftRight1(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkIntVectorShiftRight2(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{100, 200}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkIntVectorShiftRight6(b *testing.B) {
	b.StopTimer()
		s1 := IntVector{37, 101, 3, 5, 2, 2}
		s2 := s1.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ { s1.ShiftRight(0, s2) }
}

func BenchmarkIntVectorInvert1(b *testing.B) {
	s := IntVector{100}
	for i := 0; i < b.N; i++ { s.Invert(0, 1) }
}

func BenchmarkIntVectorInvert2(b *testing.B) {
	s := IntVector{100, 200}
	for i := 0; i < b.N; i++ { s.Invert(0, 2) }
}

func BenchmarkIntVectorInvert6(b *testing.B) {
	s := IntVector{37, 101, 3, 5, 2, 2}
	for i := 0; i < b.N; i++ { s.Invert(0, 6) }
}