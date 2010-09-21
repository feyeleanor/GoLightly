package storage

import "testing"
import . "golightly/test"

func TestIntBuffer(t *testing.T) {
	NewTest(t).
	Run("Memory", func(TC *Test) {
		b := IntBuffer{0, 1, 2, 3, 4, 5}
		b1 := b.Clone()

		b1 = b.Replicate(2)
		TC.	Different(b, b1).
			Different(cap(b1), cap(b)).
			Identical(cap(b1), cap(b) * 2).
			Identical(len(b1), len(b) * 2)

		b1 = b.Clone()
		b1.Copy(0, 1)
		TC.	Different(b, b1).
			Identical(b1[0], b1[1])

		b1 = b.Clone()
		b1.Swap(2, 4)
		TC.	Different(b, b1).
			Identical(b1[2], b[4]).
			Identical(b1[4], b[2])

		b1 = b.Clone()
		b1.Clear(0, 6)
		TC.	Different(b, b1).
			Identical(b1, make(IntBuffer, 6))
	}).
	Run("Resizing", func(TC *Test) {
		b := IntBuffer{0, 1, 2, 3, 4, 5}
		b1 := b.Clone()
		b1.Resize(10)
		TC.	Different(b, b1).
			Identical(cap(b1), 10).
			Identical(len(b1), 10)

		b1.Resize(6)
		TC.	Identical(b, b1).
			Different(cap(b1), cap(b)).
			Identical(cap(b1), cap(b) + 4).
			Identical(len(b1), len(b))

		b1.Extend(0)
		TC.	Identical(b, b1).
			Identical(cap(b1), cap(b) + 4).
			Identical(len(b1), len(b))

		b1.Extend(-2)
		TC.	Different(b, b1).
			Identical(cap(b1), cap(b) + 4).
			Identical(len(b1), len(b) - 2)

		b1.Extend(6)
		TC.	Different(b, b1).
			Identical(cap(b1), cap(b) + 4).
			Identical(len(b1), len(b) + 4)

		b1.Shrink(4)
		TC.	Identical(b, b1).
			Identical(cap(b1), cap(b) + 4).
			Identical(len(b1), len(b))
	}).
	Run("Enumerators", func(TC *Test) {
		b := IntBuffer{0, 1, 2, 3, 4, 5}
		b1 := b.Collect(func(x int) int {
			return x * 2
		})
		TC. Identical(b1, IntBuffer{0, 2, 4, 6, 8, 10})

		n := b.Inject(0, func(memo, x int) int {
			return memo + x
		})
		TC. Identical(n, 15)

		n = b.Cycle(3, func(i, x int) {})
		TC.	Identical(n, 3)

		n = b.Cycle(5, func(i, x int) {})
		TC.	Identical(n, 5)

		b1 = b.Combine(b.Clone(), func(x, y int) int {
			return x + y
		})
		TC.	Identical(b1, IntBuffer{0, 2, 4, 6, 8, 10}).
			Erroneous(func() {
				IntBuffer{1, 2, 3}.Combine(b, func(x, y int) int {
					return 0
				})
			})
	}).
	Run("Arithmetic", func(TC *Test) {
		b := IntBuffer{5}
		b.Negate(0)
		TC.	Identical(b[0], -5)
		b.Negate(0)
		TC.	Identical(b[0], 5)

		b = IntBuffer{0}
		b.Increment(0)
		TC.	Identical(b[0], 1)

		b.Decrement(0)
		TC.	Identical(b[0], 0)

		b = IntBuffer{0, 1}
		b.Add(0, 1)
		TC.	Identical(b[0], 1)

		b.Subtract(0, 1)
		TC.	Identical(b[0], 0)

		b = IntBuffer{2, 5}
		b.Multiply(0, 1)
		TC.	Identical(b[0], 10)
		b.Divide(0, 1)
		TC.	Identical(b[0], 2)

		b.Remainder(0, 1)
		TC.	Identical(b[0], 2)
		b.Remainder(1, 0)
		TC.	Identical(b[1], 1)
	}).
	Run("Bit Manipulation", func(TC *Test) {
		b := IntBuffer{2, 5, 6}
		b.And(0, 1)
		TC.Identical(b[0], 0)
		b.And(1, 2)
		TC.Identical(b[1], 4)

		b = IntBuffer{2, 5, 6}
		b.Or(0, 1)
		TC.Identical(b[0], 7)
		b.Or(1, 2)
		TC.Identical(b[1], 7)

		b = IntBuffer{2, 5, 6}
		b.Xor(0, 1)
		TC.Identical(b[0], 7)
		b.Xor(0, 1)
		TC.Identical(b[0], 2)
		b.Xor(1, 2)
		TC.Identical(b[1], 3)
		b.Xor(1, 2)
		TC.Identical(b[1], 5)

		b = IntBuffer{128, 4, 6}
		b.ShiftRight(0, 1)
		TC.Identical(b[0], 8)
		b.ShiftLeft(0, 1)
		TC.Identical(b[0], 128)
		b.Invert(0)
		TC.	Identical(b[0], ^128).
			Identical(b[0], -129)
	}).
	Run("Comparisons", func(TC *Test) {
		b := IntBuffer{-5, 0, 17}
		TC.	Confirm(b.Less(0, 1)).
			Refute(b.Equal(0, 1)).
			Refute(b.Greater(0, 1)).
			Confirm(b.ZeroLess(0)).
			Refute(b.ZeroEqual(0)).
			Confirm(b.ZeroEqual(1)).
			Confirm(b.ZeroGreater(2))

		b1 := b.Clone()
		b1.Copy(1, 2)
		TC.	Different(b, b1).
			Identical(b1[1], 17).
			Confirm(b1.Less(0, 1)).
			Confirm(b1.Equal(1, 2)).
			Refute(b1.Greater(0, 2)).
			Refute(b1.ZeroLess(1)).
			Refute(b1.ZeroEqual(1)).
			Confirm(b1.ZeroGreater(1))
	}).
	Run("Reslicing", func(TC *Test) {
		b := make(IntBuffer, 6)
		h := b.SliceHeader(_INT_SIZE)
		TC.	Identical(cap(b), h.Cap).
			Identical(len(b), h.Len)

		b1 := b.Clone()
		TC.	Identical(AsByteSlice(b1), AsByteSlice(b)).
			Identical(AsIntBuffer(b1), AsIntBuffer(b), b).
			Identical(AsUintBuffer(b1), AsUintBuffer(b)).
			Identical(AsFloatBuffer(b1), AsFloatBuffer(b)).
			Identical(AsPointerBuffer(b1), AsPointerBuffer(b))

		b1[0] = 1
		TC.	Different(b1, b).
			Identical(AsIntBuffer(b), b).
			Different(AsIntBuffer(b1), AsIntBuffer(b)).
			Different(AsUintBuffer(b1), AsUintBuffer(b)).
			Different(AsFloatBuffer(b1), AsFloatBuffer(b)).
			Different(AsPointerBuffer(b1), AsPointerBuffer(b))
	}).
	Run("Finders", func(TC *Test) {
		b := IntBuffer{0, 1, 2, 3, 4, 5}
		f1 := func(x int) bool { return x >= 0 }
		f2 := func(x int) bool { return x > 1 }
		f3 := func(x int) bool { return x < 0 }
		f4 := func(x int) bool { return x == 0 }

		TC.	Identical(6, b.Count(f1)).
			Identical(4, b.Count(f2)).
			Confirm(b.Any(f1)).
			Refute(b.Any(f3)).
			Confirm(b.All(f1)).
			Refute(b.All(f2)).
			Confirm(b.None(f3)).
			Refute(b.None(f1)).
			Confirm(b.One(f4)).
			Refute(b.One(f1))
	}).
	Run("To Do", func(TC *Test) {
		TC.	Untested(".Feed").
			Untested(".Pipe")
	})
}
