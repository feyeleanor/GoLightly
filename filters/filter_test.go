package filters

import "testing"
import . "golightly/test"

func TestFilter(t *testing.T) {
	NewTest(t).
	Run("Each", func(TC *Test) {
		r1 := []interface{}{ 1, 2, 3, "hello", 27, "dog", 3.754, uint(9) }
		r2 := uintContainer{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
		r3 := []int{ -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5 }
		r4 := []string { "cat", "dog", "rabbit" }
		n := 0
		f1 := func(x uint) {
			if x > 4 { n++ }
		}
		f2 := func(x interface{}) {
			switch x := x.(type) {
			case int:			if x > 4 { n++ }
			case uint:			if x > 4 { n++ }
			}
		}
		f3 := func(x int) {
			if x <= 2 { n++ }
		}
		NewTestTable(func(x, y interface{}) interface{} {
			n = 0
			Each(x, y)
			return n
		}).
		X(			r1,		r2,		r3,		r4		).
		Y(	f1,		0,		5,		0,		0		).
		Y(	f2,		2,		5,		1,		0		).
		Y(	f3,		0,		0,		8,		0		).
		Assess(TC)
	}).
	Run("Count", func(TC *Test) {
		r1 := Container{ 1, 2, 3, "hello", 27, "dog", 3.754 }
		r2 := Container{ 3, 4, 5, 6 }
		r3 := Container{ "hello", "dog", "dog" }
		r4 := append(r3, "cat" )
		f1 := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x > 2 && x < 10
				}
				return
			})
		f2 := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(string); ok {
					status = x == "hello" || x == "dog"
				}
				return
			})

		NewTestTable(func(x, y interface{}) interface{} {
			return Count(x, y)
		}).
		X(			r1,		r2,		r3,		r4		).
		Y(	f1,		1,		4,		0,		0		).
		Y(	f2,		2,		0,		3,		3		).
		Assess(TC)
	}).
	Run("Container Tests", func(TC *Test) {
		r1 := Container{ 1, 2, 3, "hello", 27, "dog", 3.754 }
		r2 := Container{ 3, 4, 5, 6 }
		r3 := Container{ "hello", "dog", "dog" }
		r4 := append(r3, "cat" )
		f1 := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x > 2 && x < 10
				}
				return
			})
		f2 := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(string); ok {
					status = x == "hello" || x == "dog"
				}
				return
			})

		NewTestTable(func(x, y interface{}) interface{} {
			return x.(Filterable).Count(y.(Predicate))
		}).
		X(			r1,		r2,		r3,		r4		).
		Y(	f1,		1,		4,		0,		0		).
		Y(	f2,		2,		0,		3,		3		).
		Assess(TC)

		NewTestTable(func(x, y interface{}) interface{} {
			return x.(Filterable).All(y.(Predicate))
		}).
		X(			r1,		r2,		r3,		r4		).
		Y(	f1,		false,	true,	false,	false	).
		Y(	f2,		false,	false,	true,	false	).
		Assess(TC)
	}).
	Run("Enumerator", func(TC *Test) {
		r1 := uintContainer{ 1, 2, 3, 27 }
		r2 := uintContainer{ 3, 4, 5, 6 }
		r3 := append(r2, r2...)
		r4 := append(r3, 7, 7, 11)
		n := 0
		f1 := func(x uint) {
			if x > 2 && x < 10 {
				n++
			}
		}
		f2 := func(x uint) {
			if x == 3 || x == 7 {
				n++
			}
		}
		tf := func(x, y interface{}) interface{} {
			n = 0
			Count(x, y)
			return n
		}
		NewTestTable(tf).
		X(			r1,		r2,		r3,		r4	).
		Y(	f1,		1,		4,		8,		10	).
		Y(	f2,		1,		1,		2,		4	).
		Assess(TC)
	}).
	Run("Other Tests", func(TC *Test) {
		
	})
}