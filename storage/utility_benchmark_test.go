package storage

import "testing"

func BenchmarkUtilityFloatsToInts1(b *testing.B) {
	b.StopTimer()
		a := make([]float, 1, 1)
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkUtilityFloatsToInts10(b *testing.B) {
	b.StopTimer()
		a := make([]float, 10, 10)
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkUtilityFloatsToInts100(b *testing.B) {
	b.StopTimer()
		a := make([]float, 100, 100)
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkUtilityFloatsToInts1000(b *testing.B) {
	b.StopTimer()
		a := make([]float, 1000, 1000)
	b.StartTimer()
    for i := 0; i < b.N; i++ { floatsToInts(a) }
}

func BenchmarkUtilityFloatsFromInts1(b *testing.B) {
	b.StopTimer()
		a := make([]int, 1, 1)
	b.StartTimer()
    for i := 0; i < b.N; i++ { intsToFloats(a) }
}

func BenchmarkUtilityFloatsFromInts10(b *testing.B) {
	b.StopTimer()
		a := make([]int, 10, 10)
	b.StartTimer()
    for i := 0; i < b.N; i++ { intsToFloats(a) }
}

func BenchmarkUtilityFloatsFromInts100(b *testing.B) {
	b.StopTimer()
		a := make([]int, 100, 100)
	b.StartTimer()
    for i := 0; i < b.N; i++ { intsToFloats(a) }
}

func BenchmarkUtilityFloatsFromInts1000(b *testing.B) {
	b.StopTimer()
		a := make([]int, 1000, 1000)
	b.StartTimer()
    for i := 0; i < b.N; i++ { intsToFloats(a) }
}
