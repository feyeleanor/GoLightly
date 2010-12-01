package baseline

import "testing"

func BenchmarkBaselineCastInt32ToInt(b *testing.B) {
	for i := 0; i < b.N; i++ { x = int(x32) }
}

func BenchmarkBaselineCastIntToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ { x32 = int32(x) }
}

func BenchmarkBaselineCastInt64ToUint64(b *testing.B) {
	for i := 0; i < b.N; i++ { u64 = uint64(x64) }
}

func BenchmarkBaselineCastUint64ToInt64(b *testing.B) {
	for i := 0; i < b.N; i++ { x64 = int64(u64) }
}

func BenchmarkBaselineVariableGet(b *testing.B) {
	for i := 0; i < b.N; i++ { x = x }
}

func BenchmarkBaselineVariableSet(b *testing.B) {
	for i := 0; i < b.N; i++ { x = 1 }
}

func BenchmarkBaselineVariableGetInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { in = in }
}

func BenchmarkBaselineVariableSetInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { in = 1 }
}

func BenchmarkBaselineVariableIncrement(b *testing.B) {
	for i := 0; i < b.N; i++ { x++ }
}

func BenchmarkBaselineVariableDecrement(b *testing.B) {
	for i := 0; i < b.N; i++ { x-- }
}

func BenchmarkBaselineFieldGet(b *testing.B) {
	for i := 0; i < b.N; i++ { x = dummy.i }
}

func BenchmarkBaselineFieldSet(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.i = 1 }
}

func BenchmarkBaselineSliceGet(b *testing.B) {
	for i := 0; i < b.N; i++ { x = s[0] }
}

func BenchmarkBaselineSliceSet(b *testing.B) {
	for i := 0; i < b.N; i++ { s[0] = 1 }
}

func BenchmarkBaselineMapGet(b *testing.B) {
	for i := 0; i < b.N; i++ { x = h[0] }
}

func BenchmarkBaselineMapSet(b *testing.B) {
	for i := 0; i < b.N; i++ { h[0] = 1 }
}

func BenchmarkBaselineIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i < 1000 {}
	}
}

func BenchmarkBaselineIfElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i < 1000 {} else {}
	}
}

func BenchmarkBaselineForLoopIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1; j++ {  }
	}
}

func BenchmarkBaselineForLoopIteration10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {  }
	}
}

func BenchmarkBaselineForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s {}
	}
}

func BenchmarkBaselineForRange10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s10 {}
	}
}

func BenchmarkBaselineForSliceLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s); j++ {}
	}
}

func BenchmarkBaselineForSliceLength10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s10); j++ {}
	}
}

func BenchmarkBaselineFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ { f() }
}

func BenchmarkBaselineFunctionCallArg(b *testing.B) {
	for i := 0; i < b.N; i++ { farg(1) }
}

func BenchmarkBaselineFunctionCall5VarArgs(b *testing.B) {
	for i := 0; i < b.N; i++ { fvarargs(1, 2, 3, 4, 5) }
}

func BenchmarkBaselineFunctionCallInt(b *testing.B) {
	for i := 0; i < b.N; i++ { fint(1) }
}

func BenchmarkBaselineFunctionCall5VarInts(b *testing.B) {
	for i := 0; i < b.N; i++ { fvarints(1, 2, 3, 4, 5) }
}

func BenchmarkBaselineMethodCallDirect(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1() }
}

func BenchmarkBaselineMethodCallDirect1Arg(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1arg(1) }
}

func BenchmarkBaselineMethodCallDirect1Int(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1int(1) }
}

func BenchmarkBaselineMethodCallDirect5Args(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1varargs(1, 2, 3, 4, 5) }
}

func BenchmarkBaselineMethodCallDirect5Ints(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1varints(1, 2, 3, 4, 5) }
}

func BenchmarkBaselineMethodCallIndirect(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2() }
}

func BenchmarkBaselineMethodCallIndirect1Arg(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2arg(1) }
}

func BenchmarkBaselineMethodCallIndirect1Int(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2int(1) }
}

func BenchmarkBaselineMethodCallIndirect5Args(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2varargs(1, 2, 3, 4, 5) }
}

func BenchmarkBaselineMethodCallIndirect5Ints(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2varints(1, 2, 3, 4, 5) }
}

func BenchmarkBaselineTypeAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ { in = in.(int) }
}

func BenchmarkBaselineTypeAssertionEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { in = in.(interface{}) }
}

func BenchmarkBaselineTypeAssertionInterface1(b *testing.B) {
	for i := 0; i < b.N; i++ { di = di.(dummyInterface1) }
}

func BenchmarkBaselineTypeAssertionInterface2(b *testing.B) {
	for i := 0; i < b.N; i++ { di = di.(dummyInterface2) }
}

func BenchmarkBaselineTypeCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := in.(int); ok {}
	}
}

func BenchmarkBaselineTypeCheckEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := in.(interface{}); ok {}
	}
}

func BenchmarkBaselineTypeCheckInterface1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := di.(dummyInterface1); ok {}
	}
}

func BenchmarkBaselineTypeCheckInterface2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := di.(dummyInterface2); ok {}
	}
}

func BenchmarkBaselineTypeSwitchOneCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch in := in.(type) {
		case int:
		}
	}
}

func BenchmarkBaselineTypeSwitchBasicTypesCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch in := in.(type) {
		case uint8:
		case uint16:
		case uint32:
		case uint64:
		case uint:
		case float32:
		case float64:
		case float:
		case complex64:
		case complex128:
		case int8:
		case int16:
		case int32:
		case int64:
		case int:
		}
	}
}

func BenchmarkBaselineTypeSwitchEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch in := in.(type) {
		case interface{}:
		}
	}
}

func BenchmarkBaselineTypeSwitchInterface1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch di := di.(type) {
		case dummyInterface1:
		}
	}
}

func BenchmarkBaselineTypeSwitchInterface2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch di := di.(type) {
		case dummyInterface2:
		}
	}
}

func BenchmarkBaselineNewStructureLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy = dummyStructure{} }
}

func BenchmarkBaselineNewStructure(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy = *new(dummyStructure) }
}

func BenchmarkBaselineNewSliceLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ { s = []int{0, 1} }
}

func BenchmarkBaselineNewSlice(b *testing.B) {
	for i := 0; i < b.N; i++ { s = make([]int, 2) }
}

func BenchmarkBaselineNewMapLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ { h = map [int]int{} }
}

func BenchmarkBaselineNewMap(b *testing.B) {
	for i := 0; i < b.N; i++ { h = make(map [int]int) }
}

func BenchmarkBaselineSliceCopy(b *testing.B) {
	for i := 0; i < b.N; i++ { copy(d10, s10) }
}

func BenchmarkBaselineNewSliceAppendElement1(b *testing.B) {
	a := []int{}
	for i := 0; i < b.N; i++ { a = append(a, 1) }
}

func BenchmarkBaselineNewSliceAppendElement10(b *testing.B) {
	a := []int{}
	for i := 0; i < b.N; i++ { a = append(a, s10...) }
}