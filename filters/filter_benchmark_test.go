package filters

import "testing"

func BenchmarkAll1(b *testing.B) {
	b.StopTimer()
		c := Container{1}
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.All(f) }
}

func BenchmarkAll10(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.All(f) }
}

func BenchmarkAll100(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(10)
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.All(f) }
}

func BenchmarkAll1000(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(100)
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.All(f) }
}

func BenchmarkAll10000(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(1000)
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.All(f) }
}

func BenchmarkCount1(b *testing.B) {
	b.StopTimer()
		c := Container{1}
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkCount10(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkCount100(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(10)
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkCount1000(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(100)
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkCount10000(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(1000)
		f := Predicate(
			func(x interface{}) (status bool) {
				if x, ok := x.(int); ok {
					status = x >= 0
				}
				return
			})
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkCountOperation1(b *testing.B) {
	b.StopTimer()
		c := Container{1}
		n := 0
		f := func(x interface{}) {
			if x, ok := x.(int); ok {
				if x > 0 { n++ }
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkCountOperation10(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		n := 0
		f := func(x interface{}) {
			if x, ok := x.(int); ok {
				if x > 0 { n++ }
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkCountOperation100(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(10)
		n := 0
		f := func(x interface{}) {
			if x, ok := x.(int); ok {
				if x > 0 { n++ }
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkCountOperation1000(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(100)
		n := 0
		f := func(x interface{}) {
			if x, ok := x.(int); ok {
				if x > 0 { n++ }
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkCountOperation10000(b *testing.B) {
	b.StopTimer()
		c := Container{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(1000)
		n := 0
		f := func(x interface{}) {
			if x, ok := x.(int); ok {
				if x > 0 { n++ }
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkUintCount1(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1}
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkUintCount10(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkUintCount100(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(10)
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkUintCount1000(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(100)
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkUintCount10000(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(1000)
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Count(c, f) }
}

func BenchmarkEach1(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1}
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkEach10(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkEach100(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(10)
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkEach1000(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(100)
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkEach10000(b *testing.B) {
	b.StopTimer()
		c := uintContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(1000)
		n := 0
		f := func(x uint) {
			if x > 0 {
				n++
			}
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ { Each(c, f) }
}

func BenchmarkProveint1(b *testing.B) {
	b.StopTimer()
		c := intContainer{1}
		f := func(x int) bool { return x >= 0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.Prove(f) }
}

func BenchmarkProveint10(b *testing.B) {
	b.StopTimer()
		c := intContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		f := func(x int) bool { return x >= 0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.Prove(f) }
}

func BenchmarkProveint100(b *testing.B) {
	b.StopTimer()
		c := intContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(10)
		f := func(x int) bool { return x >= 0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.Prove(f) }
}

func BenchmarkProveint1000(b *testing.B) {
	b.StopTimer()
		c := intContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(100)
		f := func(x int) bool { return x >= 0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.Prove(f) }
}

func BenchmarkProveint10000(b *testing.B) {
	b.StopTimer()
		c := intContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(1000)
		f := func(x int) bool { return x >= 0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.Prove(f) }
}

func BenchmarkProveint1prove(b *testing.B) {
	b.StopTimer()
		c := intContainer{1}
		f := intPredicate(func(x int) bool { return x >= 0 })
	b.StartTimer()
	for i := 0; i < b.N; i++ { f.Prove(c) }
}

func BenchmarkProveint10prove(b *testing.B) {
	b.StopTimer()
		c := intContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		f := intPredicate(func(x int) bool { return x >= 0 })
	b.StartTimer()
	for i := 0; i < b.N; i++ { f.Prove(c) }
}

func BenchmarkProveint100prove(b *testing.B) {
	b.StopTimer()
		c := intContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(10)
		f := intPredicate(func(x int) bool { return x >= 0 })
	b.StartTimer()
	for i := 0; i < b.N; i++ { f.Prove(c) }
}

func BenchmarkProveint1000prove(b *testing.B) {
	b.StopTimer()
		c := intContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(100)
		f := intPredicate(func(x int) bool { return x >= 0 })
	b.StartTimer()
	for i := 0; i < b.N; i++ { f.Prove(c) }
}

func BenchmarkProveint10000prove(b *testing.B) {
	b.StopTimer()
		c := intContainer{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}.Repeat(1000)
		f := intPredicate(func(x int) bool { return x >= 0 })
	b.StartTimer()
	for i := 0; i < b.N; i++ { f.Prove(c) }
}

func BenchmarkProvefloat1(b *testing.B) {
	b.StopTimer()
		c := floatContainer{1.0}
		f := func(x float) bool { return x >= 0.0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.Prove(f) }
}

func BenchmarkProvefloat10(b *testing.B) {
	b.StopTimer()
		c := floatContainer{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 0.0}
		f := func(x float) bool { return x >= 0.0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ { c.Prove(f) }
}