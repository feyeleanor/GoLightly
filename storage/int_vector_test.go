package storage

import "testing"
import . "golightly/test"

func TestVector(t *testing.T) {
	v1 := IntVector{37, 101, 3, 5, 2, 2}

	NewTest(t).
	Run("Creation", func(TC *Test) {
		v2 := v1.Clone()
		TC.	Identical(v1, IntVector{37, 101, 3, 5, 2, 2}).
			Identical(v1, v2)
	}).
	Run("Integer Maths", func(TC *Test) {
		s := v1[1:]
		TC.	Identical(len(s), len(v1) - 1).
			Identical(cap(s), cap(v1) - 1).
			Identical(s, IntVector{101, 3, 5, 2, 2})

		v2 := v1.Clone()
		v2.Add(0, v1)
		TC.	Identical(v2, IntVector{74, 202, 6, 10, 4, 4})

		v2 = v1.Clone()
		v2.Subtract(0, v1)
		TC.	Identical(v2, IntVector{0, 0, 0, 0, 0, 0})

		v2 = v1.Clone()
		v2.Multiply(0, v1)
		TC.	Identical(v2, IntVector{1369, 10201, 9, 25, 4, 4})

		v2 = v1.Clone()
		v2.Divide(0, v1)
		TC.	Identical(v2, IntVector{1, 1, 1, 1, 1, 1})

		v2 = v1.Clone()
		v2.Negate(0, len(v2))
		TC.	Identical(v2, IntVector{-37, -101, -3, -5, -2, -2})

		v2 = v1.Clone()
		v2.Increment(0, len(v2))
		TC.	Identical(v2, IntVector{38, 102, 4, 6, 3, 3})

		v2 = v1.Clone()
		v2.Decrement(0, len(v2))
		TC.	Identical(v2, IntVector{36, 100, 2, 4, 1, 1})
	}).
	Run("Bit manipulation", func(TC *Test) {
		v2 := v1.Clone()
		v2.ShiftRight(0, IntVector{5})
		TC.	Identical(v2[0], 1)

		v2.ShiftLeft(0, IntVector{5})
		TC.	Identical(v2[0], 32)

		v2.Invert(0, 1)
		TC.	Identical(v2[0], ^32, -33)
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
		b := IntBuffer(v2)
		TC.	Confirm(b.Less(2, 3)).
			Refute(b.Equal(2, 3)).
			Refute(b.Greater(2, 3)).
			Refute(b.ZeroLess(2)).
			Refute(b.ZeroEqual(2)).
			Confirm(b.ZeroGreater(2))

		v2[1] = v2[2]
		TC.	Different(v1, v2).
			Identical(v2[1], 3).
			Confirm(b.Less(1, 3)).
			Confirm(b.Equal(1, 2)).
			Refute(b.Greater(1, 3)).
			Refute(b.ZeroLess(1)).
			Refute(b.ZeroEqual(1)).
			Confirm(b.ZeroGreater(1))

		v2[1] = 0
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