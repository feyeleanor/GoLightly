package vm

import "testing"
import . "golightly/test"

func TestInstructionSet(t *testing.T) {
	NewTest(t).
	Run("Creation", func(TC *Test) {
		var REGISTER int

		i := new(InstructionSet)
		i.Init()
		i.Define("zero",	func (b *Buffer) { REGISTER = 0 })
		i.Define("one",		func (b *Buffer) { REGISTER = 1 })
		i.Define("two",		func (b *Buffer) { REGISTER = 2 })
		i.Define("three",	func (b *Buffer) { REGISTER = 3 })
		i.Define("four",	func (b *Buffer) { REGISTER = 4 })
		TC.	Identical(i.Len(), 5).
			Confirm(i.Exists("zero")).
			Confirm(i.Exists("one")).
			Confirm(i.Exists("two")).
			Confirm(i.Exists("three")).
			Confirm(i.Exists("four"))

		for j, f := range i.ops {
			o := OpCode{code: j}
			f.(func (b *Buffer))(&o.data)
			TC.Identical(REGISTER, j)
		}
		for j := 0; j < i.Len(); j++ {
			i.Invoke(&OpCode{code: j})
			TC.Identical(REGISTER, j)
		}

		TC.	Identical(i.Code("zero"), 0).
			Identical(i.Code("two"), 2).
			Identical(i.Code("five"), -1)


		ZERO_NIL	:= i.OpCode("zero", nil)
		ZERO_ZERO	:= i.OpCode("zero", &Buffer{0})
		ZERO_ONE 	:= i.OpCode("zero", &Buffer{1})
		ONE_NIL		:= i.OpCode("one", nil)
		ONE_ZERO	:= i.OpCode("one", &Buffer{0})
		ONE_ONE		:= i.OpCode("one", &Buffer{1})

		NewTestTable(func(y, x interface{}) interface{} {
			return y.(*OpCode).Identical(x.(*OpCode))
		}).
		X(				ZERO_NIL,	ZERO_ZERO,	ZERO_ONE,	ONE_NIL,	ONE_ZERO,	ONE_ONE	).
		Y(	ZERO_NIL,	true,		false,		false,		false,		false,		false	).
		Y(	ZERO_ZERO,	false,		true,		false,		false,		false,		false	).
		Y(	ZERO_ONE,	false,		false,		true,		false,		false,		false	).
		Y(	ONE_NIL,	false,		false,		false,		true,		false,		false	).
		Y(	ONE_ZERO,	false,		false,		false,		false,		true,		false	).
		Y(	ONE_ONE,	false,		false,		false,		false,		false,		true	).
		Assess(TC)

		NewTestTable(func(y, x interface{}) interface{} {
			return y.(*OpCode).Similar(x.(*OpCode))
		}).
		X(				ZERO_NIL,	ZERO_ZERO,	ZERO_ONE,	ONE_NIL,	ONE_ZERO,	ONE_ONE	).
		Y(	ZERO_NIL,	true,		false,		false,		false,		false,		false	).
		Y(	ZERO_ZERO,	false,		true,		true,		false,		false,		false	).
		Y(	ZERO_ONE,	false,		true,		true,		false,		false,		false	).
		Y(	ONE_NIL,	false,		false,		false,		true,		false,		false	).
		Y(	ONE_ZERO,	false,		false,		false,		false,		true,		true	).
		Y(	ONE_ONE,	false,		false,		false,		false,		true,		true	).
		Assess(TC)
	})
}
