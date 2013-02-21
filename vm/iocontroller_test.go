package vm

import "testing"
import "github.com/feyeleanor/slices"
import . "golightly/test"

func TestIOController(t *testing.T) {
	NewTest(t).
	Run("Creation", func(TC *Test) {
		i := IOController{}
		TC.	Identical(len(i), cap(i), 0)

		i = append(i, make(chan slices.ISlice))
		TC.	Identical(len(i), cap(i), 1)

		i = append(i, make(chan slices.ISlice, 256))
		TC.	Identical(len(i), cap(i), 2)

		ioc := i.Clone()
		TC.	Identical(len(ioc), cap(ioc), 2)
	}).
	Run("Traffic", func(TC *Test) {
		i := IOController{}
		i = append(i, make(chan slices.ISlice))
		i = append(i, make(chan slices.ISlice, 256))

		b := slices.ISlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		i.Send(0, b)
		TC.Identical(b, i.Receive(0))
	})
}
