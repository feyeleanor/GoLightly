package vm

import "testing"
import . "golightly/storage"
import . "golightly/test"

func TestIOController(t *testing.T) {
	NewTest(t).
	Run("Creation", func(TC *Test) {
		i := new(IOController)
		i.Init()
		TC.	Identical(i.Len(), i.Cap(), 0)

		i.Open(make(chan *Vector))
		TC.	Identical(i.Len(), i.Cap(), 1)

		i.Open(make(chan *Vector, 256))
		TC.	Identical(i.Len(), i.Cap(), 2)

		ioc := i.Clone()
		TC.	Identical(ioc.Len(), ioc.Cap(), 2)
	}).
	Run("Traffic", func(TC *Test) {
		i := new(IOController)
		i.Init()
		i.Open(make(chan *Vector))
		i.Open(make(chan *Vector, 256))

		b := &Vector{IntBuffer{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
		i.Send(0, b)
		TC.Identical(b, i.Receive(0))
	})
}
