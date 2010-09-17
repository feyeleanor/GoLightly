package vm

import "testing"
import . "golightly/test"

func TestBuffer(t *testing.T) {
	NewTest(t).
	Run("Creation", func(TC *Test) {
		b1 := make(Buffer, 6)
		TC.	Identical(b1, Buffer{0, 0, 0, 0, 0, 0})
		b1 = Buffer{1, 2, 3, 4}
		b2 := *(b1.Clone())
		TC.	Identical(b1, b2).
			Identical(b1[0], b1.At(0)).
			Identical(b1[1], b1.At(1)).
			Identical(b1[2], b1.At(2)).
			Identical(b1[3], b1.At(3))
	}).
	Run("Buffer Manipulation", func(TC *Test) {
		b := make(Buffer, 6)
		TC.	Identical(b, Buffer{0, 0, 0, 0, 0, 0})

		b.Set(0, 1, 2, 3, 4, 5, 6)
		TC.	Identical(b, Buffer{1, 2, 3, 4, 5, 6})

		b.FSet(0, 2.0, 4.0, 6.0, 8.0)
		TC.	Identical(b.FAt(0), 2.0).
			Identical(b.FAt(1), 4.0).
			Identical(b.FAt(2), 6.0).
			Identical(b.FAt(3), 8.0)
	}).
	Run("Replication", func(TC *Test) {
		b1 := Buffer{20, 30}
		b2 := *(b1.Replicate(3))
		TC.	Identical(len(b2),	b2.Len(),	cap(b2),	b2.Cap(),	6).
			Identical(b1[0],	b2[0],		b2[2],		b2[4]).
			Identical(b1[1],	b2[1],		b2[3],		b2[5]).
			Identical(b1.At(0),	b2.At(0),	b2.At(2),	b2.At(4)).
			Identical(b1.At(1),	b2.At(1),	b2.At(3),	b2.At(5))
	}).
	Run("Slicing", func(TC *Test) {
		b1 := Buffer{0, 1, 2, 3}
		b2 := *(b1.Slice(1, 3))
		TC. Identical(b2.Len(),	2).
			Identical(b2.Cap(),	2).
			Identical(b2[0],	1).
			Identical(b2[1],	2)
	}).
	Run("Integer Maths", func(TC *Test) {
		b := Buffer{0}
		b.Increment(0)
		TC.Identical(b[0], 1)

		b.Decrement(0)
		TC.Identical(b[0], 0)

		b = Buffer{0, 1}
		b.Add(0, 1)
		TC.Identical(b[0], 1)

		b.Subtract(0, 1)
		TC.Identical(b[0], 0)

		b = Buffer{5}
		b.Negate(0)
		TC.Identical(b[0], -5)
		b.Negate(0)
		TC.Identical(b[0], 5)

		b = Buffer{2, 5}
		b.Multiply(0, 1)
		TC.Identical(b[0], 10)
		b.Divide(0, 1)
		TC.Identical(b[0], 2)
	}).
	Run("Bit Manipulation", func(TC *Test) {
		b := Buffer{2, 5, 6}
		b.And(0, 1)
		TC.Identical(b[0], 0)
		b.And(1, 2)
		TC.Identical(b[1], 4)

		b = Buffer{2, 5, 6}
		b.Or(0, 1)
		TC.Identical(b[0], 7)
		b.Or(1, 2)
		TC.Identical(b[1], 7)

		b = Buffer{2, 5, 6}
		b.Xor(0, 1)
		TC.Identical(b[0], 7)
		b.Xor(0, 1)
		TC.Identical(b[0], 2)
		b.Xor(1, 2)
		TC.Identical(b[1], 3)
		b.Xor(1, 2)
		TC.Identical(b[1], 5)

		b = Buffer{128, 4, 6}
		b.ShiftRight(0, 1)
		TC.Identical(b[0], 8)
		b.ShiftLeft(0, 1)
		TC.Identical(b[0], 128)
		b.Invert(0)
		TC.	Identical(b[0], ^128).
			Identical(b[0], -129)
	}).
	Run("Integer Logic", func(TC *Test) {
		b1 := Buffer{-5, 0, 17}
		TC.	Confirm(b1.LessThan(0, 1)).
			Refute(b1.Equals(0, 1)).
			Refute(b1.GreaterThan(0, 1)).
			Confirm(b1.LessThanZero(0)).
			Refute(b1.EqualsZero(0)).
			Confirm(b1.EqualsZero(1)).
			Confirm(b1.GreaterThanZero(2))

		b2 := *(b1.Clone())
		b2.Copy(1, 2)
		TC.	Different(b1, b2).
			Identical(b2[1], 17).
			Confirm(b2.LessThan(0, 1)).
			Confirm(b2.Equals(1, 2)).
			Refute(b2.GreaterThan(0, 2)).
			Refute(b2.LessThanZero(1)).
			Refute(b2.EqualsZero(1)).
			Confirm(b2.GreaterThanZero(1))

		b2.Set(0, -5, 0, 17)
		TC.	Identical(b2, b1)
	}).
	Run("Floating-Point Maths", func(TC *Test) {
		b := make(Buffer, 2)
		b.FSet(0, 37.0, 101.0)
		TC.Identical(b.FAt(0), 37.0)
		TC.Identical(b.FAt(1), 101.0)

		b.FAdd(0, 1)
		TC.	Identical(b.FAt(0), 138.0)
		b.FSubtract(0, 1)
		TC.Identical(b.FAt(0), 37.0)

		b.FMultiply(0, 1)
		TC.Identical(b.FAt(0), 3737.0)
		b.FDivide(0, 1)
		TC.Identical(b.FAt(0), 37.0)

		b.FNegate(0)
		TC.Identical(b.FAt(0), -37.0)
	}).
	Run("Floating-point Logic", func(TC *Test) {
		b := make(Buffer, 2)
		b.FSet(0, 37.0, 101.0)
		TC.	Confirm(b.FLessThan(0, 1)).
			Refute(b.FEquals(0, 1, 0.001)).
			Refute(b.FGreaterThan(0, 1)).
			Refute(b.FLessThanZero(0)).
			Refute(b.FEqualsZero(0, 0.001)).
			Confirm(b.FGreaterThanZero(0))
	}).
	Run("To Do", func(TC *Test) {
		TC.	Untested("GetBuffer").
			Untested("PutBuffer").
			Untested("Clear")
	})
}
