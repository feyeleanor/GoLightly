package vm

import "testing"
import . "golightly/storage"
import . "golightly/test"

func TestVector(t *testing.T) {
	b := IntBuffer{37, 101, 3, 5, 2, 2}
	v1 := &Vector{b}

	NewTest(t).
	Run("Creation", func(TC *Test) {
		v2 := v1.Clone()
		TC.	Identical(v1, &Vector{IntBuffer{37, 101, 3, 5, 2, 2}}).
			Identical(v1, v2)
	}).
	Run("Integer Maths", func(TC *Test) {
		s := v1.Slice(1, 3)
		TC.	Identical(s.Len(), 2).
			Identical(s.Cap(), v1.Cap() - 1).
			Identical(s, &Vector{IntBuffer{101, 3}})

		v2 := v1.Clone()
		v2.Add(0, v1)
		TC.	Identical(v2, &Vector{IntBuffer{74, 202, 6, 10, 4, 4}})

		v2 = v1.Clone()
		v2.Subtract(0, v1)
		TC.	Identical(v2, &Vector{IntBuffer{0, 0, 0, 0, 0, 0}})

		v2 = v1.Clone()
		v2.Multiply(0, v1)
		TC.	Identical(v2, &Vector{IntBuffer{1369, 10201, 9, 25, 4, 4}})

		v2 = v1.Clone()
		v2.Divide(0, v1)
		TC.	Identical(v2, &Vector{IntBuffer{1, 1, 1, 1, 1, 1}})

		v2 = v1.Clone()
		v2.Negate(0, v2.Len())
		TC.	Identical(v2, &Vector{IntBuffer{-37, -101, -3, -5, -2, -2}})

		v2 = v1.Clone()
		v2.Increment(0, v2.Len())
		TC.	Identical(v2, &Vector{IntBuffer{38, 102, 4, 6, 3, 3}})

		v2 = v1.Clone()
		v2.Decrement(0, v2.Len())
		TC.	Identical(v2, &Vector{IntBuffer{36, 100, 2, 4, 1, 1}})
	}).
	Run("Bit manipulation", func(TC *Test) {
		v2 := v1.Clone()
		v2.IntBuffer.ShiftRight(0, 5)
		TC.	Identical(v2.IntBuffer[0], 9)

		v2.IntBuffer.ShiftLeft(0, 5)
		TC.	Identical(v2.IntBuffer[0], 36)

		v2.IntBuffer.Invert(0)
		TC.	Identical(v2.IntBuffer[0], ^36, -37)
	}).
//	Run("Floating-point Maths", func(TC *Test) {
//		v2 := v1.Clone()
//		v2.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
//		v3 := v2.Clone()
//		v3.FAdd(0, v2)
//		TC.	Identical(v3.FAt(0), 74.0).
//			Identical(v3.FAt(1), 202.0).
//			Identical(v3.FAt(2), 7.4).
//			Identical(v3.FAt(3), 10.0).
//			Identical(v3.FAt(4), 4.0).
//			Identical(v3.FAt(5), 4.0)
//
//		v3 = v2.Clone()
//		v3.FSubtract(0, v2)
//		TC.	Identical(v3.FAt(0), 0.0).
//			Identical(v3.FAt(1), 0.0).
//			Identical(v3.FAt(2), 0.0).
//			Identical(v3.FAt(3), 0.0).
//			Identical(v3.FAt(4), 0.0).
//			Identical(v3.FAt(5), 0.0)
//
//		v3 = v2.Clone()
//		v3.FMultiply(0, v2)
//		TC.	Identical(v3.FAt(0), 1369.0).
//			Identical(v3.FAt(1), 10201.0).
//			Identical(v3.FAt(2), 13.690001).
//			Identical(v3.FAt(3), 25.0).
//			Identical(v3.FAt(4), 4.0).
//			Identical(v3.FAt(5), 4.0)
//	}).
	Run("Integer Logic", func(TC *Test) {
		v2 := v1.Clone()
		b := v2.IntBuffer
		TC.	Confirm(b.Less(2, 3)).
			Refute(b.Equal(2, 3)).
			Refute(b.Greater(2, 3)).
			Refute(b.ZeroLess(2)).
			Refute(b.ZeroEqual(2)).
			Confirm(b.ZeroGreater(2))

		v2.Copy(1, 2)
		TC.	Different(v1, v2).
			Identical(v2.IntBuffer[1], 3).
			Confirm(b.Less(1, 3)).
			Confirm(b.Equal(1, 2)).
			Refute(b.Greater(1, 3)).
			Refute(b.ZeroLess(1)).
			Refute(b.ZeroEqual(1)).
			Confirm(b.ZeroGreater(1))

		v2.IntBuffer[1] = 0
		TC.	Different(v1, v2).
			Confirm(b.Less(1, 3)).
			Refute(b.Equal(1, 3)).
			Refute(b.Greater(1, 3)).
			Refute(b.ZeroLess(1)).
			Confirm(b.ZeroEqual(1)).
			Refute(b.ZeroGreater(1))
	}).
	Run("To Do", func(TC *Test) {
		TC.	Untested("GetIntBuffer").
			Untested("PutIntBuffer").
			Untested("Clear").
			ToDo("Improve tests for Vector-level operations")
	})
}
