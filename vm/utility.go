package vm

import "unsafe"

func intsToFloats(i []int) []float { return *(*[]float)(unsafe.Pointer(&i)) }
func floatsToInts(f []float) []int { return *(*[]int)(unsafe.Pointer(&f)) }
