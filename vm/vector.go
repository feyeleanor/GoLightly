//	TODO:	Storing and retrieving of pointers
//	TODO:	Shouldn't support Extend or Expand
//	TODO:	Interface?

package vm

import . "golightly/storage"
import "reflect"
import "unsafe"
import "math"

type Vector struct {
	IntBuffer
}

func (s *Vector) Slice(i, j int) *Vector				{ return &Vector{s.IntBuffer[i:j]} }
func (s *Vector) Clone() *Vector						{ return &Vector{s.IntBuffer.Clone()} }
func (s *Vector) Identical(o *Vector) bool				{ return reflect.DeepEqual(s.IntBuffer, o.IntBuffer) }
//func (s *Vector) FIdentical(o *Vector, t float) bool	{ return s.IntBuffer.FIdentical(&o.IntBuffer, t) }
func (s *Vector) Replace(o *Vector)						{ s.IntBuffer = o.IntBuffer }

func (s *Vector) Add(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] += o.IntBuffer[i] }
}

func (s *Vector) Subtract(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] -= o.IntBuffer[i] }
}

func (s *Vector) Multiply(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] *= o.IntBuffer[i] }
}

func (s *Vector) Divide(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] /= o.IntBuffer[i] }
}

func (s *Vector) FAdd(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.IntBuffer))
	for i := 0; i < limit; i++ { sfb[i + offset] += ofb[i] }
}

func (s *Vector) FSubtract(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.IntBuffer))
	for i := 0; i < limit; i++ { sfb[i + offset] -= ofb[i] }
}

func (s *Vector) FMultiply(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.IntBuffer))
	for i := 0; i < limit; i++ { sfb[i + offset] *= ofb[i] }
}

func (s *Vector) FDivide(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.IntBuffer))
	for i := 0; i < limit; i++ { sfb[i + offset] /= ofb[i] }
}

func (s *Vector) And(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] &= o.IntBuffer[i] }
}

func (s *Vector) Or(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] |= o.IntBuffer[i] }
}

func (s *Vector) Xor(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] ^= o.IntBuffer[i] }
}

func (s *Vector) Clear(offset, count int) {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.IntBuffer[i] = 0 }
}

func (s *Vector) ClearAll() {
	for i := range s.IntBuffer { s.IntBuffer[i] = 0 }
}

func (s *Vector) Increment(offset, count int) {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.IntBuffer[i] += 1 }
}

func (s *Vector) IncrementAll() {
	for i := range s.IntBuffer { s.IntBuffer[i] += 1 }
}

func (s *Vector) Decrement(offset, count int) {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.IntBuffer[i] -= 1 }
}

func (s *Vector) DecrementAll() {
	for i := range s.IntBuffer { s.IntBuffer[i] -= 1 }
}

func (s *Vector) Negate(offset, count int) {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.IntBuffer[i] = -s.IntBuffer[i] }
}

func (s *Vector) NegateAll() {
	for i, e := range s.IntBuffer { s.IntBuffer[i] = -e }
}

func (s *Vector) FNegate(offset, count int) {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	for i := offset; i < limit; i++ { sfb[i] = -sfb[i] }
}

func (s *Vector) FNegateAll() {
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	for i, _ := range s.IntBuffer { sfb[i] = -sfb[i] }
}

func (s *Vector) ShiftLeft(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] >>= uint(o.IntBuffer[i]) }
}

func (s *Vector) ShiftRight(offset int, o *Vector) {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	for i := 0; i < limit; i++ { s.IntBuffer[i + offset] <<= uint(o.IntBuffer[i]) }
}

func (s *Vector) Invert(offset, count int) {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.IntBuffer[i] = ^s.IntBuffer[i] }
}

func (s *Vector) Equals(offset int, o *Vector) bool {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	t := true
	for i := 0; i < limit; i++ { t = t && s.IntBuffer[i + offset] == o.IntBuffer[i] }
	return t
}

func (s *Vector) EqualsZero(offset, count int) bool {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	t := true
	for i := offset; t && i < limit; i++ { t = t && s.IntBuffer[i] == 0 }
	return t
}

func (s *Vector) LessThan(offset int, o *Vector) bool {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	t := true
	for i := 0; i < limit; i++ { t = t && s.IntBuffer[i + offset] < o.IntBuffer[i] }
	return t
}

func (s *Vector) LessThanZero(offset, count int) bool {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	t := true
	for i := offset; t && i < limit; i++ { t = t && s.IntBuffer[i] < 0 }
	return t
}

func (s *Vector) GreaterThan(offset int, o *Vector) bool {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	t := true
	for i := 0; i < limit; i++ { t = t && s.IntBuffer[i + offset] > o.IntBuffer[i] }
	return t
}

func (s *Vector) GreaterThanZero(offset, count int) bool {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	t := true
	for i := offset; t && i < limit; i++ { t = t && s.IntBuffer[i] > 0 }
	return t
}

func (s *Vector) FEquals(offset int, o *Vector, tolerance float) bool {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	t := true
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.IntBuffer))
	for i := 0; i < limit; i++ { t = t && math.Fabs(float64(sfb[i + offset] - ofb[i])) < float64(tolerance) }
	return t
}

func (s *Vector) FEqualsZero(offset, count int, tolerance float) bool {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	t := true
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	for i := offset; t && i < limit; i++ { t = t && math.Fabs(float64(sfb[i])) < float64(tolerance) }
	return t
}

func (s *Vector) FLessThan(offset int, o *Vector) bool {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.IntBuffer))
	t := true
	for i := 0; i < limit; i++ { t = t && sfb[i + offset] < ofb[i] }
	return t
}

func (s *Vector) FLessThanZero(offset, count int) bool {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	t := true
	for i := offset; t && i < limit; i++ { t = t && sfb[i] < 0 }
	return t
}

func (s *Vector) FGreaterThan(offset int, o *Vector) bool {
	var limit int
	if len(s.IntBuffer) < (len(o.IntBuffer) + offset) {
		limit = len(s.IntBuffer)
	} else {
		limit = len(o.IntBuffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.IntBuffer))
	t := true
	for i := 0; i < limit; i++ { t = t && sfb[i + offset] > ofb[i] }
	return t
}

func (s *Vector) FGreaterThanZero(offset, count int) bool {
	var limit int
	if len(s.IntBuffer) < (offset + count) {
		limit = len(s.IntBuffer)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.IntBuffer))
	t := true
	for i := offset; t && i < limit; i++ { t = t && sfb[i] > 0 }
	return t
}
