package vm

import "testing"
import . "golightly/test"

func TestVector(t *testing.T) {
	b := Buffer{37, 101, 3, 5, 2, 2}
	v1 := &Vector{b}

	NewTest(t).
	Run("Creation", func(TC *Test) {
		v2 := v1.Clone()
		TC.	Identical(v1, &Vector{Buffer{37, 101, 3, 5, 2, 2}}).
			Identical(v1, v2)
	}).
	Run("Integer Maths", func(TC *Test) {
		s := v1.Slice(1, 3)
		TC.	Identical(s.Len(), s.Cap(), 2).
			Identical(s, &Vector{Buffer{101, 3}})

		v2 := v1.Clone()
		v2.Add(0, v1)
		TC.	Identical(v2, &Vector{Buffer{74, 202, 6, 10, 4, 4}})

		v2 = v1.Clone()
		v2.Subtract(0, v1)
		TC.	Identical(v2, &Vector{Buffer{0, 0, 0, 0, 0, 0}})

		v2 = v1.Clone()
		v2.Multiply(0, v1)
		TC.	Identical(v2, &Vector{Buffer{1369, 10201, 9, 25, 4, 4}})

		v2 = v1.Clone()
		v2.Divide(0, v1)
		TC.	Identical(v2, &Vector{Buffer{1, 1, 1, 1, 1, 1}})

		v2 = v1.Clone()
		v2.Negate(0, v2.Len())
		TC.	Identical(v2, &Vector{Buffer{-37, -101, -3, -5, -2, -2}})

		v2 = v1.Clone()
		v2.Increment(0, v2.Len())
		TC.	Identical(v2, &Vector{Buffer{38, 102, 4, 6, 3, 3}})

		v2 = v1.Clone()
		v2.Decrement(0, v2.Len())
		TC.	Identical(v2, &Vector{Buffer{36, 100, 2, 4, 1, 1}})
	}).
	Run("Bit manipulation", func(TC *Test) {
		v2 := v1.Clone()
		v2.Buffer.ShiftRight(0, 5)
		TC.	Identical(v2.At(0), 9)

		v2.Buffer.ShiftLeft(0, 5)
		TC.	Identical(v2.At(0), 36)

		v2.Buffer.Invert(0)
		TC.	Identical(v2.At(0), ^36, -37)
	}).
	Run("Floating-point Maths", func(TC *Test) {
		v2 := v1.Clone()
		v2.FSet(0, 37.0, 101.0, 3.7, 5.0, 2.0, 2.0)
		v3 := v2.Clone()
		v3.FAdd(0, v2)
		TC.	Identical(v3.FAt(0), 74.0).
			Identical(v3.FAt(1), 202.0).
			Identical(v3.FAt(2), 7.4).
			Identical(v3.FAt(3), 10.0).
			Identical(v3.FAt(4), 4.0).
			Identical(v3.FAt(5), 4.0)

		v3 = v2.Clone()
		v3.FSubtract(0, v2)
		TC.	Identical(v3.FAt(0), 0.0).
			Identical(v3.FAt(1), 0.0).
			Identical(v3.FAt(2), 0.0).
			Identical(v3.FAt(3), 0.0).
			Identical(v3.FAt(4), 0.0).
			Identical(v3.FAt(5), 0.0)

		v3 = v2.Clone()
		v3.FMultiply(0, v2)
		TC.	Identical(v3.FAt(0), 1369.0).
			Identical(v3.FAt(1), 10201.0).
			Identical(v3.FAt(2), 13.690001).
			Identical(v3.FAt(3), 25.0).
			Identical(v3.FAt(4), 4.0).
			Identical(v3.FAt(5), 4.0)
	}).
	Run("Integer Logic", func(TC *Test) {
		v2 := v1.Clone()
		b := v2.Buffer
		TC.	Confirm(b.LessThan(2, 3)).
			Refute(b.Equals(2, 3)).
			Refute(b.GreaterThan(2, 3)).
			Refute(b.LessThanZero(2)).
			Refute(b.EqualsZero(2)).
			Confirm(b.GreaterThanZero(2))

		v2.Copy(1, 2)
		TC.	Different(v1, v2).
			Identical(v2.At(1), 3).
			Confirm(b.LessThan(1, 3)).
			Confirm(b.Equals(1, 2)).
			Refute(b.GreaterThan(1, 3)).
			Refute(b.LessThanZero(1)).
			Refute(b.EqualsZero(1)).
			Confirm(b.GreaterThanZero(1))

		v2.Set(1, 0)
		TC.	Different(v1, v2).
			Confirm(b.LessThan(1, 3)).
			Refute(b.Equals(1, 3)).
			Refute(b.GreaterThan(1, 3)).
			Refute(b.LessThanZero(1)).
			Confirm(b.EqualsZero(1)).
			Refute(b.GreaterThanZero(1))
	}).
	Run("To Do", func(TC *Test) {
		TC.	Untested("GetBuffer").
			Untested("PutBuffer").
			Untested("Clear").
			ToDo("Improve tests for Vector-level operations")
	})
}
