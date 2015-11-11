package vm

import (
	"testing"
	. "golightly/test"
)

func TestInstructionSet(t *testing.T) {
	NewTest(t).
	Run("Creation", func(TC *Test) {
		var REGISTER int

		i := new(InstructionSet)
		i.Init()
		i.Operator("zero",	func (b []int) { REGISTER = 0 })
		i.Operator("one",	func (b []int) { REGISTER = 1 })
		i.Operator("two",	func (b []int) { REGISTER = 2 })
		i.Operator("three",	func (b []int) { REGISTER = 3 })
		i.Operator("four",	func (b []int) { REGISTER = 4 })
		TC.	Identical(i.Len(), 5)

		NewTestTable(func(x, y interface{}) interface{} {
			return y.(*InstructionSet).Exists(x.(string))
		}).
		X(			"zero",		"one",		"two",		"three",	"four",		"five"	).
		Y(	i,	true,			true,			true,			true,			true,			false	).
		Assess(TC)

		for j, f := range i.ops {
			f.(func (b []int))([]int{})
			TC.Identical(j, REGISTER)
			i.Invoke(&OpCode{Code: j, data: []int{}})
			TC.Identical(j, REGISTER)
		}

		NewTestTable(func(x, y interface{}) interface{} {
			return y.(*InstructionSet).Instruction(x.(string)).op
		}).
		X(			"zero",		"one",		"two",		"three",	"four",		"five"	).
		Y(	i,	0,				1,				2,				3,				4,		   	nil	).
		Assess(TC)

		ZERO_NIL	:= i.Assemble("zero", nil)
		ZERO_ZERO	:= i.Assemble("zero", []int{0})
		ZERO_ONE 	:= i.Assemble("zero", []int{1})
		ONE_NIL		:= i.Assemble("one", nil)
		ONE_ZERO	:= i.Assemble("one", []int{0})
		ONE_ONE		:= i.Assemble("one", []int{1})

		NewTestTable(func(x, y interface{}) interface{} {
			return x.(OpCode).Identical(y.(OpCode))
		}).
		X(							ZERO_NIL,	ZERO_ZERO,	ZERO_ONE,	ONE_NIL,	ONE_ZERO,	ONE_ONE	).
		Y(	ZERO_NIL,		true,			false,			false,		false,		false,		false	).
		Y(	ZERO_ZERO,	false,		true,				false,		false,		false,		false	).
		Y(	ZERO_ONE,		false,		false,			true,			false,		false,		false	).
		Y(	ONE_NIL,		false,		false,			false,		true,			false,		false	).
		Y(	ONE_ZERO,		false,		false,			false,		false,		true,			false	).
		Y(	ONE_ONE,		false,		false,			false,		false,		false,		true	).
		Assess(TC)

		NewTestTable(func(x, y interface{}) interface{} {
			return x.(OpCode).Similar(y.(OpCode))
		}).
		X(							ZERO_NIL,		ZERO_ZERO,	ZERO_ONE,	ONE_NIL,	ONE_ZERO,	ONE_ONE	).
		Y(	ZERO_NIL,		true,				false,			false,		false,		false,		false	).
		Y(	ZERO_ZERO,	false,			true,				true,			false,		false,		false	).
		Y(	ZERO_ONE,		false,			true,				true,			false,		false,		false	).
		Y(	ONE_NIL,		false,			false,			false,		true,			false,		false	).
		Y(	ONE_ZERO,		false,			false,			false,		false,		true,			true	).
		Y(	ONE_ONE,		false,			false,			false,		false,		true,			true	).
		Assess(TC)
	})
}